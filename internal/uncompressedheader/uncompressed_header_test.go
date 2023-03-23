package uncompressedheader

import (
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/stretchr/testify/assert"
)

func TestFrameReferenceModeTrue(t *testing.T) {
	var data = []byte{0b10000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.FrameIsIntra = false
	u.frameReferenceMode(&b)

	assert.True(t, u.ReferenceSelect)
}

func TestFrameReferenceModeFalse(t *testing.T) {
	var data = []byte{0b01111111}
	b := bitstream.NewBitStream(data)
	u := UncompressedHeader{}
	u.FrameIsIntra = false
	u.frameReferenceMode(&b)

	assert.False(t, u.ReferenceSelect)
}

func TestReadTxModeCodedLosslessTrue(t *testing.T) {
	var data = []byte{0b00000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.CodedLossless = true
	u.readTxMode(&b)

	assert.Equal(t, ONLY_4X4, u.TxMode)
}

func TestReadTxModeSelect(t *testing.T) {
	var data = []byte{0b10000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.CodedLossless = false
	u.readTxMode(&b)

	assert.Equal(t, shared.TX_MODE_SELECT, u.TxMode)
}

func TestReadTxModeLargest(t *testing.T) {
	var data = []byte{0b00000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.CodedLossless = false
	u.readTxMode(&b)

	assert.Equal(t, shared.TX_MODE_LARGEST, u.TxMode)
}

func TestDeltaLfParamsNotPresent(t *testing.T) {
	var data = []byte{0b00000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.DeltaQPresent = false
	u.deltaLfParams(&b)

	assert.False(t, u.DeltaLfPresent)
	assert.Equal(t, 0, u.DeltaLfRes)
	assert.Equal(t, 0, u.DeltaLfMulti)

}

func TestDeltaLfParamsPresent(t *testing.T) {
	var data = []byte{0b11010000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.DeltaQPresent = true
	u.AllowIntraBc = false
	u.deltaLfParams(&b)

	assert.True(t, u.DeltaLfPresent)
	assert.Equal(t, 2, u.DeltaLfRes)
	assert.Equal(t, 1, u.DeltaLfMulti)
}

func TestDeltaQParamsNotPresent(t *testing.T) {
	var data = []byte{0b00000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.BaseQIdx = 0
	u.deltaQParams(&b)

	assert.False(t, u.DeltaQPresent)
	assert.Equal(t, 0, u.DeltaQRes)
}

func TestDeltaQParamsPresent(t *testing.T) {
	var data = []byte{0b11100000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.BaseQIdx = 1
	u.deltaQParams(&b)

	assert.True(t, u.DeltaQPresent)
	assert.Equal(t, 3, u.DeltaQRes)
}
