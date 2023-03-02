package uncompressedheader

import (
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
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

	assert.Equal(t, TX_MODE_SELECT, u.TxMode)
}

func TestReadTxModeLargest(t *testing.T) {
	var data = []byte{0b00000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.CodedLossless = false
	u.readTxMode(&b)

	assert.Equal(t, TX_MODE_LARGEST, u.TxMode)
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

func TestGetRelativeDistEnableOrderHintfalse(t *testing.T) {
	u := UncompressedHeader{}
	u.EnableOrderHint = false

	x := 0
	y := 0
	assert.Equal(t, 0, u.getRelativeDist(x, y))
}

func TestGetRelativeDist(t *testing.T) {
	u := UncompressedHeader{}
	u.EnableOrderHint = true
	u.State.SequenceHeader.OrderHintBits = 2

	x := 10
	y := 5
	assert.Equal(t, 1, u.getRelativeDist(x, y))
}

func TestSuperResparamsSuperResDisabled(t *testing.T) {
	var data = []byte{0b00000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.superResParams(&b)

	assert.False(t, u.UseSuperRes)
	assert.Equal(t, 8, u.SuperResDenom)
	assert.Equal(t, 0, u.FrameWidth)
}

func TestSuperResparams(t *testing.T) {
	var data = []byte{0b10000000}
	b := bitstream.NewBitStream(data)

	u := UncompressedHeader{}
	u.State.SequenceHeader.EnableSuperRes = true

	u.superResParams(&b)

	assert.True(t, u.UseSuperRes)
	assert.Equal(t, 9, u.SuperResDenom)
	assert.Equal(t, 0, u.FrameWidth)
}

func TestComputeImageSize(t *testing.T) {
	u := UncompressedHeader{}
	u.FrameWidth = 3
	u.FrameHeight = 2
	u.computeImageSize()

	assert.Equal(t, 2, u.State.MiCols)
	assert.Equal(t, 2, u.State.MiRows)
}
