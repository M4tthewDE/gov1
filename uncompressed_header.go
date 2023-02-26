package main

const NUM_REF_FRAMES = 8
const REFS_PER_FRAME = 7
const KEY_FRAME = 0
const PRIMARY_REF_NONE = 7
const MAX_SEGMENTS = 8
const SEG_LVL_MAX = 8
const SEG_LVL_SKIP = 6
const SEG_LVL_REF_FRAME = 5
const SEG_LVL_GLOBALMV = 7

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

const MAX_LOOP_FILTER = 63

const IDENTITY = 0
const TRANSLATION = 1
const ROTZOOM = 2
const AFFINE = 3
const WARPEDMODEL_PREC_BITS = 16
const WARPEDMODEL_NONDIAGAFFINE_CLAMP = 1 << 13
const WARPEDMODEL_TRANS_CLAMP = 1 << 23

const GM_TRANS_ONLY_PREC_BITS = 3
const GM_TRANS_PREC_BITS = 6
const GM_ABS_TRANS_ONLY_BITS = 9
const GM_ABS_ALPHA_BITS = 12
const GM_ABS_TRANS_BITS = 12
const GM_ALPHA_PREC_BITS = 15

var Segmentation_Feature_Bits = []int{8, 6, 6, 6, 6, 3, 0, 0}
var Segmentation_Feature_Signed = []int{1, 1, 1, 1, 1, 0, 0, 0}
var Segmentation_Feature_Max = []int{255, MAX_LOOP_FILTER, MAX_LOOP_FILTER, MAX_LOOP_FILTER, MAX_LOOP_FILTER, 7, 0, 0}

const MI_SIZE_LOG2 = 2

var Mi_Width_Log2 = []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 0, 2, 1, 3, 2, 4}
var Mi_Height_Log2 = []int{0, 1, 0, 1, 2, 1, 2, 3, 2, 3, 4, 3, 4, 5, 4, 5, 2, 0, 3, 1, 4, 2}

type UncompressedHeader struct {
	ShowExistingFrame          bool
	ShowableFrame              bool
	RefreshImageFlags          int
	DisplayFrameId             int
	FrameIdNumbersPresent      bool
	AllowHighPrecisionMv       bool
	AllowIntraBc               bool
	LastFrameIdx               int
	GoldFrameIdx               int
	IsMotionModeSwitchable     bool
	DisableFrameEndUpdateCdf   bool
	AllLossless                bool
	AllowWarpedMotion          bool
	ReducedTxSet               bool
	FrameIsIntra               bool
	ReferenceSelect            bool
	CodedLossless              bool
	TxMode                     int
	DeltaQPresent              bool
	DeltaQRes                  int
	DeltaLfPresent             bool
	DeltaLfRes                 int
	DeltaLfMulti               int
	BaseQIdx                   int
	EnableOrderHint            bool
	OrderHintBits              int
	UseSuperRes                bool
	SuperResDenom              int
	FrameWidth                 int
	FrameHeight                int
	UpscaledWidth              int
	UpscaledHeight             int
	FrameSizeOverrideFlag      bool
	DeltaQYDc                  int
	DeltaQUAc                  int
	DeltaQUDc                  int
	DeltaQVAc                  int
	DeltaQVDc                  int
	UsingQMatrix               bool
	Qmy                        int
	Qmu                        int
	Qmv                        int
	RenderWidth                int
	RenderHeight               int
	InterpolationFilter        int
	PrevFrameId                int
	CurrentFrameId             int
	RefFrameId                 []int
	RefValid                   []int
	TileInfo                   TileInfo
	FramePresentationTime      int
	PrimaryRefFrame            int
	SegIdPreSkip               int
	LastActiveSegId            int
	SegmentationEnabled        int
	SegmentationTemporalUpdate int
	SegmentationUpdateMap      int
	SegmentationUpdateData     int
	LosslessArray              []bool
	GmParams                   [][]int
	ForceIntegerMv             bool
	AllowScreenContentTools    int
	RefOrderHint               []int
	ref_frame_idx              []int
	OrderHint                  int
	SkipModeFrame              []int
	SkipModePresent            int
}

func NewUncompressedHeader(p *Parser) UncompressedHeader {
	u := UncompressedHeader{}

	u.Build(p)
	return u
}

