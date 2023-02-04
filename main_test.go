package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObuHeader(t *testing.T) {
	var input byte = 0b10001101
	header := ParseObuHeader(input)

	assert.Equal(t, true, header.ForbiddenBit)
	assert.Equal(t, ObuSequenceHeader, header.Type)
	assert.Equal(t, true, header.ExtensionFlag)
	assert.Equal(t, false, header.HasSizeField)
	assert.Equal(t, true, header.ReservedBit)
}

func TestObujExtensionHeader(t *testing.T) {
	var input byte = 0b01101110
	extensionHeader := ParseObuExtensionHeader(input)
	assert.Equal(t, 3, extensionHeader.TemporalID)
	assert.Equal(t, 1, extensionHeader.SpatialID)
	assert.Equal(t, 6, extensionHeader.Reserved3Bits)
}
