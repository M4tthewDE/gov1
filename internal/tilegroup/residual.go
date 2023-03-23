package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// residual( )
func (t *TileGroup) residual(sh sequenceheader.SequenceHeader, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) {
	widthChunks := util.Max(1, state.BlockWidth[state.MiSize]>>6)
	heightChunks := util.Max(1, state.BlockHeight[state.MiSize]>>6)

	miSizeChunk := state.MiSize
	if widthChunks > 1 || heightChunks > 1 {
		miSizeChunk = shared.BLOCK_64X64
	}

	for chunkY := 0; chunkY < heightChunks; chunkY++ {
		for chunkX := 0; chunkX < widthChunks; chunkX++ {
			miRowChunk := state.MiRow + (chunkY << 4)
			miColChunk := state.MiCol + (chunkX << 4)

			for plane := 0; plane < 1+util.Int(t.HasChroma)*2; plane++ {
				txSz := TX_4X4
				if !t.Lossless {
					txSz = getTxSize(plane, t.TxSize, sh, state)
				}

				stepX := Tx_Width[txSz] >> 2
				stepY := Tx_Height[txSz] >> 2
				planeSz := t.getPlaneResidualSize(miSizeChunk, plane, sh)
				num4x4W := shared.NUM_4X4_BLOCKS_WIDE[planeSz]
				num4x4H := shared.NUM_4X4_BLOCKS_HIGH[planeSz]

				subX := 0
				subY := 0
				if plane > 0 {
					subX = util.Int(sh.ColorConfig.SubsamplingX)
					subY = util.Int(sh.ColorConfig.SubsamplingY)
				}

				baseX := (miColChunk >> subX) * MI_SIZE
				baseY := (miRowChunk >> subY) * MI_SIZE

				if util.Bool(t.IsInter) && !t.Lossless && !util.Bool(plane) {
					t.transformTree(baseX, baseY, num4x4W*4, num4x4H*4, state, sh, b, uh)
				} else {
					baseXBlock := (state.MiCol >> subX) * MI_SIZE
					baseYBlock := (state.MiRow >> subY) * MI_SIZE

					for y := 0; y < num4x4H; y += stepY {
						for x := 0; x < num4x4W; x += stepX {
							t.transformBlock(plane,
								baseXBlock,
								baseYBlock,
								txSz, x+((chunkX<<4)>>subX),
								y+((chunkY<<4)>>subY),
								sh,
								state,
								b,
								uh,
							)
						}

					}
				}
			}
		}
	}
}

var TX_WIDTH_LOG2 = []int{
	2, 3, 4, 5, 6, 2, 3, 3, 4, 4, 5, 5, 6, 2, 4, 3, 5, 4, 6,
}

var TX_HEIGHT_LOG2 = []int{
	2, 3, 4, 5, 6, 3, 2, 4, 3, 5, 4, 6, 5, 4, 2, 5, 3, 6, 4,
}