func (u *UncompressedHeader) Build(p *Parser) {
	var idLen int
	if p.sequenceHeader.FrameIdNumbersPresent {
		idLen = p.sequenceHeader.AdditionalFrameIdLengthMinusOne +
			p.sequenceHeader.DeltaFrameIdLengthMinusTwo + 3
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
	if p.sequenceHeader.ReducedStillPictureHeader {
		u.ShowExistingFrame = false
		frameType = KEY_FRAME
		u.FrameIsIntra = true

		showFrame = true
		u.ShowableFrame = false
	} else {
		showExistingFrame := p.f(1) != 0

		if showExistingFrame {
			frameToShowMapIdx := p.f(3)

			if p.sequenceHeader.DecoderModelInfoPresent && !p.sequenceHeader.TimingInfo.EqualPictureInterval {
				u.TemporalPointInfo(p)
			}

			u.RefreshImageFlags = 0
			if p.sequenceHeader.FrameIdNumbersPresent {
				u.DisplayFrameId = p.f(idLen)
			}

			frameType := refFrameType[frameToShowMapIdx]

			// KEY_FRAME
			if frameType == KEY_FRAME {
				u.RefreshImageFlags = allFrames
			}

			if p.sequenceHeader.FilmGrainParamsPresent {
				p.LoadGrainParams(frameToShowMapIdx)
			}
		}

		frameType = p.f(2)

		u.FrameIsIntra = (frameType == 2 || frameType == 0)

		showFrame = p.f(1) != 0

		if showFrame && p.sequenceHeader.DecoderModelInfoPresent && !p.sequenceHeader.TimingInfo.EqualPictureInterval {
			u.TemporalPointInfo(p)
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
			refValid = SliceAssign(refValid, i, 0)
			refOrderHint = SliceAssign(refOrderHint, i, 0)
		}

		for i := 0; i < REFS_PER_FRAME; i++ {
			orderHints = SliceAssign(orderHints, LAST_FRAME+1, 0)
		}
	}

	disableCdfUpdate := p.f(1) != 0

	if p.sequenceHeader.SeqForceScreenContentTools == 2 {
		u.AllowScreenContentTools = p.f(1)
	} else {
		u.AllowScreenContentTools = 1
	}

	if Bool(u.AllowScreenContentTools) {
		if p.sequenceHeader.SeqForceIntegerMv == 2 {
			u.ForceIntegerMv = p.f(1) != 0
		} else {
			u.ForceIntegerMv = true
		}
	} else {
		u.ForceIntegerMv = false
	}

	if u.FrameIsIntra {
		u.ForceIntegerMv = true
	}

	if p.sequenceHeader.FrameIdNumbersPresent {
		u.PrevFrameId = u.CurrentFrameId
		u.CurrentFrameId = p.f(idLen)
		u.markRefFrames(idLen, p)
	} else {
		u.CurrentFrameId = 0
	}

	if frameType == SWITCH_FRAME {
		u.FrameSizeOverrideFlag = true
	} else if p.sequenceHeader.ReducedStillPictureHeader {
		u.FrameSizeOverrideFlag = false

	} else {
		u.FrameSizeOverrideFlag = p.f(1) != 0
	}

	orderHint := p.f(p.sequenceHeader.OrderHintBits)
	u.OrderHint = orderHint

	if u.FrameIsIntra || errorResilientMode {
		u.PrimaryRefFrame = PRIMARY_REF_NONE
	} else {
		u.PrimaryRefFrame = p.f(3)
	}

	if p.sequenceHeader.DecoderModelInfoPresent {
		bufferRemovalTimePresent := p.f(1) != 0

		if bufferRemovalTimePresent {
			for opNum := 0; opNum <= p.sequenceHeader.OperatingPointsCountMinusOne; opNum++ {
				if p.sequenceHeader.DecoderModelPresentForThisOp[opNum] {
					opPtIdc := p.sequenceHeader.OperatingPointIdc[opNum]
					inTemporalLayer := ((opPtIdc >> p.header.ExtensionHeader.TemporalID) & 1) != 0
					inSpatialLayer := ((opPtIdc >> p.header.ExtensionHeader.SpatialID) & 1) != 0

					if opPtIdc == 0 || (inTemporalLayer && inSpatialLayer) {
						n := p.sequenceHeader.DecoderModelInfo.BufferRemovalTimeLengthMinusOne + 1
						bufferRemovalTime = SliceAssign(bufferRemovalTime, opNum, p.f(n))
					}
				}
			}
		}
	}

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
		if errorResilientMode && p.sequenceHeader.EnableOrderHint {
			for i := 0; i < NUM_REF_FRAMES; i++ {
				ref_order_hint = SliceAssign(ref_order_hint, i, p.f(p.sequenceHeader.OrderHintBits))

				if ref_order_hint[i] != u.RefOrderHint[i] {
					RefValid = SliceAssign(RefValid, i, 0)
				}
			}
		}
	}

	expectedFrameId := []int{}

	if u.FrameIsIntra {
		u.frameSize(p)
		u.renderSize(p)

		if Bool(u.AllowScreenContentTools) && u.UpscaledWidth == u.FrameWidth {
			u.AllowIntraBc = p.f(1) != 0
		}
	} else {

		var frameRefsShortSignaling bool
		if !p.sequenceHeader.EnableOrderHint {
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
				u.ref_frame_idx[i] = p.f(3)
				u.ref_frame_idx = SliceAssign(u.ref_frame_idx, i, p.f(3))
			}

			if p.sequenceHeader.FrameIdNumbersPresent {
				n := p.sequenceHeader.DeltaFrameIdLengthMinusTwo + 2
				deltaFrameIdMinusOne := p.f(n)
				DeltaFrameId := deltaFrameIdMinusOne + 1
				expectedFrameId = SliceAssign(expectedFrameId, i, (u.CurrentFrameId+(1<<idLen)-DeltaFrameId)%(1<<idLen))
			}
		}

		if u.FrameSizeOverrideFlag && !errorResilientMode {
			p.frameSizeWithRefs()
		} else {
			u.frameSize(p)
			u.renderSize(p)

		}
		if u.ForceIntegerMv {
			u.AllowHighPrecisionMv = false
		} else {

			u.AllowHighPrecisionMv = p.f(1) != 0
		}

		u.readInterpolationFilter(p)
		u.IsMotionModeSwitchable = p.f(1) != 0

		if errorResilientMode || !p.sequenceHeader.EnableRefFrameMvs {
			useRefFrameMvs = false

		} else {
			useRefFrameMvs = p.f(1) != 0
		}

		OrderHints := []int{}
		RefFrameSignBias := []bool{}

		for i := 0; i < REFS_PER_FRAME; i++ {
			refFrame := LAST_FRAME + 1
			hint := u.RefOrderHint[u.ref_frame_idx[i]]
			OrderHints = SliceAssign(OrderHints, refFrame, hint)
			if !p.sequenceHeader.EnableOrderHint {
				RefFrameSignBias = SliceAssign(RefFrameSignBias, refFrame, false)
			} else {
				RefFrameSignBias = SliceAssign(RefFrameSignBias, refFrame, u.getRelativeDist(hint, u.OrderHint, p) > 0)
			}
		}
	}

	if p.sequenceHeader.ReducedStillPictureHeader || disableCdfUpdate {
		u.DisableFrameEndUpdateCdf = true
	} else {
		u.DisableFrameEndUpdateCdf = p.f(1) != 0
	}

	if u.PrimaryRefFrame == PRIMARY_REF_NONE {

		p.initNonCoeffCdfs()
		p.setupPastIndependence()
	} else {
		p.loadCdfs(u.ref_frame_idx[u.PrimaryRefFrame])
		p.loadPrevious()
	}

	if useRefFrameMvs {
		p.motionFieldEstimation()
	}

	u.TileInfo = NewTileInfo(p, p.sequenceHeader)
	u.quantizationParams(p)
	u.segmentationParams(p)
	u.deltaQParams(p)
	u.deltaLfParams(p)

	if u.PrimaryRefFrame == PRIMARY_REF_NONE {

		p.initCoeffCdfs()
	} else {
		p.loadPreviousSegementIds()
	}

	u.CodedLossless = true

	SegQMLevel := [][]int{}

	for segmentId := 0; segmentId < MAX_SEGMENTS; segmentId++ {
		qIndex := p.getQIndex(1, segmentId)
		u.LosslessArray = SliceAssign(u.LosslessArray, segmentId, qIndex == 0 && u.DeltaQYDc == 0 && u.DeltaQUAc == 0 && u.DeltaQUDc == 0 && u.DeltaQVAc == 0 && u.DeltaQVDc == 0)

		if !u.LosslessArray[segmentId] {
			u.CodedLossless = false
		}

		if u.UsingQMatrix {
			if u.LosslessArray[segmentId] {
				SliceAssignNested(SegQMLevel, 0, segmentId, 15)
				SliceAssignNested(SegQMLevel, 1, segmentId, 15)
				SliceAssignNested(SegQMLevel, 2, segmentId, 15)
			} else {
				SliceAssignNested(SegQMLevel, 0, segmentId, u.Qmy)
				SliceAssignNested(SegQMLevel, 1, segmentId, u.Qmy)
				SliceAssignNested(SegQMLevel, 2, segmentId, u.Qmy)

			}
		}
	}

	u.AllLossless = u.CodedLossless && (u.FrameWidth == u.UpscaledWidth)

	p.loopFilterParams()
	p.cdefParams()
	p.lrParams()
	u.readTxMode(p)
	u.frameReferenceMode(p)
	u.skipModeParams(p)

	if u.FrameIsIntra || errorResilientMode || !p.sequenceHeader.EnableWarpedMotion {
		u.AllowWarpedMotion = false
	} else {
		u.AllowWarpedMotion = p.f(1) != 0
	}

	u.ReducedTxSet = p.f(1) != 0

	u.globalMotionParams(p)
	p.filmGrainParams()
}

