package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.14. Loop filter process
func (t *TileGroup) loopFilter(state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	for plane := 0; plane < sh.ColorConfig.NumPlanes; plane++ {
		if plane == 0 || util.Bool(uh.LoopFilterLevel[1+plane]) {
			for pass := 0; pass < 2; pass++ {
				rowStep := 1
				colStep := 1
				if plane == 0 {
					rowStep = 1 << util.Int(sh.ColorConfig.SubsamplingX)
					colStep = 1 << util.Int(sh.ColorConfig.SubsamplingY)
				}

				for row := 0; row < state.MiRows; row += rowStep {
					for col := 0; col < state.MiCols; col += colStep {
						t.loopFilterEdge(plane, util.Bool(pass), row, col, sh, uh, state)
					}
				}
			}
		}
	}
}

// 7.14.2. Edge loop filter process
func (t *TileGroup) loopFilterEdge(plane int, pass bool, row int, col int, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader, state *state.State) {
	subX := 0
	subY := 0
	if plane != 0 {
		subX = util.Int(sh.ColorConfig.SubsamplingX)
		subY = util.Int(sh.ColorConfig.SubsamplingY)
	}

	var dx int
	var dy int
	if !pass {
		dx = 1
		dy = 0
	} else {
		dy = 1
		dx = 0
	}

	x := col * MI_SIZE
	y := row * MI_SIZE

	row = row | subY
	col = col | subX

	var onScreen bool
	if x >= uh.FrameWidth {
		onScreen = false
	} else if y >= uh.FrameHeight {
		onScreen = false
	} else if !pass && x == 0 {
		onScreen = false
	} else if pass && y == 0 {
		onScreen = false
	} else {
		onScreen = true
	}

	if !onScreen {
		return
	}

	xP := x >> subX
	yP := y >> subY

	prevRow := row - (dy << subY)
	prevCol := col - (dx << subX)

	state.MiSize = state.MiSizes[row][col]
	txSz := t.LoopFilterTxSizes[plane][row>>subY][col>>subX]
	planeSize := t.getPlaneResidualSize(state.MiSize, plane, sh)
	skip := t.Skips[row][col]
	isIntra := state.RefFrames[row][col][0] <= shared.INTRA_FRAME
	prevTxSz := t.LoopFilterTxSizes[plane][prevRow>>subY][prevCol>>subX]

	var isBlockEdge bool
	if !pass && xP%shared.BLOCK_WIDTH[planeSize] == 0 {
		isBlockEdge = true
	} else if pass && yP%shared.BLOCK_HEIGHT[planeSize] == 0 {
		isBlockEdge = true
	} else {
		isBlockEdge = false
	}

	var isTxEdge bool
	if !pass && xP%Tx_Width[txSz] == 0 {
		isTxEdge = true
	} else if pass && yP%Tx_Height[txSz] == 0 {
		isTxEdge = true
	} else {
		isTxEdge = false
	}

	var applyFilter bool
	if !isTxEdge {
		applyFilter = false
	} else if isBlockEdge || skip == 0 || isIntra {
		applyFilter = true
	} else {
		applyFilter = false
	}

	filterSize := t.filterSize(txSz, prevTxSz, pass, plane)
	lvl, limit, blimit, thresh := t.adaptiveFilterStrength(row, col, plane, pass, state, uh)

	if lvl == 0 {
		lvl, limit, blimit, thresh = t.adaptiveFilterStrength(row, col, plane, pass, state, uh)
	}

	for i := 0; i < MI_SIZE; i++ {
		if applyFilter && lvl > 0 {
			t.sampleFiltering(xP+dy*i, yP+dx*i, plane, limit, blimit, thresh, dx, dy, filterSize, state, sh)
		}
	}
}

