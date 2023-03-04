package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
)

// inter_frame_mode_info()
func (t *TileGroup) interFrameModeInfo(b *bitstream.BitStream) {
	t.useIntrabc = 0

	if t.State.AvailL {
		t.LeftRefFrame[0] = t.State.RefFrames[t.State.MiRow][t.State.MiCol-1][0]
		t.LeftRefFrame[1] = t.State.RefFrames[t.State.MiRow][t.State.MiCol-1][1]
	} else {
		t.LeftRefFrame[0] = shared.INTRA_FRAME
		t.LeftRefFrame[1] = shared.NONE
	}

	if t.State.AvailU {
		t.AboveRefFrame[0] = t.State.RefFrames[t.State.MiRow-1][t.State.MiCol][0]
		t.AboveRefFrame[1] = t.State.RefFrames[t.State.MiRow-1][t.State.MiCol][1]
	} else {
		t.AboveRefFrame[0] = shared.INTRA_FRAME
		t.AboveRefFrame[1] = shared.NONE
	}

	t.LeftIntra = t.LeftRefFrame[0] <= shared.INTRA_FRAME
	t.AboveIntra = t.AboveRefFrame[0] <= shared.INTRA_FRAME
	t.LeftSingle = t.LeftRefFrame[1] <= shared.INTRA_FRAME
	t.AboveSingle = t.AboveRefFrame[1] <= shared.INTRA_FRAME

	t.Skip = 0
	t.interSegmentId(1, b)
	t.readSkipMode(b)

	if util.Bool(t.SkipMode) {
		t.Skip = 1
	} else {
		t.readSkip(b)
	}

	if !util.Bool(t.State.UncompressedHeader.SegIdPreSkip) {
		t.interSegmentId(0, b)
	}

	t.Lossless = t.State.UncompressedHeader.LosslessArray[t.SegmentId]
	t.readCdef(b)
	t.readDeltaQIndex(b)
	t.readDeltaLf(b)
	t.State.ReadDeltas = false
	t.readIsInter(b)

	if util.Bool(t.IsInter) {
		t.interBlockModeInfo(b)
	} else {
		t.intraBlockModeInfo(b)
	}
}

// intra_block_mode_info()
func (t *TileGroup) intraBlockModeInfo(b *bitstream.BitStream) {
	t.State.RefFrame[0] = shared.INTRA_FRAME
	t.State.RefFrame[1] = shared.NONE
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
	if t.State.MiSize >= shared.BLOCK_8X8 &&
		t.Block_Width[t.State.MiSize] <= 64 &&
		t.Block_Height[t.State.MiSize] <= 64 &&
		util.Bool(t.State.UncompressedHeader.AllowScreenContentTools) {
		t.paletteModeInfo(b)
	}

	t.filterIntraModeInfo(b)
}

// inter_segment_id( preSkip )
func (t *TileGroup) interSegmentId(preSkip int, b *bitstream.BitStream) {
	if t.State.UncompressedHeader.SegmentationEnabled {
		predictedSegmentId := t.getSegmentId()

		if util.Bool(t.State.UncompressedHeader.SegmentationUpdateMap) {
			if util.Bool(preSkip) && !util.Bool(t.State.UncompressedHeader.SegIdPreSkip) {
				t.SegmentId = 0
				return
			}
			if !util.Bool(preSkip) {
				if util.Bool(t.Skip) {
					segIdPredicted := 0

					for i := 0; i < t.State.Num4x4BlocksWide[t.State.MiSize]; i++ {
						t.AboveSegPredContext[t.State.MiCol+i] = segIdPredicted
					}
					for i := 0; i < t.State.Num4x4BlocksHigh[t.State.MiSize]; i++ {
						t.AboveSegPredContext[t.State.MiRow+i] = segIdPredicted
					}
					t.readSegmentId(b)
					return
				}
			}

			if t.State.UncompressedHeader.SegmentationTemporalUpdate == 1 {
				segIdPredicted := b.S()
				if util.Bool(segIdPredicted) {
					t.SegmentId = predictedSegmentId
				} else {
					t.readSegmentId(b)
				}

				for i := 0; i < t.State.Num4x4BlocksWide[t.State.MiSize]; i++ {
					t.AboveSegPredContext[t.State.MiCol+i] = segIdPredicted
				}
				for i := 0; i < t.State.Num4x4BlocksHigh[t.State.MiSize]; i++ {
					t.AboveSegPredContext[t.State.MiRow+i] = segIdPredicted
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
func (t *TileGroup) getSegmentId() int {
	bw4 := t.State.Num4x4BlocksWide[t.State.MiSize]
	bh4 := t.State.Num4x4BlocksHigh[t.State.MiSize]
	xMis := util.Min(t.State.MiCols-t.State.MiCol, bw4)
	yMis := util.Min(t.State.MiRows-t.State.MiRow, bh4)
	seg := 7

	for y := 0; y < yMis; y++ {
		for x := 0; x < xMis; x++ {
			seg = util.Min(seg, t.State.PrevSegmentIds[t.State.MiRow+y][t.State.MiCol+x])
		}
	}

	return seg
}

// read_skip_mode()
func (t *TileGroup) readSkipMode(b *bitstream.BitStream) {
	if t.segFeatureActive(shared.SEG_LVL_SKIP) || t.segFeatureActive(shared.SEG_LVL_REF_FRAME) || t.segFeatureActive(shared.SEG_LVL_GLOBALMV) || !util.Bool(t.State.UncompressedHeader.SkipModePresent) || t.Block_Width[t.State.MiSize] < 8 || t.Block_Height[t.State.MiSize] < 8 {
		t.SkipMode = 0
	} else {
		t.SkipMode = b.S()
	}
}

// read_is_inter()
func (t *TileGroup) readIsInter(b *bitstream.BitStream) {
	if util.Bool(t.SkipMode) {
		t.IsInter = 1
	} else if t.segFeatureActive(shared.SEG_LVL_REF_FRAME) {
		t.IsInter = util.Int(t.State.FeatureData[t.SegmentId][shared.SEG_LVL_REF_FRAME] != shared.INTRA_FRAME)
	} else if t.segFeatureActive(shared.SEG_LVL_GLOBALMV) {
		t.IsInter = 0
	} else {
		t.IsInter = b.S()
	}
}
