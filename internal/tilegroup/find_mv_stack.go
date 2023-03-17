package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

const REF_CAT_LEVEL = 640
const MV_BORDER = 128

// 7.10.2. Find MV stack process
// find_mv_stack( isCompound )
func (t *TileGroup) findMvStack(isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]

	// 1.
	t.NumMvFound = 0

	// 2.
	t.NewMvCount = 0

	// 3.
	t.GlobalMvs[0] = t.setupGlobalMvProcess(0, state, uh)

	// 4.
	if util.Bool(isCompound) {
		t.GlobalMvs[1] = t.setupGlobalMvProcess(1, state, uh)
	}

	// 5.
	t.FoundMatch = 0

	// 6.
	t.scanRowProcess(-1, isCompound, state, uh)

	// 7.
	foundAboveMatch := t.FoundMatch
	t.FoundMatch = 0

	// 8.
	t.scanColProcess(-1, isCompound, state, uh)

	// 9.
	foundLeftMatch := t.FoundMatch
	t.FoundMatch = 0

	// 10.
	if util.Max(bw4, bh4) <= 16 {
		t.scanPointProcess(-1, bw4, isCompound, state, uh)
	}

	// 11.
	if t.FoundMatch == 1 {
		foundAboveMatch = 1
	}

	// 12.
	t.CloseMatches = foundAboveMatch + foundLeftMatch

	// 13.
	numNearest := t.NumMvFound

	// 14.
	numNew := t.NewMvCount

	// 15.
	if numNearest > 0 {
		for idx := 0; idx < numNearest-1; idx++ {
			t.WeightStack[idx] += REF_CAT_LEVEL
		}
	}

	// 16.
	t.ZeroMvContext = 0

	// 17.
	if uh.UseRefFrameMvs {
		t.temporalScanProcess(isCompound, state, uh)
	}

	// 18.
	t.scanPointProcess(-1, -1, isCompound, state, uh)

	// 19.
	if t.FoundMatch == 1 {
		foundAboveMatch = 1
	}

	// 20.
	t.FoundMatch = 0

	// 21.
	t.scanRowProcess(-3, isCompound, state, uh)

	// 22.
	if t.FoundMatch == 1 {
		foundAboveMatch = 1
	}

	// 23.
	t.FoundMatch = 0

	// 24.
	t.scanColProcess(-3, isCompound, state, uh)

	// 25.
	if t.FoundMatch == 1 {
		foundLeftMatch = 1
	}

	// 26.
	t.FoundMatch = 0

	// 27.
	if bh4 > 1 {
		t.scanRowProcess(-5, isCompound, state, uh)
	}

	// 28.
	if t.FoundMatch == 1 {
		foundAboveMatch = 1
	}

	// 29.
	t.FoundMatch = 0

	// 30.
	if bw4 > 1 {
		t.scanColProcess(-5, isCompound, state, uh)
	}

	// 31.
	if t.FoundMatch == 1 {
		foundAboveMatch = 1
	}

	// 32.
	t.TotalMatches = foundAboveMatch + foundLeftMatch

	// 33.
	t.sortingProcess(0, numNearest, isCompound)

	// 34.
	t.sortingProcess(numNearest, t.NumMvFound, isCompound)

	// 35.
	if t.NumMvFound < 2 {
		t.extraSearchProcess(isCompound, state)
	}

	// 36.
	t.contextAndClampingProcess(isCompound, numNew, state)
}

