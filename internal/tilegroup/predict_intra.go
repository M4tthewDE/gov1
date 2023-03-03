package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.11.2 Intra prediction process
// predict_intra( plane, x, y, haveLeft, haveAbove, haveAboveRight, haveBelowLeft, mode, log2W, log2H )
func (t *TileGroup) predictIntra(plane int, x int, y int, haveLeft bool, haveAbove bool, haveAboveRight int, haveBelowLeft int, mode int, log2W int, log2H int) {
	w := 1 << log2W
	h := 1 << log2H
	maxX := (t.State.MiCols * MI_SIZE) - 1
	maxY := (t.State.MiRows * MI_SIZE) - 1

	if plane > 0 {
		maxX = ((t.State.MiCols * MI_SIZE) >> util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)) - 1
		maxY = ((t.State.MiRows * MI_SIZE) >> util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)) - 1
	}

	for i := 0; i < w+h-1; i++ {
		if util.Int(haveAbove) == 0 && util.Int(haveLeft) == 1 {
			t.AboveRow[i] = t.State.CurrFrame[plane][y][x-1]
		} else if util.Int(haveAbove) == 0 && util.Int(haveLeft) == 0 {
			t.AboveRow[i] = (1 << (t.State.SequenceHeader.ColorConfig.BitDepth - 1)) - 1

		} else {
			aboveLimit := util.Min(maxX, x+w-1)
			if util.Bool(haveAboveRight) {
				aboveLimit = util.Min(maxX, x+2*w-1)
			}
			t.AboveRow[i] = t.State.CurrFrame[plane][y-1][util.Min(aboveLimit, x+i)]
		}
	}

	for i := 0; i < w+h-1; i++ {
		if util.Int(haveLeft) == 0 && util.Int(haveAbove) == 1 {
			t.LeftCol[i] = t.State.CurrFrame[plane][y-1][x]
		} else if util.Int(haveLeft) == 0 && util.Int(haveAbove) == 0 {
			t.AboveRow[i] = (1 << (t.State.SequenceHeader.ColorConfig.BitDepth - 1)) + 1

		} else {
			leftLimit := util.Min(maxY, y+h-1)
			if util.Bool(haveBelowLeft) {
				leftLimit = util.Min(maxY, y+2*h-1)
			}
			t.AboveRow[i] = t.State.CurrFrame[plane][util.Min(leftLimit, y+i)][x-1]
		}
	}

	if util.Int(haveAbove) == 1 && util.Int(haveLeft) == 1 {
		t.AboveRow[len(t.AboveRow)-1] = t.State.CurrFrame[plane][y-1][x-1]
	} else if util.Int(haveAbove) == 1 {
		t.AboveRow[len(t.AboveRow)-1] = t.State.CurrFrame[plane][y-1][x]
	} else if util.Int(haveLeft) == 1 {
		t.AboveRow[len(t.AboveRow)-1] = t.State.CurrFrame[plane][y][x-1]
	} else {
		t.AboveRow[len(t.AboveRow)-1] = 1 << (t.State.SequenceHeader.ColorConfig.BitDepth - 1)
	}

	t.LeftCol[len(t.LeftCol)-1] = t.AboveRow[len(t.AboveRow)-1]

	var pred [][]int
	if plane == 0 && util.Bool(t.UseFilterIntra) {
		pred = t.recursiveIntraPredictionProcess(w, h)
	} else if t.isDirectionalMode(mode) {
		pred = t.directionalIntraPredictionProcess(plane, x, y, util.Int(haveLeft), util.Int(haveAbove), mode, w, h, maxX, maxY)
	} else if mode == SMOOTH_PRED || mode == SMOOTH_V_PRED || mode == SMOOTH_H_PRED {
		pred = t.smoothIntraPredictionProcess(mode, log2W, log2H, w, h)
	} else if mode == DC_PRED {
		pred = t.dcIntraPredictionProcess(haveLeft, haveAbove, log2W, log2H, w, h)
	} else {
		pred = t.basicIntraPredictionProcess(w, h)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			t.State.CurrFrame[plane][y+i][x+j] = pred[i][j]
		}
	}
}

