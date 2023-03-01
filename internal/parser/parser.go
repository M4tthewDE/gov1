package parser

import (
	"math"

	"github.com/m4tthewde/gov1/internal"
	"github.com/m4tthewde/gov1/internal/util"
)

type Parser struct {
	data     []byte
	position int

	OperatingPointIdc    int
	seenFrameHeader      bool
	leb128Bytes          int
	renderWidth          int
	renderHeight         int
	upscaledWidth        int
	upscaledHeight       int
	TileNum              int
	MiCols               int
	MiRows               int
	MiColStarts          []int
	MiRowStarts          []int
	MiCol                int
	MiRow                int
	MiSize               int
	MiSizes              [][]int
	MiRowStart           int
	MiColStart           int
	MiRowEnd             int
	MiColEnd             int
	TileColsLog2         int
	TileRowsLog2         int
	TileCols             int
	TileRows             int
	TileSizeBytes        int
	CurrentQIndex        int
	DeltaLF              []int
	RefLrWiener          [][][]int
	Num4x4BlocksWide     []int
	Num4x4BlocksHigh     []int
	ReadDeltas           bool
	Cdef                 internal.Cdef
	BlockDecoded         [][][]int
	FrameRestorationType []int
	LoopRestorationSize  []int
	AvailU               bool
	AvailL               bool
	AvailUChroma         bool
	AvailLChroma         bool
	FeatureEnabled       [][]int
	FeatureData          [][]int
	RefFrame             []int
	RefFrames            [][][]int
	RefFrameWidth        []int
	RefFrameHeight       []int
	GmType               []int
	PrevGmParams         [][]int
	PrevSegmentIds       [][]int
	CurrFrame            [][][]int
	SymbolMaxBits        int
}

func NewParser(data []byte) Parser {
	return Parser{
		data:              data,
		position:          0,
		OperatingPointIdc: 0,
		seenFrameHeader:   false,
		leb128Bytes:       0,
	}
}

// f(n)
func (p *Parser) F(n int) int {
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

// temporal_unit( sz )
func (p *Parser) temporalUnit(sz int) {
	for sz > 0 {
		frameUnitSize := p.leb128()
		sz -= p.leb128Bytes
		p.frameUnit(frameUnitSize)
		sz -= frameUnitSize
	}
}

// frame_unit( sz )
func (p *Parser) frameUnit(sz int) {
	for sz > 0 {
		obuLength := p.leb128()
		sz -= p.leb128Bytes
		p.parseObu(obuLength)
		sz -= obuLength

	}
}

func (p *Parser) moreDataInBistream() bool {
	return p.position/8 != len(p.data)
}

// uvlc()
func (p *Parser) Uvlc() int {
	leadingZeros := 0

	for {
		done := p.F(1) != 0
		if done {
			break
		}
		leadingZeros++
	}

	if leadingZeros >= 32 {
		return (1 << 32) - 1
	}

	return p.F(leadingZeros) + (1 << leadingZeros) - 1
}

// leb128()
func (p *Parser) leb128() int {
	value := 0
	for i := 0; i < 8; i++ {
		leb128_byte := p.F(8)

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
	p.F(1)
	nbBits--

	for nbBits > 0 {
		//trailingZeroBit
		p.F(1)
		nbBits--
	}
}

// byte_alignment()
func (p *Parser) byteAlignment() {
	for p.position&7 != 0 {
		p.F(1)
	}
}

// su()
func (p *Parser) su(n int) int {
	value := p.F(n)
	signMask := 1 << (n - 1)

	if (value & signMask) != 0 {
		value = value - 2*signMask
	}

	return value
}

// ns( n )
func (p *Parser) ns(n int) int {
	w := util.FloorLog2(n) + 1
	m := (1 << w) - n
	v := p.F(w - 1)
	if v < m {
		return v
	}
	extraBit := p.F(1)
	return (v << 1) - m + extraBit
}

// le(n)
func (p *Parser) le(n int) int {
	t := 0
	for i := 0; i < n; i++ {
		byte := p.F(8)
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

// S()
func (p *Parser) S() int {
	panic("not implemented: S()")
	return 0
}

// L()
func (p *Parser) L(a int) int {
	panic("not implemented: L()")
	return 0
}

// NS( n )
func (p *Parser) NS(n int) int {
	w := util.FloorLog2(n) + 1
	m := (1 << w) - n
	v := p.L(w - 1)
	if v < m {
		return v
	}
	extraBit := p.L(1)
	return (v << 1) - m + extraBit
}

func (p *Parser) isInside(candidateR int, candidateC int) bool {
	return candidateC >= p.MiColStart && candidateC < p.MiColEnd && candidateR >= p.MiRowStart && candidateR < p.MiRowEnd
}

// choose_operating_point()
func (p *Parser) ChooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