// 7.10.2.1 Setup global MV process
func (t *TileGroup) setupGlobalMvProcess(refList int, state *state.State, uh uncompressedheader.UncompressedHeader) []int {
	ref := state.RefFrame[refList]

	var typ int
	if ref != shared.INTRA_FRAME {
		typ = state.GmType[ref]
	}

	bw := t.Block_Width[state.MiSize]
	bh := t.Block_Height[state.MiSize]

	var xc int
	var yc int
	mv := []int{}
	if ref == shared.INTRA_FRAME || typ == shared.IDENTITY {
		mv[0] = 0
		mv[1] = 0
	} else if typ == shared.TRANSLATION {
		mv[0] = uh.GmParams[ref][0] >> (shared.WARPEDMODEL_PREC_BITS - 3)
		mv[1] = uh.GmParams[ref][1] >> (shared.WARPEDMODEL_PREC_BITS - 3)
	} else {
		x := state.MiCol*MI_SIZE + bw/2 - 1
		y := state.MiRow*MI_SIZE + bh/2 - 1

		xc = (uh.GmParams[ref][2]-(1<<shared.WARPEDMODEL_PREC_BITS))*x + uh.GmParams[ref][3]*y + uh.GmParams[ref][0]
		yc = uh.GmParams[ref][4]*x + (uh.GmParams[ref][5]-(1<<shared.WARPEDMODEL_PREC_BITS))*y + uh.GmParams[ref][1]

		if uh.AllowHighPrecisionMv {
			mv[0] = util.Round2Signed(yc, shared.WARPEDMODEL_PREC_BITS-3)
			mv[1] = util.Round2Signed(xc, shared.WARPEDMODEL_PREC_BITS-3)
		} else {
			mv[0] = util.Round2Signed(yc, shared.WARPEDMODEL_PREC_BITS-2) * 2
			mv[1] = util.Round2Signed(xc, shared.WARPEDMODEL_PREC_BITS-2) * 2
		}
	}
	mv = t.lowerPrecisionProcess(mv, uh)

	return mv
}

// 7.10.2.2 Scan row process
func (t *TileGroup) scanRowProcess(deltaRow int, isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	end4 := util.Min(util.Min(bw4, state.MiCols-state.MiCol), 16)
	deltaCol := 0
	useStep16 := bw4 >= 16

	if util.Abs(deltaRow) > 1 {
		deltaRow += state.MiRow & 1
		deltaCol = 1 - (state.MiCol & 1)
	}

	i := 0

	for i < end4 {
		mvRow := state.MiRow + deltaRow
		mvCol := state.MiCol + deltaCol + i

		if !t.isInside(mvRow, mvCol, state) {
			break
		}

		len := util.Min(bw4, shared.NUM_4X4_BLOCKS_WIDE[state.MiSizes[mvRow][mvCol]])
		if util.Abs(deltaRow) > 1 {
			len = util.Max(2, len)
		}
		if useStep16 {
			len = util.Max(4, len)
		}
		weight := len * 2
		t.addRefMvCandidate(mvRow, mvCol, isCompound, weight, state, uh)
		i += len
	}
}

// 7.10.2.3 Scan col process
func (t *TileGroup) scanColProcess(deltaCol int, isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]
	end4 := util.Min(util.Min(bh4, state.MiRows-state.MiRow), 16)
	deltaRow := 0
	useStep16 := bh4 >= 16

	if util.Abs(deltaCol) > 1 {
		deltaRow = 1 - (state.MiRow & 1)
		deltaCol += state.MiCol & 1
	}

	i := 0

	for i < end4 {
		mvRow := state.MiRow + deltaRow + i
		mvCol := state.MiCol + deltaCol

		if !t.isInside(mvRow, mvCol, state) {
			break
		}

		len := util.Min(bh4, shared.NUM_4X4_BLOCKS_HIGH[state.MiSizes[mvRow][mvCol]])
		if util.Abs(deltaCol) > 1 {
			len = util.Max(2, len)
		}
		if useStep16 {
			len = util.Max(4, len)
		}
		weight := len * 2
		t.addRefMvCandidate(mvRow, mvCol, isCompound, weight, state, uh)
		i += len
	}
}

