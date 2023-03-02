package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/util"
)

// palette_tokens()
func (t *TileGroup) paletteTokens(b *bitstream.BitStream) {
	blockHeight := t.Block_Height[t.State.MiSize]
	blockWidth := t.Block_Width[t.State.MiSize]
	onscreenHeight := util.Min(blockHeight, (t.State.MiRows-t.State.MiRow)*MI_SIZE)
	onscreenWidth := util.Min(blockWidth, (t.State.MiCols-t.State.MiCol)*MI_SIZE)

	if util.Bool(t.PaletteSizeY) {
		colorIndexMapY := b.Ns(t.PaletteSizeY)
		t.ColorMapY[0][0] = colorIndexMapY

		for i := 1; i < onscreenHeight+onscreenWidth-1; i++ {
			for j := util.Min(i, onscreenWidth-1); j >= util.Max(0, i-onscreenHeight+1); j-- {
				t.getPaletteColorContext(t.ColorMapY, (i - j), j, t.PaletteSizeY)
				paletteColorIdxY := b.S()
				t.ColorMapY[i-j][j] = t.ColorOrder[paletteColorIdxY]
			}
		}
		for i := 0; i < onscreenHeight; i++ {
			for j := onscreenWidth; j < blockWidth; j++ {
				t.ColorMapY[i][j] = t.ColorMapY[i][onscreenWidth-1]
			}
		}
		for i := onscreenHeight; i < blockHeight; i++ {
			for j := 0; j < blockWidth; j++ {
				t.ColorMapY[i][j] = t.ColorMapY[onscreenHeight-1][j]
			}
		}
	}

	if util.Bool(t.PaletteSizeUV) {
		colorIndexMapUv := b.Ns(t.PaletteSizeUV)
		t.ColorMapUV[0][0] = colorIndexMapUv
		blockHeight = blockHeight >> util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
		blockWidth = blockWidth >> util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		onscreenHeight = onscreenHeight >> util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		onscreenWidth = onscreenWidth >> util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)

		if blockWidth < 4 {
			blockWidth += 2
			onscreenWidth += 2
		}

		if blockHeight < 4 {
			blockHeight += 2
			onscreenHeight += 2
		}
		for i := 1; i < onscreenHeight+onscreenWidth-1; i++ {
			for j := util.Min(i, onscreenWidth-1); j >= util.Max(0, i-onscreenHeight+1); j-- {
				t.getPaletteColorContext(t.ColorMapUV, (i - j), j, t.PaletteSizeUV)
				paletteColorIdxUv := b.S()
				t.ColorMapUV[i-j][j] = t.ColorOrder[paletteColorIdxUv]
			}
		}
		for i := 0; i < onscreenHeight; i++ {
			for j := onscreenWidth; j < blockWidth; j++ {
				t.ColorMapUV[i][j] = t.ColorMapUV[i][onscreenWidth-1]
			}
		}
		for i := onscreenHeight; i < blockHeight; i++ {
			for j := 0; j < blockWidth; j++ {
				t.ColorMapUV[i][j] = t.ColorMapUV[onscreenHeight-1][j]
			}
		}
	}
}

// get_palette_color_context( colorMap, r, c, n )
func (t *TileGroup) getPaletteColorContext(colorMap [][]int, r int, c int, n int) {
	var scores []int
	for i := 0; i < PALETTE_COLORS; i++ {
		scores[i] = 0
		t.ColorOrder[i] = i
	}

	var neighbor int
	if c > 0 {
		neighbor = colorMap[r][c-1]
		scores[neighbor] += 2
	}

	if r > 0 && c > 0 {
		neighbor = colorMap[r-1][c-1]
		scores[neighbor] += 1
	}
	if r > 0 {
		neighbor = colorMap[r-1][c]
		scores[neighbor] += 1
	}

	for i := 0; i < PALETTE_NUM_NEIGHBORS; i++ {
		maxScore := scores[i]
		maIdx := i
		for j := i + 1; j < n; j++ {
			if scores[j] > maxScore {
				maxScore = scores[j]
				maIdx = j
			}
		}
		if maIdx != i {
			maxScore = scores[maIdx]
			maxColorOrder := t.ColorOrder[maIdx]
			for k := maIdx; k > i; k-- {
				scores[k] = scores[k-1]
				t.ColorOrder[k] = t.ColorOrder[k-1]
			}
			scores[i] = maxScore
			t.ColorOrder[i] = maxColorOrder
		}
	}

	t.ColorContextHash = 0
	for i := 0; i < PALETTE_NUM_NEIGHBORS; i++ {
		t.ColorContextHash += scores[i] * Palette_Color_Hash_Multipliers[i]
	}
}