// transform_block(plane, baseX, baseY, txSz, x, y)
func (t *TileGroup) transformBlock(plane int, baseX int, baseY int, txSz int, x int, y int, sh sequenceheader.SequenceHeader, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) {
	startX := baseX + 4*x
	startY := baseY + 4*y

	subX := 0
	subY := 0
	if plane > 0 {
		subX = util.Int(sh.ColorConfig.SubsamplingX)
		subY = util.Int(sh.ColorConfig.SubsamplingY)
	}

	row := (startY << subY) >> shared.MI_SIZE_LOG2
	col := (startY << subY) >> shared.MI_SIZE_LOG2

	sbMask := 15
	if sh.Use128x128SuperBlock {
		sbMask = 21
	}

	subBlockMiRow := row & sbMask
	subBlockMiCol := col & sbMask
	stepX := Tx_Width[txSz] >> shared.MI_SIZE_LOG2
	stepY := Tx_Height[txSz] >> shared.MI_SIZE_LOG2
	maxX := (state.MiCols * MI_SIZE) >> subX
	maxY := (state.MiRows * MI_SIZE) >> subY

	if startX >= maxX || startY >= maxY {
		return
	}

	if !util.Bool(t.IsInter) {
		if (plane == 0) && util.Bool(t.PaletteSizeY) || (plane != 0) && util.Bool(t.PaletteSizeUV) {
			t.palettePrediction(plane, startX, startY, x, y, txSz, state)
		} else {
			isCfl := (plane > 0 && t.UVMode == UV_CFL_PRED)
			var mode int
			if plane == 0 {
				mode = t.YMode
			} else {
				if isCfl {
					mode = DC_PRED
				} else {
					mode = t.UVMode
				}
			}
			log2W := TX_WIDTH_LOG2[txSz]
			log2H := TX_HEIGHT_LOG2[txSz]

			conditionL := state.AvailL
			conditionU := state.AvailU
			if plane == 0 {
				conditionL = state.AvailLChroma
				conditionU = state.AvailUChroma
			}

			// this assumes negative indeces are intended
			x1 := (subBlockMiCol >> subX) + stepX
			x2 := (subBlockMiCol >> subX) - 1
			y1 := (subBlockMiRow >> subY) + stepY
			y2 := (subBlockMiRow >> subY) - 1

			if x2 < 0 {
				x2 = len(state.BlockDecoded[plane][y1]) + x2
			}

			if y2 < 0 {
				y2 = len(state.BlockDecoded[plane]) + y2
			}

			t.predictIntra(
				plane,
				startX,
				startY,
				conditionL || x > 0,
				conditionU || y > 0,
				state.BlockDecoded[plane][y2][x1],
				state.BlockDecoded[plane][y1][x2],
				mode,
				log2W,
				log2H,
				state,
				sh)
			if isCfl {
				t.predictChromaFromLuma(plane, startX, startY, txSz, sh, state)
			}
		}

		if plane == 0 {
			t.MaxLumaW = startX + stepX*4
			t.MaxLumaH = startY + stepY*4
		}
	}

	if !util.Bool(t.Skip) {
		eob := t.coeffs(plane, startX, startY, txSz, b, uh, state, sh)
		if eob > 0 {
			t.reconstruct(plane, startX, startY, txSz, sh, uh, state)
		}
	}

	for i := 0; i < stepY; i++ {
		for j := 0; j < stepY; j++ {
			t.LoopFilterTxSizes[plane][(row>>subY)+i][(col>>subX + j)] = txSz
			state.BlockDecoded[plane][(subBlockMiRow>>subY)+i][(subBlockMiCol>>subX)+j] = true
		}
	}
}

// 7.11.4. Palette prediction process
func (t *TileGroup) palettePrediction(plane int, startX int, startY int, x int, y int, txSz int, state *state.State) {
	w := Tx_Width[txSz]
	h := Tx_Height[txSz]

	var palette []int
	if plane == 0 {
		palette = t.PaletteColorsY
	} else if plane == 1 {
		palette = t.PaletteColorsU
	} else {
		palette = t.PaletteColorsV
	}

	// map is a reserved keyword
	var mapp [][]int
	if plane == 0 {
		mapp = t.ColorMapY
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			state.CurrFrame[plane][startY+i][startX+j] = palette[mapp[y*4+i][x*4+j]]
		}
	}
}

