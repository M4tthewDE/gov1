package sequenceheader

import (
	"github.com/m4tthewde/gov1/internal"
	"github.com/m4tthewde/gov1/internal/parser"
)

const CP_UNSPECIFIED = 2
const TC_UNSPECIFIED = 2
const MC_UNSPECIFIED = 2

type SequenceHeader struct {
	SeqProfile                          int
	StillPicture                        bool
	ReducedStillPictureHeader           bool
	TimingInfo                          TimingInfo
	DecoderModelInfoPresent             bool
	DecoderModelInfo                    DecoderModelInfo
	InitialDisplayDelayPresent          bool
	OperatingPointsCountMinusOne        int
	SeqLevelIdx                         []int
	SeqTier                             []int
	DecoderModelPresentForThisOp        []bool
	InitialDisplayDelayPresentForThisOp []bool
	InitialDisplayDelayMinusOne         []int
	OperatingPointIdc                   []int
	FrameWidthBitsMinusOne              int
	FrameHeightBitsMinusOne             int
	MaxFrameWidthMinusOne               int
	MaxFrameHeightMinusOne              int
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
	BitDepth                int
}

type OperatingParametersInfo struct {
	DecoderBufferDelay []int
	EncoderBufferDelay []int
	LowDelayModeFlag   []bool
}

func NewSequenceHeader(p *parser.Parser) SequenceHeader {
	s := SequenceHeader{}
	s.Build(p)
	return s
}

// sequence_header_obu()
func (s *SequenceHeader) Build(p *parser.Parser) {
	s.SeqProfile = p.F(3)
	s.StillPicture = p.F(1) != 0
	s.ReducedStillPictureHeader = p.F(1) != 0

	s.InitialDisplayDelayPresent = false
	s.OperatingPointsCountMinusOne = 0

	operatingPointIdcArray := []int{0}

	s.SeqLevelIdx = []int{p.F(5)}
	s.SeqTier = []int{0}
	s.DecoderModelPresentForThisOp = []bool{false}
	s.InitialDisplayDelayPresentForThisOp = []bool{false}
	s.InitialDisplayDelayMinusOne = []int{}

	if !s.ReducedStillPictureHeader {
		timingInfoPresent := p.F(1) != 0
		if timingInfoPresent {
			s.TimingInfo = NewTimingInfo(p)
			s.DecoderModelInfoPresent = p.F(1) != 0

			if s.DecoderModelInfoPresent {
				s.DecoderModelInfo = DecoderModelInfo{
					BufferDelayLengthMinusOne:           p.F(5),
					NumUnitsInDecodingTick:              p.F(32),
					BufferRemovalTimeLengthMinusOne:     p.F(5),
					FramePresentationTimeLengthMinusOne: p.F(5),
				}
			}

			s.InitialDisplayDelayPresent = internal.Bool(p.F(1))
			s.OperatingPointsCountMinusOne = p.F(5)

			for i := 0; i <= s.OperatingPointsCountMinusOne; i++ {
				operatingPointIdcArray[i] = p.F(12)
				s.SeqLevelIdx[i] = p.F(12)

				if s.SeqLevelIdx[i] > 7 {
					s.SeqTier[i] = p.F(1)
				} else {
					s.SeqTier[i] = 0
				}
				if s.DecoderModelInfoPresent {
					s.DecoderModelPresentForThisOp[i] = internal.Bool(p.F(1))

					if s.DecoderModelPresentForThisOp[i] {
						// TODO: what are we doing with this?
						_ = NewOperatingParametersInfo(p, i)
					}
				} else {
					s.DecoderModelPresentForThisOp[i] = false
				}

				if s.InitialDisplayDelayPresent {
					s.InitialDisplayDelayPresentForThisOp[i] = internal.Bool(p.F(1))

					if s.InitialDisplayDelayPresentForThisOp[i] {
						s.InitialDisplayDelayMinusOne[i] = p.F(4)
					}
				}
			}
		}
		operatingPoint := p.ChooseOperatingPoint()
		p.OperatingPointIdc = operatingPointIdcArray[operatingPoint]

		s.FrameWidthBitsMinusOne = p.F(4)
		s.FrameHeightBitsMinusOne = p.F(4)

		n := s.FrameWidthBitsMinusOne + 1
		s.MaxFrameWidthMinusOne = p.F(n)

		n = s.FrameHeightBitsMinusOne + 1
		s.MaxFrameHeightMinusOne = p.F(n)

		s.FrameIdNumbersPresent = false

		if s.ReducedStillPictureHeader {
			s.FrameIdNumbersPresent = false
		} else {
			s.FrameIdNumbersPresent = p.F(1) != 0
		}

		if s.FrameIdNumbersPresent {
			s.DeltaFrameIdLengthMinusTwo = p.F(4)
			s.AdditionalFrameIdLengthMinusOne = p.F(3)
		}

		s.Use128x128SuperBlock = p.F(1) != 0
		s.EnableFilterIntra = p.F(1) != 0
		s.EnableIntraEdgeFilter = p.F(1) != 0
		s.EnableInterIntraCompound = false
		s.EnableMaskedCompound = false
		s.EnableWarpedMotion = false
		s.EnableDualFilter = false
		s.EnableOrderHint = false
		s.EnableJntComp = false
		s.EnableRefFrameMvs = false

		// SELECT_SCREEN_CONTENT_TOOLS
		s.SeqForceScreenContentTools = 2

		// SELECT_INTEGER_MV
		s.SeqForceIntegerMv = 2

		s.OrderHintBits = 0

		if !s.ReducedStillPictureHeader {
			s.EnableInterIntraCompound = p.F(1) != 0
			s.EnableMaskedCompound = p.F(1) != 0
			s.EnableWarpedMotion = p.F(1) != 0
			s.EnableDualFilter = p.F(1) != 0
			s.EnableOrderHint = p.F(1) != 0
			if s.EnableOrderHint {
				s.EnableJntComp = p.F(1) != 0
				s.EnableRefFrameMvs = p.F(1) != 0
			}

			s.SeqChooseScreenContentTools = p.F(1) != 0
			if s.SeqChooseScreenContentTools {
				s.SeqForceScreenContentTools = 2
			} else {
				s.SeqForceScreenContentTools = p.F(1)
			}

			if s.SeqForceScreenContentTools > 0 {
				seqChooseIntegerMv := p.F(1) != 0

				if seqChooseIntegerMv {
					s.SeqForceIntegerMv = 2
				} else {
					s.SeqForceIntegerMv = p.F(1)
				}
			} else {
				s.SeqForceIntegerMv = 2
			}

			if s.EnableOrderHint {
				s.OrderHintBits = p.F(3) + 1
			}
		}
	}

	s.EnableSuperRes = p.F(1) != 0
	s.EnableCdef = p.F(1) != 0
	s.EnableRestoration = p.F(1) != 0
	s.ColorConfig = NewColorConfig(p, s.SeqProfile)
	s.FilmGrainParamsPresent = p.F(1) != 0
}

