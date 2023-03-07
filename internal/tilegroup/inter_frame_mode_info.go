package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
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
		t.readSkip(b)
	}

	if !util.Bool(uh.SegIdPreSkip) {
		t.interSegmentId(0, b, uh, state)
	}

	t.Lossless = uh.LosslessArray[t.SegmentId]
	t.readCdef(b)
	t.readDeltaQIndex(b)
	t.readDeltaLf(b)
	state.ReadDeltas = false
	t.readIsInter(b, state)

	if util.Bool(t.IsInter) {
		t.interBlockModeInfo(b, state, uh, sh)
	} else {
		t.intraBlockModeInfo(b, state, uh)
	}
}

// intra_block_mode_info()
func (t *TileGroup) intraBlockModeInfo(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	state.RefFrame[0] = shared.INTRA_FRAME
	state.RefFrame[1] = shared.NONE
	yMode := b.S()
	t.YMode = yMode
	t.intraAngleInfoY(b)

	if t.HasChroma {
		uvMode := b.S()
		t.UVMode = uvMode

		if t.UVMode == UV_CFL_PRED {
			t.readCflAlphas(b)
		}

		t.intraAngleInfoUv(b)
	}

	t.PaletteSizeY = 0
	t.PaletteSizeUV = 0
	if state.MiSize >= shared.BLOCK_8X8 &&
		t.Block_Width[state.MiSize] <= 64 &&
		t.Block_Height[state.MiSize] <= 64 &&
		util.Bool(uh.AllowScreenContentTools) {
		t.paletteModeInfo(b)
	}

	t.filterIntraModeInfo(b)
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

					for i := 0; i < state.Num4x4BlocksWide[state.MiSize]; i++ {
						t.AboveSegPredContext[state.MiCol+i] = segIdPredicted
					}
					for i := 0; i < state.Num4x4BlocksHigh[state.MiSize]; i++ {
						t.AboveSegPredContext[state.MiRow+i] = segIdPredicted
					}
					t.readSegmentId(b)
					return
				}
			}

			if uh.SegmentationTemporalUpdate == 1 {
				segIdPredicted := b.S()
				if util.Bool(segIdPredicted) {
					t.SegmentId = predictedSegmentId
				} else {
					t.readSegmentId(b)
				}

				for i := 0; i < state.Num4x4BlocksWide[state.MiSize]; i++ {
					t.AboveSegPredContext[state.MiCol+i] = segIdPredicted
				}
				for i := 0; i < state.Num4x4BlocksHigh[state.MiSize]; i++ {
					t.AboveSegPredContext[state.MiRow+i] = segIdPredicted
				}

			} else {
				t.readSegmentId(b)
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
	bw4 := state.Num4x4BlocksWide[state.MiSize]
	bh4 := state.Num4x4BlocksHigh[state.MiSize]
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
	if t.segFeatureActive(shared.SEG_LVL_SKIP) || t.segFeatureActive(shared.SEG_LVL_REF_FRAME) || t.segFeatureActive(shared.SEG_LVL_GLOBALMV) || !util.Bool(uh.SkipModePresent) || t.Block_Width[state.MiSize] < 8 || t.Block_Height[state.MiSize] < 8 {
		t.SkipMode = 0
	} else {
		t.SkipMode = b.S()
	}
}

// read_is_inter()
func (t *TileGroup) readIsInter(b *bitstream.BitStream, state *state.State) {
	if util.Bool(t.SkipMode) {
		t.IsInter = 1
	} else if t.segFeatureActive(shared.SEG_LVL_REF_FRAME) {
		t.IsInter = util.Int(state.FeatureData[t.SegmentId][shared.SEG_LVL_REF_FRAME] != shared.INTRA_FRAME)
	} else if t.segFeatureActive(shared.SEG_LVL_GLOBALMV) {
		t.IsInter = 0
	} else {
		t.IsInter = b.S()
	}
}
