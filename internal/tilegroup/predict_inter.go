package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
	"github.com/m4tthewde/gov1/internal/wedgemask"
)

// 7.11.3 Inter prediction process
func (t *TileGroup) predictInter(plane int, x int, y int, w int, h int, candRow int, candCol int) {
	isCompound := t.State.RefFrames[candRow][candCol][1] > INTRA_FRAME

	t.roundVariablesDerivationProcess(isCompound)

	if plane == 0 && t.MotionMode == LOCALWARP {
		t.warpEstimationProcess()
	}

	if plane == 0 && t.MotionMode == LOCALWARP && t.LocalValid {
		t.LocalValid, _, _, _, _ = t.setupShearProcess(t.LocalWarpParams)
	}

	refList := 0
	refFrame := t.State.RefFrames[candRow][candCol][refList]

	var globalValid bool
	if t.YMode == shared.GLOBALMV || t.YMode == shared.GLOBAL_GLOBALMV && t.State.GmType[refFrame] > shared.TRANSLATION {
		globalValid, _, _, _, _ = t.setupShearProcess(t.State.UncompressedHeader.GmParams[refFrame])
	}

	useWarp := 0
	if w < 8 || h < 8 {
		useWarp = 0
	} else if t.State.UncompressedHeader.ForceIntegerMv {
		useWarp = 0
	} else if t.MotionMode == LOCALWARP && t.LocalValid {
		useWarp = 1
	} else if (t.YMode == shared.GLOBALMV || t.YMode == shared.GLOBAL_GLOBALMV) && t.State.GmType[refFrame] > shared.TRANSLATION && !t.isScaled(refFrame) && globalValid {
		useWarp = 2
	}

	mv := t.Mvs[candRow][candCol][refList]

	var refIdx int
	if !util.Bool(t.useIntrabc) {
		refIdx = t.State.UncompressedHeader.RefFrameIdx[refFrame-shared.LAST_FRAME]
	} else {
		refIdx = -1
		t.State.RefFrameWidth[len(t.State.RefFrameWidth)-1] = t.State.UncompressedHeader.FrameWidth
		t.State.RefFrameHeight[len(t.State.RefFrameHeight)-1] = t.State.UncompressedHeader.FrameHeight
		t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = t.State.UncompressedHeader.UpscaledWidth
	}

	startX, startY, stepX, stepY := t.motionVectorScalingProcess(plane, refIdx, x, y, mv)

	if util.Bool(t.useIntrabc) {
		t.State.RefFrameWidth[len(t.State.RefFrameWidth)-1] = t.State.MiCols * MI_SIZE
		t.State.RefFrameHeight[len(t.State.RefFrameHeight)-1] = t.State.MiRows * MI_SIZE
		t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = t.State.MiCols * MI_SIZE
	}

	var preds [][][]int
	if useWarp != 0 {
		for i8 := 0; i8 <= ((h - 1) >> 3); i8++ {
			for j8 := 0; j8 <= ((w - 1) >> 3); j8++ {
				// TODO: what exactly is supposed to happen here
				preds[refList] = t.blockWarpProcess(useWarp, plane, refList, x, y, i8, j8, w, h)
			}
		}
	}

	if useWarp == 0 {
		preds[refList] = t.blockInterPredictionProcess(plane, refIdx, startX, startY, stepX, stepY, w, h, candRow, candCol)
	}

	if isCompound {
		refList = 1

		refFrame := t.State.RefFrames[candRow][candCol][refList]

		var globalValid bool
		if t.YMode == shared.GLOBALMV || t.YMode == shared.GLOBAL_GLOBALMV && t.State.GmType[refFrame] > shared.TRANSLATION {
			globalValid, _, _, _, _ = t.setupShearProcess(t.State.UncompressedHeader.GmParams[refFrame])
		}

		useWarp := 0
		if w < 8 || h < 8 {
			useWarp = 0
		} else if t.State.UncompressedHeader.ForceIntegerMv {
			useWarp = 0
		} else if t.MotionMode == LOCALWARP && t.LocalValid {
			useWarp = 1
		} else if (t.YMode == shared.GLOBALMV || t.YMode == shared.GLOBAL_GLOBALMV) && t.State.GmType[refFrame] > shared.TRANSLATION && !t.isScaled(refFrame) && globalValid {
			useWarp = 2
		}

		mv := t.Mvs[candRow][candCol][refList]

		var refIdx int
		if !util.Bool(t.useIntrabc) {
			refIdx = t.State.UncompressedHeader.RefFrameIdx[refFrame-shared.LAST_FRAME]
		} else {
			refIdx = -1
			t.State.RefFrameWidth[len(t.State.RefFrameWidth)-1] = t.State.UncompressedHeader.FrameWidth
			t.State.RefFrameHeight[len(t.State.RefFrameHeight)-1] = t.State.UncompressedHeader.FrameHeight
			t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = t.State.UncompressedHeader.UpscaledWidth
		}

		startX, startY, stepX, stepY := t.motionVectorScalingProcess(plane, refIdx, x, y, mv)

		if util.Bool(t.useIntrabc) {
			t.State.RefFrameWidth[len(t.State.RefFrameWidth)-1] = t.State.MiCols * MI_SIZE
			t.State.RefFrameHeight[len(t.State.RefFrameHeight)-1] = t.State.MiRows * MI_SIZE
			t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = t.State.MiCols * MI_SIZE
		}

		var preds [][][]int
		if useWarp != 0 {
			for i8 := 0; i8 <= ((h - 1) >> 3); i8++ {
				for j8 := 0; j8 <= ((w - 1) >> 3); j8++ {
					// TODO: what exactly is supposed to happen here
					preds[refList] = t.blockWarpProcess(useWarp, plane, refList, x, y, i8, j8, w, h)
				}
			}
		}

		if useWarp == 0 {
			preds[refList] = t.blockInterPredictionProcess(plane, refIdx, startX, startY, stepX, stepY, w, h, candRow, candCol)
		}
	}

	if t.CompoundType == COMPOUND_WEDGE && plane == 0 {
		t.wedgeMaskProcess(w, h)
	} else if t.CompoundType == COMPOUND_INTRA {
		t.intraModeVariantMaskProcess(w, h)
	} else if t.CompoundType == COMPOUND_DIFFWTD {
		t.differenceWeightMaskProcess(preds, w, h)
	}

	if t.CompoundType == COMPOUND_DISTANCE {
		t.distanceWeightsProcess(candRow, candCol)
	}

	if !isCompound && !t.IsInterIntra {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				t.State.CurrFrame[plane][y+i][x+i] = util.Clip1(preds[0][i][j], t.State.SequenceHeader.ColorConfig.BitDepth)
			}
		}
	} else if t.CompoundType == COMPOUND_AVERAGE {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				t.State.CurrFrame[plane][y+i][x+i] = util.Clip1(util.Round2(preds[0][i][j]+preds[1][i][j], 1+t.InterPostRound), t.State.SequenceHeader.ColorConfig.BitDepth)
			}
		}
	} else if t.CompoundType == COMPOUND_DISTANCE {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				t.State.CurrFrame[plane][y+i][x+i] = util.Clip1(util.Round2(t.FwdWeight*preds[0][i][j]+t.BckWeight*preds[1][i][j], 4+t.InterPostRound), t.State.SequenceHeader.ColorConfig.BitDepth)
			}
		}
	} else {
		t.maskBlendProcess(preds, plane, x, y, w, h)
	}

	if t.MotionMode == OBMC {
		t.overlappedMotionCompensationProcess(plane, w, h)
	}
}

