package main

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
func NewHeader(p *Parser) Header {
	forbiddenBit := p.f(1) != 0
	obuType := p.f(4)
	extensionFlag := p.f(1) != 0
	hasSizeField := p.f(1) != 0
	reservedBit := p.f(1) != 0

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
func NewExtensionHeader(p *Parser) ExtensionHeader {
	return ExtensionHeader{
		TemporalID:    p.f(3),
		SpatialID:     p.f(2),
		Reserved3Bits: p.f(3),
	}
}
