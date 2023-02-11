package main

type ObuSequenceHeader struct {
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
}

type OperatingParametersInfo struct {
	DecoderBufferDelay []int
	EncoderBufferDelay []int
	LowDelayModeFlag   []bool
}

// sequence_header_obu()
func (s ObuSequenceHeader) Build(p *Parser) {
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
			s.TimingInfo = p.parseTimingInfo()
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
						_ = p.parseOperatingParametersInfo(i)
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
	s.ColorConfig = p.parseColorConfig(s.SeqProfile)
	s.FilmGrainParamsPresent = p.f(1) != 0
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

// operating_parameters_inf( op )
func (p *Parser) parseOperatingParametersInfo(bufferDelayLengthMinusOne int) OperatingParametersInfo {
	n := bufferDelayLengthMinusOne + 1

	return OperatingParametersInfo{
		DecoderBufferDelay: []int{p.f(n)},
		EncoderBufferDelay: []int{p.f(n)},
		LowDelayModeFlag:   []bool{p.f(n) != 0},
	}
}