// 7.11.3.2 Rounding variables derivation process
func (t *TileGroup) roundVariablesDerivationProcess(isCompound bool) {
	t.InterRound0 = 3
	t.InterRound1 = 3
	if isCompound {
		t.InterRound1 = 7
	} else {
		t.InterRound1 = 11
	}

	if t.State.SequenceHeader.ColorConfig.BitDepth == 12 {
		t.InterRound0 = t.InterRound0 + 2
	}

	if t.State.SequenceHeader.ColorConfig.BitDepth == 12 && !isCompound {
		t.InterRound1 = t.InterRound1 - 2
	}

}

// 7.11.3.3  Motion vector scaling process
func (t *TileGroup) motionVectorScalingProcess(plane int, refIdx int, x int, y int, mv []int) (int, int, int, int) {
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (t.State.UncompressedHeader.FrameWidth / 2)) / t.State.UncompressedHeader.FrameWidth
	yScale := ((t.RefUpscaledHeight[refIdx] << REF_SCALE_SHIFT) + (t.State.UncompressedHeader.FrameHeight / 2)) / t.State.UncompressedHeader.FrameHeight

	subX := 0
	subY := 0
	if plane != 0 {
		subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
	}

	halfSample := (1 << (SUBPEL_BITS - 1))
	origX := ((x << SUBPEL_BITS) + ((2 * mv[1]) >> subX) + halfSample)
	origY := ((y << SUBPEL_BITS) + ((2 * mv[0]) >> subY) + halfSample)

	baseX := (origX*xScale - (halfSample << REF_SCALE_SHIFT))
	baseY := (origY*yScale - (halfSample << REF_SCALE_SHIFT))

	off := ((1 << (SCALE_SUBPEL_BITS - SUBPEL_BITS)) / 2)

	startX := (util.Round2Signed(baseX, REF_SCALE_SHIFT+SUBPEL_BITS-SCALE_SUBPEL_BITS) + off)
	startY := (util.Round2Signed(baseY, REF_SCALE_SHIFT+SUBPEL_BITS-SCALE_SUBPEL_BITS) + off)

	stepX := util.Round2Signed(xScale, REF_SCALE_SHIFT-SCALE_SUBPEL_BITS)
	stepY := util.Round2Signed(yScale, REF_SCALE_SHIFT-SCALE_SUBPEL_BITS)

	return startX, startY, stepX, stepY
}

