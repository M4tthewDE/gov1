package main

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

func NewSequenceHeader(p *Parser) SequenceHeader {
	s := SequenceHeader{}
	s.Build(p)
	return s
}

// sequence_header_obu()
func (s *SequenceHeader) Build(p *Parser) {
	s.SeqProfile = p.f(3)
	s.StillPicture = p.f(1) != 0
	s.ReducedStillPictureHeader = p.f(1) != 0

	s.InitialDisplayDelayPresent = false
	s.OperatingPointsCountMinusOne = 0

	operatingPointIdcArray := []int{0}

	s.SeqLevelIdx = []int{p.f(5)}
	s.SeqTier = []int{0}
	s.DecoderModelPresentForThisOp = []bool{false}
	s.InitialDisplayDelayPresentForThisOp = []bool{false}
	s.InitialDisplayDelayMinusOne = []int{}

	if !s.ReducedStillPictureHeader {
		timingInfoPresent := p.f(1) != 0
		if timingInfoPresent {
			s.TimingInfo = NewTimingInfo(p)
			s.DecoderModelInfoPresent = p.f(1) != 0

			if s.DecoderModelInfoPresent {
				s.DecoderModelInfo = DecoderModelInfo{
					BufferDelayLengthMinusOne:           p.f(5),
					NumUnitsInDecodingTick:              p.f(32),
					BufferRemovalTimeLengthMinusOne:     p.f(5),
					FramePresentationTimeLengthMinusOne: p.f(5),
				}
			}

			s.InitialDisplayDelayPresent = p.f(1) != 0
			s.OperatingPointsCountMinusOne = p.f(5)

			for i := 0; i <= s.OperatingPointsCountMinusOne; i++ {
				if len(operatingPointIdcArray) >= i {
					operatingPointIdcArray = append(operatingPointIdcArray, p.f(12))
				} else {
					operatingPointIdcArray[i] = p.f(12)
				}

				if len(s.SeqLevelIdx) >= i {
					s.SeqLevelIdx = append(s.SeqLevelIdx, p.f(12))
				} else {
					s.SeqLevelIdx[i] = p.f(12)
				}

				if s.SeqLevelIdx[i] > 7 {
					s.SeqTier = append(s.SeqTier, p.f(1))
				} else {
					s.SeqTier = append(s.SeqTier, 0)
				}
				if s.DecoderModelInfoPresent {
					s.DecoderModelPresentForThisOp[i] = p.f(1) != 0
					if len(s.DecoderModelPresentForThisOp) >= i {
						s.DecoderModelPresentForThisOp = append(s.DecoderModelPresentForThisOp, p.f(1) != 0)
					} else {
						s.DecoderModelPresentForThisOp[i] = p.f(1) != 0
					}
					if s.DecoderModelPresentForThisOp[i] {
						// TODO: what are we doing with this?
						_ = s.parseOperatingParametersInfo(p, i)
					}
				} else {
					if len(s.DecoderModelPresentForThisOp) >= i {
						s.DecoderModelPresentForThisOp = append(s.DecoderModelPresentForThisOp, false)
					} else {
						s.DecoderModelPresentForThisOp[i] = false
					}
				}

				if s.InitialDisplayDelayPresent {
					if len(s.InitialDisplayDelayPresentForThisOp) >= i {
						s.InitialDisplayDelayPresentForThisOp = append(s.InitialDisplayDelayPresentForThisOp, p.f(1) != 0)
					} else {
						s.InitialDisplayDelayPresentForThisOp[i] = p.f(1) != 0
					}

					s.InitialDisplayDelayPresentForThisOp[i] = p.f(1) != 0
					if s.InitialDisplayDelayPresentForThisOp[i] {
						if len(s.InitialDisplayDelayMinusOne) >= i {
							s.InitialDisplayDelayMinusOne = append(s.InitialDisplayDelayMinusOne, p.f(4))

						} else {
							s.InitialDisplayDelayMinusOne[i] = p.f(4)
						}
					}
				}
			}
		}
		operatingPoint := p.chooseOperatingPoint()
		p.operatingPointIdc = operatingPointIdcArray[operatingPoint]

		s.FrameWidthBitsMinusOne = p.f(4)
		s.FrameHeightBitsMinusOne = p.f(4)

		n := s.FrameWidthBitsMinusOne + 1
		s.MaxFrameWidthMinusOne = p.f(n)

		n = s.FrameHeightBitsMinusOne + 1
		s.MaxFrameHeightMinusOne = p.f(n)

		s.FrameIdNumbersPresent = false

		if s.ReducedStillPictureHeader {
			s.FrameIdNumbersPresent = false
		} else {
			s.FrameIdNumbersPresent = p.f(1) != 0
		}

		if s.FrameIdNumbersPresent {
			s.DeltaFrameIdLengthMinusTwo = p.f(4)
			s.AdditionalFrameIdLengthMinusOne = p.f(3)
		}

		s.Use128x128SuperBlock = p.f(1) != 0
		s.EnableFilterIntra = p.f(1) != 0
		s.EnableIntraEdgeFilter = p.f(1) != 0
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
			s.EnableInterIntraCompound = p.f(1) != 0
			s.EnableMaskedCompound = p.f(1) != 0
			s.EnableWarpedMotion = p.f(1) != 0
			s.EnableDualFilter = p.f(1) != 0
			s.EnableOrderHint = p.f(1) != 0
			if s.EnableOrderHint {
				s.EnableJntComp = p.f(1) != 0
				s.EnableRefFrameMvs = p.f(1) != 0
			}

			s.SeqChooseScreenContentTools = p.f(1) != 0
			if s.SeqChooseScreenContentTools {
				s.SeqForceScreenContentTools = 2
			} else {
				s.SeqForceScreenContentTools = p.f(1)
			}

			if s.SeqForceScreenContentTools > 0 {
				seqChooseIntegerMv := p.f(1) != 0

				if seqChooseIntegerMv {
					s.SeqForceIntegerMv = 2
				} else {
					s.SeqForceIntegerMv = p.f(1)
				}
			} else {
				s.SeqForceIntegerMv = 2
			}

			if s.EnableOrderHint {
				s.OrderHintBits = p.f(3) + 1
			}
		}
	}

	s.EnableSuperRes = p.f(1) != 0
	s.EnableCdef = p.f(1) != 0
	s.EnableRestoration = p.f(1) != 0
	s.ColorConfig = NewColorConfig(p, s.SeqProfile)
	s.FilmGrainParamsPresent = p.f(1) != 0
}