// 7.10.2.4 Scan point proces
func (t *TileGroup) scanPointProcess(deltaRow int, deltaCol int, isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	mvRow := state.MiRow + deltaRow
	mvCol := state.MiCol + deltaCol
	weight := 4

	// "has been written to" - what does this mean?
	if t.isInside(mvRow, mvCol, state) && state.RefFrames[mvRow][mvCol][0] != 0 {
		t.addRefMvCandidate(mvRow, mvCol, isCompound, weight, state, uh)
	}
}

// 7.10.2.5 Temporal scan process
func (t *TileGroup) temporalScanProcess(isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]

	stepW4 := 2
	if bw4 >= 16 {
		stepW4 = 4
	}
	stepH4 := 2
	if bh4 >= 16 {
		stepH4 = 4
	}

	for deltaRow := 0; deltaRow < util.Min(bh4, 16); deltaRow += stepH4 {
		for deltaCol := 0; deltaCol < util.Min(bw4, 16); deltaCol += stepW4 {
			t.temporalSampleProcess(deltaRow, deltaCol, isCompound, state, uh)
		}
	}

}

// 7.10.2.6 Temporal sample process
func (t *TileGroup) temporalSampleProcess(deltaRow int, deltaCol int, isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	mvRow := (state.MiRow + deltaRow) | 1
	mvCol := (state.MiCol + deltaCol) | 1

	if !t.isInside(mvRow, mvCol, state) {
		return
	}

	x8 := mvCol >> 1
	y8 := mvRow >> 1

	if deltaRow == 0 && deltaCol == 0 {
		t.ZeroMvContext = 1
	}

	if !util.Bool(isCompound) {
		candMv := t.MotionFieldMvs[state.RefFrame[0]][y8][x8]
		if candMv[0] == -1<<15 {
			return
		}

		t.lowerPrecisionProcess(candMv, uh)

		if deltaRow == 0 && deltaCol == 0 {
			if util.Abs(candMv[0]-t.GlobalMvs[0][0]) >= 16 ||
				util.Abs(candMv[1]-t.GlobalMvs[0][1]) >= 16 {
				t.ZeroMvContext = 1
			} else {
				t.ZeroMvContext = 0
			}
		}

		var idx int
		for idx := 0; idx < t.NumMvFound; idx++ {
			if candMv[0] == t.RefStackMv[idx][0][0] &&
				candMv[1] == t.RefStackMv[idx][0][1] {
				break
			}
		}
		if idx < t.NumMvFound {
			t.WeightStack[idx] += 2
		} else if t.NumMvFound < MAX_REF_MV_STACK_SIZE {
			t.RefStackMv[t.NumMvFound][0] = candMv
			t.WeightStack[t.NumMvFound] = 2
			t.NumMvFound += 1
		}
	} else {
		candMv0 := t.MotionFieldMvs[state.RefFrame[0]][y8][x8]
		if candMv0[0] == -1<<15 {
			return
		}
		candMv1 := t.MotionFieldMvs[state.RefFrame[1]][y8][x8]
		if candMv1[0] == -1<<15 {
			return
		}
		t.lowerPrecisionProcess(candMv0, uh)
		t.lowerPrecisionProcess(candMv1, uh)

		if deltaRow == 0 && deltaCol == 0 {
			if util.Abs(candMv0[0]-t.GlobalMvs[0][0]) >= 16 ||
				util.Abs(candMv0[1]-t.GlobalMvs[0][1]) >= 16 ||
				util.Abs(candMv1[0]-t.GlobalMvs[1][0]) >= 16 ||
				util.Abs(candMv1[1]-t.GlobalMvs[1][1]) >= 16 {
				t.ZeroMvContext = 1
			} else {
				t.ZeroMvContext = 0
			}

		}
		var idx int
		for idx := 0; idx < t.NumMvFound; idx++ {
			if candMv0[0] == t.RefStackMv[idx][0][0] &&
				candMv0[1] == t.RefStackMv[idx][0][1] &&
				candMv1[0] == t.RefStackMv[idx][1][0] &&
				candMv1[1] == t.RefStackMv[idx][1][1] {
				break
			}
		}

		if idx < t.NumMvFound {
			t.WeightStack[idx] += 2
		} else if t.NumMvFound < MAX_REF_MV_STACK_SIZE {
			t.RefStackMv[t.NumMvFound][0] = candMv0
			t.RefStackMv[t.NumMvFound][1] = candMv1
			t.WeightStack[t.NumMvFound] = 2
			t.NumMvFound += 1
		}

	}

}