// 7.11.3.4 Block inter prediction process
func (t *TileGroup) blockInterPredictionProcess(plane int, refIdx int, x int, y int, xStep int, yStep int, w int, h int, candRow int, candCol int) [][]int {
	var ref [][][]int
	if refIdx == -1 {
		ref = t.State.CurrFrame
	} else {
		ref = t.FrameStore[refIdx]
	}

	subX := 0
	subY := 0
	if plane != 0 {
		subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
	}

	lastX := ((t.RefUpscaledWidth[refIdx] + subX) >> subX) - 1
	lastY := ((t.State.RefFrameHeight[refIdx] + subY) >> subY) - 1

	intermediateHeight := (((h-1)*yStep + (1 << SCALE_SUBPEL_BITS) - 1) >> SCALE_SUBPEL_BITS) + 8

	interpFilter := t.InterpFilters[candRow][candCol][1]
	if w <= 4 {
		if interpFilter == shared.EIGHTTAP || interpFilter == shared.EIGHTTAP_SHARP {
			interpFilter = 4
		} else if interpFilter == shared.EIGHTTAP_SMOOTH {
			interpFilter = 5
		}
	}

	var intermediate [][]int
	for r := 0; r < intermediateHeight; r++ {
		for c := 0; c < w; c++ {
			s := 0
			p := x + xStep*c
			for t := 0; t < 8; t++ {
				s += Subpel_Filters[interpFilter][(p>>6)*SUBPEL_MASK][t] * ref[plane][util.Clip3(0, lastY, (y>>10)+r-3)][util.Clip3(0, lastX, (p>>10)+t-3)]
			}
			intermediate[r][c] = util.Round2(s, t.InterRound0)
		}
	}

	interpFilter = t.InterpFilters[candRow][candCol][0]
	if h <= 4 {
		if interpFilter == shared.EIGHTTAP || interpFilter == shared.EIGHTTAP_SHARP {
			interpFilter = 4
		} else if interpFilter == shared.EIGHTTAP_SMOOTH {
			interpFilter = 5
		}
	}

	var pred [][]int
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			s := 0
			p := (y & 1023) + yStep*r
			for t := 0; t < 8; t++ {
				s += Subpel_Filters[interpFilter][(p>>6)*SUBPEL_MASK][t] * intermediate[(p>>10)+t][c]
			}
			pred[r][c] = util.Round2(s, t.InterRound1)
		}
	}

	return pred
}

