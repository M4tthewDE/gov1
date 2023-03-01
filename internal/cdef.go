package main

type Cdef struct {
	CdefIdx  [][]int
	CdefBits int
}

func (cdef *Cdef) clear_cdef(r int, c int, p *Parser) {
	cdef.CdefIdx = SliceAssignNested(cdef.CdefIdx, r, c, -1)

	if p.sequenceHeader.Use128x128SuperBlock {
		cdefSize4 := p.Num4x4BlocksWide[BLOCK_64X64]
		cdef.CdefIdx = SliceAssignNested(cdef.CdefIdx, r, c+cdefSize4, -1)
		cdef.CdefIdx = SliceAssignNested(cdef.CdefIdx, r+cdefSize4, c, -1)
		cdef.CdefIdx = SliceAssignNested(cdef.CdefIdx, r+cdefSize4, c+cdefSize4, -1)
	}
}
