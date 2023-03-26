package sequenceheader

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/util"
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

	DecoderBufferDelay []int
	EncoderBufferDelay []int
	LowDelayModeFlag   []bool
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
	// can only be 3 or 1
	NumPlanes int
	// can only be 10 or 12
	BitDepth int
}

type OperatingParametersInfo struct {
	DecoderBufferDelay []int
	EncoderBufferDelay []int
	LowDelayModeFlag   []bool
}

// sequence_header_obu()
func NewSequenceHeader(b *bitstream.BitStream, state *state.State) SequenceHeader {
	s := SequenceHeader{}

	s.SeqProfile = b.F(3)
	s.StillPicture = b.F(1) != 0
	s.ReducedStillPictureHeader = b.F(1) != 0

	s.InitialDisplayDelayPresent = false
	s.OperatingPointsCountMinusOne = 0

	operatingPointIdcArray := []int{0}

	s.SeqLevelIdx = []int{b.F(5)}
	s.SeqTier = []int{0}
	s.DecoderModelPresentForThisOp = []bool{false}
	s.InitialDisplayDelayPresentForThisOp = []bool{false}
	s.InitialDisplayDelayMinusOne = []int{}

	if !s.ReducedStillPictureHeader {
		timingInfoPresent := b.F(1) != 0
		if timingInfoPresent {
			s.TimingInfo = NewTimingInfo(b)
			s.DecoderModelInfoPresent = b.F(1) != 0

			if s.DecoderModelInfoPresent {
				s.DecoderModelInfo = DecoderModelInfo{
					BufferDelayLengthMinusOne:           b.F(5),
					NumUnitsInDecodingTick:              b.F(32),
					BufferRemovalTimeLengthMinusOne:     b.F(5),
					FramePresentationTimeLengthMinusOne: b.F(5),
				}
			}

			s.InitialDisplayDelayPresent = util.Bool(b.F(1))
			s.OperatingPointsCountMinusOne = b.F(5)

			for i := 0; i <= s.OperatingPointsCountMinusOne; i++ {
				operatingPointIdcArray[i] = b.F(12)
				s.SeqLevelIdx[i] = b.F(12)

				if s.SeqLevelIdx[i] > 7 {
					s.SeqTier[i] = b.F(1)
				} else {
					s.SeqTier[i] = 0
				}
				if s.DecoderModelInfoPresent {
					s.DecoderModelPresentForThisOp[i] = util.Bool(b.F(1))

					if s.DecoderModelPresentForThisOp[i] {
						s.operatingParametersInfo(b, i)
					}
				} else {
					s.DecoderModelPresentForThisOp[i] = false
				}

				if s.InitialDisplayDelayPresent {
					s.InitialDisplayDelayPresentForThisOp[i] = util.Bool(b.F(1))

					if s.InitialDisplayDelayPresentForThisOp[i] {
						s.InitialDisplayDelayMinusOne[i] = b.F(4)
					}
				}
			}
		}
		operatingPoint := chooseOperatingPoint()
		state.OperatingPointIdc = operatingPointIdcArray[operatingPoint]

		s.FrameWidthBitsMinusOne = b.F(4)
		s.FrameHeightBitsMinusOne = b.F(4)

		n := s.FrameWidthBitsMinusOne + 1
		s.MaxFrameWidthMinusOne = b.F(n)

		n = s.FrameHeightBitsMinusOne + 1
		s.MaxFrameHeightMinusOne = b.F(n)

		s.FrameIdNumbersPresent = false

		if s.ReducedStillPictureHeader {
			s.FrameIdNumbersPresent = false
		} else {
			s.FrameIdNumbersPresent = b.F(1) != 0
		}

		if s.FrameIdNumbersPresent {
			s.DeltaFrameIdLengthMinusTwo = b.F(4)
			s.AdditionalFrameIdLengthMinusOne = b.F(3)
		}

		s.Use128x128SuperBlock = b.F(1) != 0
		s.EnableFilterIntra = b.F(1) != 0
		s.EnableIntraEdgeFilter = b.F(1) != 0
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
			s.EnableInterIntraCompound = b.F(1) != 0
			s.EnableMaskedCompound = b.F(1) != 0
			s.EnableWarpedMotion = b.F(1) != 0
			s.EnableDualFilter = b.F(1) != 0
			s.EnableOrderHint = b.F(1) != 0
			if s.EnableOrderHint {
				s.EnableJntComp = b.F(1) != 0
				s.EnableRefFrameMvs = b.F(1) != 0
			}

			s.SeqChooseScreenContentTools = b.F(1) != 0
			if s.SeqChooseScreenContentTools {
				s.SeqForceScreenContentTools = 2
			} else {
				s.SeqForceScreenContentTools = b.F(1)
			}

			if s.SeqForceScreenContentTools > 0 {
				seqChooseIntegerMv := b.F(1) != 0

				if seqChooseIntegerMv {
					s.SeqForceIntegerMv = 2
				} else {
					s.SeqForceIntegerMv = b.F(1)
				}
			} else {
				s.SeqForceIntegerMv = 2
			}

			if s.EnableOrderHint {
				s.OrderHintBits = b.F(3) + 1
			}
		}
	}

	s.EnableSuperRes = b.F(1) != 0
	s.EnableCdef = b.F(1) != 0
	s.EnableRestoration = b.F(1) != 0
	s.ColorConfig = NewColorConfig(b, s.SeqProfile)
	s.FilmGrainParamsPresent = b.F(1) != 0

	return s
}