// 7.11.3.5 Block warp process
func (t *TileGroup) blockWarpProcess(useWarp int, plane int, refList int, x int, y int, i8 int, j8 int, w int, h int) [][]int {
	var pred [][]int

	refIdx := t.State.UncompressedHeader.RefFrameIdx[t.State.RefFrame[refList]-shared.LAST_FRAME]
	ref := t.FrameStore[refIdx]

	subX := 0
	subY := 0
	if plane != 0 {
		subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
	}

	lastX := ((t.RefUpscaledWidth[refIdx] + subX) >> subX) - 1
	lastY := ((t.State.RefFrameHeight[refIdx] + subY) >> subY) - 1

	srcX := (x + j8*8 + 4) << subX
	srcY := (y + i8*8 + 4) << subY

	var warpParams []int
	if useWarp == 1 {
		warpParams = t.LocalWarpParams
	} else {
		warpParams = t.State.UncompressedHeader.GmParams[t.State.RefFrame[refList]]
	}

	dstX := warpParams[2]*srcX + warpParams[3]*srcY + warpParams[0]
	dstY := warpParams[4]*srcX + warpParams[5]*srcY + warpParams[1]

	_, alpha, beta, gamma, delta := t.setupShearProcess(warpParams)

	x4 := dstX >> subX
	y4 := dstY >> subY
	ix4 := x4 >> shared.WARPEDMODEL_PREC_BITS
	sx4 := x4 & ((1 << shared.WARPEDMODEL_PREC_BITS) - 1)
	iy4 := y4 >> shared.WARPEDMODEL_PREC_BITS
	sy4 := y4 & ((1 << shared.WARPEDMODEL_PREC_BITS) - 1)

	var intermediate [][]int
	for i1 := -7; i1 < 8; i1++ {
		for i2 := -4; i2 < 4; i2++ {
			sx := sx4 + alpha*i2 + beta*i1
			offs := util.Round2(sx, shared.WARPEDDIFF_PREC_BITS) + shared.WARPEDPIXEL_PREC_SHIFTS
			s := 0

			for i3 := 0; i3 < 8; i3++ {
				s += Warped_Filters[offs][i3] * ref[plane][util.Clip3(0, lastY, iy4+i1)][util.Clip3(0, lastX, ix4+i2-3+i3)]
			}
			intermediate[(i1 + 7)][(i2 + 4)] = util.Round2(s, t.InterRound0)
		}

	}

	for i1 := -4; i1 < util.Min(4, h-i8*8-4); i1++ {
		for i2 := -4; i2 < util.Min(4, w-j8*8-4); i2++ {
			sy := sy4 + gamma*i2 + delta*i1
			offs := util.Round2(sy, shared.WARPEDDIFF_PREC_BITS) + shared.WARPEDPIXEL_PREC_SHIFTS
			s := 0

			for i3 := 0; i3 < 8; i3++ {
				s += Warped_Filters[offs][i3] * intermediate[i1+i3+4][i2+4]
			}
			pred[i8*8+i1+4][j8*8+i2+4] = util.Round2(s, t.InterRound1)
		}
	}

	return pred
}

