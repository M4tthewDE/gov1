package cdef

type Cdef struct {
	CdefIdx  [][]int
	CdefBits int
}

func (cdef *Cdef) ClearCdef(r int, c int, Use128x128SuperBlock bool, cdefSize4 int) {
	for cap(cdef.CdefIdx) < r+1 {
		cdef.CdefIdx = append(cdef.CdefIdx, make([]int, 0))
	}

	for cap(cdef.CdefIdx[r]) < c+1 {
		cdef.CdefIdx[r] = append(cdef.CdefIdx[r], 0)
	}

	cdef.CdefIdx[r][c] = -1

	if Use128x128SuperBlock {
		cdef.CdefIdx[r][c+cdefSize4] = -1
		cdef.CdefIdx[r+cdefSize4][c] = -1
		cdef.CdefIdx[r+cdefSize4][c+cdefSize4] = -1
	}
}
