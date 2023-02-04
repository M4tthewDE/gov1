package main

import (
	"math"
)

type ObuType int

const (
	SequenceHeader       ObuType = 1
	TemporalLimiter      ObuType = 2
	FrameHeader          ObuType = 3
	TileGroup            ObuType = 4
	Metadata             ObuType = 5
	Frame                ObuType = 6
	RedundantFrameHeader ObuType = 7
	TileList             ObuType = 8
	Padding              ObuType = 15
)

type Obu struct {
	Header ObuHeader
	Size   int
}

type ObuHeader struct {
	ForbiddenBit  bool
	Type          ObuType
	ExtensionFlag bool
	HasSizeField  bool
	ReservedBit   bool
}

type ObuExtensionHeader struct {
	TemporalID    int
	SpatialID     int
	Reserved3Bits int
}

type ObuSequenceHeader struct {
	SeqProfile                          int
	StillPicture                        bool
	ReducedStillPictureHeader           bool
	TimingInfoPresent                   bool
	DecoderModelInfoPresent             bool
	OperatingPointsCountMinusOne        bool
	OperatingPointIdc                   []int
	SeqLevelIdx                         []int
	SeqTier                             []int
	DecoderModelPresentForThisOp        []int
	InitialDisplayDelayPresentForThisOp []int
}

type TimingInfo struct {
	NumUnitsInDisplayTick     int
	TimeScale                 int
	EqualPictureInterval      bool
	NumTicksPerMinuteMinusOne int
}

type DecoderModelInfo struct {
	BufferDelayLengthMinusOne           int
	NumUnitsInDecodingTick              int
	BufferRemovalTimeLengthMinusOne     int
	FramePresentationTimeLengthMinusOne int
}

type OperatingParametersInfo struct {
	DecoderBufferDelay []int
	EncoderBufferDelay []int
	LowDelayModeFlag   []bool
}

type Parser struct {
	data     []byte
	position int
}

func NewParser(data []byte) Parser {
	return Parser{
		data:     data,
		position: 0,
	}
}

// TODO: add current bit position
func (p *Parser) currentBit() {
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

// open_bitstream_unit(sz)
func (p *Parser) Parse() Obu {
	obu := Obu{}

	obu.Header = p.ParseObuHeader()

	if obu.Header.HasSizeField {
		obu.Size = p.leb128()
	} else {
		panic("not implemented")
	}

	// TODO: implement rest

	return obu
}

func (p *Parser) ParseEndToEnd() {
}

// obu_header()
func (p *Parser) ParseObuHeader() ObuHeader {
	return ObuHeader{
		ForbiddenBit:  p.f(1) != 0,
		Type:          ObuType(p.f(4)),
		ExtensionFlag: p.f(1) != 0,
		HasSizeField:  p.f(1) != 0,
		ReservedBit:   p.f(1) != 0,
	}
}

// obu_extension(header)
func (p *Parser) ParseObuExtensionHeader() ObuExtensionHeader {
	return ObuExtensionHeader{
		TemporalID:    p.f(3),
		SpatialID:     p.f(2),
		Reserved3Bits: p.f(3),
	}
}

// sequence_header_obu()
func (p *Parser) ParseObuSequenceHeader() ObuSequenceHeader {
	seqProfile := p.f(3)
	stillPicture := p.f(1) != 0
	reducedStillPictureHeader := p.f(1) != 0

	timingInfoPresent := false
	decoderModelInfoPresent := false
	initialDisplayDelayPresent := false
	operatingPointsCountMinusOne := false
	operatingPointIdc := []int{0}
	seqLevelIdx := []int{p.f(5)}
	seqTier := make([]int, 0)
	decoderModelPresentForThisOp := []bool{}
	initialDisplayDelayPresentForThisOp := []bool{}

	if !reducedStillPictureHeader {
		timingInfoPresent := p.f(1) != 0
		if timingInfoPresent {
			timingInfo := p.parseTimingInfo()
			decoderModelInfoPresent := p.f(1) != 0

			if decoderModelInfoPresent {
				decoderModelInfo := DecoderModelInfo{
					BufferDelayLengthMinusOne:           p.f(5),
					NumUnitsInDecodingTick:              p.f(32),
					BufferRemovalTimeLengthMinusOne:     p.f(5),
					FramePresentationTimeLengthMinusOne: p.f(5),
				}
			}
			initialDisplayDelayPresent := p.f(1) != 0
			operatingPointsCountMinusOne := p.f(5)

			for i := 0; i <= operatingPointsCountMinusOne; i++ {
				operatingPointIdc[i] = p.f(12)
				seqLevelIdx[i] = p.f(12)
				if seqLevelIdx[i] > 7 {
					seqTier[i] = p.f(1)
				} else {
					seqTier[i] = 0
				}
				if decoderModelInfoPresent {
					decoderModelPresentForThisOp[i] = p.f(1) != 0
					if decoderModelPresentForThisOp[i] {
						// TODO: what are we doing with this?
						_ = p.parseOperatingParametersInfo()
					}
				}
			}
		}
	}

	return ObuSequenceHeader{
		SeqProfile:                seqProfile,
		StillPicture:              stillPicture,
		ReducedStillPictureHeader: reducedStillPictureHeader,
	}
}

// operating_parameters_inf( op )
func (p *Parser) parseOperatingParametersInfo(bufferDelayLengthMinusOne int) OperatingParametersInfo {
	n := bufferDelayLengthMinusOne + 1

	return OperatingParametersInfo{
		DecoderBufferDelay: []int{p.f(n)},
		EncoderBufferDelay: []int{p.f(n)},
		LowDelayModeFlag:   []bool{p.f(n) != 0},
	}
}

func (p *Parser) parseTimingInfo() TimingInfo {
	numUnitsInDisplayTick := p.f(32)
	timeScale := p.f(32)
	equalPictureInterval := p.f(1) != 0
	numTicksPerPictureMinusOne := 0

	if equalPictureInterval {
		numTicksPerPictureMinusOne = p.uvlc()
	}

	return TimingInfo{
		NumUnitsInDisplayTick:     numUnitsInDisplayTick,
		TimeScale:                 timeScale,
		EqualPictureInterval:      equalPictureInterval,
		NumTicksPerMinuteMinusOne: numTicksPerPictureMinusOne,
	}
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
	Leb128Bytes := 0

	for i := 0; i < 8; i++ {
		leb128_byte := p.f(8)

		value |= int((leb128_byte & 127) << (i * 7))
		Leb128Bytes += 1
		if (leb128_byte & 0x80) == 0 {
			break
		}

	}

	return value
}