// 7.11.3.6 Setup shear process
func (t *TileGroup) setupShearProcess(warpParams []int) (bool, int, int, int, int) {
	alpha0 := util.Clip3(-32768, 32767, warpParams[2]-(1<<shared.WARPEDMODEL_PREC_BITS))
	beta0 := util.Clip3(-32768, 32767, warpParams[3])

	divShift, divFactor := t.resolveDivisorProcess(warpParams[2])

	v := warpParams[4] << shared.WARPEDMODEL_PREC_BITS
	gamma0 := util.Clip3(-32768, 32767, util.Round2Signed(v*divFactor, divShift))
	w := warpParams[3] * warpParams[4]
	delta0 := util.Clip3(-32768, 32767, warpParams[5]-util.Round2Signed(w*divFactor, divShift)-1<<shared.WARPEDMODEL_PREC_BITS)

	alpha := util.Round2Signed(alpha0, shared.WARP_PARAM_REDUCE_BITS) << shared.WARP_PARAM_REDUCE_BITS
	beta := util.Round2Signed(beta0, shared.WARP_PARAM_REDUCE_BITS) << shared.WARP_PARAM_REDUCE_BITS
	gamma := util.Round2Signed(gamma0, shared.WARP_PARAM_REDUCE_BITS) << shared.WARP_PARAM_REDUCE_BITS
	delta := util.Round2Signed(delta0, shared.WARP_PARAM_REDUCE_BITS) << shared.WARP_PARAM_REDUCE_BITS

	warpValid := true
	if 4*util.Abs(alpha)+7*util.Abs(beta) >= (1 << shared.WARPEDMODEL_PREC_BITS) {
		warpValid = false
	}

	if 4*util.Abs(gamma)+4*util.Abs(delta) >= (1 << shared.WARPEDMODEL_PREC_BITS) {
		warpValid = false
	}

	return warpValid, alpha, beta, gamma, delta
}

// 7.11.3.7 Resolve divisor process
func (t *TileGroup) resolveDivisorProcess(d int) (int, int) {
	n := util.FloorLog2(util.Abs(d))
	e := util.Abs(d) - (1 << n)

	var f int
	if n > DIV_LUT_BITS {
		f = util.Round2(e, n-DIV_LUT_BITS)
	} else {
		f = e << (DIV_LUT_BITS - n)
	}

	divShift := n + DIV_LUT_PREC_BITS
	var divFactor int
	if d < 0 {
		divFactor = -Div_Lut[f]
	} else {
		divFactor = Div_Lut[f]
	}

	return divShift, divFactor
}

