package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/util"
)

// decode_block( r, c, subSize)
func (t *TileGroup) decodeBlock(r int, c int, subSize int, b *bitstream.BitStream) {
	t.State.MiRow = r
	t.State.MiCol = c
	t.State.MiSize = subSize
	bw4 := t.State.Num4x4BlocksWide[subSize]
	bh4 := t.State.Num4x4BlocksHigh[subSize]

	if bh4 == 1 && t.State.SequenceHeader.ColorConfig.SubsamplingY && (t.State.MiRow&1) == 0 {
		t.HasChroma = false
	} else if bw4 == 1 && t.State.SequenceHeader.ColorConfig.SubsamplingX && (t.State.MiCol&1) == 0 {
		t.HasChroma = false
	} else {
		t.HasChroma = t.State.SequenceHeader.ColorConfig.NumPlanes > 1
	}

	t.State.AvailU = t.isInside(r-1, c)
	t.State.AvailL = t.isInside(r, c-1)
	t.State.AvailUChroma = t.State.AvailU
	t.State.AvailLChroma = t.State.AvailL

	if t.HasChroma {
		if t.State.SequenceHeader.ColorConfig.SubsamplingY && bh4 == 1 {
			t.State.AvailUChroma = t.isInside(r-2, c)
		}
		if t.State.SequenceHeader.ColorConfig.SubsamplingX && bw4 == 1 {
			t.State.AvailLChroma = t.isInside(r, c-2)
		}
	} else {
		t.State.AvailUChroma = false
		t.State.AvailLChroma = false
	}

	t.modeInfo(b)
	t.paletteTokens(b)
	t.readBlockTxSize(b)

	if util.Bool(t.State.Skip) {
		t.resetBlockContext(bw4, bh4, b)
	}

	isCompound := t.State.RefFrame[1] > INTRA_FRAME

	for y := 0; y < bh4; y++ {
		for x := 0; x < bw4; x++ {
			t.YModes[r+y][c+x] = t.YMode

			if t.State.RefFrame[0] == INTRA_FRAME && t.HasChroma {
				t.UVModes[r+y][c+x] = t.UVMode
			}

			for refList := 0; refList < 2; refList++ {
				t.State.RefFrames[r+y][c+x][refList] = t.State.RefFrame[refList]
			}

			if util.Bool(t.State.IsInter) {
				if !util.Bool(t.useIntrabc) {
					t.CompGroupIdxs[r+y][c+x] = t.CompGroupIdx
					t.CompoundIdxs[r+y][c+x] = t.CompoundIdx
				}
				for dir := 0; dir < 2; dir++ {
					t.InterpFilters[r+y][c+x][dir] = t.InterpFilter[dir]
				}
				for refList := 0; refList < 1+util.Int(isCompound); refList++ {
					t.Mvs[r+y][c+x][refList] = t.Mv[refList]
				}
			}
		}
	}

	t.computePrediction(b)
}

// mode_info()
func (t *TileGroup) modeInfo(b *bitstream.BitStream) {
	if t.State.UncompressedHeader.FrameIsIntra {
		t.intraFrameModeInfo(b)
	} else {
		t.interFrameModeInfo(b)
	}
}

// reset_block_context( bw4, bh4 )
func (t *TileGroup) resetBlockContext(bw4 int, bh4 int, b *bitstream.BitStream) {
	for plane := 0; plane < 1+2*util.Int(t.HasChroma); plane++ {
		subX := 0
		subY := 0
		if plane > 0 {
			subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
			subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
		}

		for i := t.State.MiCol >> subX; i < ((t.State.MiCol + bw4) >> subX); i++ {
			t.AboveLevelContext[plane][i] = 0
			t.AboveDcContext[plane][i] = 0
		}

		for i := t.State.MiRow >> subY; i < ((t.State.MiRow + bh4) >> subY); i++ {
			t.LeftLevelContext[plane][i] = 0
			t.LeftDcContext[plane][i] = 0
		}
	}

}
