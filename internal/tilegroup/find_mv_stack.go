package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.10.2. Find MV stack process
// find_mv_stack( isCompound )
func (t *TileGroup) findMvStack(isCompound int) {
	// 1.
	t.NumMvFound = 0

	// 2.
	t.NewMvCount = 0

	// 3.
	t.GlobalMvs[0] = t.setupGlobalMvProcess(0)

	// 4.
	if util.Bool(isCompound) {
		t.GlobalMvs[1] = t.setupGlobalMvProcess(1)
	}

	// 5.
	t.FoundMatch = 0

	// 6.
	t.scanRowProcess(-1, isCompound)

	// TODO: I feel like something is missing here
}

func (t *TileGroup) scanRowProcess(deltaRow int, isCompound int) {
	bw4 := t.State.Num4x4BlocksWide[t.State.MiSize]
	end4 := util.Min(util.Min(bw4, t.State.MiCols-t.State.MiCol), 16)
	deltaCol := 0
	useStep16 := bw4 >= 16

	if util.Abs(deltaRow) > 1 {
		deltaRow += t.State.MiRow & 1
		deltaCol = 1 - (t.State.MiCol & 1)
	}

	i := 0

	for i < end4 {
		mvRow := t.State.MiRow + deltaRow
		mvCol := t.State.MiCol + deltaCol + i

		if !t.isInside(mvRow, mvCol) {
			break
		}

		len := util.Min(bw4, t.State.Num4x4BlocksWide[t.State.MiSizes[mvRow][mvCol]])
		if util.Abs(deltaRow) > 1 {
			len = util.Max(2, len)
		}
		if useStep16 {
			len = util.Max(4, len)
		}
		weight := len * 2
		t.addRefMvCandidate(mvRow, mvCol, isCompound, weight)
		i += len
	}
}

// 7.10.2.7. Add reference motion vector process
func (t *TileGroup) addRefMvCandidate(mvRow int, mvCol int, isCompound int, weight int) {
	if t.IsInters[mvRow][mvCol] == 0 {
		return
	}

	// TODO: not sure if this loop is correct here
	for candList := 0; candList < 2; candList++ {
		if isCompound == 0 {
			if t.State.RefFrames[mvRow][mvCol][candList] == t.State.RefFrame[0] {
				t.searchStackProcess(mvRow, mvCol, candList, weight)
			}

		} else {
			if t.State.RefFrames[mvRow][mvCol][0] == t.State.RefFrame[0] && t.State.RefFrames[mvRow][mvCol][1] == t.State.RefFrame[1] {
				t.compoundSearchStackProcess(mvRow, mvCol, weight)
			}
		}
	}
}

// 7.10.2.8. Search stack process
func (t *TileGroup) searchStackProcess(mvRow int, mvCol int, candList int, weight int) {
	candMode := t.YModes[mvRow][mvCol]
	candSize := t.State.MiSizes[mvRow][mvCol]
	large := util.Min(t.Block_Width[candSize], t.Block_Height[candSize]) >= 8

	var candMv []int
	if (candMode == shared.GLOBALMV || candMode == shared.GLOBAL_GLOBALMV) && (t.State.GmType[t.State.RefFrame[0]] > shared.TRANSLATION) && large {
		candMv = t.GlobalMvs[0]
	} else {
		candMv = t.Mvs[mvRow][mvCol][candList]
	}

	candMv = t.lowerPrecisionProcess(candMv)
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
func (t *TileGroup) setupGlobalMvProcess(refList int) []int {
	ref := t.State.RefFrame[refList]

	var typ int
	if ref != INTRA_FRAME {
		typ = t.State.GmType[ref]
	}

	bw := t.Block_Width[t.State.MiSize]
	bh := t.Block_Height[t.State.MiSize]

	var xc int
	var yc int
	mv := []int{}
	if ref == INTRA_FRAME || typ == shared.IDENTITY {
		mv[0] = 0
		mv[1] = 0
	} else if typ == shared.TRANSLATION {
		mv[0] = t.State.UncompressedHeader.GmParams[ref][0] >> (shared.WARPEDMODEL_PREC_BITS - 3)
		mv[1] = t.State.UncompressedHeader.GmParams[ref][1] >> (shared.WARPEDMODEL_PREC_BITS - 3)
	} else {
		x := t.State.MiCol*MI_SIZE + bw/2 - 1
		y := t.State.MiRow*MI_SIZE + bh/2 - 1

		xc = (t.State.UncompressedHeader.GmParams[ref][2]-(1<<shared.WARPEDMODEL_PREC_BITS))*x + t.State.UncompressedHeader.GmParams[ref][3]*y + t.State.UncompressedHeader.GmParams[ref][0]
		yc = t.State.UncompressedHeader.GmParams[ref][4]*x + (t.State.UncompressedHeader.GmParams[ref][5]-(1<<shared.WARPEDMODEL_PREC_BITS))*y + t.State.UncompressedHeader.GmParams[ref][1]

		if t.State.UncompressedHeader.AllowHighPrecisionMv {
			mv[0] = util.Round2Signed(yc, shared.WARPEDMODEL_PREC_BITS-3)
			mv[1] = util.Round2Signed(xc, shared.WARPEDMODEL_PREC_BITS-3)
		} else {
			mv[0] = util.Round2Signed(yc, shared.WARPEDMODEL_PREC_BITS-2) * 2
			mv[1] = util.Round2Signed(xc, shared.WARPEDMODEL_PREC_BITS-2) * 2
		}
	}
	mv = t.lowerPrecisionProcess(mv)

	return mv
}

// 7.10.2.10. Lower precision process
func (t *TileGroup) lowerPrecisionProcess(candMv []int) []int {
	if t.State.UncompressedHeader.AllowHighPrecisionMv {
		return candMv
	}

	for i := 0; i < 2; i++ {
		if t.State.UncompressedHeader.ForceIntegerMv {
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