// 7.11.2.2 Basic intra prediction process
func (t *TileGroup) basicIntraPredictionProcess(w int, h int) [][]int {
	var pred [][]int

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			base := t.AboveRow[j] + t.LeftCol[i] - t.AboveRow[len(t.AboveRow)-1]
			pLeft := util.Abs(base - t.LeftCol[i])
			pTop := util.Abs(base - t.AboveRow[j])
			pTopLeft := util.Abs(base - t.AboveRow[len(t.AboveRow)-1])

			if pLeft <= pTop && pLeft <= pTopLeft {
				pred[i][j] = t.LeftCol[i]
			} else if pTop <= pTopLeft {
				pred[i][j] = t.AboveRow[j]
			} else {
				pred[i][j] = t.AboveRow[len(t.AboveRow)-1]
			}
		}
	}

	return pred
}

// 7.11.2.3. Recursive intra prediction process
func (t *TileGroup) recursiveIntraPredictionProcess(w int, h int) [][]int {
	w4 := w >> 2
	h2 := h >> 1

	var pred [][]int
	for i2 := 0; i2 <= h2-1; i2++ {
		for j4 := 0; j4 <= w4-1; j4++ {
			var p []int
			for i := 0; i <= 6; i++ {
				if i < 5 {
					if i2 == 0 {
						p[i] = t.AboveRow[(j4<<2)+i-1]
					} else if j4 == 0 && i == 0 {
						p[i] = t.LeftCol[(i2<<1)-1]
					} else {
						p[i] = pred[(i2<<1)-1][(j4<<2)+i-1]
					}
				} else {
					if j4 == 0 {
						p[i] = t.LeftCol[(i2<<1)+i-5]
					} else {
						p[i] = pred[(i2<<1)+i-5][(j4<<2)-1]

					}
				}
			}

			var pr int
			for i1 := 0; i1 <= 1; i1++ {
				for j1 := 0; j1 <= 3; j1++ {
					pr = 0
					for i := 0; i <= 6; i++ {
						pr += Intra_Filter_Taps[t.FilterIntraMode][(i1<<2)+j1][i] * p[i]
					}
					pred[(i2<<1)+i1][(j4<<2)+j1] = util.Clip1(util.Round2Signed(pr, INTRA_FILTER_SCALE_BITS), t.State.SequenceHeader.ColorConfig.BitDepth)
				}

			}
		}
	}

	return pred
}