func NewTimingInfo(b *bitstream.BitStream) TimingInfo {
	numUnitsInDisplayTick := b.F(32)
	timeScale := b.F(32)
	equalPictureInterval := b.F(1) != 0
	numTicksPerPictureMinusOne := 0

	if equalPictureInterval {
		numTicksPerPictureMinusOne = b.Uvlc()
	}

	return TimingInfo{
		NumUnitsInDisplayTick:     numUnitsInDisplayTick,
		TimeScale:                 timeScale,
		EqualPictureInterval:      equalPictureInterval,
		NumTicksPerMinuteMinusOne: numTicksPerPictureMinusOne,
	}
}

func NewColorConfig(b *bitstream.BitStream, seqProfile int) ColorConfig {
	c := ColorConfig{}

	c.HighBitDepth = b.F(1) != 0

	if seqProfile == 2 && c.HighBitDepth {
		c.TwelveBit = b.F(1) != 0
		c.BitDepth = 10
		if c.TwelveBit {
			c.BitDepth = 12
		}
	} else if seqProfile <= 2 {
		c.BitDepth = 8
		if c.HighBitDepth {
			c.BitDepth = 10
		}
	}

	c.MonoChrome = false
	if seqProfile != 1 {
		c.MonoChrome = b.F(1) != 0
	}

	c.NumPlanes = 3
	if c.MonoChrome {
		c.NumPlanes = 1
	}

	c.ColorDescriptionPresent = b.F(1) != 0

	c.ColorPrimaries = CP_UNSPECIFIED
	c.TransferCharacteristics = TC_UNSPECIFIED
	c.MatrixCoefficients = MC_UNSPECIFIED

	if c.ColorDescriptionPresent {
		c.ColorPrimaries = b.F(8)
		c.TransferCharacteristics = b.F(8)
		c.MatrixCoefficients = b.F(8)
	}

	if c.MonoChrome {
		c.ColorRange = b.F(1) != 0
		c.SubsamplingX = true
		c.SubsamplingY = true

		//CSP_UNKNOWN
		c.ChromaSamplePosition = 0
		c.SeparateUvDeltaQ = false

		return c
	} else if c.ColorPrimaries == 1 && c.TransferCharacteristics == 13 && c.MatrixCoefficients == 0 {
		c.ColorRange = true
		c.SubsamplingX = false
		c.SubsamplingY = false
	} else {
		c.ColorRange = b.F(1) != 0
		if seqProfile == 0 {
			c.SubsamplingX = true
			c.SubsamplingY = true

		} else if seqProfile == 1 {
			c.SubsamplingX = false
			c.SubsamplingY = false

		} else {
			if c.BitDepth == 12 {
				c.SubsamplingX = b.F(1) != 0
				if c.SubsamplingX {
					c.SubsamplingY = b.F(1) != 0
				} else {
					c.SubsamplingY = false
				}

			} else {
				c.SubsamplingX = true
				c.SubsamplingY = false
			}

		}
		if c.SubsamplingX && c.SubsamplingY {
			c.ChromaSamplePosition = b.F(2)
		}

	}

	c.SeparateUvDeltaQ = b.F(1) != 0
	return c
}

// operating_parameters_info( op )
func (sh *SequenceHeader) operatingParametersInfo(b *bitstream.BitStream, op int) {
	n := sh.DecoderModelInfo.BufferDelayLengthMinusOne + 1

	sh.DecoderBufferDelay[op] = b.F(n)
	sh.EncoderBufferDelay[op] = b.F(n)
	sh.LowDelayModeFlag[op] = util.Bool(b.F(1))
}

// choose_operating_point()
func chooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
