package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadbit(t *testing.T) {
	var data = []byte{0b10001101, 0b11001100}
	p := NewParser(data)

	// byte 0
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 1, p.readBit())
	p.position++

	// byte 1
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 1, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
	p.position++
	assert.Equal(t, 0, p.readBit())
}

func TestF(t *testing.T) {
	var data = []byte{0b10001101, 0b11001100}
	p := NewParser(data)

	assert.Equal(t, 1, p.f(1))
	assert.Equal(t, 0, p.f(2))
	assert.Equal(t, 3, p.f(3))
	assert.Equal(t, 7, p.f(4))
	assert.Equal(t, 6, p.f(5))
	assert.Equal(t, 0, p.f(1))
}

func TestObuHeader(t *testing.T) {
	var data = []byte{0b10001101, 0b11101001}
	p := NewParser(data)

	header := p.ParseObuHeader()

	assert.Equal(t, true, header.ForbiddenBit)
	assert.Equal(t, SequenceHeader, header.Type)
	assert.Equal(t, true, header.ExtensionFlag)
	assert.Equal(t, false, header.HasSizeField)
	assert.Equal(t, true, header.ReservedBit)
}

func TestObuExtensionHeader(t *testing.T) {
	var data = []byte{0b01101110}
	p := NewParser(data)

	extensionHeader := p.ParseObuExtensionHeader()

	assert.Equal(t, 3, extensionHeader.TemporalID)
	assert.Equal(t, 1, extensionHeader.SpatialID)
	assert.Equal(t, 6, extensionHeader.Reserved3Bits)
}

// TODO: add assertions
func TestParseEndToEnd(t *testing.T) {
	fileName := "testdata/argon_coveragetool_av1_base_and_extended_profiles_v2.1/profile0_core/streams/test1228.obu"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	p := NewParser(data)
	obu := p.Parse()
	fmt.Printf("%+v", obu)
}