func NewTimingInfo(p *parser.Parser) TimingInfo {
	numUnitsInDisplayTick := p.F(32)
	timeScale := p.F(32)
	equalPictureInterval := p.F(1) != 0
	numTicksPerPictureMinusOne := 0

	if equalPictureInterval {
		numTicksPerPictureMinusOne = p.Uvlc()
	}

	return TimingInfo{
		NumUnitsInDisplayTick:     numUnitsInDisplayTick,
		TimeScale:                 timeScale,
		EqualPictureInterval:      equalPictureInterval,
		NumTicksPerMinuteMinusOne: numTicksPerPictureMinusOne,
	}
}

func NewColorConfig(p *parser.Parser, seqProfile int) ColorConfig {
	c := ColorConfig{}
	c.build(p, seqProfile)
	return c
}

func (c *ColorConfig) build(p *parser.Parser, seqProfile int) {
	c.HighBitDepth = p.F(1) != 0

	if seqProfile == 2 && c.HighBitDepth {
		c.TwelveBit = p.F(1) != 0
		c.BitDepth = 10
		if c.TwelveBit {
			c.BitDepth = 12
		}
	}

	c.MonoChrome = false
	if seqProfile != 1 {
		c.MonoChrome = p.F(1) != 0
	}

	c.NumPlanes = 3
	if c.MonoChrome {
		c.NumPlanes = 1
	}

	c.ColorDescriptionPresent = p.F(1) != 0

	c.ColorPrimaries = CP_UNSPECIFIED
	c.TransferCharacteristics = TC_UNSPECIFIED
	c.MatrixCoefficients = MC_UNSPECIFIED

	if c.ColorDescriptionPresent {
		c.ColorPrimaries = p.F(8)
		c.TransferCharacteristics = p.F(8)
		c.MatrixCoefficients = p.F(8)
	}

	if c.MonoChrome {
		c.ColorRange = p.F(1) != 0
		c.SubsamplingX = true
		c.SubsamplingY = true

		//CSP_UNKNOWN
		c.ChromaSamplePosition = 0
		c.SeparateUvDeltaQ = false

		return

	} else if c.ColorPrimaries == 1 && c.TransferCharacteristics == 13 && c.MatrixCoefficients == 0 {
		c.ColorRange = true
		c.SubsamplingX = false
		c.SubsamplingY = false
	} else {
		c.ColorRange = p.F(1) != 0
		if seqProfile == 0 {
			c.SubsamplingX = true
			c.SubsamplingY = true

		} else if seqProfile == 1 {
			c.SubsamplingX = false
			c.SubsamplingY = false

		} else {
			if c.BitDepth == 12 {
				c.SubsamplingX = p.F(1) != 0
				if c.SubsamplingX {
					c.SubsamplingY = p.F(1) != 0
				} else {
					c.SubsamplingY = false
				}

			} else {
				c.SubsamplingX = true
				c.SubsamplingY = false
			}

		}
		if c.SubsamplingX && c.SubsamplingY {
			c.ChromaSamplePosition = p.F(2)
		}

	}

	c.SeparateUvDeltaQ = p.F(1) != 0
}

// operating_parameters_info( op )
func NewOperatingParametersInfo(p *parser.Parser, bufferDelayLengthMinusOne int) OperatingParametersInfo {
	n := bufferDelayLengthMinusOne + 1

	return OperatingParametersInfo{
		DecoderBufferDelay: []int{p.F(n)},
		EncoderBufferDelay: []int{p.F(n)},
		LowDelayModeFlag:   []bool{p.F(n) != 0},
	}
}
