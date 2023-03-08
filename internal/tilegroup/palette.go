package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/util"
)

// palette_mode_info()
func (t *TileGroup) paletteModeInfo(b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader) {
	// TODO: this is used for initilization of has_palette_y I think
	//bSizeCtx := Mi_Width_Log2[p.MiSize] + Mi_Height_Log2[p.MiSize] - 2

	if t.YMode == DC_PRED {
		hasPaletteY := b.S()

		if util.Bool(hasPaletteY) {
			paletteSizeYMinus2 := b.S()
			t.PaletteSizeY = paletteSizeYMinus2 + 2
			cacheN := t.getPaletteCache(0, state)
			idx := 0

			for i := 0; i < cacheN && idx < t.PaletteSizeY; i++ {
				usePaletteColorCacheY := b.L(1)

				if util.Bool(usePaletteColorCacheY) {
					t.PaletteColorsY[idx] = t.PaletteCache[i]
					idx++
				}
			}

			if idx < t.PaletteSizeY {
				t.PaletteColorsY[idx] = b.L(sh.ColorConfig.BitDepth)
				idx++
			}

			var paletteBits int
			if idx < t.PaletteSizeY {
				minBits := sh.ColorConfig.BitDepth - 1
				paletteNumExtraBitsY := b.L(2)
				paletteBits = minBits + paletteNumExtraBitsY
			}

			for idx < t.PaletteSizeY {
				paletteDeltaY := b.L(paletteBits)
				paletteDeltaY++
				t.PaletteColorsY[idx] = util.Clip1(t.PaletteColorsY[idx-1]+paletteDeltaY, sh.ColorConfig.BitDepth)
				rangE := (1 << sh.ColorConfig.BitDepth) - t.PaletteColorsY[idx] - 1
				paletteBits = util.Min(paletteBits, util.CeilLog2(rangE))
				idx++
			}
			t.PaletteColorsY = util.Sort(t.PaletteColorsY, 0, t.PaletteSizeY-1)
		}
	}

	if t.HasChroma && t.UVMode == DC_PRED {
		hasPaletteUv := b.S()
		if util.Bool(hasPaletteUv) {
			paletteSizeUvMinus2 := b.S()
			t.PaletteSizeUV = paletteSizeUvMinus2 + 2
			cacheN := t.getPaletteCache(1, state)
			idx := 0

			for i := 0; i < cacheN && idx < t.PaletteSizeUV; i++ {
				usePaletteColorCacheU := b.L(1)

				if util.Bool(usePaletteColorCacheU) {
					t.PaletteColorsY[idx] = t.PaletteCache[i]
					idx++
				}
			}

			if idx < t.PaletteSizeUV {
				t.PaletteColorsU[idx] = b.L(sh.ColorConfig.BitDepth)
				idx++
			}

			var paletteBits int
			if idx < t.PaletteSizeUV {
				minBits := sh.ColorConfig.BitDepth - 3
				paletteNumExtraBitsU := b.L(2)
				paletteBits = minBits + paletteNumExtraBitsU
			}

			for idx < t.PaletteSizeUV {
				paletteDeltaU := b.L(paletteBits)
				t.PaletteColorsU[idx] = util.Clip1(t.PaletteColorsU[idx-1]+paletteDeltaU, sh.ColorConfig.BitDepth)
				rangE := (1 << sh.ColorConfig.BitDepth) - t.PaletteColorsU[idx] - 1
				paletteBits = util.Min(paletteBits, util.CeilLog2(rangE))
				idx++
			}
			t.PaletteColorsU = util.Sort(t.PaletteColorsU, 0, t.PaletteSizeUV-1)

			deltaEncodePaletteColorsv := b.L(1)

			if util.Bool(deltaEncodePaletteColorsv) {
				minBits := sh.ColorConfig.BitDepth - 4
				maxVal := 1 << sh.ColorConfig.BitDepth
				paletteNumExtraBitsv := b.L(2)
				paletteBits = minBits + paletteNumExtraBitsv
				t.PaletteColorsV[0] = b.L(sh.ColorConfig.BitDepth)

				for idx := 1; idx < t.PaletteSizeUV; idx++ {
					paletteDeltaV := b.L(paletteBits)
					if util.Bool(paletteDeltaV) {
						paletteDeltaSignBitV := b.L(1)
						if util.Bool(paletteDeltaSignBitV) {
							paletteDeltaV = -paletteDeltaV
						}
					}

					val := t.PaletteColorsV[idx-1] + paletteDeltaV
					if val < 0 {
						val += maxVal
					}
					if val >= maxVal {
						val -= maxVal
					}
					t.PaletteColorsV[idx] = util.Clip1(val, sh.ColorConfig.BitDepth)
				}
			} else {
				for idx := 0; idx < t.PaletteSizeUV; idx++ {
					t.PaletteColorsV[idx] = b.L(sh.ColorConfig.BitDepth)
				}
			}
		}
	}
}

// get_palette_cache( plane )
func (t *TileGroup) getPaletteCache(plane int, state *state.State) int {
	aboveN := 0

	if util.Bool((state.MiRow * MI_SIZE) % 64) {
		aboveN = t.PaletteSizes[plane][state.MiRow-1][state.MiCol]
	}

	leftN := 0
	if state.AvailL {
		leftN = t.PaletteSizes[plane][state.MiRow][state.MiCol-1]
	}

	aboveIdx := 0
	leftIdx := 0
	n := 0

	for aboveIdx < aboveN && leftIdx < leftN {
		aboveC := t.PaletteColors[plane][state.MiRow-1][state.MiCol][aboveIdx]
		leftC := t.PaletteColors[plane][state.MiRow][state.MiCol-1][leftIdx]

		if leftC < aboveC {
			if n == 0 || leftC != t.PaletteCache[n-1] {
				t.PaletteCache[n] = leftC
				n++
			}
			leftIdx++
		} else {
			if n == 0 || aboveC != t.PaletteCache[n-1] {
				t.PaletteCache[n] = aboveC
				n++
			}
			aboveIdx++
			if leftC == aboveC {
				leftIdx++
			}
		}
	}

	for aboveIdx < aboveN {
		val := t.PaletteColors[plane][state.MiRow-1][state.MiCol][aboveIdx]
		aboveIdx++
		if n == 0 || val != t.PaletteCache[n-1] {
			t.PaletteCache[n] = val
			n++
		}
	}
	return n
}