// 7.11.2.4. Directional intra prediction process
func (t *TileGroup) directionalIntraPredictionProcess(plane int, x int, y int, haveLeft int, haveAbove int, mode int, w int, h int, maxX int, maxY int) [][]int {
	var pred [][]int

	angleDelta := t.AngleDeltaUV
	if plane == 0 {
		angleDelta = t.AngleDeltaY
	}

	pAngle := Mode_To_Angle[mode] + angleDelta*ANGLE_STEP
	upsampleAbove := false
	upsampleLeft := false

	if util.Int(t.State.SequenceHeader.EnableIntraEdgeFilter) == 1 {
		var filterType int
		if pAngle != 90 && pAngle != 180 {
			if pAngle > 90 && pAngle < 180 && (w+h) >= 24 {
				t.LeftCol[len(t.LeftCol)] = t.filterCornerProcess()
				t.AboveRow[len(t.AboveRow)] = t.filterCornerProcess()
			}
			filterType = util.Int(t.getFilterType(plane))

			if haveAbove == 1 {
				strength := t.intraEdgeFilterStrengthSelectionProcess(w, h, filterType, pAngle-90)
				sumPart := 0
				if pAngle < 90 {
					sumPart = h
				}
				numPx := util.Min(w, (maxX-x+1)) + sumPart
				t.intraEdgeFilterProcess(numPx, strength, 0)
			}

			if haveLeft == 1 {
				strength := t.intraEdgeFilterStrengthSelectionProcess(w, h, filterType, pAngle-180)
				sumPart := 0
				if pAngle > 180 {
					sumPart = w
				}
				numPx := util.Min(h, (maxY-y+1)) + sumPart
				t.intraEdgeFilterProcess(numPx, strength, 1)
			}
		}

		upsampleAbove = t.intraEdgeUpsampleSelectionProcess(w, h, filterType, pAngle-90)

		sumPart := 0
		if pAngle < 90 {
			sumPart = h
		}
		numPx := w + sumPart

		if upsampleAbove {
			t.intraEdgeUpsampleProcess(numPx, false)
		}

		upsampleLeft = t.intraEdgeUpsampleSelectionProcess(w, h, filterType, pAngle-180)

		sumPart = 0
		if pAngle > 180 {
			sumPart = w
		}
		numPx = h + sumPart

		if upsampleLeft {
			t.intraEdgeUpsampleProcess(numPx, true)
		}

	}

	var dx int
	if pAngle < 90 {
		dx = Dr_Intra_Derivative[pAngle]
	} else if pAngle > 90 && pAngle < 180 {
		dx = Dr_Intra_Derivative[180-pAngle]
	}

	var dy int
	if pAngle > 90 && pAngle < 180 {
		dy = Dr_Intra_Derivative[pAngle-90]
	} else if pAngle > 180 {
		dy = Dr_Intra_Derivative[270-pAngle]
	}

	if pAngle < 90 {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				idx := (i + i) * dx
				base := (idx >> (6 - util.Int(upsampleAbove))) + (j << util.Int(upsampleAbove))
				shift := ((idx << util.Int(upsampleAbove)) >> 1) & 0x1F
				maxBaseX := (w + h - 1) << util.Int(upsampleAbove)

				if base < maxBaseX {
					pred[i][j] = util.Round2(t.AboveRow[base]*(32-shift)+t.AboveRow[base+1]*shift, 5)
				} else {
					pred[i][j] = t.AboveRow[maxBaseX]
				}
			}

		}
	} else if pAngle > 90 && pAngle < 180 {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				idx := (j << 6) - (i+1)*dx
				base := (idx >> (6 - util.Int(upsampleAbove)))

				if base >= -(1 << util.Int(upsampleAbove)) {
					shift := ((idx << util.Int(upsampleAbove)) >> 1) & 0x1F
					pred[i][j] = util.Round2(t.AboveRow[base]*(32-shift)+t.AboveRow[base+1]*shift, 5)
				} else {
					idx = (i << 6) - (j+1)*dy
					base = idx >> (6 - util.Int(upsampleLeft))
					shift := ((idx << util.Int(upsampleLeft)) >> 1) & 0x1F
					pred[i][j] = util.Round2(t.LeftCol[base]*(32-shift)+t.LeftCol[base+1]*shift, 5)
				}
			}
		}

	} else if pAngle > 180 {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				idx := (j + 1) * dy
				base := (idx >> (6 >> util.Int(upsampleLeft))) + (i << util.Int(upsampleLeft))
				shift := ((idx << util.Int(upsampleLeft)) >> 1) & 0x1F
				pred[i][j] = util.Round2(t.LeftCol[base]*(32-shift)+t.LeftCol[base+1]*shift, 5)
			}
		}

	} else if pAngle == 90 {
		for j := 0; j < w; j++ {
			for i := 0; i < h; i++ {
				pred[i][j] = t.AboveRow[j]
			}
		}
	} else if pAngle == 180 {
		for j := 0; j < w; j++ {
			for i := 0; i < h; i++ {
				pred[i][j] = t.LeftCol[j]
			}
		}
	}

	return pred
}

// 7.11.2.5 DC intra prediction process
func (t *TileGroup) dcIntraPredictionProcess(haveLeft bool, haveAbove bool, log2W int, log2H int, w int, h int) [][]int {
	var pred [][]int
	if haveLeft && haveAbove {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				sum := 0
				for k := 0; k < h; k++ {
					sum += t.LeftCol[k]
				}
				for k := 0; k < w; k++ {
					sum += t.AboveRow[k]
				}
				sum += (w + h) >> 1
				avg := sum / (w + h)
				pred[i][j] = avg
			}
		}
	} else if haveLeft && !haveAbove {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				sum := 0
				for k := 0; k < h; k++ {
					sum += t.LeftCol[k]
				}
				leftAvg := util.Clip1((sum+(h>>1))>>log2H, t.State.SequenceHeader.ColorConfig.BitDepth)
				pred[i][j] = leftAvg
			}
		}

	} else if !haveLeft && !haveAbove {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				pred[i][j] = 1 << (t.State.SequenceHeader.ColorConfig.BitDepth - 1)
			}
		}

	}

	return pred
}