// 7.11.3.8 Warp estimation process
func (t *TileGroup) warpEstimationProcess() {
	var A [][]int
	var Bx []int
	var By []int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			A[i][j] = 0
		}

		Bx[i] = 0
		By[i] = 0
	}

	// TODO: matrix is symmetrical, so an entry is omitted!
	w4 := t.State.Num4x4BlocksWide[t.State.MiSize]
	h4 := t.State.Num4x4BlocksHigh[t.State.MiSize]
	midY := t.State.MiRow*4 + h4*2 - 1
	midX := t.State.MiCol*4 + w4*2 - 1
	suy := midY * 8
	sux := midX * 8
	duy := suy + t.Mv[0][0]
	dux := sux + t.Mv[0][1]

	for i := 0; i < t.NumSamples; i++ {
		sy := t.CandList[i][0] - suy
		sx := t.CandList[i][1] - sux
		dy := t.CandList[i][2] - duy
		dx := t.CandList[i][3] - dux

		if util.Abs(sx-dx) < LS_MV_MAX && util.Abs(sy-dy) < LS_MV_MAX {
			A[0][0] += util.LsProduct(sx, sx) + 8
			A[0][1] += util.LsProduct(sx, sy) + 4
			A[1][1] += util.LsProduct(sy, sx) + 8
			Bx[0] += util.LsProduct(sx, dx) + 8
			Bx[1] += util.LsProduct(sy, dx) + 4
			Bx[0] += util.LsProduct(sx, dy) + 4
			Bx[1] += util.LsProduct(sy, dy) + 8
		}
	}

	det := A[0][0]*A[1][1] - A[0][1]*A[0][1]

	if det == 0 {
		t.LocalValid = false
		return
	} else {
		t.LocalValid = true
	}

	divShift, divFactor := t.resolveDivisorProcess(det)

	divShift -= shared.WARPEDMODEL_PREC_BITS

	if divShift < 0 {
		divFactor = divFactor << (divShift)
		divShift = 0
	}

	t.LocalWarpParams[2] = t.diag(A[1][1]*Bx[0]-A[0][1]*Bx[1], divFactor, divShift)
	t.LocalWarpParams[3] = t.nondiag(-A[1][1]*Bx[0]+A[0][0]*Bx[1], divFactor, divShift)
	t.LocalWarpParams[4] = t.nondiag(A[1][1]*By[0]-A[0][1]*By[1], divFactor, divShift)
	t.LocalWarpParams[5] = t.diag(-A[1][1]*By[0]+A[0][0]*By[1], divFactor, divShift)

	mvx := t.Mv[0][1]
	mvy := t.Mv[0][0]
	vx := mvx*(1<<(shared.WARPEDMODEL_PREC_BITS-3)) - (midX*(t.LocalWarpParams[2]-(1<<shared.WARPEDMODEL_PREC_BITS)) + midY*t.LocalWarpParams[3])
	vy := mvy*(1<<(shared.WARPEDMODEL_PREC_BITS-3)) - (midX * (t.LocalWarpParams[4] + midY + (t.LocalWarpParams[5]) - (1 << shared.WARPEDMODEL_PREC_BITS)))

	t.LocalWarpParams[0] = util.Clip3(-shared.WARPEDMODEL_TRANS_CLAMP, shared.WARPEDMODEL_TRANS_CLAMP-1, vx)
	t.LocalWarpParams[1] = util.Clip3(-shared.WARPEDMODEL_TRANS_CLAMP, shared.WARPEDMODEL_TRANS_CLAMP-1, vy)
}

