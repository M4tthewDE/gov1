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
	TimingInfo                          TimingInfo
	DecoderModelInfo                    DecoderModelInfo
	InitialDisplayDelayPresent          bool
	OperatingPointsCountMinusOne        int
	SeqLevelIdx                         []int
	SeqTier                             []int
	DecoderModelPresentForThisOp        []bool
	InitialDisplayDelayPresentForThisOp []bool
	InitialDisplayDelayMinusOne         []int
	OperatingPointIdc                   []int
	FrameWidthbitsMinusOne              int
	FrameHeightbitsMinusOne             int
	MaxFrameWidthMinusOne               int
	MaxFrameHeightinusOne               int
	FrameIdNumbersPresent               bool
	AdditionalFrameIdLengthMinusOne     int
	DeltaFrameIdLengthMinusTwo          int
	Use128x128SuperBlock                bool
	EnableFilterIntra                   bool
	EnableIntraEdgeFilter               bool
	EnableInterIntraCompound            bool
	EnableMaskedCompound                bool
	EnableWarpedMotion                  bool
	EnableOrderHint                     bool
	EnableDualFilter                    bool
	EnableJntComp                       bool
	EnableRefFrameMvs                   bool
	SeqChooseScreenContentTools         bool
	SeqForceScreenContentTools          int
	SeqChooseIntegerMv                  bool
	SeqForceIntegerMv                   int
	OrderHintBits                       int
	EnableSuperRes                      bool
	EnableCdef                          bool
	EnableRestoration                   bool
	ColorConfig                         ColorConfig
	FilmGrainParamsPresent              bool
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