// mark_ref_frames( idLen)
func (u *UncompressedHeader) markRefFrames(idLen int, p *Parser) {
	diffLen := p.sequenceHeader.DeltaFrameIdLengthMinusTwo + 2

	for i := 0; i < NUM_REF_FRAMES; i++ {
		if u.CurrentFrameId > (1 << diffLen) {
			if u.RefFrameId[i] > u.CurrentFrameId ||
				u.RefFrameId[i] < (u.CurrentFrameId-(1<<diffLen)) {
				u.RefValid = SliceAssign(u.RefValid, i, 0)
			}
		} else {
			if u.RefFrameId[i] > u.CurrentFrameId && u.RefFrameId[i] < ((1<<idLen)+u.CurrentFrameId-(1<<diffLen)) {
				u.RefValid = SliceAssign(u.RefValid, i, 0)
			}
		}
	}
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
		n := p.sequenceHeader.FrameWidthBitsMinusOne + 1
		frameWidthMinusOne := p.f(n)

		n = p.sequenceHeader.FrameHeightBitsMinusOne + 1
		frameHeightMinusOne := p.f(n)

		u.FrameWidth = frameWidthMinusOne + 1
		u.FrameHeight = frameHeightMinusOne + 1
	} else {
		u.FrameWidth = p.sequenceHeader.MaxFrameWidthMinusOne + 1
		u.FrameHeight = p.sequenceHeader.MaxFrameHeightMinusOne + 1
	}

	u.superResParams(p)
	u.computeImageSize(p)
}