// 7.11.3.9 Overlapped motion compensation process
func (t *TileGroup) overlappedMotionCompensationProcess(plane int, w int, h int) {
	subX := 0
	subY := 0
	if plane != 0 {
		subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
	}

	if t.State.AvailU {
		if t.getPlaneResidualSize(t.State.MiSize, plane) >= shared.BLOCK_8X8 {
			pass := 0
			w4 := t.State.Num4x4BlocksWide[t.State.MiSize]
			x4 := t.State.MiCol
			y4 := t.State.MiRow
			nCount := 0
			nLimit := util.Min(4, shared.MI_WIDTH_LOG2[t.State.MiSize])

			for nCount < nLimit && x4 < util.Min(t.State.MiCols, t.State.MiCol+w4) {
				candRow := t.State.MiRow - 1
				candCol := x4 | 1
				candSz := t.State.MiSizes[candRow][candCol]
				step4 := util.Clip3(2, 16, t.State.Num4x4BlocksWide[candSz])
				if t.State.RefFrames[candRow][candCol][0] > INTRA_FRAME {
					nCount += 1
					predW := util.Min(w, (step4*MI_SIZE)>>subX)
					predH := util.Min(h>>1, 32>>subY)
					mask := util.GetObmcMask(predW)

					// predict_overlap( )
					mv := t.Mvs[candRow][candCol][0]
					refIdx := t.State.UncompressedHeader.RefFrameIdx[t.State.RefFrames[candRow][candCol][0]-shared.LAST_FRAME]
					predX := (x4 * 4) >> subX
					predY := (y4 * 4) >> subY
					startX, startY, stepX, stepY := t.motionVectorScalingProcess(plane, refIdx, predX, predY, mv)
					obmcPred := t.blockInterPredictionProcess(plane, refIdx, startX, startY, stepX, stepY, predW, predH, candRow, candCol)

					for i := 0; i < predH; i++ {
						for j := 0; j < predW; j++ {
							obmcPred[i][j] = util.Clip1(obmcPred[i][j], t.State.SequenceHeader.ColorConfig.BitDepth)
						}
					}

					t.overlapBlendingProcess(plane, predX, predY, predW, predH, util.Bool(pass), obmcPred, mask)
				}
				x4 += step4
			}
		}
	}

	if t.State.AvailL {
		pass := 0
		h4 := t.State.Num4x4BlocksHigh[t.State.MiSize]
		x4 := t.State.MiCol
		y4 := t.State.MiRow
		nCount := 0
		nLimit := util.Min(4, shared.MI_HEIGHT_LOG2[t.State.MiSize])

		for nCount < nLimit && y4 < util.Min(t.State.MiRows, t.State.MiRow+h4) {
			candCol := t.State.MiCol - 1
			candRow := y4 | 1
			candSz := t.State.MiSizes[candRow][candCol]
			step4 := util.Clip3(2, 16, t.State.Num4x4BlocksHigh[candSz])
			if t.State.RefFrames[candRow][candCol][0] > INTRA_FRAME {
				nCount += 1
				predW := util.Min(w>>1, 32>>subX)
				predH := util.Min(h, (step4*MI_SIZE)>>subY)
				mask := util.GetObmcMask(predW)

				// predict_overlap( )
				mv := t.Mvs[candRow][candCol][0]
				refIdx := t.State.UncompressedHeader.RefFrameIdx[t.State.RefFrames[candRow][candCol][0]-shared.LAST_FRAME]
				predX := (x4 * 4) >> subX
				predY := (y4 * 4) >> subY
				startX, startY, stepX, stepY := t.motionVectorScalingProcess(plane, refIdx, predX, predY, mv)
				obmcPred := t.blockInterPredictionProcess(plane, refIdx, startX, startY, stepX, stepY, predW, predH, candRow, candCol)

				for i := 0; i < predH; i++ {
					for j := 0; j < predW; j++ {
						obmcPred[i][j] = util.Clip1(obmcPred[i][j], t.State.SequenceHeader.ColorConfig.BitDepth)
					}
				}

				t.overlapBlendingProcess(plane, predX, predY, predW, predH, util.Bool(pass), obmcPred, mask)
			}
			y4 += step4
		}
	}
}

// 7.11.3.10 Overlap blending process
func (t *TileGroup) overlapBlendingProcess(plane int, predX int, predY int, predW int, predH int, pass bool, obmcPred [][]int, mask []int) {
	for i := 0; i < predH; i++ {
		var m int
		for j := 0; j < predW; j++ {
			if !pass {
				m = mask[i]
			} else {
				m = mask[j]
			}

			t.State.CurrFrame[plane][predY+i][predX+j] = util.Round2(m*t.State.CurrFrame[plane][predY+i][predX+j]+(64-m)*obmcPred[i][j], 6)
		}
	}
}

// 7.11.3.11 Wedge mask process
func (t *TileGroup) wedgeMaskProcess(w int, h int) {
	inputState := t.State.newWedgeMaskState()
	wedgemask.InitialiseWedgeMaskTable(inputState)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			t.Mask[i][j] = wedgemask.WedgeMasks[t.State.MiSize][t.WedgeSign][t.WedgeIndex][i][j]
		}
	}
}

// 7.11.3.12 Difference weight mask process
func (t *TileGroup) differenceWeightMaskProcess(preds [][][]int, w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			diff := util.Abs(preds[0][i][j] - preds[1][i][j])
			diff = util.Round2(diff, (t.State.SequenceHeader.ColorConfig.BitDepth-8)+t.InterPostRound)
			m := util.Clip3(0, 64, 38+diff/16)
			if util.Bool(t.MaskType) {
				t.Mask[i][j] = 64 - m
			} else {
				t.Mask[i][j] = m
			}
		}
	}
}

