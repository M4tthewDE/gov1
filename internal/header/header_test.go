package header

import (
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/stretchr/testify/assert"
)

func TestNewHeader(t *testing.T) {
	var data = []byte{0b10001101, 0b11101001}
	b := bitstream.NewBitStream(data)

	header := NewHeader(&b)

	assert.Equal(t, true, header.ForbiddenBit)
	assert.Equal(t, OBU_SEQUENCE_HEADER, header.Type)
	assert.Equal(t, true, header.ExtensionFlag)
	assert.Equal(t, false, header.HasSizeField)
	assert.Equal(t, true, header.ReservedBit)
}

func TestExtensionHeader(t *testing.T) {
	var data = []byte{0b01101110}
	b := bitstream.NewBitStream(data)

	extensionHeader := NewExtensionHeader(&b)

	assert.Equal(t, 3, extensionHeader.TemporalID)
	assert.Equal(t, 1, extensionHeader.SpatialID)
	assert.Equal(t, 6, extensionHeader.Reserved3Bits)
}
