package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.10.2. Find MV stack process
// find_mv_stack( isCompound )
func (t *TileGroup) findMvStack(isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
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

	// TODO: I feel like something is missing here
}

func (t *TileGroup) scanRowProcess(deltaRow int, isCompound int, state *state.State, uh uncompressedheader.UncompressedHeader) {
	bw4 := state.Num4x4BlocksWide[state.MiSize]
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

		if !t.isInside(mvRow, mvCol) {
			break
		}

		len := util.Min(bw4, state.Num4x4BlocksWide[state.MiSizes[mvRow][mvCol]])
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
					// TODO: does this work?/
					candMv[i]--
				} else {
					candMv[i]++
				}
			}
		}
	}

	return candMv
}
