package parser

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/logger"
	"github.com/m4tthewde/gov1/internal/obu"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"go.uber.org/zap"
)

type Parser struct {
	bitStream          *bitstream.BitStream
	SequenceHeader     sequenceheader.SequenceHeader
	UncompressedHeader uncompressedheader.UncompressedHeader
	obus               []obu.Obu
}

func NewParser(b *bitstream.BitStream) Parser {
	logger.Initialize()

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
	state := state.State{}
	for sz > 0 {
		obuLength := p.bitStream.Leb128()
		sz -= p.bitStream.Leb128Bytes

		logger.Logger.Info("Starting obu parsing process...", zap.Int("sz", sz))

		o := obu.NewObu(obuLength, &state, p.bitStream, p.SequenceHeader, p.UncompressedHeader)
		p.obus = append(p.obus, o)

		switch o.Header.Type {
		case header.OBU_SEQUENCE_HEADER:
			p.SequenceHeader = o.SequenceHeader
		case header.OBU_FRAME_HEADER:
			p.UncompressedHeader = o.UncompressedHeader
		}

		// TODO: update state for further obus
		sz -= obuLength
	}
}

// bitstream( )
func (p *Parser) bitstream() []obu.Obu {
	logger.Logger.Info("Starting bitstream...")
	for p.bitStream.MoreDataInBistream() {
		temporalUnitSize := p.bitStream.Leb128()
		p.temporalUnit(temporalUnitSize)
	}

	return p.obus
}

// choose_operating_point()
func (p *Parser) ChooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
