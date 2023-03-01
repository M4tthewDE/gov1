package parser

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/obu"
)

type Parser struct {
	state     State
	bitStream bitstream.BitStream
}

func NewParser() Parser {
	return Parser{}
}

// temporal_unit( sz )
func (p *Parser) temporalUnit(sz int) {
	for sz > 0 {
		frameUnitSize := p.bitStream.Leb128()
		sz -= p.bitStream.Leb128Bytes
		p.frameUnit(frameUnitSize)
		sz -= frameUnitSize
	}
}

// frame_unit( sz )
func (p *Parser) frameUnit(sz int) {
	for sz > 0 {
		obuLength := p.bitStream.Leb128()
		sz -= p.bitStream.Leb128Bytes

		_, _ = obu.NewObu(obuLength, &p.bitStream)
		sz -= obuLength

	}
}

// init_symbol( x )
func (p *Parser) initSymbol(a int) {
	panic("not implemented: init_symbol()")
}

// clear_above_context()
func (p *Parser) clearAboveContext() {
	panic("not implemented: clear_above_context()")
}

// clear_left_context( x )
func (p *Parser) clearLeftContext() {
	panic("not implemented: clear_left_context()")
}

func (p *Parser) isInside(candidateR int, candidateC int) bool {
	return candidateC >= p.state.miColStart &&
		candidateC < p.state.miColEnd &&
		candidateR >= p.state.miRowStart &&
		candidateR < p.state.miRowEnd
}

// choose_operating_point()
func (p *Parser) ChooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
