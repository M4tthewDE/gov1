package bitstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadbit(t *testing.T) {
	var data = []byte{0b10001101, 0b11001100}
	b := NewBitStream(data)

	// byte 0
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 1, b.readBit())
	b.Position++

	// byte 1
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 1, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
	b.Position++
	assert.Equal(t, 0, b.readBit())
}

func TestF(t *testing.T) {
	var data = []byte{0b10001101, 0b11001100}
	b := NewBitStream(data)

	assert.Equal(t, 1, b.F(1))
	assert.Equal(t, 0, b.F(2))
	assert.Equal(t, 3, b.F(3))
	assert.Equal(t, 7, b.F(4))
	assert.Equal(t, 6, b.F(5))
	assert.Equal(t, 0, b.F(1))
}
