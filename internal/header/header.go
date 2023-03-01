package header

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
)

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
func NewHeader(b *bitstream.BitStream) Header {
	forbiddenBit := b.F(1) != 0
	obuType := b.F(4)
	extensionFlag := b.F(1) != 0
	hasSizeField := b.F(1) != 0
	reservedBit := b.F(1) != 0

	if extensionFlag {
		extensionHeader := NewExtensionHeader(b)
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
func NewExtensionHeader(b *bitstream.BitStream) ExtensionHeader {
	return ExtensionHeader{
		TemporalID:    b.F(3),
		SpatialID:     b.F(2),
		Reserved3Bits: b.F(3),
	}
}
