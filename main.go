package main

import (
	"encoding/json"
	"fmt"
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
	ForbiddenBit       bool
	Type               ObuType
	ExtensionFlag      bool
	HasSizeField       bool
	ReservedBit        bool
	ObuExtensionHeader ObuExtensionHeader
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
	DecoderModelInfoPresent             bool
	DecoderModelInfo                    DecoderModelInfo
	InitialDisplayDelayPresent          bool
	OperatingPointsCountMinusOne        int
	SeqLevelIdx                         []int
	SeqTier                             []int
	DecoderModelPresentForThisOp        []bool
	InitialDisplayDelayPresentForThisOp []bool
	InitialDisplayDelayMinusOne         []int
	OperatingPointIdc                   int
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

type ObuFrame struct{}

type UncompressedHeader struct {
	ShowExistingFrame bool
}

type Parser struct {
	data              []byte
	position          int
	operatingPointIdc int
	seenFrameHeader   bool
	leb128Bytes       int
	tileNum           int
}

func NewParser(data []byte) Parser {
	return Parser{
		data:              data,
		position:          0,
		operatingPointIdc: 0,
		seenFrameHeader:   false,
		leb128Bytes:       0,
		tileNum:           0,
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
		p.ParseObu(obuLength)
		sz -= obuLength

	}
}

// open_bitstream_unit(sz)
func (p *Parser) ParseObu(sz int) {
	obu := Obu{}

	obu.Header = p.ParseObuHeader()

	if obu.Header.HasSizeField {
		obu.Size = p.leb128()
	} else {
		extensionFlagInt := 0
		if obu.Header.ExtensionFlag {
			extensionFlagInt = 1
		}
		obu.Size = sz - 1 - extensionFlagInt
	}

	startPosition := p.position

	if obu.Header.Type != SequenceHeader &&
		obu.Header.Type != TemporalLimiter &&
		p.operatingPointIdc != 0 &&
		obu.Header.ExtensionFlag {
		inTemporalLayer := (p.operatingPointIdc >> obu.Header.ObuExtensionHeader.TemporalID) & 1
		inSpatialLayer := (p.operatingPointIdc >> (obu.Header.ObuExtensionHeader.SpatialID + 8)) & 1

		if !(inTemporalLayer != 0) || !(inSpatialLayer != 0) {
			//drop_obu()
			p.position = p.position + obu.Size*8
			return
		}
	}

	x, _ := json.MarshalIndent(obu, "", "	")
	fmt.Printf("%s\n", string(x))

	switch obu.Header.Type {
	case SequenceHeader:
		sequenceHeader := p.ParseObuSequenceHeader()
		x, _ := json.MarshalIndent(sequenceHeader, "", "	")
		fmt.Printf("%s\n", string(x))
	case Frame:
		frame := p.ParseFrame(obu.Size)
		x, _ := json.MarshalIndent(frame, "", "	")
		fmt.Printf("%s\n", string(x))
	}

	payloadBits := p.position - startPosition

	fmt.Println("----------------------------------------")
	fmt.Printf("p.position: %d\n", p.position)
	fmt.Printf("startPosition: %d\n", startPosition)
	fmt.Printf("payloadBits: %d\n", payloadBits)
	fmt.Printf("obu.Size*8 - payloadBits: %d\n", obu.Size*8-payloadBits)
	fmt.Println("----------------------------------------")

	if obu.Size > 0 &&
		obu.Header.Type != TileGroup &&
		obu.Header.Type != TileList &&
		obu.Header.Type != Frame {
		p.trailingBits(obu.Size*8 - payloadBits)
	}

}

// frame_obu( sz )
func (p *Parser) ParseFrame(sz int) ObuFrame {
	startBitPos := p.position

	return ObuFrame{}
}

// frame_header_obu()
func (p *Parser) ParseFrameHeader(sequenceHeader ObuSequenceHeader) {
	if p.seenFrameHeader {
		p.FrameHeaderCopy()
	} else {
		p.seenFrameHeader = true
		uncompressedHeader := p.UncompressedHeader(sequenceHeader)

		if uncompressedHeader.ShowExistingFrame {
			p.DecodeFrameWrapup()
			p.seenFrameHeader = false
		} else {
			p.tileNum = 0
			p.seenFrameHeader = true
		}

	}
}

// frame_header_copy()
func (p *Parser) FrameHeaderCopy() {
	panic("not implemented")
}

// uncompressed_header()
func (p *Parser) UncompressedHeader(sequenceHeader ObuSequenceHeader, extensionHeader ObuExtensionHeader) UncompressedHeader {
	var idLen int
	if sequenceHeader.FrameIdNumbersPresent {
		idLen = sequenceHeader.AdditionalFrameIdLengthMinusOne +
			sequenceHeader.DeltaFrameIdLengthMinusTwo + 3
	}

	var showExistingFrame bool
	var frameType int
	var showFrame bool

	// NUM_REF_FRAMES
	allFrames := ((1 << 8) - 1) != 0
	if sequenceHeader.ReducedStillPictureHeader {
		showExistingFrame = false
		// KEY_FRAME
		frameType = 0
		frameIsIntra := true
		showFrame := true
		showableFrame := false
	} else {
		showExistingFrame := p.f(1) != 0

		if showExistingFrame {
			frameToShowMapIdx := p.f(3)

			if sequenceHeader.DecoderModelInfoPresent && !sequenceHeader.TimingInfo.EqualPictureInterval {
				temporalPointInfo := p.TemporalPointInfo()
			}

			refreshImageFlags := false

			if sequenceHeader.FrameIdNumbersPresent {
				displayFrameId := p.f(idLen)
			}

			frameType := RefFrameType[frameToShowMapIdx]

			// KEY_FRAME
			if frameType == 0 {
				refreshImageFlags = allFrames
			}

			if sequenceHeader.FilmGrainParamsPresent {
				p.LoadGrainParams(frameToShowMapIdx)
			}

			// TODO: fill
			return UncompressedHeader{}
		}
		frameType = p.f(2)

		frameIsIntra := (frameType == 2 || frameType == 0)

		showFrame = p.f(1) != 0

		if showFrame && sequenceHeader.DecoderModelInfoPresent && !sequenceHeader.TimingInfo.EqualPictureInterval {
			temporalPointInfo := p.TemporalPointInfo()
		}

		var showableFrame bool

		if showFrame {
			showableFrame = frameType != 0
		} else {
			showableFrame = p.f(1) != 0
		}

		var errorResilientMode bool

		if frameType == 3 || frameType == 0 && showFrame {
			errorResilientMode = true
		} else {
			errorResilientMode = p.f(1) != 0
		}
	}

	if frameType == 0 && showFrame {
		for i := 0; i < NUM_REF_FRAMES; i++ {
			RefValid[i] = 0
			RefOrderHint[i] = 0
		}

		for i := 0; i < REFS_PER_FRAME; i++ {
			OrderHints[LAST_FRAME+1] = 0
		}
	}

	disableCdfUpdate := p.f(1) != 0

	var allowScreenContentTools bool
	if sequenceHeader.SeqForceScreenContentTools == 2 {
		allowScreenContentTools = p.f(1) != 0
	} else {
		allowScreenContentTools = true
	}

	var forceIntegerMv bool
	if allowScreenContentTools {
		if sequenceHeader.SeqForceIntegerMv == 2 {
			forceIntegerMv = p.f(1) != 0
		} else {
			forceIntegerMv = true
		}
	} else {
		forceIntegerMv = false
	}

	if FrameIsIntra {
		forceIntegerMv = true
	}

	var currentFrameId int

	if frameIdNumbersPresent {
		PrevFrameId = currentFrameId
		currentFrameId := p.f(idLen)
		p.markRefFrames(idLen)
	} else {
		currentFrameId = 0
	}

	var frameSizeOverrideFlag bool

	// SWITCH_FRAME
	if frameType == 3 {
		frameSizeOverrideFlag = true
	} else if reducedStillPictureHeader {
		frameSizeOverrideFlag = false

	} else {
		frameSizeOverrideFlag = p.f(1) != 0
	}

	orderHint := p.f(OrderHintBits)
	OrderHint := orderHint

	var primaryRefName int
	if frameIsIntra || errorResilientMode {
		primaryRefName = 7
	} else {
		primaryRefName := p.f(3)
	}

	if sequenceHeader.DecoderModelInfoPresent {
		bufferRemovalTimePresent := p.f(1) != 0

		if bufferRemovalTimePresent {
			for opNum := 0; opNum <= sequenceHeader.OperatingPointsCountMinusOne; opNum++ {
				if sequenceHeader.DecoderModelPresentForThisOp[opNum] {
					opPtIdc := sequenceHeader.OperatingPointIdc[opNum]
					inTemporalLayer := (opPtIdc >> extensionHeader.TemporalID) & 1
					inSpatialLayer := (opPtIdc >> extensionHeader.SpatialID) & 1

					if opPtIdc == 0 || (inTemporalLayer && inSpatialLayer) {
						n := sequenceHeader.DecoderModelInfo.BufferRemovalTimeLengthMinusOne + 1
						bufferRemovalTime[opNum] = p.f(n)
					}
				}
			}
		}
	}

	allowHighPrecisionMv := false
	useRefFrameMvs := false
	allowIntrabc := false
	var refreshFrameFlags int

	if frameType == 3 || frameType == 0 || showFrame {
		refreshFrameFlags = allFrames
	} else {
		refreshFrameFlags = p.f(8)
	}

	if !FrameIsIntra || refreshFrameFlags != allFrames {
		if errorResilientMode && enableOrderHint {
			for i := 0; i < NUM_REF_FRAMES; i++ {
				ref_order_hint[i] = p.f(OrderHintBits)

				if ref_order_hint[i] != RefOrderHint[i] {
					RefValid[i] = 0

				}
			}
		}
	}

	if FrameIsIntra {
		p.frameSize()
		p.renderSize()

		if allowScreenContentTools && UpscaledWidth == FrameWidth {
			allowIntrabc = p.f(1) != 0
		}
	} else {

		var frameRefsShortSignaling bool
		if !enableOrderHint {
			frameRefsShortSignaling = false
		} else {
			frameRefsShortSignaling = p.f(1) != 0
			if frameRefsShortSignaling {
				lastFrameIdx := p.f(3)
				goldFrameIdx := p.f(3)
				p.setFrameRefs()
			}
		}

		for i := 0; i < REFS_PER_FRAME; i++ {
			if !frameRefsShortSignaling {
				ref_frame_idx[i] = p.f(3)
			}

			if frameIdNumbersPresent {
				n := sequenceHeader.DeltaFrameIdLengthMinusTwo + 2
				deltaFrameIdMinusOne := p.f(n)
				DeltaFrameId := deltaFrameIdMinusOne + 1
				expectedFrameId[i] = ((currentFrameId + (1 << idLen) - DeltaFrameId) % (1 << idLen))
			}
		}

		if frameSizeOverrideFlag && !errorResilientMode {
			p.frameSizeWithRefs()
		} else {
			p.frameSize()
			p.renderSize()

		}
		if forceIntegerMv {
			allowHighPrecisionMv = false
		} else {

			allowHighPrecisionMv := p.f(1) != 0
		}

		p.readInterpolationFilter()
		isMotionModeSwitchable := p.f(1) != 0

		if errorResilientMode || !sequenceHeader.EnableRefFrameMvs {
			useRefFrameMvs = false

		} else {
			useRefFrameMvs = p.f(1) != 0
		}

		for i := 0; i < REFS_PER_FRAME; i++ {
			refFrame := LAST_FRAME + 1
			hint = RefOrderHint[ref_frame_idx[i]]
			OrderHints[refFrame] = hint
			if !enableOrderHint {
				RefFrameSignBias[refFrame] = 0
			} else {
				RefFrameSignBias[refFrame] = p.getRelativeDist(hint, OrderHint) > 0
			}
		}
	}

	var disableFrameEndUpdateCdf bool
	if reducedStillPictureHeader || disableCdfUpdate {
		disableFrameEndUpdateCdf = true
	} else {
		disableFrameEndUpdateCdf = p.f(1) != 0
	}

	if primaryRefName == PRIMARY_REF_NONE {

		p.initNonCoeffCdfs()
		p.setupPastIndpendence()
	} else {
		p.loadCdfs(ref_frame_idx[primaryRefName])
		p.loadPrevious()
	}

	if useRefFrameMvs == 1 {
		p.motionFieldEstimation()
	}

	p.tileInfo()
	p.quantizationParams()
	p.segmentationParams()
	p.deltaQParams()
	p.deltaLfParams()

	if primaryRefName == primaryRefName {

		p.initCoeffCdfs()
	} else {
		p.loadPreviousSegementIds()
	}

	CodedLossless = 1

	for segmentId := 0; segmentId < MAX_SEGMENTS; segmentId++ {
		qIndex := p.getQIndex(1, segmentId)
		LosslessArray[segmentId] = qIndex == 0 && DeltaQYDc == 0 && DeltaQUAc == 0 && DeltaQUDc && DeltaQVAc == 0 && DeltaQVDc == 0

		if !LosslessArray[segmentId] {
			CodedLossless = 0
		}

		if usingQMtraix {
			if LosslessArray[segmentId] {
				SegQMLevel[0][segmentId] = 15
				SegQMLevel[1][segmentId] = 15
				SegQMLevel[2][segmentId] = 15
			} else {
				SegQMLevel[0][segmentId] = qm_y
				SegQMLevel[1][segmentId] = qm_y
				SegQMLevel[2][segmentId] = qm_y

			}
		}
	}

	AllLossless = CodedLossless && (FrameWidth == UpscaledWidth)

	p.loopFilterParams()
	p.cdefParams()
	p.lrParams()
	p.readTxMode()
	p.frameReferenceMode()
	p.skipModeParams()

	var allowWarpedMotion bool
	if FrameIsIntra || errorResilientMode || !sequenceHeader.EnableWarpedMotion {
		allowWarpedMotion = false
	} else {
		allowWarpedMotion = p.f(1) != 0
	}

	reducedTxSet := p.f(1) != 0

	p.globalMotionParams()
	p.filmGrainParams()

	return UncompressedHeader{}
}

func (p *Parser) markRefFrames(a int) {
	panic("not implemented")
}

func (p *Parser) setFrameRefs() {
	panic("not implemented")
}

func (p *Parser) frameSizeWithRefs() {
	panic("not implemented")
}

func (p *Parser) frameSize() {
	panic("not implemented")
}

func (p *Parser) renderSize() {
	panic("not implemented")
}

func (p *Parser) readInterpolationFilter() {
	panic("not implemented")
}

func (p *Parser) getRelativeDist() {
	panic("not implemented")
}

func (p *Parser) initNonCoeffCdfs() {
	panic("not implemented")
}

func (p *Parser) setupPastIndpendence() {
	panic("not implemented")
}

func (p *Parser) loadCdfs() {
	panic("not implemented")
}

func (p *Parser) loadPrevious() {
	panic("not implemented")
}

func (p *Parser) motionFieldEstimation() {
	panic("not implemented")
}

func (p *Parser) tileInfo() {
	panic("not implemented")
}

func (p *Parser) quantizationParams() {
	panic("not implemented")
}

func (p *Parser) segmentationParams() {
	panic("not implemented")
}

func (p *Parser) deltaQParams() {
	panic("not implemented")
}

func (p *Parser) deltaLfParams() {
	panic("not implemented")
}

func (p *Parser) initCoeffCdfs() {
	panic("not implemented")
}

func (p *Parser) loadPreviousSegementIds() {
	panic("not implemented")
}

func (p *Parser) getQIndex(a int, b int) int {
	panic("not implemented")
}

func (p *Parser) loopFilterParams() {
	panic("not implemented")
}

func (p *Parser) cdefParams() {
	panic("not implemented")
}

func (p *Parser) lrParams() {
	panic("not implemented")
}

func (p *Parser) readTxMode() {
	panic("not implemented")
}

func (p *Parser) frameReferenceMode() {
	panic("not implemented")
}

func (p *Parser) skipModeParams() {
	panic("not implemented")
}

func (p *Parser) globalMotionParams() {
	panic("not implemented")
}

func (p *Parser) filmGrainParams() {
	panic("not implemented")
}

// decode_frame_wrapup()
func (p *Parser) DecodeFrameWrapup() {
	panic("not implemented")
}

// temporal_point_info()
func (p *Parser) TemporalPointInfo() int {
	panic("not implemented")
}

// load_grain_params( idx )
func (p *Parser) LoadGrainParams(idx int) {
	panic("not implemented")
}

// obu_header()
func (p *Parser) ParseObuHeader() ObuHeader {
	forbiddenBit := p.f(1) != 0
	obuType := ObuType(p.f(4))
	extensionFlag := p.f(1) != 0
	hasSizeField := p.f(1) != 0
	reservedBit := p.f(1) != 0

	if extensionFlag {
		extensionHeader := p.ParseObuExtensionHeader()
		return ObuHeader{
			ForbiddenBit:       forbiddenBit,
			Type:               obuType,
			ExtensionFlag:      extensionFlag,
			HasSizeField:       hasSizeField,
			ReservedBit:        reservedBit,
			ObuExtensionHeader: extensionHeader,
		}
	}

	return ObuHeader{
		ForbiddenBit:  forbiddenBit,
		Type:          obuType,
		ExtensionFlag: extensionFlag,
		HasSizeField:  hasSizeField,
		ReservedBit:   reservedBit,
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

	var operatingPointIdc int

	seqProfile := p.f(3)
	stillPicture := p.f(1) != 0
	reducedStillPictureHeader := p.f(1) != 0

	timingInfoPresent := false
	decoderModelInfoPresent := false
	initialDisplayDelayPresent := false
	operatingPointsCountMinusOne := 0
	operatingPointIdcArray := []int{0}
	seqLevelIdx := []int{p.f(5)}
	seqTier := []int{0}
	decoderModelPresentForThisOp := []bool{false}
	initialDisplayDelayPresentForThisOp := []bool{false}
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
				if len(operatingPointIdcArray) >= i {
					operatingPointIdcArray = append(operatingPointIdcArray, p.f(12))
				} else {
					operatingPointIdcArray[i] = p.f(12)
				}

				if len(seqLevelIdx) >= i {
					seqLevelIdx = append(seqLevelIdx, p.f(12))
				} else {
					seqLevelIdx[i] = p.f(12)
				}

				if seqLevelIdx[i] > 7 {
					seqTier = append(seqTier, p.f(1))
				} else {
					seqTier = append(seqTier, 0)
				}
				if decoderModelInfoPresent {
					decoderModelPresentForThisOp[i] = p.f(1) != 0
					if len(decoderModelPresentForThisOp) >= i {
						decoderModelPresentForThisOp = append(decoderModelPresentForThisOp, p.f(1) != 0)
					} else {
						decoderModelPresentForThisOp[i] = p.f(1) != 0
					}
					if decoderModelPresentForThisOp[i] {
						// TODO: what are we doing with this?
						_ = p.parseOperatingParametersInfo(i)
					}
				} else {
					if len(decoderModelPresentForThisOp) >= i {
						decoderModelPresentForThisOp = append(decoderModelPresentForThisOp, false)
					} else {
						decoderModelPresentForThisOp[i] = false
					}
				}

				if initialDisplayDelayPresent {
					if len(initialDisplayDelayPresentForThisOp) >= i {
						initialDisplayDelayPresentForThisOp = append(initialDisplayDelayPresentForThisOp, p.f(1) != 0)
					} else {
						initialDisplayDelayPresentForThisOp[i] = p.f(1) != 0
					}
					initialDisplayDelayPresentForThisOp[i] = p.f(1) != 0
					if initialDisplayDelayPresentForThisOp[i] {
						if len(initialDisplayDelayMinusOne) >= i {
							initialDisplayDelayMinusOne = append(initialDisplayDelayMinusOne, p.f(4))

						} else {
							initialDisplayDelayMinusOne[i] = p.f(4)
						}
					}
				}
			}
		}
		operatingPoint := p.chooseOperatingPoint()
		// FIXME: what does this mean
		operatingPointIdc = operatingPointIdcArray[operatingPoint]

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
		DecoderModelInfoPresent:             decoderModelInfoPresent,
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