// 7.11.5 Predict chroma from luma process
func (t *TileGroup) predictChromaFromLuma(plane int, startX int, startY int, txSz int, sh sequenceheader.SequenceHeader, state *state.State) {
	w := Tx_Width[txSz]
	h := Tx_Height[txSz]
	subX := util.Int(sh.ColorConfig.SubsamplingX)
	subY := util.Int(sh.ColorConfig.SubsamplingY)
	var alpha int
	if plane == 0 {
		alpha = t.CflAlphaU
	} else {
		alpha = t.CflAlphaV
	}

	L := [][]int{}

	lumaAvg := 0
	for i := 0; i < h; i++ {
		lumaY := (startY + i) << subY
		lumaY = util.Min(lumaY, t.MaxLumaH-(1<<subY))
		for j := 0; j < w; j++ {
			lumaX := (startX + j) << subX
			lumaX = util.Min(lumaX, t.MaxLumaW-(1<<subX))
			t := 0
			for dy := 0; dy <= subY; dy += 1 {
				for dx := 0; dx <= subX; dx += 1 {
					t += state.CurrFrame[0][lumaY+dy][lumaX+dx]
				}
			}
			v := t << (3 - subX - subY)
			L[i][j] = v
			lumaAvg += v
		}
	}
	lumaAvg = util.Round2(lumaAvg, TX_WIDTH_LOG2[txSz]+TX_HEIGHT_LOG2[txSz])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dc := state.CurrFrame[plane][startY+i][startX+j]
			scaledLuma := util.Round2Signed(alpha*(L[i][j]-lumaAvg), 6)
			state.CurrFrame[plane][startY+i][startX+j] = util.Clip1(dc+scaledLuma, sh.ColorConfig.BitDepth)
		}
	}
}

// transform_tree( startX, startY, w, h )
func (t *TileGroup) transformTree(startX int, startY int, w int, h int, state *state.State, sh sequenceheader.SequenceHeader, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) {
	maxX := state.MiCols * MI_SIZE
	maxY := state.MiRows * MI_SIZE
	if startX >= maxX || startY >= maxY {
		return
	}
	row := startY >> shared.MI_SIZE_LOG2
	col := startX >> shared.MI_SIZE_LOG2
	lumaTxSz := t.InterTxSizes[row][col]
	lumaW := Tx_Width[lumaTxSz]
	lumaH := Tx_Height[lumaTxSz]
	if w <= lumaW && h <= lumaH {
		txSz := findTxSize(w, h)
		t.transformBlock(0, startX, startY, txSz, 0, 0, sh, state, b, uh)
	} else {
		if w > h {
			t.transformTree(startX, startY, w/2, h, state, sh, b, uh)
			t.transformTree(startX+w/2, startY, w/2, h, state, sh, b, uh)
		} else if w < h {
			t.transformTree(startX, startY, w, h/2, state, sh, b, uh)
			t.transformTree(startX, startY+h/2, w, h/2, state, sh, b, uh)
		} else {
			t.transformTree(startX, startY, w/2, h/2, state, sh, b, uh)
			t.transformTree(startX+w/2, startY, w/2, h/2, state, sh, b, uh)
			t.transformTree(startX, startY+h/2, w/2, h/2, state, sh, b, uh)
			t.transformTree(startX+w/2, startY+h/2, w/2, h/2, state, sh, b, uh)
		}
	}
}

const TX_SIZES_ALL = 19

// find_tx_size( w, h )
func findTxSize(w int, h int) int {
	var txSz int
	for txSz = 0; txSz < TX_SIZES_ALL; txSz++ {
		if Tx_Width[txSz] == w && Tx_Height[txSz] == h {
			break
		}
	}
	return txSz
}

// get_tx_size( plane, txSz )
func getTxSize(plane int, txSz int, sh sequenceheader.SequenceHeader, state *state.State) int {
	if plane == 0 {
		return txSz
	}

	uvTx := Max_Tx_Size_Rect[getPlaneResidualSize(state.MiSize, plane, sh)]
	if Tx_Width[uvTx] == 64 || Tx_Height[uvTx] == 64 {
		if Tx_Width[uvTx] == 16 {
			return TX_16X32
		}

		if Tx_Height[uvTx] == 16 {
			return TX_32X16
		}

		return TX_32X32
	}

	return uvTx
}

func getPlaneResidualSize(subsize int, plane int, sh sequenceheader.SequenceHeader) int {
	subx := 0
	suby := 0

	if plane > 0 {
		subx = util.Int(sh.ColorConfig.SubsamplingX)
		suby = util.Int(sh.ColorConfig.SubsamplingY)
	}

	return shared.Subsampled_Size[subsize][subx][suby]
}
