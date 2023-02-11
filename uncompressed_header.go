package main

const NUM_REF_FRAMES = 8
const REFS_PER_FRAME = 7
const KEY_FRAME = 0
const LAST_FRAME = 1
const PRIMARY_REF_NONE = 7
const MAX_SEGMENTS = 8

const EIGHTTAP = 0
const EIGHTTAP_SMOOTH = 1
const EIGHTTAP_SHARP = 2
const BILINEAR = 3
const SWITCHABLE = 4

const ONLY_4X4 = 0
const TX_MODE_LARGEST = 1
const TX_MODE_SELECT = 2

const SUPERRES_DENOM_BITS = 3
const SUPERRES_DENOM_MIN = 9
const SUPERRES_NUM = 8

const SWITCH_FRAME = 3

type UncompressedHeader struct {
	SequenceHeader           ObuSequenceHeader
	ShowExistingFrame        bool
	TemporalPointInfo        int
	ShowableFrame            bool
	RefreshImageFlags        int
	DisplayFrameId           int
	FrameIdNumbersPresent    bool
	AllowHighPrecisionMv     bool
	AllowIntraBc             bool
	LastFrameIdx             int
	GoldFrameIdx             int
	IsMotionModeSwitchable   bool
	DisableFrameEndUpdateCdf bool
	AllLossless              bool
	AllowWarpedMotion        bool
	ReducedTxSet             bool
	FrameIsIntra             bool
	ReferenceSelect          bool
	CodedLossless            bool
	TxMode                   int
	DeltaQPresent            bool
	DeltaQRes                int
	DeltaLfPresent           bool
	DeltaLfRes               int
	DeltaLfMulti             int
	BaseQIdx                 int
	EnableOrderHint          bool
	OrderHintBits            int
	UseSuperRes              bool
	SuperResDenom            int
	FrameWidth               int
	FrameHeight              int
	UpscaledWidth            int
	MiCols                   int
	MiRows                   int
	FrameSizeOverrideFlag    bool
	DeltaQYDc                int
	DeltaQUAc                int
	DeltaQUDc                int
	DeltaQVAc                int
	DeltaQVDc                int
	UsingQMatrix             bool
	Qmy                      int
	Qmu                      int
	Qmv                      int
}

