package bitstream

import (
	"math"

	"github.com/m4tthewde/gov1/internal/util"
)

type BitStream struct {
	Data        []byte
	Position    int
	Leb128Bytes int
}

func NewBitStream(data []byte) BitStream {
	return BitStream{
		Data:     data,
		Position: 0,
	}
}

// f(n)
func (b *BitStream) F(n int) int {
	x := 0
	for i := 0; i < n; i++ {
		x = 2*x + b.readBit()
		b.Position++
	}

	return x
}

// read_bit()
func (b *BitStream) readBit() int {
	return int((b.Data[int(math.Floor(float64(b.Position)/8))] >> (8 - b.Position%8 - 1)) & 1)
}

func (b *BitStream) MoreDataInBistream() bool {
	return b.Position/8 != len(b.Data)
}

// uvlc()
func (b *BitStream) Uvlc() int {
	leadingZeros := 0

	for {
		done := b.F(1) != 0
		if done {
			break
		}
		leadingZeros++
	}

	if leadingZeros >= 32 {
		return (1 << 32) - 1
	}

	return b.F(leadingZeros) + (1 << leadingZeros) - 1
}

// leb128()
func (b *BitStream) Leb128() int {
	value := 0
	b.Leb128Bytes = 0
	for i := 0; i < 8; i++ {
		leb128_byte := b.F(8)

		value |= int((leb128_byte & 0x7F) << (i * 7))
		b.Leb128Bytes += 1
		if !util.Bool(leb128_byte & 0x80) {
			break
		}

	}

	return value
}

// trailing_bits( nbBits )
func (b *BitStream) TrailingBits(nbBits int) {
	// trailingOneBit
	b.F(1)
	nbBits--

	for nbBits > 0 {
		//trailingZeroBit
		b.F(1)
		nbBits--
	}
}

// byte_alignment()
func (b *BitStream) ByteAlignment() {
	for b.Position&7 != 0 {
		b.F(1)
	}
}

// su()
func (b *BitStream) Su(n int) int {
	value := b.F(n)
	signMask := 1 << uint((n - 1))

	if (value & signMask) != 0 {
		value = value - 2*signMask
	}

	return value
}

// ns( n )
func (b *BitStream) Ns(n int) int {
	w := util.FloorLog2(n) + 1
	m := (1 << w) - n
	v := b.F(w - 1)
	if v < m {
		return v
	}
	extraBit := b.F(1)
	return (v << 1) - m + extraBit
}

// le(n)
func (b *BitStream) Le(n int) int {
	t := 0
	for i := 0; i < n; i++ {
		byte := b.F(8)
		t += (byte << (i * 8))
	}
	return t
}

// NS( n )
func (b *BitStream) NS(n int) int {
	w := util.FloorLog2(n) + 1
	m := (1 << w) - n
	v := b.L(w - 1)
	if v < m {
		return v
	}
	extraBit := b.L(1)
	return (v << 1) - m + extraBit
}

// S()
func (b *BitStream) S() int {
	panic("not implemented: S()")
	return 0
}

// L()
func (b *BitStream) L(a int) int {
	panic("not implemented: L()")
	return 0
}
