package main

import "math"

type Parser struct {
	data     []byte
	position int

	header             Header
	sequenceHeader     SequenceHeader
	uncompressedHeader UncompressedHeader

	operatingPointIdc int
	seenFrameHeader   bool
	leb128Bytes       int
	renderWidth       int
	renderHeight      int
	upscaledWidth     int
	upscaledHeight    int
	TileNum           int
	MiCols            int
	MiRows            int
	MiColStarts       []int
	MiRowStarts       []int
	MiRowStart        int
	MiColStart        int
	MiRowEnd          int
	MiColEnd          int
	TileColsLog2      int
	TileRowsLog2      int
	TileCols          int
	TileRows          int
	TileSizeBytes     int
	CurrentQIndex     int
	DeltaLF           []int
	RefLrWiener       [][][]int
	Num4x4BlocksWide  []int
	ReadDeltas        bool
	Cdef              Cdef
}

func NewParser(data []byte) Parser {
	return Parser{
		data:              data,
		position:          0,
		operatingPointIdc: 0,
		seenFrameHeader:   false,
		leb128Bytes:       0,
	}
}

// f(n)
func (p *Parser) f(n int) int {
	x := 0
	for i := 0; i < n; i++ {
		x = 2*x + p.readBit()
		p.position++
	}

	return x
}

// read_bit()
func (p *Parser) readBit() int {
	return int((p.data[int(math.Floor(float64(p.position)/8))] >> (8 - p.position%8 - 1)) & 1)
}

// bitstream()
func (p *Parser) bitStream() {
	for p.moreDataInBistream() {
		temporalUnitSize := p.leb128()
		p.temporalUnit(temporalUnitSize)
	}
}

func (p *Parser) moreDataInBistream() bool {
	return p.position/8 != len(p.data)
}

// uvlc()
func (p *Parser) uvlc() int {
	leadingZeros := 0

	for {
		done := p.f(1) != 0
		if done {
			break
		}
		leadingZeros++
	}

	if leadingZeros >= 32 {
		return (1 << 32) - 1
	}

	return p.f(leadingZeros) + (1 << leadingZeros) - 1
}

// leb128()
func (p *Parser) leb128() int {
	value := 0
	for i := 0; i < 8; i++ {
		leb128_byte := p.f(8)

		value |= int((leb128_byte & 127) << (i * 7))
		p.leb128Bytes += 1
		if (leb128_byte & 0x80) == 0 {
			break
		}

	}

	return value
}

// trailing_bits( nbBits )
func (p *Parser) trailingBits(nbBits int) {
	// trailingOneBit
	p.f(1)
	nbBits--

	for nbBits > 0 {
		//trailingZeroBit
		p.f(1)
		nbBits--
	}
}

// byte_alignment()
func (p *Parser) byteAlignment() {
	for p.position&7 != 0 {
		p.f(1)
	}
}

// su()
func (p *Parser) su(n int) int {
	value := p.f(n)
	signMask := 1 << (n - 1)

	if (value & signMask) != 0 {
		value = value - 2*signMask
	}

	return value
}

// ns( n )
func (p *Parser) ns(n int) int {
	w := FloorLog2(n) + 1
	m := (1 << w) - n
	v := p.f(w - 1)
	if v < m {
		return v
	}
	extraBit := p.f(1)
	return (v << 1) - m + extraBit
}

// le(n)
func (p *Parser) le(n int) int {
	t := 0
	for i := 0; i < n; i++ {
		byte := p.f(8)
		t += (byte << (i * 8))
	}
	return t
}

// init_symbol( x )
func (p *Parser) initSymbol(a int) {
	panic("not implemented: init_symbol()")
}

// clear_above_context()
func (p *Parser) clearAboveContext() {
	panic("not implemented: clear_above_context()")
}

// clear_left_context( x )
func (p *Parser) clearLeftContext() {
	panic("not implemented: clear_left_context()")
}