// 7.10.2.7. Add reference motion vector process
func (t *TileGroup) addRefMvCandidate(mvRow int, mvCol int, isCompound int, weight int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	if t.IsInters[mvRow][mvCol] == 0 {
		return
	}

	// TODO: not sure if this loop is correct here
	for candList := 0; candList < 2; candList++ {
		if isCompound == 0 {
			if state.RefFrames[mvRow][mvCol][candList] == state.RefFrame[0] {
				t.searchStackProcess(mvRow, mvCol, candList, weight, state, uh)
			}

		} else {
			if state.RefFrames[mvRow][mvCol][0] == state.RefFrame[0] && state.RefFrames[mvRow][mvCol][1] == state.RefFrame[1] {
				t.compoundSearchStackProcess(mvRow, mvCol, weight)
			}
		}
	}
}

// 7.10.2.8. Search stack process
func (t *TileGroup) searchStackProcess(mvRow int, mvCol int, candList int, weight int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	candMode := t.YModes[mvRow][mvCol]
	candSize := state.MiSizes[mvRow][mvCol]
	large := util.Min(t.Block_Width[candSize], t.Block_Height[candSize]) >= 8

	var candMv []int
	if (candMode == shared.GLOBALMV || candMode == shared.GLOBAL_GLOBALMV) && (state.GmType[state.RefFrame[0]] > shared.TRANSLATION) && large {
		candMv = t.GlobalMvs[0]
	} else {
		candMv = t.Mvs[mvRow][mvCol][candList]
	}

	candMv = t.lowerPrecisionProcess(candMv, uh)
	if util.HasNewmv(candMode) {
		t.NewMvCount += 1
	}

	t.FoundMatch = 1

	for idx := 0; idx < t.NumMvFound; idx++ {
		if util.Equals(candMv, t.RefStackMv[idx][0]) {
			t.WeightStack[idx] += weight
			return
		}
	}

	if t.NumMvFound < MAX_REF_MV_STACK_SIZE {
		t.RefStackMv[t.NumMvFound][0] = candMv
		t.WeightStack[t.NumMvFound] = weight
		t.NumMvFound += 1
	}
}

// 7.10.2.9. Compound search stack process
func (t *TileGroup) compoundSearchStackProcess(mvRow int, mvCol int, weight int) {
	// TODO: implement
}

// 7.10.2.10. Lower precision process
func (t *TileGroup) lowerPrecisionProcess(candMv []int, uh uncompressedheader.UncompressedHeader) []int {
	if uh.AllowHighPrecisionMv {
		return candMv
	}

	for i := 0; i < 2; i++ {
		if uh.ForceIntegerMv {
			a := util.Abs(candMv[i])
			aInt := (a + 3) >> 3

			if candMv[i] > 0 {
				candMv[i] = aInt << 3
			} else {
				candMv[i] = -(aInt << 3)
			}
		} else {
			if util.Bool(candMv[i] & 1) {
				if candMv[i] > 0 {
					candMv[i]--
				} else {
					candMv[i]++
				}
			}
		}
	}

	return candMv
}

// 7.10.2.11 Sorting proces
func (t *TileGroup) sortingProcess(start int, end int, isCompound int) {
	for end > start {
		newEnd := start
		for idx := start + 1; idx < end; idx++ {
			if t.WeightStack[idx-1] < t.WeightStack[idx] {
				t.swapStack(idx-1, idx, isCompound)
				newEnd = idx
			}
		}
		end = newEnd
	}
}