// 7.11.2.6 Smooth intra prediction process
func (t *TileGroup) smoothIntraPredictionProcess(mode int, log2W int, log2H int, w int, h int) [][]int {
	var pred [][]int
	if mode == SMOOTH_PRED {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var smWeightsX []int
				switch log2W {
				case 2:
					smWeightsX = Sm_Weights_Tx_4x4
				case 3:
					smWeightsX = Sm_Weights_Tx_8x8
				case 4:
					smWeightsX = Sm_Weights_Tx_16x16
				case 5:
					smWeightsX = Sm_Weights_Tx_32x32
				case 6:
					smWeightsX = Sm_Weights_Tx_64x64
				}

				var smWeightsY []int
				switch log2H {
				case 2:
					smWeightsY = Sm_Weights_Tx_4x4
				case 3:
					smWeightsY = Sm_Weights_Tx_8x8
				case 4:
					smWeightsY = Sm_Weights_Tx_16x16
				case 5:
					smWeightsY = Sm_Weights_Tx_32x32
				case 6:
					smWeightsY = Sm_Weights_Tx_64x64
				}

				smoothPred := smWeightsY[i]*t.AboveRow[j] + (256-smWeightsY[i])*t.LeftCol[h-1] + smWeightsX[j]*t.LeftCol[i] + (256-smWeightsX[j])*t.AboveRow[w-1]
				pred[i][j] = util.Round2(smoothPred, 9)
			}
		}
	} else if mode == SMOOTH_V_PRED {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var smWeights []int
				switch log2H {
				case 2:
					smWeights = Sm_Weights_Tx_4x4
				case 3:
					smWeights = Sm_Weights_Tx_8x8
				case 4:
					smWeights = Sm_Weights_Tx_16x16
				case 5:
					smWeights = Sm_Weights_Tx_32x32
				case 6:
					smWeights = Sm_Weights_Tx_64x64
				}

				smoothPred := smWeights[i]*t.AboveRow[j] + (256-smWeights[i])*t.LeftCol[h-1]
				pred[i][j] = util.Round2(smoothPred, 8)
			}
		}
	} else if mode == SMOOTH_H_PRED {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var smWeights []int
				switch log2W {
				case 2:
					smWeights = Sm_Weights_Tx_4x4
				case 3:
					smWeights = Sm_Weights_Tx_8x8
				case 4:
					smWeights = Sm_Weights_Tx_16x16
				case 5:
					smWeights = Sm_Weights_Tx_32x32
				case 6:
					smWeights = Sm_Weights_Tx_64x64
				}

				smoothPred := smWeights[j]*t.LeftCol[i] + (256-smWeights[j])*t.AboveRow[w-1]
				pred[i][j] = util.Round2(smoothPred, 8)
			}
		}
	}

	return pred
}

// 7.11.2.7. Filter corner process
func (t *TileGroup) filterCornerProcess() int {
	s := t.LeftCol[0]*5 + t.AboveRow[len(t.AboveRow)-1]*6 + t.AboveRow[0]*5
	return util.Round2(s, 4)
}

// 7.11.2.8. Intra filter type process
func (t *TileGroup) getFilterType(plane int) bool {
	aboveSmooth := false
	leftSmooth := false

	condition := t.State.AvailUChroma
	if plane == 0 {
		condition = t.State.AvailU
	}

	if condition {
		r := t.State.MiRow - 1
		c := t.State.MiCol

		if plane > 0 {
			if t.State.SequenceHeader.ColorConfig.SubsamplingX && util.Bool(t.State.MiCol&1) {
				c++
			}
			if t.State.SequenceHeader.ColorConfig.SubsamplingY && util.Bool(t.State.MiRow&1) {
				r--
			}
		}
		aboveSmooth = t.isSmooth(r, c, plane)
	}

	condition = t.State.AvailLChroma
	if plane == 0 {
		condition = t.State.AvailL
	}

	if condition {
		r := t.State.MiRow
		c := t.State.MiCol - 1

		if plane > 0 {
			if t.State.SequenceHeader.ColorConfig.SubsamplingX && util.Bool(t.State.MiCol&1) {
				c--
			}
			if t.State.SequenceHeader.ColorConfig.SubsamplingY && util.Bool(t.State.MiRow&1) {
				r++
			}
		}
		aboveSmooth = t.isSmooth(r, c, plane)
	}

	return aboveSmooth || leftSmooth
}

