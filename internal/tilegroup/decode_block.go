package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// decode_block( r, c, subSize)
func (t *TileGroup) decodeBlock(r int, c int, subSize int, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	state.MiRow = r
	state.MiCol = c
	state.MiSize = subSize
	bw4 := state.Num4x4BlocksWide[subSize]
	bh4 := state.Num4x4BlocksHigh[subSize]

	if bh4 == 1 && sh.ColorConfig.SubsamplingY && (state.MiRow&1) == 0 {
		t.HasChroma = false
	} else if bw4 == 1 && sh.ColorConfig.SubsamplingX && (state.MiCol&1) == 0 {
		t.HasChroma = false
	} else {
		t.HasChroma = sh.ColorConfig.NumPlanes > 1
	}

	state.AvailU = t.isInside(r-1, c, state)
	state.AvailL = t.isInside(r, c-1, state)
	state.AvailUChroma = state.AvailU
	state.AvailLChroma = state.AvailL

	if t.HasChroma {
		if sh.ColorConfig.SubsamplingY && bh4 == 1 {
			state.AvailUChroma = t.isInside(r-2, c, state)
		}
		if sh.ColorConfig.SubsamplingX && bw4 == 1 {
			state.AvailLChroma = t.isInside(r, c-2, state)
		}
	} else {
		state.AvailUChroma = false
		state.AvailLChroma = false
	}

	t.modeInfo(b, uh, sh, state)
	t.paletteTokens(b, state, sh)
	t.readBlockTxSize(b, state, uh)

	if util.Bool(state.Skip) {
		t.resetBlockContext(bw4, bh4, b, state, sh)
	}

	isCompound := state.RefFrame[1] > shared.INTRA_FRAME

	for y := 0; y < bh4; y++ {
		for x := 0; x < bw4; x++ {
			t.YModes[r+y][c+x] = t.YMode

			if state.RefFrame[0] == shared.INTRA_FRAME && t.HasChroma {
				t.UVModes[r+y][c+x] = t.UVMode
			}

			for refList := 0; refList < 2; refList++ {
				state.RefFrames[r+y][c+x][refList] = state.RefFrame[refList]
			}

			if util.Bool(state.IsInter) {
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

	t.computePrediction(state, sh, uh)
	t.residual(sh, state, b, uh)

	for y := 0; y < bh4; y++ {
		for x := 0; x < bw4; x++ {
			t.IsInters[r+y][c+x] = t.IsInter
			t.SkipModes[r+y][c+x] = t.SkipMode
			t.Skips[r+y][c+x] = t.Skip
			t.TxSizes[r+y][c+x] = t.TxSize
			state.MiSizes[r+y][c+x] = state.MiSize
			t.PaletteSizes[0][r+y][c+x] = t.PaletteSizeY
			t.PaletteSizes[1][r+y][c+x] = t.PaletteSizeUV

			for i := 0; i < t.PaletteSizeY; i++ {
				t.PaletteColors[0][r+y][c+x][i] = t.PaletteColorsY[i]
			}
			for i := 0; i < t.PaletteSizeUV; i++ {
				t.PaletteColors[1][r+y][c+x][i] = t.PaletteColorsU[i]
			}
			for i := 0; i < shared.FRAME_LF_COUNT; i++ {
				state.DeltaLFs[r+y][c+x][i] = state.DeltaLF[i]
			}
		}
	}
}

// mode_info()
func (t *TileGroup) modeInfo(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader, state *state.State) {
	if uh.FrameIsIntra {
		t.intraFrameModeInfo(b, uh, state, sh)
	} else {
		t.interFrameModeInfo(b, state, uh, sh)
	}
}

// reset_block_context( bw4, bh4 )
func (t *TileGroup) resetBlockContext(bw4 int, bh4 int, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader) {
	for plane := 0; plane < 1+2*util.Int(t.HasChroma); plane++ {
		subX := 0
		subY := 0
		if plane > 0 {
			subX = util.Int(sh.ColorConfig.SubsamplingX)
			subY = util.Int(sh.ColorConfig.SubsamplingY)
		}

		for i := state.MiCol >> subX; i < ((state.MiCol + bw4) >> subX); i++ {
			t.AboveLevelContext[plane][i] = 0
			t.AboveDcContext[plane][i] = 0
		}

		for i := state.MiRow >> subY; i < ((state.MiRow + bh4) >> subY); i++ {
			t.LeftLevelContext[plane][i] = 0
			t.LeftDcContext[plane][i] = 0
		}
	}

}