func (t *TileGroup) swapStack(i int, j int, isCompound int) {
	temp := t.WeightStack[i]
	t.WeightStack[i] = t.WeightStack[j]
	t.WeightStack[j] = temp

	for list := 0; list < 1+isCompound; list++ {
		for comp := 0; comp < 2; comp++ {
			t.RefStackMv[j][list][comp] = t.RefStackMv[j][list][comp]
			t.RefStackMv[j][list][comp] = temp
		}

	}
}

// 7.10.2.12 Extra search process
func (t *TileGroup) extraSearchProcess(isCompound int, state *state.State) {
	for list := 0; list < 2; list++ {
		t.RefIdCount[list] = 0
		t.RefDiffCount[list] = 0
	}

	w4 := util.Min(16, shared.NUM_4X4_BLOCKS_WIDE[state.MiSize])
	h4 := util.Min(16, shared.NUM_4X4_BLOCKS_HIGH[state.MiSize])
	w4 = util.Min(w4, state.MiCols-state.MiCol)
	h4 = util.Min(h4, state.MiRows-state.MiRow)
	num4x4 := util.Min(w4, h4)

	for pass := 0; pass < 2; pass++ {
		idx := 0
		for idx < num4x4 && t.NumMvFound < 2 {
			var mvRow int
			var mvCol int
			if pass == 0 {
				mvRow = state.MiRow - 1
				mvCol = state.MiCol + idx
			} else {
				mvRow = state.MiRow + idx
				mvCol = state.MiCol - 1
			}

			if t.isInside(mvRow, mvCol, state) {
				break
			}

			t.addExtraMvCandidateProcess(mvRow, mvCol, isCompound, state)
			if pass == 0 {
				idx += shared.NUM_4X4_BLOCKS_WIDE[state.MiSizes[mvRow][mvCol]]
			} else {
				idx += shared.NUM_4X4_BLOCKS_HIGH[state.MiSizes[mvRow][mvCol]]
			}
		}
	}

	var combinedMvs [][][]int
	if isCompound == 1 {
		for list := 0; list < 2; list++ {
			compCount := 0
			for idx := 0; idx < t.RefIdCount[list]; idx++ {
				combinedMvs[compCount][list] = t.RefIdMvs[list][idx]
				compCount++
			}
			for idx := 0; idx < t.RefDiffCount[list] && compCount < 2; idx++ {
				combinedMvs[compCount][list] = t.RefDiffMvs[list][idx]
				compCount++
			}
			for compCount < 2 {
				combinedMvs[compCount][list] = t.GlobalMvs[list]
				compCount++
			}
		}
		if t.NumMvFound == 1 {
			if util.Equals(combinedMvs[0][0], t.RefStackMv[0][0]) &&
				util.Equals(combinedMvs[0][1], t.RefStackMv[0][1]) {
				t.RefStackMv[t.NumMvFound][0] = combinedMvs[1][0]
				t.RefStackMv[t.NumMvFound][1] = combinedMvs[1][1]
			} else {
				t.RefStackMv[t.NumMvFound][0] = combinedMvs[0][0]
				t.RefStackMv[t.NumMvFound][1] = combinedMvs[0][1]
			}
			t.WeightStack[t.NumMvFound] = 2
			t.NumMvFound++
		} else {
			for idx := 0; idx < 2; idx++ {
				t.RefStackMv[t.NumMvFound][0] = combinedMvs[idx][0]
				t.RefStackMv[t.NumMvFound][1] = combinedMvs[idx][1]
				t.WeightStack[t.NumMvFound] = 2
				t.NumMvFound++
			}
		}
	}

	if isCompound == 0 {
		for idx := t.NumMvFound; idx < 2; idx++ {
			t.RefStackMv[idx][0] = t.GlobalMvs[0]
		}
	}
}