// 7.11.2.9. Intra edge filter strength selection process
func (t *TileGroup) intraEdgeFilterStrengthSelectionProcess(w int, h int, filterType int, delta int) int {
	d := util.Abs(delta)
	blkWh := w + h

	strength := 0
	if filterType == 0 {
		if blkWh <= 8 {
			if d >= 56 {
				strength = 1
			}
		} else if blkWh <= 12 {
			if d >= 40 {
				strength = 1
			}
		} else if blkWh <= 16 {
			if d >= 8 {
				strength = 1
			}
			if d >= 16 {
				strength = 2
			}
			if d >= 32 {
				strength = 3
			}
		} else if blkWh <= 32 {
			strength = 1
			if d >= 4 {
				strength = 2
			}
			if d >= 4 {
				strength = 3
			}
		}
	} else {
		if blkWh <= 8 {
			if d >= 40 {
				strength = 1
			}
			if d >= 64 {
				strength = 2
			}
		} else if blkWh <= 16 {
			if d >= 20 {
				strength = 1
			}
			if d >= 48 {
				strength = 2
			}
		} else if blkWh <= 24 {
			if d >= 4 {
				strength = 3
			}
		} else {
			strength = 3
		}
	}

	return strength
}

// 7.11.2.10 Intra edge upsample selection process
func (t *TileGroup) intraEdgeUpsampleSelectionProcess(w int, h int, filterType int, delta int) bool {
	d := util.Abs(delta)
	blkWh := w + h
	var useUpsample bool

	if d <= 0 || d >= 40 {
		useUpsample = false
	} else if filterType == 0 {
		useUpsample = blkWh <= 16
	} else {
		useUpsample = blkWh <= 8
	}

	return useUpsample
}

// 7.11.2.11 Intra edge upsample process
func (t *TileGroup) intraEdgeUpsampleProcess(numPx int, dir bool) {
	// does this actually modify those arrays?
	var buf []int
	if !dir {
		buf = t.AboveRow
	} else {
		buf = t.LeftCol
	}

	var dup []int
	dup[0] = buf[len(buf)-1]
	for i := -1; i < numPx; i++ {
		dup[i+2] = buf[i]
	}
	dup[numPx+2] = buf[numPx-1]

	buf[len(buf)-2] = dup[0]
	for i := 0; i < numPx; i++ {
		s := -dup[i] + (9 * dup[i+1]) + (9 * dup[i+2]) - dup[i+3]
		s = util.Clip1(util.Round2(s, 4), t.State.SequenceHeader.ColorConfig.BitDepth)
		buf[2*i-1] = s
		buf[2*i] = dup[i+2]
	}
}

// 7.11.2.12 Intra edge filter process
func (t *TileGroup) intraEdgeFilterProcess(sz int, strength int, left int) {
	if strength == 0 {
		return
	}

	var edge []int
	for i := 0; i < sz; i++ {
		if util.Bool(left) {
			edge[i] = t.LeftCol[i-1]
		} else {
			edge[i] = t.AboveRow[i-1]
		}
	}

	for i := 0; i < sz; i++ {
		s := 0
		for j := 0; j < INTRA_EDGE_TAPS; j++ {
			k := util.Clip3(0, sz-12, i-2+j)
			s += Intra_Edge_Kernel[strength-1][j] * edge[k]
			if left == 1 {
				t.LeftCol[i-1] = (s + 8) >> 4
			}
			if left == 0 {
				t.AboveRow[i-1] = (s + 8) >> 4
			}
		}
	}
}

// is_smooth( row, col, plane )
func (t *TileGroup) isSmooth(row int, col int, plane int) bool {
	var mode int
	if plane == 0 {
		mode = t.YModes[row][col]
	} else {
		if t.State.RefFrames[row][col][0] > INTRA_FRAME {
			return false
		}
		mode = t.UVModes[row][col]
	}

	return mode == SMOOTH_PRED || mode == SMOOTH_V_PRED || mode == SMOOTH_H_PRED
}