// 7.14.3 Filter size process
func (t *TileGroup) filterSize(txSz int, prevTxSz int, pass bool, plane int) int {
	var baseSize int
	if !pass {
		baseSize = util.Min(Tx_Width[prevTxSz], Tx_Width[txSz])
	} else {
		baseSize = util.Min(Tx_Height[prevTxSz], Tx_Height[prevTxSz])

	}

	if plane == 0 {
		return util.Min(16, baseSize)
	} else {
		return util.Min(8, baseSize)
	}
}

// 7.14.4 Adaptive filter strength process
func (t *TileGroup) adaptiveFilterStrength(row int, col int, plane int, pass bool, state *state.State, uh uncompressedheader.UncompressedHeader) (int, int, int, int) {
	segment := t.SegmentIds[row][col]
	ref := state.RefFrames[row][col][0]
	mode := t.YModes[row][col]

	var modeType int
	if mode >= shared.NEARESTMV && mode != shared.GLOBALMV {
		modeType = 1
	} else {
		modeType = 0
	}

	var deltaLF int
	if uh.DeltaLfMulti == 0 {
		deltaLF = state.DeltaLFs[row][col][0]
	} else {
		if plane == 0 {
			deltaLF = state.DeltaLFs[row][col][util.Int(pass)]
		} else {
			deltaLF = state.DeltaLFs[row][col][plane+1]
		}
	}

	lvl := t.adaptiveFilterStrengthSelection(segment, ref, modeType, deltaLF, plane, pass, uh, state)

	var shift int
	if uh.LoopFilterSharpness > 4 {
		shift = 2
	} else if uh.LoopFilterSharpness > 0 {
		shift = 2
	} else {
		shift = 0
	}

	var limit int
	if uh.LoopFilterSharpness > 0 {
		limit = util.Clip3(1, 9-uh.LoopFilterSharpness, lvl>>shift)
	} else {
		limit = util.Max(1, lvl>>shift)
	}

	blimit := 2*(lvl+2) + limit
	thresh := lvl >> 4

	return lvl, limit, blimit, thresh
}

// 7.14.5 Adaptive filter strength selection process
func (t *TileGroup) adaptiveFilterStrengthSelection(segment int, ref int, modeType int, deltaLF int, plane int, pass bool, uh uncompressedheader.UncompressedHeader, state *state.State) int {
	var i int
	if plane == 0 {
		i = util.Int(pass)
	} else {
		i = plane + 1
	}

	baseFilterLevel := util.Clip3(0, shared.MAX_LOOP_FILTER, deltaLF+uh.LoopFilterLevel[i])

	lvlSeg := baseFilterLevel
	feature := shared.SEG_LVL_ALT_LF_Y_V + i
	if t.segFeatureActiveIdx(segment, feature, uh, state) {
		lvlSeg = state.FeatureData[segment][feature] + lvlSeg
		lvlSeg = util.Clip3(0, shared.MAX_LOOP_FILTER, lvlSeg)
	}

	if uh.LoopFilterDeltaEnabled {
		nShift := lvlSeg >> 5
		if ref == shared.INTRA_FRAME {
			lvlSeg = lvlSeg + (uh.LoopFilterRefDeltas[shared.INTRA_FRAME] << nShift)
		} else {
			lvlSeg = lvlSeg + (uh.LoopFilterRefDeltas[ref] << nShift) + (uh.LoopFilterModeDeltas[modeType] << nShift)
		}
		lvlSeg = util.Clip3(0, shared.MAX_LOOP_FILTER, lvlSeg)
	}

	return lvlSeg
}

// 7.14.6 Sample filtering
func (t *TileGroup) sampleFiltering(x int, y int, plane int, limit int, blimit int, thresh int, dx int, dy int, filterSize int, state *state.State, sh sequenceheader.SequenceHeader) {
	hevMask, filterMask, flatMask, flatMask2 := t.filterMask(x, y, plane, limit, blimit, thresh, dx, dy, filterSize, state, sh)

	if filterMask == 0 {
		return
	} else if filterSize == 4 || flatMask == 0 {
		t.narrowFilter(hevMask, x, y, plane, dx, dy, state, sh)
	} else if filterSize == 8 || flatMask2 == 0 {
		t.wideFilter(x, y, plane, dx, dy, 3, state)
	} else {
		t.wideFilter(x, y, plane, dx, dy, 4, state)
	}
}

