package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/symbol"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// inter_frame_mode_info()
func (t *TileGroup) interFrameModeInfo(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	t.useIntrabc = 0

	if state.AvailL {
		t.LeftRefFrame[0] = state.RefFrames[state.MiRow][state.MiCol-1][0]
		t.LeftRefFrame[1] = state.RefFrames[state.MiRow][state.MiCol-1][1]
	} else {
		t.LeftRefFrame[0] = shared.INTRA_FRAME
		t.LeftRefFrame[1] = shared.NONE
	}

	if state.AvailU {
		t.AboveRefFrame[0] = state.RefFrames[state.MiRow-1][state.MiCol][0]
		t.AboveRefFrame[1] = state.RefFrames[state.MiRow-1][state.MiCol][1]
	} else {
		t.AboveRefFrame[0] = shared.INTRA_FRAME
		t.AboveRefFrame[1] = shared.NONE
	}

	t.LeftIntra = t.LeftRefFrame[0] <= shared.INTRA_FRAME
	t.AboveIntra = t.AboveRefFrame[0] <= shared.INTRA_FRAME
	t.LeftSingle = t.LeftRefFrame[1] <= shared.INTRA_FRAME
	t.AboveSingle = t.AboveRefFrame[1] <= shared.INTRA_FRAME

	t.Skip = 0
	t.interSegmentId(1, b, uh, state)
	t.readSkipMode(b, state, uh)

	if util.Bool(t.SkipMode) {
		t.Skip = 1
	} else {
		t.readSkip(b, uh, state)
	}

	if !util.Bool(uh.SegIdPreSkip) {
		t.interSegmentId(0, b, uh, state)
	}

	t.Lossless = uh.LosslessArray[t.SegmentId]
	t.readCdef(b, uh, sh, state)
	t.readDeltaQIndex(b, sh, uh, state)
	t.readDeltaLf(b, sh, uh, state)
	state.ReadDeltas = false
	t.readIsInter(b, state, uh)

	if util.Bool(t.IsInter) {
		t.interBlockModeInfo(b, state, uh, sh)
	} else {
		t.intraBlockModeInfo(b, state, uh, sh)
	}
}

// intra_block_mode_info()
func (t *TileGroup) intraBlockModeInfo(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	state.RefFrame[0] = shared.INTRA_FRAME
	state.RefFrame[1] = shared.NONE
	yMode := b.S()
	t.YMode = yMode
	t.intraAngleInfoY(b, state)

	if t.HasChroma {
		uvMode := b.S()
		t.UVMode = uvMode

		if t.UVMode == UV_CFL_PRED {
			t.readCflAlphas(b)
		}

		t.intraAngleInfoUv(b, state)
	}

	t.PaletteSizeY = 0
	t.PaletteSizeUV = 0
	if state.MiSize >= shared.BLOCK_8X8 &&
		shared.BLOCK_WIDTH[state.MiSize] <= 64 &&
		shared.BLOCK_HEIGHT[state.MiSize] <= 64 &&
		util.Bool(uh.AllowScreenContentTools) {
		t.paletteModeInfo(b, state, sh, uh)
	}

	t.filterIntraModeInfo(b, sh, state)
}

// inter_segment_id( preSkip )
func (t *TileGroup) interSegmentId(preSkip int, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State) {
	if uh.SegmentationEnabled {
		predictedSegmentId := t.getSegmentId(state)

		if util.Bool(uh.SegmentationUpdateMap) {
			if util.Bool(preSkip) && !util.Bool(uh.SegIdPreSkip) {
				t.SegmentId = 0
				return
			}
			if !util.Bool(preSkip) {
				if util.Bool(t.Skip) {
					segIdPredicted := 0

					for i := 0; i < shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]; i++ {
						t.AboveSegPredContext[state.MiCol+i] = segIdPredicted
					}
					for i := 0; i < shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]; i++ {
						t.AboveSegPredContext[state.MiRow+i] = segIdPredicted
					}
					t.readSegmentId(b, uh, state)
					return
				}
			}

			if uh.SegmentationTemporalUpdate == 1 {
				segIdPredicted := b.S()
				if util.Bool(segIdPredicted) {
					t.SegmentId = predictedSegmentId
				} else {
					t.readSegmentId(b, uh, state)
				}

				for i := 0; i < shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]; i++ {
					t.AboveSegPredContext[state.MiCol+i] = segIdPredicted
				}
				for i := 0; i < shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]; i++ {
					t.AboveSegPredContext[state.MiRow+i] = segIdPredicted
				}

			} else {
				t.readSegmentId(b, uh, state)
			}
		} else {
			t.SegmentId = predictedSegmentId
		}
	} else {
		t.SegmentId = 0
	}
}

// get_segment_id( )
func (t *TileGroup) getSegmentId(state *state.State) int {
	bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]
	xMis := util.Min(state.MiCols-state.MiCol, bw4)
	yMis := util.Min(state.MiRows-state.MiRow, bh4)
	seg := 7

	for y := 0; y < yMis; y++ {
		for x := 0; x < xMis; x++ {
			seg = util.Min(seg, state.PrevSegmentIds[state.MiRow+y][state.MiCol+x])
		}
	}

	return seg
}

// read_skip_mode()
func (t *TileGroup) readSkipMode(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	if t.segFeatureActive(shared.SEG_LVL_SKIP, uh, state) || t.segFeatureActive(shared.SEG_LVL_REF_FRAME, uh, state) || t.segFeatureActive(shared.SEG_LVL_GLOBALMV, uh, state) || !util.Bool(uh.SkipModePresent) || shared.BLOCK_WIDTH[state.MiSize] < 8 || shared.BLOCK_HEIGHT[state.MiSize] < 8 {
		t.SkipMode = 0
	} else {
		t.SkipMode = b.S()
	}
}

// read_is_inter()
func (t *TileGroup) readIsInter(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	if util.Bool(t.SkipMode) {
		t.IsInter = 1
	} else if t.segFeatureActive(shared.SEG_LVL_REF_FRAME, uh, state) {
		t.IsInter = util.Int(state.FeatureData[t.SegmentId][shared.SEG_LVL_REF_FRAME] != shared.INTRA_FRAME)
	} else if t.segFeatureActive(shared.SEG_LVL_GLOBALMV, uh, state) {
		t.IsInter = 0
	} else {
		// S()
		var ctx int
		if state.AvailU && state.AvailL {
			if t.LeftIntra && t.AboveIntra {
				ctx = 3
			} else {
				ctx = util.Int(t.LeftIntra || t.AboveIntra)
			}

		} else if state.AvailU || state.AvailL {
			if state.AvailU {
				ctx = 2 * util.Int(t.AboveIntra)
			} else {
				ctx = 2 * util.Int(t.LeftIntra)
			}
		} else {
			ctx = 0
		}

		t.IsInter = symbol.ReadSymbol(state.TileIsInterCdf[ctx], state, b, uh)
	}
}
