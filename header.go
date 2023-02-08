package main

type ObuHeader struct {
	ForbiddenBit       bool
	Type               ObuType
	ExtensionFlag      bool
	HasSizeField       bool
	ReservedBit        bool
	ObuExtensionHeader ObuExtensionHeader
}

type ObuExtensionHeader struct {
	TemporalID    int
	SpatialID     int
	Reserved3Bits int
}

// obu_header()
func (p *Parser) ParseObuHeader() ObuHeader {
	forbiddenBit := p.f(1) != 0
	obuType := ObuType(p.f(4))
	extensionFlag := p.f(1) != 0
	hasSizeField := p.f(1) != 0
	reservedBit := p.f(1) != 0

	if extensionFlag {
		extensionHeader := p.ParseObuExtensionHeader()
		return ObuHeader{
			ForbiddenBit:       forbiddenBit,
			Type:               obuType,
			ExtensionFlag:      extensionFlag,
			HasSizeField:       hasSizeField,
			ReservedBit:        reservedBit,
			ObuExtensionHeader: extensionHeader,
		}
	}

	return ObuHeader{
		ForbiddenBit:  forbiddenBit,
		Type:          obuType,
		ExtensionFlag: extensionFlag,
		HasSizeField:  hasSizeField,
		ReservedBit:   reservedBit,
	}
}

// obu_extension(header)
func (p *Parser) ParseObuExtensionHeader() ObuExtensionHeader {
	return ObuExtensionHeader{
		TemporalID:    p.f(3),
		SpatialID:     p.f(2),
		Reserved3Bits: p.f(3),
	}
}