// 7.14.6.2 Filter mask process
func (t *TileGroup) filterMask(x int, y int, plane int, limit int, blimit int, thresh int, dx int, dy int, filterSize int, state *state.State, sh sequenceheader.SequenceHeader) (int, int, int, int) {
	q0 := state.CurrFrame[plane][y][x]
	q1 := state.CurrFrame[plane][y+dy][x+dx]
	q2 := state.CurrFrame[plane][y+dy*2][x+dx*2]
	q3 := state.CurrFrame[plane][y+dy*3][x+dx*3]
	q4 := state.CurrFrame[plane][y+dy*4][x+dx*4]
	q5 := state.CurrFrame[plane][y+dy*5][x+dx*5]
	q6 := state.CurrFrame[plane][y+dy*6][x+dx*6]
	p0 := state.CurrFrame[plane][y-dy][x-dx]
	p1 := state.CurrFrame[plane][y-dy*2][x-dx*2]
	p2 := state.CurrFrame[plane][y-dy*3][x-dx*3]
	p3 := state.CurrFrame[plane][y-dy*4][x-dx*4]
	p4 := state.CurrFrame[plane][y-dy*5][x-dx*5]
	p5 := state.CurrFrame[plane][y-dy*6][x-dx*6]
	p6 := state.CurrFrame[plane][y-dy*7][x-dx*7]

	hevMask := 0
	threshBd := thresh << (sh.ColorConfig.BitDepth - 8)
	hevMask |= util.Int((util.Abs(p1-p0) > threshBd))
	hevMask |= util.Int((util.Abs(q1-q0) > threshBd))

	var filterLen int
	if filterSize == 4 {
		filterLen = 4
	} else if plane != 0 {
		filterLen = 6
	} else if filterSize == 8 {
		filterLen = 8
	} else {
		filterLen = 16
	}

	limitBd := limit << (sh.ColorConfig.BitDepth - 8)
	blimitBd := blimit << (sh.ColorConfig.BitDepth - 8)
	mask := 0
	mask |= util.Int((util.Abs(p1-p0) > limitBd))
	mask |= util.Int((util.Abs(q1-q0) > limitBd))
	mask |= util.Int((util.Abs(p0-q0)*2+util.Abs(p1-q1)/2 > blimitBd))
	if filterLen >= 6 {
		mask |= util.Int((util.Abs(p2-p1) > limitBd))
		mask |= util.Int((util.Abs(q2-q1) > limitBd))
	}
	if filterLen >= 8 {
		mask |= util.Int((util.Abs(p3-p2) > limitBd))
		mask |= util.Int((util.Abs(q3-q2) > limitBd))
	}

	var flatMask int
	filterMask := util.Int((mask == 0))
	thresholdBd := 1 << (sh.ColorConfig.BitDepth - 8)
	if filterSize >= 8 {
		mask = 0
		mask |= util.Int((util.Abs(p1-p0) > thresholdBd))
		mask |= util.Int((util.Abs(q1-q0) > thresholdBd))
		mask |= util.Int((util.Abs(p2-p0) > thresholdBd))
		mask |= util.Int((util.Abs(q2-q0) > thresholdBd))
		if filterLen >= 8 {
			mask |= util.Int((util.Abs(p3-p0) > thresholdBd))
			mask |= util.Int((util.Abs(q3-q0) > thresholdBd))
		}
		flatMask = util.Int((mask == 0))
	}

	var flatMask2 int
	thresholdBd = 1 << (sh.ColorConfig.BitDepth - 8)
	if filterSize >= 16 {
		mask = 0
		mask |= util.Int((util.Abs(p6-p0) > thresholdBd))
		mask |= util.Int((util.Abs(q6-q0) > thresholdBd))
		mask |= util.Int((util.Abs(p5-p0) > thresholdBd))
		mask |= util.Int((util.Abs(q5-q0) > thresholdBd))
		mask |= util.Int((util.Abs(p4-p0) > thresholdBd))
		mask |= util.Int((util.Abs(q4-q0) > thresholdBd))
		flatMask2 = util.Int((mask == 0))
	}

	return hevMask, filterMask, flatMask, flatMask2
}

