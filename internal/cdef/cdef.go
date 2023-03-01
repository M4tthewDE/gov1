package cdef

import "github.com/m4tthewde/gov1/internal/parser"

type Cdef struct {
	CdefIdx  [][]int
	CdefBits int
}

func (cdef *Cdef) clear_cdef(r int, c int, p *parser.Parser) {
	cdef.CdefIdx[r][c] = -1

	if p.sequenceHeader.Use128x128SuperBlock {
		cdefSize4 := p.Num4x4BlocksWide[BLOCK_64X64]
		cdef.CdefIdx[r][c+cdefSize4] = -1
		cdef.CdefIdx[r+cdefSize4][c] = -1
		cdef.CdefIdx[r+cdefSize4][c+cdefSize4] = -1
	}
}