func NewTimingInfo(p *Parser) TimingInfo {
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

func NewColorConfig(p *Parser, seqProfile int) ColorConfig {
	c := ColorConfig{}
	c.build(p, seqProfile)
	return c
}

func (c *ColorConfig) build(p *Parser, seqProfile int) {
	c.HighBitDepth = p.f(1) != 0

	if seqProfile == 2 && c.HighBitDepth {
		c.TwelveBit = p.f(1) != 0
		c.BitDepth = 10
		if c.TwelveBit {
			c.BitDepth = 12
		}
	}

	c.MonoChrome = false
	if seqProfile != 1 {
		c.MonoChrome = p.f(1) != 0
	}

	c.NumPlanes = 3
	if c.MonoChrome {
		c.NumPlanes = 1
	}

	c.ColorDescriptionPresent = p.f(1) != 0

	c.ColorPrimaries = CP_UNSPECIFIED
	c.TransferCharacteristics = TC_UNSPECIFIED
	c.MatrixCoefficients = MC_UNSPECIFIED

	if c.ColorDescriptionPresent {
		c.ColorPrimaries = p.f(8)
		c.TransferCharacteristics = p.f(8)
		c.MatrixCoefficients = p.f(8)
	}

	if c.MonoChrome {
		c.ColorRange = p.f(1) != 0
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
		c.ColorRange = p.f(1) != 0
		if seqProfile == 0 {
			c.SubsamplingX = true
			c.SubsamplingY = true

		} else if seqProfile == 1 {
			c.SubsamplingX = false
			c.SubsamplingY = false

		} else {
			if c.BitDepth == 12 {
				c.SubsamplingX = p.f(1) != 0
				if c.SubsamplingX {
					c.SubsamplingY = p.f(1) != 0
				} else {
					c.SubsamplingY = false
				}

			} else {
				c.SubsamplingX = true
				c.SubsamplingY = false
			}

		}
		if c.SubsamplingX && c.SubsamplingY {
			c.ChromaSamplePosition = p.f(2)
		}

	}

	c.SeparateUvDeltaQ = p.f(1) != 0
}

// operating_parameters_inf( op )
func (s *SequenceHeader) parseOperatingParametersInfo(p *Parser, bufferDelayLengthMinusOne int) OperatingParametersInfo {
	n := bufferDelayLengthMinusOne + 1

	return OperatingParametersInfo{
		DecoderBufferDelay: []int{p.f(n)},
		EncoderBufferDelay: []int{p.f(n)},
		LowDelayModeFlag:   []bool{p.f(n) != 0},
	}
}