// 7.14.6.3 Narrow filter process
func (t *TileGroup) narrowFilter(hevMask int, x int, y int, plane int, dx int, dy int, state *state.State, sh sequenceheader.SequenceHeader) {
	q0 := state.CurrFrame[plane][y][x]
	q1 := state.CurrFrame[plane][y+dy][x+dx]
	p0 := state.CurrFrame[plane][y-dy][x-dx]
	p1 := state.CurrFrame[plane][y-dy*2][x-dx*2]
	ps1 := p1 - (0x80 << (sh.ColorConfig.BitDepth - 8))
	ps0 := p0 - (0x80 << (sh.ColorConfig.BitDepth - 8))
	qs0 := q0 - (0x80 << (sh.ColorConfig.BitDepth - 8))
	qs1 := q1 - (0x80 << (sh.ColorConfig.BitDepth - 8))
	var filter int
	if util.Bool(hevMask) {
		filter = filter4Clamp(ps1-qs1, sh)
	} else {
		filter = 0
	}
	filter = filter4Clamp(filter+3*(qs0-ps0), sh)
	filter1 := filter4Clamp(filter+4, sh) >> 3
	filter2 := filter4Clamp(filter+3, sh) >> 3
	oq0 := filter4Clamp(qs0-filter1, sh) + (0x80 << (sh.ColorConfig.BitDepth - 8))
	op0 := filter4Clamp(ps0+filter2, sh) + (0x80 << (sh.ColorConfig.BitDepth - 8))
	state.CurrFrame[plane][y][x] = oq0
	state.CurrFrame[plane][y-dy][x-dx] = op0
	if !util.Bool(hevMask) {
		filter = util.Round2(filter1, 1)
		oq1 := filter4Clamp(qs1-filter, sh) + (0x80 << (sh.ColorConfig.BitDepth - 8))
		op1 := filter4Clamp(ps1+filter, sh) + (0x80 << (sh.ColorConfig.BitDepth - 8))
		state.CurrFrame[plane][y+dy][x+dx] = oq1
		state.CurrFrame[plane][y-dy*2][x-dx*2] = op1
	}
}

func filter4Clamp(value int, sh sequenceheader.SequenceHeader) int {
	return util.Clip3(-(1 << (sh.ColorConfig.BitDepth - 1)), (1<<(sh.ColorConfig.BitDepth-1))-1, value)
}

// 7.14.6.4. Wide filter process
func (t *TileGroup) wideFilter(x int, y int, plane int, dx int, dy int, log2Size int, state *state.State) {
	var n int
	if log2Size == 4 {
		n = 6
	} else if plane == 0 {
		n = 3
	} else {
		n = 2
	}

	var n2 int
	if log2Size == 3 && plane == 0 {
		n2 = 0
	} else {
		n2 = 1
	}

	var F []int
	for i := -n; i < n; i++ {
		t := 0
		for j := -n; j <= n; j++ {
			p := util.Clip3(-(n + 1), n, i+j)
			var tap int
			if util.Abs(j) <= n2 {
				tap = 2
			} else {
				tap = 1
			}
			t += state.CurrFrame[plane][y+p*dy][x+p*dx] * tap
		}
		F[i] = util.Round2(t, log2Size)
	}
	for i := -n; i < n; i++ {
		state.CurrFrame[plane][y+i*dy][x+i*dx] = F[i]
	}
}
