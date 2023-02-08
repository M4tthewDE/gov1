package main

type UncompressedHeader struct {
	ShowExistingFrame     bool
	TemporalPointInfo     int
	ShowableFrame         bool
	RefreshImageFlags     int
	DisplayFrameId        int
	FrameIdNumbersPresent bool
}

const NUM_REF_FRAMES = 8
const REFS_PER_FRAME = 7
const KEY_FRAME = 0
const LAST_FRAME = 1
const PRIMARY_REF_NONE = 7

// uncompressed_header()
func (p *Parser) UncompressedHeader(sequenceHeader ObuSequenceHeader, extensionHeader ObuExtensionHeader, inTemporalLayer bool) UncompressedHeader {
	var idLen int
	if sequenceHeader.FrameIdNumbersPresent {
		idLen = sequenceHeader.AdditionalFrameIdLengthMinusOne +
			sequenceHeader.DeltaFrameIdLengthMinusTwo + 3
	}

	var showExistingFrame bool
	var frameType int
	var showFrame bool
	var showableFrame bool
	var frameIsIntra bool

	var temporalPointInfo int
	var errorResilientMode bool

	var refreshImageFlags int
	var displayFrameId int

	refFrameType := []int{}
	refValid := []int{}
	refOrderHint := []int{}
	orderHints := []int{}
	bufferRemovalTime := []int{}

	allFrames := ((1 << NUM_REF_FRAMES) - 1)
	if sequenceHeader.ReducedStillPictureHeader {
		showExistingFrame = false
		frameType = KEY_FRAME
		frameIsIntra = true

		showFrame = true
		showableFrame = false
	} else {
		showExistingFrame := p.f(1) != 0

		if showExistingFrame {
			frameToShowMapIdx := p.f(3)

			if sequenceHeader.DecoderModelInfoPresent && !sequenceHeader.TimingInfo.EqualPictureInterval {
				temporalPointInfo = p.TemporalPointInfo()
			}

			refreshImageFlags = 0

			if sequenceHeader.FrameIdNumbersPresent {
				displayFrameId = p.f(idLen)
			}

			frameType := refFrameType[frameToShowMapIdx]

			// KEY_FRAME
			if frameType == KEY_FRAME {
				refreshImageFlags = allFrames
			}

			if sequenceHeader.FilmGrainParamsPresent {
				p.LoadGrainParams(frameToShowMapIdx)
			}

			return UncompressedHeader{
				ShowExistingFrame: showExistingFrame,
				TemporalPointInfo: temporalPointInfo,
				ShowableFrame:     showableFrame,
				RefreshImageFlags: refreshImageFlags,
				DisplayFrameId:    displayFrameId,
			}
		}
		frameType = p.f(2)

		frameIsIntra = (frameType == 2 || frameType == 0)

		showFrame = p.f(1) != 0

		if showFrame && sequenceHeader.DecoderModelInfoPresent && !sequenceHeader.TimingInfo.EqualPictureInterval {
			temporalPointInfo = p.TemporalPointInfo()
		}

		if showFrame {
			showableFrame = frameType != 0
		} else {
			showableFrame = p.f(1) != 0
		}

		if frameType == 3 || frameType == 0 && showFrame {
			errorResilientMode = true
		} else {
			errorResilientMode = p.f(1) != 0
		}
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

	allowHighPrecisionMv := false
	useRefFrameMvs := false
	allowIntrabc := false
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

	if frameIsIntra {
		p.frameSize()
		p.renderSize()

		if allowScreenContentTools && UpscaledWidth == FrameWidth {
			allowIntrabc = p.f(1) != 0
		}
	} else {

		var frameRefsShortSignaling bool
		if !sequenceHeader.EnableOrderHint {
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
			if !sequenceHeader.EnableOrderHint {
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

	if primaryRefFrame == PRIMARY_REF_NONE {

		p.initNonCoeffCdfs()
		p.setupPastIndpendence()
	} else {
		p.loadCdfs(ref_frame_idx[primaryRefFrame])
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

	if primaryRefFrame == PRIMARY_REF_NONE {

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
	if frameIsIntra || errorResilientMode || !sequenceHeader.EnableWarpedMotion {
		allowWarpedMotion = false
	} else {
		allowWarpedMotion = p.f(1) != 0
	}

	reducedTxSet := p.f(1) != 0

	p.globalMotionParams()
	p.filmGrainParams()

	return UncompressedHeader{
		ShowExistingFrame: showExistingFrame,
		TemporalPointInfo: temporalPointInfo,
	}
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

// choose_operating_point()
func (p *Parser) chooseOperatingPoint() int {
	// TODO: implement
	// can be chose by implementation!
	return 0
}