// 7.10.2.13 Add extra mv candidate process
func (t *TileGroup) addExtraMvCandidateProcess(mvRow int, mvCol int, isCompound int, state *state.State) {
	var candRef int
	var candMv []int
	if util.Bool(isCompound) {
		for candList := 0; candList < 2; candList++ {
			candRef = state.RefFrames[mvRow][mvCol][candList]
			if candRef > shared.INTRA_FRAME {
				for list := 0; list < 2; list++ {
					candMv = t.Mvs[mvRow][mvCol][candList]
					if candRef == state.RefFrame[list] && t.RefIdCount[list] < 2 {
						t.RefIdMvs[list][t.RefIdCount[list]] = candMv
						t.RefIdCount[list]++
					} else if t.RefDiffCount[list] < 2 {
						if t.RefFrameSignBias[candRef] != t.RefFrameSignBias[state.RefFrame[list]] {
							candMv[0] *= -1
							candMv[1] *= -1
						}
						t.RefDiffMvs[list][t.RefDiffCount[list]] = candMv
						t.RefDiffCount[list]++
					}
				}
			}
		}
	} else {
		for candList := 0; candList < 2; candList++ {
			candRef = state.RefFrames[mvRow][mvCol][candList]
			if candRef > shared.INTRA_FRAME {
				candMv = t.Mvs[mvRow][mvCol][candList]
				if t.RefFrameSignBias[candRef] != t.RefFrameSignBias[state.RefFrame[0]] {
					candMv[0] *= -1
					candMv[1] *= -1
				}
				var idx int
				for idx := 0; idx < t.NumMvFound; idx++ {
					if util.Equals(candMv, t.RefStackMv[idx][0]) {
						break
					}
				}
				if idx == t.NumMvFound {
					t.RefStackMv[idx][0] = candMv
					t.WeightStack[idx] = 2
					t.NumMvFound++
				}
			}
		}
	}
}

// 7.10.2.14 Context and claping process
func (t *TileGroup) contextAndClampingProcess(isCompound int, numNew int, state *state.State) {
	bw := t.Block_Width[state.MiSize]
	bh := t.Block_Height[state.MiSize]

	numLists := 1
	if util.Bool(isCompound) {
		numLists = 2
	}

	for idx := 0; idx < t.NumMvFound; idx++ {
		z := 0
		if idx+1 < t.NumMvFound {
			w0 := t.WeightStack[idx]
			w1 := t.WeightStack[idx+1]
			if w0 >= REF_CAT_LEVEL {
				if w1 < REF_CAT_LEVEL {
					z = 1
				}
			} else {
				z = 2
			}
		}
		t.DrlCtxStack[idx] = z
	}

	for list := 0; list < numLists; list++ {
		for idx := 0; idx < t.NumMvFound; idx++ {
			refMv := t.RefStackMv[idx][list]
			refMv[0] = t.clampMvRow(refMv[0], MV_BORDER+bh*8, state)
			refMv[1] = t.clampMvCol(refMv[1], MV_BORDER+bw*8, state)
			t.RefStackMv[idx][list] = refMv
		}
	}
}

// clamp_mv_row( mvec, border)
func (t *TileGroup) clampMvRow(mvec int, border int, state *state.State) int {
	bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]
	mbToTopEdge := -((state.MiRow * MI_SIZE) * 8)
	mbToBottomEdge := ((state.MiRows - bh4 - state.MiRow) * MI_SIZE) * 8
	return util.Clip3(mbToTopEdge-border, mbToBottomEdge+border, mvec)
}

// clamp_mv_col( mvec, border)
func (t *TileGroup) clampMvCol(mvec int, border int, state *state.State) int {
	bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	mbToLeftEdge := -((state.MiCol * MI_SIZE) * 8)
	mbToRightEdge := ((state.MiCols - bw4 - state.MiCol) * MI_SIZE) * 8
	return util.Clip3(mbToLeftEdge-border, mbToRightEdge+border, mvec)
}
