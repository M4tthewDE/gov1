package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrameReferenceModeTrue(t *testing.T) {
	var data = []byte{0b10000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.FrameIsIntra = false
	u.frameReferenceMode(&p)

	assert.True(t, u.ReferenceSelect)
}

func TestFrameReferenceModeFalse(t *testing.T) {
	var data = []byte{0b01111111}
	p := NewParser(data)
	u := UncompressedHeader{}
	u.FrameIsIntra = false
	u.frameReferenceMode(&p)

	assert.False(t, u.ReferenceSelect)
}

func TestReadTxModeCodedLosslessTrue(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.CodedLossless = true
	u.readTxMode(&p)

	assert.Equal(t, ONLY_4X4, u.TxMode)
}

func TestReadTxModeSelect(t *testing.T) {
	var data = []byte{0b10000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.CodedLossless = false
	u.readTxMode(&p)

	assert.Equal(t, TX_MODE_SELECT, u.TxMode)
}

func TestReadTxModeLargest(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.CodedLossless = false
	u.readTxMode(&p)

	assert.Equal(t, TX_MODE_LARGEST, u.TxMode)
}

func TestDeltaLfParamsNotPresent(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.DeltaQPresent = false
	u.deltaLfParams(&p)

	assert.False(t, u.DeltaLfPresent)
	assert.Equal(t, 0, u.DeltaLfRes)
	assert.Equal(t, 0, u.DeltaLfMulti)

}

func TestDeltaLfParamsPresent(t *testing.T) {
	var data = []byte{0b11010000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.DeltaQPresent = true
	u.AllowIntraBc = false
	u.deltaLfParams(&p)

	assert.True(t, u.DeltaLfPresent)
	assert.Equal(t, 2, u.DeltaLfRes)
	assert.Equal(t, 1, u.DeltaLfMulti)
}

func TestDeltaQParamsNotPresent(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.BaseQIdx = 0
	u.deltaQParams(&p)

	assert.False(t, u.DeltaQPresent)
	assert.Equal(t, 0, u.DeltaQRes)
}

func TestDeltaQParamsPresent(t *testing.T) {
	var data = []byte{0b11100000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.BaseQIdx = 1
	u.deltaQParams(&p)

	assert.True(t, u.DeltaQPresent)
	assert.Equal(t, 3, u.DeltaQRes)
}

func TestGetRelativeDistEnableOrderHintfalse(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)
	u := UncompressedHeader{}

	u.EnableOrderHint = false

	a := 0
	b := 0
	assert.Equal(t, 0, u.getRelativeDist(a, b, &p))
}

func TestGetRelativeDist(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)
	u := UncompressedHeader{}

	u.EnableOrderHint = true
	p.sequenceHeader.OrderHintBits = 2

	a := 10
	b := 5
	assert.Equal(t, 1, u.getRelativeDist(a, b, &p))
}

func TestSuperResparamsSuperResDisabled(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)

	u := UncompressedHeader{}

	u.superResParams(&p)

	assert.False(t, u.UseSuperRes)
	assert.Equal(t, 8, u.SuperResDenom)
	assert.Equal(t, 0, u.FrameWidth)
}

func TestSuperResparams(t *testing.T) {
	var data = []byte{0b10000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	p.sequenceHeader.EnableSuperRes = true

	u.superResParams(&p)

	assert.True(t, u.UseSuperRes)
	assert.Equal(t, 9, u.SuperResDenom)
	assert.Equal(t, 0, u.FrameWidth)
}

func TestComputeImageSize(t *testing.T) {
	var data = []byte{0b00000000}
	p := NewParser(data)

	u := UncompressedHeader{}
	u.FrameWidth = 3
	u.FrameHeight = 2
	u.computeImageSize(&p)

	assert.Equal(t, 2, p.MiCols)
	assert.Equal(t, 2, p.MiRows)
}
