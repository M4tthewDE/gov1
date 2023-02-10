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
