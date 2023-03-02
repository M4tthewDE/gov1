package parser

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/obu"
)

type Parser struct {
	state     State
	bitStream *bitstream.BitStream
}

func NewParser(b *bitstream.BitStream) Parser {
	return Parser{
		bitStream: b,
	}
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

		inputState := obu.NewState()
		_ = obu.NewObu(obuLength, inputState, p.bitStream)
		// TODO: update state for further obus
		sz -= obuLength

	}
}

// bitstream( )
func (p *Parser) bitstream() {
	for p.bitStream.MoreDataInBistream() {
		temporalUnitSize := p.bitStream.Leb128()
		p.temporalUnit(temporalUnitSize)
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
	return candidateC >= p.state.MiColStart &&
		candidateC < p.state.MiColEnd &&
		candidateR >= p.state.MiRowStart &&
		candidateR < p.state.MiRowEnd
}

// choose_operating_point()
func (p *Parser) ChooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
