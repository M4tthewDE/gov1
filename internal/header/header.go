package header

import "github.com/m4tthewde/gov1/internal/parser"

type Header struct {
	ForbiddenBit    bool
	Type            int
	ExtensionFlag   bool
	HasSizeField    bool
	ReservedBit     bool
	ExtensionHeader ExtensionHeader
}

type ExtensionHeader struct {
	TemporalID    int
	SpatialID     int
	Reserved3Bits int
}

// obu_header()
func NewHeader(p *parser.Parser) Header {
	forbiddenBit := p.F(1) != 0
	obuType := p.F(4)
	extensionFlag := p.F(1) != 0
	hasSizeField := p.F(1) != 0
	reservedBit := p.F(1) != 0

	if extensionFlag {
		extensionHeader := NewExtensionHeader(p)
		return Header{
			ForbiddenBit:    forbiddenBit,
			Type:            obuType,
			ExtensionFlag:   extensionFlag,
			HasSizeField:    hasSizeField,
			ReservedBit:     reservedBit,
			ExtensionHeader: extensionHeader,
		}
	}

	return Header{
		ForbiddenBit:  forbiddenBit,
		Type:          obuType,
		ExtensionFlag: extensionFlag,
		HasSizeField:  hasSizeField,
		ReservedBit:   reservedBit,
	}
}

// obu_extension(header)
func NewExtensionHeader(p *parser.Parser) ExtensionHeader {
	return ExtensionHeader{
		TemporalID:    p.F(3),
		SpatialID:     p.F(2),
		Reserved3Bits: p.F(3),
	}
}