type ColorConfig struct {
	HighBitDepth            bool
	TwelveBit               bool
	MonoChrome              bool
	ColorDescriptionPresent bool
	ColorPrimaries          int
	TransferCharacteristics int
	MatrixCoefficients      int
	ColorRange              bool
	SubsamplingX            bool
	SubsamplingY            bool
	ChromaSamplePosition    int
	SeparateUvDeltaQ        bool
	NumPlanes               int
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
	var timingInfo TimingInfo
	var decoderModelInfo DecoderModelInfo
	var frameWidthBitsMinusOne int
	var frameHeightBitsMinusOne int
	var maxFrameWidthMinusOne int
	var maxFrameHeightMinusOne int
	var frameIdNumbersPresent bool
	var additionalFrameIdLengthMinusOne int
	var deltaFrameIdLengthMinusTwo int
	var use128x128Superblock bool
	var enableFilterIntra bool
	var enableIntraEdgeFilter bool
	var enableInterIntraCompound bool
	var enableMaskedCompound bool
	var enableWarpedMotion bool
	var enableOrderHint bool
	var enableDualFilter bool
	var enableJntComp bool
	var enableRefFrameMvs bool
	var seqChooseScreenContentTools bool
	var seqForceScreenContentTools int
	var seqChooseIntegerMv bool
	var seqForceIntegerMv int
	var orderHintBits int

	seqProfile := p.f(3)
	stillPicture := p.f(1) != 0
	reducedStillPictureHeader := p.f(1) != 0

	timingInfoPresent := false
	decoderModelInfoPresent := false
	initialDisplayDelayPresent := false
	operatingPointsCountMinusOne := 0
	operatingPointIdc := []int{0}
	seqLevelIdx := []int{p.f(5)}
	seqTier := []int{}
	decoderModelPresentForThisOp := []bool{}
	initialDisplayDelayPresentForThisOp := []bool{}
	initialDisplayDelayMinusOne := []int{}

	if !reducedStillPictureHeader {
		timingInfoPresent = p.f(1) != 0
		if timingInfoPresent {
			timingInfo = p.parseTimingInfo()
			decoderModelInfoPresent = p.f(1) != 0

			if decoderModelInfoPresent {
				decoderModelInfo = DecoderModelInfo{
					BufferDelayLengthMinusOne:           p.f(5),
					NumUnitsInDecodingTick:              p.f(32),
					BufferRemovalTimeLengthMinusOne:     p.f(5),
					FramePresentationTimeLengthMinusOne: p.f(5),
				}
			}
			initialDisplayDelayPresent = p.f(1) != 0
			operatingPointsCountMinusOne = p.f(5)

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
						_ = p.parseOperatingParametersInfo(i)
					}
				} else {
					decoderModelPresentForThisOp[i] = false
				}

				if initialDisplayDelayPresent {
					initialDisplayDelayPresentForThisOp[i] = p.f(1) != 0
					if initialDisplayDelayPresentForThisOp[i] {
						initialDisplayDelayMinusOne[i] = p.f(4)
					}
				}
			}
		}
		operatingPoint := p.chooseOperatingPoint()
		// FIXME: what does this mean
		operatingPointIdc := operatingPointIdc[operatingPoint]

		frameWidthBitsMinusOne = p.f(4)
		frameHeightBitsMinusOne = p.f(4)

		n := frameWidthBitsMinusOne + 1
		maxFrameWidthMinusOne = p.f(n)

		n = frameHeightBitsMinusOne + 1
		maxFrameHeightMinusOne = p.f(n)

		frameIdNumbersPresent = false

		if reducedStillPictureHeader {
			frameIdNumbersPresent = false
		} else {
			frameIdNumbersPresent = p.f(1) != 0
		}

		if frameIdNumbersPresent {
			deltaFrameIdLengthMinusTwo = p.f(4)
			additionalFrameIdLengthMinusOne = p.f(3)
		}

		use128x128Superblock = p.f(1) != 0
		enableFilterIntra = p.f(1) != 0
		enableIntraEdgeFilter = p.f(1) != 0
		enableInterIntraCompound = false
		enableMaskedCompound = false
		enableWarpedMotion = false
		enableDualFilter = false
		enableOrderHint := false
		enableJntComp = false
		enableRefFrameMvs = false

		// SELECT_SCREEN_CONTENT_TOOLS
		seqForceScreenContentTools = 2

		// SELECT_INTEGER_MV
		seqForceIntegerMv = 2

		orderHintBits = 0

		if !reducedStillPictureHeader {
			enableInterIntraCompound = p.f(1) != 0
			enableMaskedCompound = p.f(1) != 0
			enableWarpedMotion = p.f(1) != 0
			enableDualFilter = p.f(1) != 0
			enableOrderHint = p.f(1) != 0
			if enableOrderHint {
				enableJntComp = p.f(1) != 0
				enableRefFrameMvs = p.f(1) != 0
			}
			seqChooseScreenContentTools = p.f(1) != 0
			if seqChooseScreenContentTools {
				seqForceScreenContentTools = 2
			} else {
				seqForceScreenContentTools = p.f(1)
			}

			if seqForceScreenContentTools > 0 {
				seqChooseIntegerMv := p.f(1) != 0

				if seqChooseIntegerMv {
					seqForceIntegerMv = 2
				} else {
					seqForceIntegerMv = p.f(1)
				}
			} else {
				seqForceIntegerMv = 2
			}

			if enableOrderHint {
				orderHintBits = p.f(3) + 1
			}
		}
	}

	enableSuperRes := p.f(1) != 0
	enableCdef := p.f(1) != 0
	enableRestoration := p.f(1) != 0
	colorConfig := p.parseColorConfig(seqProfile)
	filmGrainParamsPresent := p.f(1) != 0

	return ObuSequenceHeader{
		SeqProfile:                          seqProfile,
		StillPicture:                        stillPicture,
		ReducedStillPictureHeader:           reducedStillPictureHeader,
		TimingInfo:                          timingInfo,
		DecoderModelInfo:                    decoderModelInfo,
		InitialDisplayDelayPresent:          initialDisplayDelayPresent,
		OperatingPointsCountMinusOne:        operatingPointsCountMinusOne,
		SeqLevelIdx:                         seqLevelIdx,
		SeqTier:                             seqTier,
		DecoderModelPresentForThisOp:        decoderModelPresentForThisOp,
		InitialDisplayDelayPresentForThisOp: initialDisplayDelayPresentForThisOp,
		InitialDisplayDelayMinusOne:         initialDisplayDelayMinusOne,
		OperatingPointIdc:                   operatingPointIdc,
		FrameWidthbitsMinusOne:              frameWidthBitsMinusOne,
		FrameHeightbitsMinusOne:             frameHeightBitsMinusOne,
		MaxFrameWidthMinusOne:               maxFrameWidthMinusOne,
		MaxFrameHeightinusOne:               maxFrameHeightMinusOne,
		FrameIdNumbersPresent:               frameIdNumbersPresent,
		AdditionalFrameIdLengthMinusOne:     additionalFrameIdLengthMinusOne,
		DeltaFrameIdLengthMinusTwo:          deltaFrameIdLengthMinusTwo,
		Use128x128SuperBlock:                use128x128Superblock,
		EnableFilterIntra:                   enableFilterIntra,
		EnableIntraEdgeFilter:               enableIntraEdgeFilter,
		EnableInterIntraCompound:            enableInterIntraCompound,
		EnableMaskedCompound:                enableMaskedCompound,
		EnableWarpedMotion:                  enableWarpedMotion,
		EnableOrderHint:                     enableOrderHint,
		EnableDualFilter:                    enableDualFilter,
		EnableJntComp:                       enableJntComp,
		EnableRefFrameMvs:                   enableRefFrameMvs,
		SeqChooseScreenContentTools:         seqChooseScreenContentTools,
		SeqForceScreenContentTools:          seqForceScreenContentTools,
		SeqChooseIntegerMv:                  seqChooseIntegerMv,
		SeqForceIntegerMv:                   seqForceIntegerMv,
		OrderHintBits:                       orderHintBits,
		EnableSuperRes:                      enableSuperRes,
		EnableCdef:                          enableCdef,
		EnableRestoration:                   enableRestoration,
		ColorConfig:                         colorConfig,
		FilmGrainParamsPresent:              filmGrainParamsPresent,
	}
}