// superres_params()
func (u *UncompressedHeader) superResParams(p *Parser) {
	if p.sequenceHeader.EnableSuperRes {
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
func (u *UncompressedHeader) computeImageSize(p *Parser) {
	p.MiCols = 2 * ((u.FrameWidth + 7) >> 3)
	p.MiRows = 2 * ((u.FrameHeight + 7) >> 3)
}

// render_size()
func (u *UncompressedHeader) renderSize(p *Parser) {
	renderAndFramSizeDifferent := p.f(1) != 0

	if renderAndFramSizeDifferent {
		renderWidthMinusOne := p.f(16)
		renderHeightMinusOne := p.f(16)

		u.RenderWidth = renderWidthMinusOne + 1
		u.RenderHeight = renderHeightMinusOne + 1
	} else {
		u.RenderWidth = u.UpscaledWidth
		u.RenderHeight = u.UpscaledHeight
	}
}

func (u *UncompressedHeader) readInterpolationFilter(p *Parser) {
	isFilterSwitchable := p.f(1) != 0

	if isFilterSwitchable {
		u.InterpolationFilter = SWITCHABLE
	} else {
		u.InterpolationFilter = p.f(2)
	}
}

// get_relative_dist()
func (u *UncompressedHeader) getRelativeDist(a int, b int, p *Parser) int {
	if !u.EnableOrderHint {
		return 0
	}

	diff := a - b
	m := 1 << (p.sequenceHeader.OrderHintBits - 1)
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
	if p.sequenceHeader.ColorConfig.NumPlanes > 1 {
		if p.sequenceHeader.ColorConfig.SeparateUvDeltaQ {
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

		if !p.sequenceHeader.ColorConfig.SeparateUvDeltaQ {
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

// segmentation_params
func (u *UncompressedHeader) segmentationParams(p *Parser) {
	u.SegmentationEnabled = p.f(1)
	if u.SegmentationEnabled == 1 {
		if u.PrimaryRefFrame == PRIMARY_REF_NONE {
			u.SegmentationUpdateMap = 1
			u.SegmentationTemporalUpdate = 0
			u.SegmentationUpdateData = 1

		} else {
			u.SegmentationUpdateMap = p.f(1)
			if u.SegmentationUpdateMap == 1 {
				u.SegmentationTemporalUpdate = p.f(1)
			}
			u.SegmentationUpdateData = p.f(1)
		}

		if u.SegmentationUpdateData == 1 {
			for i := 0; i < MAX_SEGMENTS; i++ {
				for j := 0; j < SEG_LVL_MAX; i++ {
					featureValue := 0
					featureEnabled := p.f(1)
					p.FeatureEnabled[i][j] = featureEnabled
					clippedValue := 0

					if featureEnabled == 1 {
						bitsToRead := Segmentation_Feature_Bits[j]
						limit := Segmentation_Feature_Max[j]
						if Segmentation_Feature_Signed[j] == 1 {
							featureValue = p.su(1 + bitsToRead)
							clippedValue = Clip3(-limit, limit, featureValue)
						} else {
							featureValue = p.su(bitsToRead)
							clippedValue = Clip3(0, limit, featureValue)

						}
					}
					p.FeatureData[i][j] = clippedValue
				}
			}
		}
	} else {
		for i := 0; i < MAX_SEGMENTS; i++ {
			for j := 0; j < SEG_LVL_MAX; i++ {
				p.FeatureEnabled[i][j] = 0
				p.FeatureData[i][j] = 0
			}

		}
	}
	u.SegIdPreSkip = 0
	u.LastActiveSegId = 0

	for i := 0; i < MAX_SEGMENTS; i++ {
		for j := 0; j < SEG_LVL_MAX; i++ {
			if p.FeatureEnabled[i][j] == 1 {
				u.LastActiveSegId = i

				if j >= SEG_LVL_REF_FRAME {
					u.SegIdPreSkip = 1
				}
			}
		}

	}

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

// skip_mode_params()
func (u *UncompressedHeader) skipModeParams(p *Parser) {
	var skipModeAllowed = 0
	var forwardHint int
	var backwardHint int
	if u.FrameIsIntra || u.ReferenceSelect || u.EnableOrderHint {
		skipModeAllowed = 0
	} else {
		forwardIdx := -1
		backwardIdx := -1

		for i := 0; i < REFS_PER_FRAME; i++ {
			refHint := u.RefOrderHint[u.ref_frame_idx[i]]

			if u.getRelativeDist(refHint, u.OrderHint, p) < 0 {
				if forwardIdx < 0 || u.getRelativeDist(refHint, forwardHint, p) > 0 {
					forwardIdx = i
					forwardHint = refHint
				}
			} else if u.getRelativeDist(refHint, u.OrderHint, p) > 0 {
				if backwardIdx < 0 || u.getRelativeDist(refHint, backwardHint, p) > 0 {
					backwardIdx = i
					backwardHint = refHint
				}
			}
		}

		if forwardIdx < 0 {
			skipModeAllowed = 0
		} else if backwardIdx >= 0 {
			skipModeAllowed = 1
			u.SkipModeFrame[0] = LAST_FRAME + Min(forwardIdx, backwardIdx)
			u.SkipModeFrame[1] = LAST_FRAME + Max(forwardIdx, backwardIdx)
		} else {
			secondForwardIdx := -1
			var secondForwardHint int
			for i := 0; i < REFS_PER_FRAME; i++ {
				refHint := u.RefOrderHint[u.ref_frame_idx[i]]
				if u.getRelativeDist(refHint, forwardHint, p) < 0 {
					if secondForwardIdx < 0 || u.getRelativeDist(refHint, secondForwardHint, p) > 0 {
						secondForwardIdx = i
						secondForwardHint = refHint
					}
				}
			}

			if secondForwardIdx < 0 {
				skipModeAllowed = 0
			} else {
				skipModeAllowed = 1
				u.SkipModeFrame[0] = LAST_FRAME + Min(forwardIdx, secondForwardIdx)
				u.SkipModeFrame[1] = LAST_FRAME + Max(forwardIdx, secondForwardIdx)
			}
		}
	}

	if Bool(skipModeAllowed) {
		u.SkipModePresent = p.f(1)
	} else {
		u.SkipModePresent = 0
	}
}

func (u *UncompressedHeader) globalMotionParams(p *Parser) {
	for ref := LAST_FRAME; ref <= ALTREF_FRAME; ref++ {
		p.GmType[ref] = IDENTITY
		for i := 0; i < 6; i++ {
			if i%3 == 2 {
				u.GmParams[ref][i] = 1 << WARPEDMODEL_PREC_BITS

			} else {
				u.GmParams[ref][i] = 0

			}
		}
	}

	if u.FrameIsIntra {
		return
	}

	for ref := LAST_FRAME; ref <= ALTREF_FRAME; ref++ {
		var typ int
		isGlobal := p.f(1)
		if Bool(isGlobal) {
			isRotZoom := p.f(1)
			if Bool(isRotZoom) {
				typ = ROTZOOM
			} else {

				isTranslation := p.f(1)
				if Bool(isTranslation) {
					typ = TRANSLATION
				} else {
					typ = AFFINE
				}
			}
		} else {
			typ = IDENTITY
		}
		p.GmType[ref] = typ
	}

}

func (u *UncompressedHeader) readGlobalParam(typ int, ref int, idx int, p *Parser) {
	absBits := GM_ABS_ALPHA_BITS
	precBits := GM_ALPHA_PREC_BITS

	if idx < 2 {
		if typ == TRANSLATION {
			absBits = GM_ABS_TRANS_ONLY_BITS - Int(!u.AllowHighPrecisionMv)
			precBits = GM_TRANS_ONLY_PREC_BITS - Int(!u.AllowHighPrecisionMv)
		} else {
			absBits = GM_ABS_TRANS_BITS
			precBits = GM_TRANS_PREC_BITS
		}
	}

	precDiff := WARPEDMODEL_PREC_BITS - precBits

	var round int
	if idx%3 == 2 {
		round = 1 << WARPEDMODEL_PREC_BITS
	} else {
		round = 0
	}

	var sub int
	if idx%3 == 2 {
		sub = 1 << precBits
	} else {
		sub = 0
	}

	mx := 1 << absBits
	r := (p.PrevGmParams[ref][idx] >> precDiff) - sub
	u.GmParams[ref][idx] = (u.decodeSignedSubexpWithRef(mx, r, p) << precDiff) + round
}

func (u *UncompressedHeader) decodeSignedSubexpWithRef(mx int, r int, p *Parser) int {
	v := u.decodeSubexp(mx, p)
	if (r << 1) <= mx {
		return InverseRecenter(r, v)
	} else {
		return mx - 1 - InverseRecenter(mx-1-r, v)
	}
}

func (u *UncompressedHeader) decodeSubexp(numSyms int, p *Parser) int {
	i := 0
	mk := 0
	k := 3

	var b2 int
	for {
		if Bool(i) {
			b2 = k + i - 1
		} else {
			b2 = k
		}

		a := 1 << b2
		if numSyms <= mk+3*a {
			subexpFinalBits := p.ns(numSyms - mk)
			return subexpFinalBits + mk
		} else {
			subexmpMoreBits := p.f(1)
			if Bool(subexmpMoreBits) {
				i++
				mk += a
			} else {
				subexpBits := p.f(b2)
				return subexpBits + mk
			}
		}
	}
}

func (p *Parser) filmGrainParams() {
	panic("not implemented")
}

// decode_frame_wrapup()
func (p *Parser) DecodeFrameWrapup() {
	panic("not implemented")
}

// temporal_point_info()
func (u *UncompressedHeader) TemporalPointInfo(p *Parser) {
	n := p.sequenceHeader.DecoderModelInfo.FramePresentationTimeLengthMinusOne + 1
	u.FramePresentationTime = p.f(n)
}

// load_grain_params( idx )
func (p *Parser) LoadGrainParams(idx int) {
	panic("not implemented")
}

// choose_operating_point()
func (p *Parser) chooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