func (u *UncompressedHeader) Build(p *Parser, sequenceHeader ObuSequenceHeader, extensionHeader ObuExtensionHeader) {
	u.SequenceHeader = sequenceHeader

	var idLen int
	if sequenceHeader.FrameIdNumbersPresent {
		idLen = sequenceHeader.AdditionalFrameIdLengthMinusOne +
			sequenceHeader.DeltaFrameIdLengthMinusTwo + 3
	}

	var frameType int
	var showFrame bool

	var errorResilientMode bool

	refFrameType := []int{}
	refValid := []int{}
	refOrderHint := []int{}
	orderHints := []int{}
	bufferRemovalTime := []int{}

	allFrames := ((1 << NUM_REF_FRAMES) - 1)
	if sequenceHeader.ReducedStillPictureHeader {
		u.ShowExistingFrame = false
		frameType = KEY_FRAME
		u.FrameIsIntra = true

		showFrame = true
		u.ShowableFrame = false
	} else {
		showExistingFrame := p.f(1) != 0

		if showExistingFrame {
			frameToShowMapIdx := p.f(3)

			if sequenceHeader.DecoderModelInfoPresent && !sequenceHeader.TimingInfo.EqualPictureInterval {
				u.TemporalPointInfo = p.TemporalPointInfo()
			}

			u.RefreshImageFlags = 0
			if sequenceHeader.FrameIdNumbersPresent {
				u.DisplayFrameId = p.f(idLen)
			}

			frameType := refFrameType[frameToShowMapIdx]

			// KEY_FRAME
			if frameType == KEY_FRAME {
				u.RefreshImageFlags = allFrames
			}

			if sequenceHeader.FilmGrainParamsPresent {
				p.LoadGrainParams(frameToShowMapIdx)
			}
		}

		frameType = p.f(2)

		u.FrameIsIntra = (frameType == 2 || frameType == 0)

		showFrame = p.f(1) != 0

		if showFrame && sequenceHeader.DecoderModelInfoPresent && !sequenceHeader.TimingInfo.EqualPictureInterval {
			u.TemporalPointInfo = p.TemporalPointInfo()
		}

		if showFrame {
			u.ShowableFrame = frameType != 0
		} else {
			u.ShowableFrame = p.f(1) != 0
		}

		if frameType == 3 || frameType == 0 && showFrame {
			errorResilientMode = true
		} else {
			errorResilientMode = p.f(1) != 0
		}

		return
	}

	if frameType == 0 && showFrame {
		for i := 0; i < NUM_REF_FRAMES; i++ {
			refValid[i] = 0
			refOrderHint[i] = 0
		}

		for i := 0; i < REFS_PER_FRAME; i++ {
			orderHints[LAST_FRAME+1] = 0
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

	if u.FrameIsIntra {
		forceIntegerMv = true
	}

	var currentFrameId int

	if sequenceHeader.FrameIdNumbersPresent {
		panic("What is supposed to happen here?")
		//PrevFrameId = currentFrameId
		currentFrameId = p.f(idLen)
		p.markRefFrames(idLen)
	} else {
		currentFrameId = 0
	}

	if frameType == SWITCH_FRAME {
		u.FrameSizeOverrideFlag = true
	} else if sequenceHeader.ReducedStillPictureHeader {
		u.FrameSizeOverrideFlag = false

	} else {
		u.FrameSizeOverrideFlag = p.f(1) != 0
	}

	orderHint := p.f(sequenceHeader.OrderHintBits)
	OrderHint := orderHint

	var primaryRefFrame int
	if u.FrameIsIntra || errorResilientMode {
		primaryRefFrame = PRIMARY_REF_NONE
	} else {
		primaryRefFrame = p.f(3)
	}

	if sequenceHeader.DecoderModelInfoPresent {
		bufferRemovalTimePresent := p.f(1) != 0

		if bufferRemovalTimePresent {
			for opNum := 0; opNum <= sequenceHeader.OperatingPointsCountMinusOne; opNum++ {
				if sequenceHeader.DecoderModelPresentForThisOp[opNum] {
					opPtIdc := sequenceHeader.OperatingPointIdc[opNum]
					inTemporalLayer := ((opPtIdc >> extensionHeader.TemporalID) & 1) != 0
					inSpatialLayer := ((opPtIdc >> extensionHeader.SpatialID) & 1) != 0

					if opPtIdc == 0 || (inTemporalLayer && inSpatialLayer) {
						n := sequenceHeader.DecoderModelInfo.BufferRemovalTimeLengthMinusOne + 1
						bufferRemovalTime[opNum] = p.f(n)
					}
				}
			}
		}
	}

	RefOrderHint := []int{}
	RefValid := []int{}
	ref_order_hint := []int{}

	useRefFrameMvs := false
	var refreshFrameFlags int

	if frameType == 3 || frameType == 0 || showFrame {
		refreshFrameFlags = allFrames
	} else {
		refreshFrameFlags = p.f(8)
	}

	if !u.FrameIsIntra || refreshFrameFlags != allFrames {
		if errorResilientMode && sequenceHeader.EnableOrderHint {
			for i := 0; i < NUM_REF_FRAMES; i++ {
				ref_order_hint[i] = p.f(sequenceHeader.OrderHintBits)

				if ref_order_hint[i] != RefOrderHint[i] {
					RefValid[i] = 0

				}
			}
		}
	}

	ref_frame_idx := []int{}
	expectedFrameId := []int{}

	if u.FrameIsIntra {
		u.frameSize(p)
		p.renderSize()

		if allowScreenContentTools && u.UpscaledWidth == u.FrameWidth {
			u.AllowIntraBc = p.f(1) != 0
		}
	} else {

		var frameRefsShortSignaling bool
		if !sequenceHeader.EnableOrderHint {
			frameRefsShortSignaling = false
		} else {
			frameRefsShortSignaling = p.f(1) != 0
			if frameRefsShortSignaling {
				u.LastFrameIdx = p.f(3)
				u.GoldFrameIdx = p.f(3)
				p.setFrameRefs()
			}
		}

		for i := 0; i < REFS_PER_FRAME; i++ {
			if !frameRefsShortSignaling {
				ref_frame_idx[i] = p.f(3)
			}

			if sequenceHeader.FrameIdNumbersPresent {
				n := sequenceHeader.DeltaFrameIdLengthMinusTwo + 2
				deltaFrameIdMinusOne := p.f(n)
				DeltaFrameId := deltaFrameIdMinusOne + 1
				expectedFrameId[i] = ((currentFrameId + (1 << idLen) - DeltaFrameId) % (1 << idLen))
			}
		}

		if u.FrameSizeOverrideFlag && !errorResilientMode {
			p.frameSizeWithRefs()
		} else {
			u.frameSize(p)
			p.renderSize()

		}
		if forceIntegerMv {
			u.AllowHighPrecisionMv = false
		} else {

			u.AllowHighPrecisionMv = p.f(1) != 0
		}

		p.readInterpolationFilter()
		u.IsMotionModeSwitchable = p.f(1) != 0

		if errorResilientMode || !sequenceHeader.EnableRefFrameMvs {
			useRefFrameMvs = false

		} else {
			useRefFrameMvs = p.f(1) != 0
		}

		OrderHints := []int{}
		RefFrameSignBias := []bool{}

		for i := 0; i < REFS_PER_FRAME; i++ {
			refFrame := LAST_FRAME + 1
			hint := RefOrderHint[ref_frame_idx[i]]
			OrderHints[refFrame] = hint
			if !sequenceHeader.EnableOrderHint {
				RefFrameSignBias[refFrame] = false
			} else {
				RefFrameSignBias[refFrame] = u.getRelativeDist(hint, OrderHint) > 0
			}
		}
	}

	if sequenceHeader.ReducedStillPictureHeader || disableCdfUpdate {
		u.DisableFrameEndUpdateCdf = true
	} else {
		u.DisableFrameEndUpdateCdf = p.f(1) != 0
	}

	if primaryRefFrame == PRIMARY_REF_NONE {

		p.initNonCoeffCdfs()
		p.setupPastIndependence()
	} else {
		p.loadCdfs(ref_frame_idx[primaryRefFrame])
		p.loadPrevious()
	}

	if useRefFrameMvs {
		p.motionFieldEstimation()
	}

	p.tileInfo()
	u.quantizationParams(p)
	p.segmentationParams()
	u.deltaQParams(p)
	u.deltaLfParams(p)

	if primaryRefFrame == PRIMARY_REF_NONE {

		p.initCoeffCdfs()
	} else {
		p.loadPreviousSegementIds()
	}

	u.CodedLossless = true
	LosslessArray := []bool{}

	SegQMLevel := [][]int{}

	for segmentId := 0; segmentId < MAX_SEGMENTS; segmentId++ {
		qIndex := p.getQIndex(1, segmentId)
		LosslessArray[segmentId] = qIndex == 0 && u.DeltaQYDc == 0 && u.DeltaQUAc == 0 && u.DeltaQUDc == 0 && u.DeltaQVAc == 0 && u.DeltaQVDc == 0

		if !LosslessArray[segmentId] {
			u.CodedLossless = false
		}

		if u.UsingQMatrix {
			if LosslessArray[segmentId] {
				SegQMLevel[0][segmentId] = 15
				SegQMLevel[1][segmentId] = 15
				SegQMLevel[2][segmentId] = 15
			} else {
				SegQMLevel[0][segmentId] = u.Qmy
				SegQMLevel[1][segmentId] = u.Qmy
				SegQMLevel[2][segmentId] = u.Qmy

			}
		}
	}

	u.AllLossless = u.CodedLossless && (u.FrameWidth == u.UpscaledWidth)

	p.loopFilterParams()
	p.cdefParams()
	p.lrParams()
	u.readTxMode(p)
	u.frameReferenceMode(p)
	p.skipModeParams()

	if u.FrameIsIntra || errorResilientMode || !sequenceHeader.EnableWarpedMotion {
		u.AllowWarpedMotion = false
	} else {
		u.AllowWarpedMotion = p.f(1) != 0
	}

	u.ReducedTxSet = p.f(1) != 0

	p.globalMotionParams()
	p.filmGrainParams()
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

// frame_size()
func (u *UncompressedHeader) frameSize(p *Parser) {
	if u.FrameSizeOverrideFlag {
		n := u.SequenceHeader.FrameWidthbitsMinusOne + 1
		frameWidthMinusOne := p.f(n)

		n = u.SequenceHeader.FrameHeightbitsMinusOne + 1
		frameHeightMinusOne := p.f(n)

		u.FrameWidth = frameWidthMinusOne + 1
		u.FrameHeight = frameHeightMinusOne + 1
	} else {
		u.FrameWidth = u.SequenceHeader.MaxFrameWidthMinusOne + 1
		u.FrameHeight = u.SequenceHeader.MaxFrameHeightinusOne + 1
	}

	u.superResParams(p)
	u.computeImageSize()
}

// superres_params()
func (u *UncompressedHeader) superResParams(p *Parser) {
	if u.SequenceHeader.EnableSuperRes {
		u.UseSuperRes = p.f(1) != 0
	} else {
		u.UseSuperRes = false
	}

	if u.UseSuperRes {
		codedDenom := p.f(SUPERRES_DENOM_BITS)
		u.SuperResDenom = codedDenom + SUPERRES_DENOM_MIN
	} else {
		u.SuperResDenom = SUPERRES_NUM
	}

	u.UpscaledWidth = u.FrameWidth
	u.FrameWidth = (u.UpscaledWidth*SUPERRES_NUM + (u.SuperResDenom / 2)) / u.SuperResDenom
}

// compute_image_size()
func (u *UncompressedHeader) computeImageSize() {
	u.MiCols = 2 * ((u.FrameWidth + 7) >> 3)
	u.MiRows = 2 * ((u.FrameHeight + 7) >> 3)
}

// render_size()
func (p *Parser) renderSize() {
	renderAndFramSizeDifferent := p.f(1) != 0

	if renderAndFramSizeDifferent {
		renderWidthMinusOne := p.f(16)
		renderHeightMinusOne := p.f(16)

		p.renderWidth = renderWidthMinusOne + 1
		p.renderHeight = renderHeightMinusOne + 1
	} else {
		p.renderWidth = p.upscaledWidth
		p.renderHeight = p.upscaledHeight
	}
}

func (p *Parser) readInterpolationFilter() int {
	isFilterSwitchable := p.f(1) != 0

	var interpolationFilter int
	if isFilterSwitchable {
		interpolationFilter = SWITCHABLE
	} else {
		interpolationFilter = p.f(2)
	}

	return interpolationFilter
}

// get_relative_dist()
func (u *UncompressedHeader) getRelativeDist(a int, b int) int {
	if !u.EnableOrderHint {
		return 0
	}

	diff := a - b
	m := 1 << (u.SequenceHeader.OrderHintBits - 1)
	diff = (diff & (m - 1)) - (diff & m)

	return diff
}

func (p *Parser) initNonCoeffCdfs() {
	panic("not implemented")
}

func (p *Parser) setupPastIndependence() {
	panic("not implemented")
}

func (p *Parser) loadCdfs(a int) {
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

// quantization_params()
func (u *UncompressedHeader) quantizationParams(p *Parser) {
	u.BaseQIdx = p.f(8)

	u.DeltaQYDc = u.readDeltaQ(p)

	var diffUvDelta bool
	if u.SequenceHeader.ColorConfig.NumPlanes > 1 {
		if u.SequenceHeader.ColorConfig.SeparateUvDeltaQ {
			diffUvDelta = p.f(1) != 0
		} else {
			diffUvDelta = false
		}

		u.DeltaQUDc = u.readDeltaQ(p)
		u.DeltaQUAc = u.readDeltaQ(p)

		if diffUvDelta {
			u.DeltaQVDc = u.readDeltaQ(p)
			u.DeltaQVAc = u.readDeltaQ(p)

		} else {
			u.DeltaQVDc = u.DeltaQUDc
			u.DeltaQVAc = u.DeltaQUAc
		}
	} else {
		u.DeltaQUDc = 0
		u.DeltaQUAc = 0
		u.DeltaQVDc = 0
		u.DeltaQVAc = 0
	}

	u.UsingQMatrix = p.f(1) != 0
	if u.UsingQMatrix {
		u.Qmy = p.f(4)
		u.Qmu = p.f(4)

		if !u.SequenceHeader.ColorConfig.SeparateUvDeltaQ {
			u.Qmv = u.Qmu
		} else {
			u.Qmv = p.f(4)
		}
	}
}

// read_delta_q()
func (u *UncompressedHeader) readDeltaQ(p *Parser) int {
	deltaCoded := p.f(1) != 0
	if deltaCoded {
		return p.su(1 + 6)
	} else {
		return 0
	}
}

func (p *Parser) segmentationParams() {
	panic("not implemented")
}

// delta_q_parms()
func (u *UncompressedHeader) deltaQParams(p *Parser) {
	u.DeltaQRes = 0
	u.DeltaQPresent = false

	if u.BaseQIdx > 0 {
		u.DeltaQPresent = p.f(1) != 0
	}

	if u.DeltaQPresent {
		u.DeltaQRes = p.f(2)
	}
}

// delta_lf_params()
func (u *UncompressedHeader) deltaLfParams(p *Parser) {
	u.DeltaLfPresent = false
	u.DeltaLfRes = 0
	u.DeltaLfMulti = 0

	if u.DeltaQPresent {
		if !u.AllowIntraBc {
			u.DeltaLfPresent = p.f(1) != 0
		}

		if u.DeltaLfPresent {
			u.DeltaLfRes = p.f(2)
			u.DeltaLfMulti = p.f(1)
		}
	}
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

// read_tx_mode()
func (u *UncompressedHeader) readTxMode(p *Parser) {
	if u.CodedLossless {
		u.TxMode = ONLY_4X4
	} else {
		txModeSelect := p.f(1) != 0

		if txModeSelect {
			u.TxMode = TX_MODE_SELECT
		} else {
			u.TxMode = TX_MODE_LARGEST
		}
	}
}

// frame_reference_mode()
func (u *UncompressedHeader) frameReferenceMode(p *Parser) {
	if u.FrameIsIntra {
		u.ReferenceSelect = false
	} else {
		u.ReferenceSelect = p.f(1) != 0
	}

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

// choose_operating_point()
func (p *Parser) chooseOperatingPoint() int {
	// TODO: implement
	// can be chose by implementation!
	panic("not implemented")
	return 0
}
