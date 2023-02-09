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

type UncompressedHeader struct {
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
}

func (u UncompressedHeader) Build(p *Parser, sequenceHeader ObuSequenceHeader, extensionHeader ObuExtensionHeader) {
	var idLen int
	if sequenceHeader.FrameIdNumbersPresent {
		idLen = sequenceHeader.AdditionalFrameIdLengthMinusOne +
			sequenceHeader.DeltaFrameIdLengthMinusTwo + 3
	}

	var frameType int
	var showFrame bool
	var frameIsIntra bool

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
		frameIsIntra = true

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

		frameIsIntra = (frameType == 2 || frameType == 0)

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

	if frameIsIntra {
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

	var frameSizeOverrideFlag bool

	// SWITCH_FRAME
	if frameType == 3 {
		frameSizeOverrideFlag = true
	} else if sequenceHeader.ReducedStillPictureHeader {
		frameSizeOverrideFlag = false

	} else {
		frameSizeOverrideFlag = p.f(1) != 0
	}

	orderHint := p.f(sequenceHeader.OrderHintBits)
	OrderHint := orderHint

	var primaryRefFrame int
	if frameIsIntra || errorResilientMode {
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

	if !frameIsIntra || refreshFrameFlags != allFrames {
		if errorResilientMode && sequenceHeader.EnableOrderHint {
			for i := 0; i < NUM_REF_FRAMES; i++ {
				ref_order_hint[i] = p.f(sequenceHeader.OrderHintBits)

				if ref_order_hint[i] != RefOrderHint[i] {
					RefValid[i] = 0

				}
			}
		}
	}

	var UpscaledWidth int
	var FrameWidth int

	ref_frame_idx := []int{}
	expectedFrameId := []int{}

	if frameIsIntra {
		p.frameSize()
		p.renderSize()

		if allowScreenContentTools && UpscaledWidth == FrameWidth {
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

		if frameSizeOverrideFlag && !errorResilientMode {
			p.frameSizeWithRefs()
		} else {
			p.frameSize()
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
				RefFrameSignBias[refFrame] = p.getRelativeDist(hint, OrderHint) > 0
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
		p.setupPastIndpendence()
	} else {
		p.loadCdfs(ref_frame_idx[primaryRefFrame])
		p.loadPrevious()
	}

	if useRefFrameMvs {
		p.motionFieldEstimation()
	}

	p.tileInfo()
	p.quantizationParams()
	p.segmentationParams()
	p.deltaQParams()
	p.deltaLfParams()

	if primaryRefFrame == PRIMARY_REF_NONE {

		p.initCoeffCdfs()
	} else {
		p.loadPreviousSegementIds()
	}

	CodedLossless := true
	LosslessArray := []bool{}

	// TODO: use real values
	var DeltaQYDc int
	var DeltaQUAc int
	var DeltaQUDc int
	var DeltaQVAc int
	var DeltaQVDc int
	var usingQMatrix bool
	SegQMLevel := [][]int{}
	var qm_y int

	for segmentId := 0; segmentId < MAX_SEGMENTS; segmentId++ {
		qIndex := p.getQIndex(1, segmentId)
		LosslessArray[segmentId] = qIndex == 0 && DeltaQYDc == 0 && DeltaQUAc == 0 && DeltaQUDc == 0 && DeltaQVAc == 0 && DeltaQVDc == 0

		if !LosslessArray[segmentId] {
			CodedLossless = false
		}

		if usingQMatrix {
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

	u.AllLossless = CodedLossless && (FrameWidth == UpscaledWidth)

	p.loopFilterParams()
	p.cdefParams()
	p.lrParams()
	p.readTxMode()
	p.frameReferenceMode()
	p.skipModeParams()

	if frameIsIntra || errorResilientMode || !sequenceHeader.EnableWarpedMotion {
		u.AllowWarpedMotion = false
	} else {
		u.AllowWarpedMotion = p.f(1) != 0
	}

	u.ReducedTxSet = p.f(1) != 0

	p.globalMotionParams()
	p.filmGrainParams()
}

// frame_header_copy()
func (p *Parser) FrameHeaderCopy() {
	panic("not implemented")
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

func (p *Parser) getRelativeDist(hint int, OrderHint int) int {
	panic("not implemented")
}

func (p *Parser) initNonCoeffCdfs() {
	panic("not implemented")
}

func (p *Parser) setupPastIndpendence() {
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

// choose_operating_point()
func (p *Parser) chooseOperatingPoint() int {
	// TODO: implement
	// can be chose by implementation!
	return 0
}