func (p *Parser) parseColorConfig(seqProfile int) ColorConfig {
	var bitDepth int
	var twelveBit bool

	highBitDepth := p.f(1) != 0
	if seqProfile == 2 && highBitDepth {
		twelveBit := p.f(1) != 0
		bitDepth = 10
		if twelveBit {
			bitDepth = 12
		}
	}

	monoChrome := false
	if seqProfile != 1 {
		monoChrome = p.f(1) != 0
	}

	numPlanes := 3
	if monoChrome {
		numPlanes = 1
	}
	colorDescriptionPresent := p.f(1) != 0

	// CP_UNSPECIFIED
	colorPrimaries := 2
	// TC_UNSPECIFIED
	transferCharacteristics := 2
	// MC_UNSPECIFIED
	matrixCoefficientes := 2

	if colorDescriptionPresent {
		colorPrimaries = p.f(8)
		transferCharacteristics = p.f(8)
		matrixCoefficientes = p.f(8)
	}

	var subsamplingX bool
	var subsamplingY bool
	var chromaSamplePosition int
	var separateUvDeltaQ bool
	var colorRange bool

	if monoChrome {
		colorRange = p.f(1) != 0
		subsamplingX = true
		subsamplingY = true

		//CSP_UNKNOWN
		chromaSamplePosition = 0
		separateUvDeltaQ = false

		return ColorConfig{
			HighBitDepth:            highBitDepth,
			TwelveBit:               twelveBit,
			MonoChrome:              monoChrome,
			ColorDescriptionPresent: colorDescriptionPresent,
			ColorPrimaries:          colorPrimaries,
			TransferCharacteristics: transferCharacteristics,
			MatrixCoefficients:      matrixCoefficientes,
			ColorRange:              colorRange,
			SubsamplingX:            subsamplingX,
			SubsamplingY:            subsamplingY,
			ChromaSamplePosition:    chromaSamplePosition,
			SeparateUvDeltaQ:        separateUvDeltaQ,
			NumPlanes:               numPlanes,
		}

	} else if colorPrimaries == 1 && transferCharacteristics == 13 && matrixCoefficientes == 0 {
		colorRange = true
		subsamplingX = false
		subsamplingY = false
	} else {
		colorRange = p.f(1) != 0
		if seqProfile == 0 {
			subsamplingX = true
			subsamplingY = true

		} else if seqProfile == 1 {
			subsamplingX = false
			subsamplingY = false

		} else {
			if bitDepth == 12 {
				subsamplingX = p.f(1) != 0
				if subsamplingX {
					subsamplingY = p.f(1) != 0
				} else {
					subsamplingY = false
				}

			} else {
				subsamplingX = true
				subsamplingY = false
			}

		}
		if subsamplingX && subsamplingY {
			chromaSamplePosition = p.f(2)
		}

	}
	separateUvDeltaQ = p.f(1) != 0

	return ColorConfig{
		HighBitDepth:            highBitDepth,
		TwelveBit:               twelveBit,
		MonoChrome:              monoChrome,
		ColorDescriptionPresent: colorDescriptionPresent,
		ColorPrimaries:          colorPrimaries,
		TransferCharacteristics: transferCharacteristics,
		MatrixCoefficients:      matrixCoefficientes,
		ColorRange:              colorRange,
		SubsamplingX:            subsamplingX,
		SubsamplingY:            subsamplingY,
		ChromaSamplePosition:    chromaSamplePosition,
		SeparateUvDeltaQ:        separateUvDeltaQ,
		NumPlanes:               numPlanes,
	}
}

// choose_operating_point()
func (p *Parser) chooseOperatingPoint() int {
	// TODO: implement
	// can be chose by implementation!
	return 0
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