// 7.11.3.13 Intra mode variant mask proces
func (t *TileGroup) intraModeVariantMaskProcess(w int, h int) {
	sizeScale := MAX_SB_SIZE / util.Max(h, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if t.InterIntraMode == II_V_PRED {
				t.Mask[i][j] = Ii_Weights_1d[i*sizeScale]
			} else if t.InterIntraMode == II_H_PRED {
				t.Mask[i][j] = Ii_Weights_1d[j*sizeScale]
			} else if t.InterIntraMode == II_SMOOTH_PRED {
				t.Mask[i][j] = Ii_Weights_1d[util.Min(i, j)*sizeScale]
			} else {
				t.Mask[i][j] = 32
			}
		}
	}
}

// 7.11.3.14 Mask blend process
func (t *TileGroup) maskBlendProcess(preds [][][]int, plane int, dstX int, dstY int, w int, h int) {
	subX := 0
	subY := 0
	if plane != 0 {
		subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
	}

	var m int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (!util.Bool(subX) && !util.Bool(subY)) || (util.Bool(t.InterIntra) && !util.Bool(t.WedgeInterIntra)) {
				m = t.Mask[y][x]
			} else if util.Bool(subX) && !util.Bool(subY) {
				m = util.Round2(t.Mask[y][2*x]+t.Mask[y][2*x+1], 1)
			} else if !util.Bool(subX) && util.Bool(subY) {
				m = util.Round2(t.Mask[2*y][x]+t.Mask[2*y+1][x], 1)
			} else {
				m = util.Round2(t.Mask[2*y][2*x]+t.Mask[2*y][2*x+1]+t.Mask[2*y+1][2*x]+t.Mask[2*y+1][2*x+1], 2)
			}
			if util.Bool(t.InterIntra) {
				pred0 := util.Clip1(util.Round2(preds[0][y][x], t.InterPostRound), t.State.SequenceHeader.ColorConfig.BitDepth)
				pred1 := t.State.CurrFrame[plane][y+dstY][x+dstX]
				t.State.CurrFrame[plane][y+dstY][x+dstX] = util.Round2(m*pred1+(64-m)*pred0, 6)
			} else {
				pred0 := preds[0][y][x]
				pred1 := preds[1][y][x]
				t.State.CurrFrame[plane][y+dstY][x+dstX] = util.Round2(m*pred0+(64-m)*pred1, 6+t.InterPostRound)
			}
		}
	}
}

// 7.11.3.15 Distance weights process
func (t *TileGroup) distanceWeightsProcess(candRow int, candCol int) {
	var dist []int
	for refList := 0; refList < 2; refList++ {
		h := t.State.UncompressedHeader.OrderHints[t.State.RefFrames[candRow][candCol][refList]]
		dist[refList] = util.Clip3(0, MAX_FRAME_DISTANCE, util.Abs(t.State.UncompressedHeader.GetRelativeDist(h, t.State.UncompressedHeader.OrderHint)))
	}
	d0 := dist[1]
	d1 := dist[0]
	order := util.Int(d0 <= d1)

	if d0 == 0 || d1 == 0 {
		t.FwdWeight = Quant_Dist_Lookup[3][order]
		t.BckWeight = Quant_Dist_Lookup[3][1-order]
	} else {
		var i int
		for i = 0; i < 3; i++ {
			c0 := Quant_Dist_Weight[i][order]
			c1 := Quant_Dist_Weight[i][1-order]

			if util.Bool(order) {
				if d0*c0 > d1*c1 {
					break
				}
			} else {
				if d0*c0 < d1*c1 {
					break
				}
			}
		}
		t.FwdWeight = Quant_Dist_Lookup[i][order]
		t.BckWeight = Quant_Dist_Lookup[i][1-order]
	}
}

// get_plane_residual_size( subsize, plane)
func (t *TileGroup) getPlaneResidualSize(subsize int, plane int) int {
	subx := 0
	suby := 0

	if plane > 0 {
		subx = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
		suby = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
	}

	return shared.Subsampled_Size[subsize][subx][suby]
}
