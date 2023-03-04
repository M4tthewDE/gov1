package uncompressedheader

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/tileinfo"
	"github.com/m4tthewde/gov1/internal/util"
)

const ONLY_4X4 = 0

const SUPERRES_DENOM_BITS = 3
const SUPERRES_DENOM_MIN = 9
const SWITCH_FRAME = 3

type UncompressedHeader struct {
	State State

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
	TileInfo                   tileinfo.TileInfo
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
	RefFrameIdx                []int
	OrderHint                  int
	OrderHints                 []int
	SkipModeFrame              []int
	SkipModePresent            int
}

func NewUncompressedHeader(b *bitstream.BitStream, inputState State) UncompressedHeader {
	u := UncompressedHeader{
		State: inputState,
	}

	u.build(b)
	return u
}

func (u *UncompressedHeader) build(b *bitstream.BitStream) {
	var idLen int
	if u.State.SequenceHeader.FrameIdNumbersPresent {
		idLen = u.State.SequenceHeader.AdditionalFrameIdLengthMinusOne +
			u.State.SequenceHeader.DeltaFrameIdLengthMinusTwo + 3
	}

	var frameType int
	var showFrame bool

	var errorResilientMode bool

	refFrameType := []int{}
	var refValid [shared.NUM_REF_FRAMES]int
	var refOrderHint [shared.NUM_REF_FRAMES]int
	var orderHints [shared.REFS_PER_FRAME]int

	bufferRemovalTime := []int{}

	allFrames := ((1 << shared.NUM_REF_FRAMES) - 1)
	if u.State.SequenceHeader.ReducedStillPictureHeader {
		u.ShowExistingFrame = false
		frameType = shared.KEY_FRAME
		u.FrameIsIntra = true

		showFrame = true
		u.ShowableFrame = false
	} else {
		showExistingFrame := util.Bool(b.F(1))

		if showExistingFrame {
			frameToShowMapIdx := b.F(3)

			if u.State.SequenceHeader.DecoderModelInfoPresent && !u.State.SequenceHeader.TimingInfo.EqualPictureInterval {
				u.TemporalPointInfo(b)
			}

			u.RefreshImageFlags = 0
			if u.State.SequenceHeader.FrameIdNumbersPresent {
				u.DisplayFrameId = b.F(idLen)
			}

			frameType := refFrameType[frameToShowMapIdx]

			if frameType == shared.KEY_FRAME {
				u.RefreshImageFlags = allFrames
			}

			if u.State.SequenceHeader.FilmGrainParamsPresent {
				u.loadGrainParams(frameToShowMapIdx)
			}

			return
		}

		frameType = b.F(2)

		u.FrameIsIntra = (frameType == 2 || frameType == 0)

		showFrame = util.Bool(b.F(1))

		if showFrame && u.State.SequenceHeader.DecoderModelInfoPresent && !u.State.SequenceHeader.TimingInfo.EqualPictureInterval {
			u.TemporalPointInfo(b)
		}

		if showFrame {
			u.ShowableFrame = frameType != 0
		} else {
			u.ShowableFrame = util.Bool(b.F(1))
		}

		if frameType == 3 || frameType == 0 && showFrame {
			errorResilientMode = true
		} else {
			errorResilientMode = util.Bool(b.F(1))
		}
	}

	if frameType == 0 && showFrame {
		for i := 0; i < shared.NUM_REF_FRAMES; i++ {
			refValid[i] = 0
			refOrderHint[i] = 0
		}

		for i := 0; i < shared.REFS_PER_FRAME; i++ {
			orderHints[shared.LAST_FRAME+1] = 0
		}
	}

	disableCdfUpdate := util.Bool(b.F(1))

	if u.State.SequenceHeader.SeqForceScreenContentTools == 2 {
		u.AllowScreenContentTools = b.F(1)
	} else {
		u.AllowScreenContentTools = 1
	}

	if util.Bool(u.AllowScreenContentTools) {
		if u.State.SequenceHeader.SeqForceIntegerMv == 2 {
			u.ForceIntegerMv = util.Bool(b.F(1))
		} else {
			u.ForceIntegerMv = true
		}
	} else {
		u.ForceIntegerMv = false
	}

	if u.FrameIsIntra {
		u.ForceIntegerMv = true
	}

	if u.State.SequenceHeader.FrameIdNumbersPresent {
		u.PrevFrameId = u.CurrentFrameId
		u.CurrentFrameId = b.F(idLen)
		u.markRefFrames(idLen)
	} else {
		u.CurrentFrameId = 0
	}

	if frameType == SWITCH_FRAME {
		u.FrameSizeOverrideFlag = true
	} else if u.State.SequenceHeader.ReducedStillPictureHeader {
		u.FrameSizeOverrideFlag = false

	} else {
		u.FrameSizeOverrideFlag = util.Bool(b.F(1))
	}

	orderHint := b.F(u.State.SequenceHeader.OrderHintBits)
	u.OrderHint = orderHint

	if u.FrameIsIntra || errorResilientMode {
		u.PrimaryRefFrame = shared.PRIMARY_REF_NONE
	} else {
		u.PrimaryRefFrame = b.F(3)
	}

	if u.State.SequenceHeader.DecoderModelInfoPresent {
		bufferRemovalTimePresent := util.Bool(b.F(1))

		if bufferRemovalTimePresent {
			for opNum := 0; opNum <= u.State.SequenceHeader.OperatingPointsCountMinusOne; opNum++ {
				if u.State.SequenceHeader.DecoderModelPresentForThisOp[opNum] {
					opPtIdc := u.State.SequenceHeader.OperatingPointIdc[opNum]
					inTemporalLayer := ((opPtIdc >> u.State.Header.ExtensionHeader.TemporalID) & 1) != 0
					inSpatialLayer := ((opPtIdc >> u.State.Header.ExtensionHeader.SpatialID) & 1) != 0

					if opPtIdc == 0 || (inTemporalLayer && inSpatialLayer) {
						n := u.State.SequenceHeader.DecoderModelInfo.BufferRemovalTimeLengthMinusOne + 1
						bufferRemovalTime[opNum] = b.F(n)
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
		refreshFrameFlags = b.F(8)
	}

	if !u.FrameIsIntra || refreshFrameFlags != allFrames {
		if errorResilientMode && u.State.SequenceHeader.EnableOrderHint {
			for i := 0; i < shared.NUM_REF_FRAMES; i++ {
				ref_order_hint[i] = b.F(u.State.SequenceHeader.OrderHintBits)

				if ref_order_hint[i] != u.RefOrderHint[i] {
					RefValid[i] = 0
				}
			}
		}
	}

	expectedFrameId := []int{}

	if u.FrameIsIntra {
		u.frameSize(b)
		u.renderSize(b)

		if util.Bool(u.AllowScreenContentTools) && u.UpscaledWidth == u.FrameWidth {
			u.AllowIntraBc = util.Bool(b.F(1))
		}
	} else {

		var frameRefsShortSignaling bool
		if !u.State.SequenceHeader.EnableOrderHint {
			frameRefsShortSignaling = false
		} else {
			frameRefsShortSignaling = util.Bool(b.F(1))
			if frameRefsShortSignaling {
				u.LastFrameIdx = b.F(3)
				u.GoldFrameIdx = b.F(3)
				u.setFrameRefs()
			}
		}

		for i := 0; i < shared.REFS_PER_FRAME; i++ {
			if !frameRefsShortSignaling {
				u.RefFrameIdx[i] = b.F(3)
				u.RefFrameIdx[i] = b.F(3)
			}

			if u.State.SequenceHeader.FrameIdNumbersPresent {
				n := u.State.SequenceHeader.DeltaFrameIdLengthMinusTwo + 2
				deltaFrameIdMinusOne := b.F(n)
				DeltaFrameId := deltaFrameIdMinusOne + 1
				expectedFrameId[i] = (u.CurrentFrameId + (1 << idLen) - DeltaFrameId) % (1 << idLen)
			}
		}

		if u.FrameSizeOverrideFlag && !errorResilientMode {
			u.frameSizeWithRefs()
		} else {
			u.frameSize(b)
			u.renderSize(b)

		}
		if u.ForceIntegerMv {
			u.AllowHighPrecisionMv = false
		} else {

			u.AllowHighPrecisionMv = util.Bool(b.F(1))
		}

		u.readInterpolationFilter(b)
		u.IsMotionModeSwitchable = util.Bool(b.F(1))

		if errorResilientMode || !u.State.SequenceHeader.EnableRefFrameMvs {
			useRefFrameMvs = false

		} else {
			useRefFrameMvs = util.Bool(b.F(1))
		}

		RefFrameSignBias := []bool{}

		for i := 0; i < shared.REFS_PER_FRAME; i++ {
			refFrame := shared.LAST_FRAME + 1
			hint := u.RefOrderHint[u.RefFrameIdx[i]]
			u.OrderHints[refFrame] = hint
			if !u.State.SequenceHeader.EnableOrderHint {
				RefFrameSignBias[refFrame] = false
			} else {
				RefFrameSignBias[refFrame] = u.GetRelativeDist(hint, u.OrderHint) > 0
			}
		}
	}

	if u.State.SequenceHeader.ReducedStillPictureHeader || disableCdfUpdate {
		u.DisableFrameEndUpdateCdf = true
	} else {
		u.DisableFrameEndUpdateCdf = util.Bool(b.F(1))
	}

	if u.PrimaryRefFrame == shared.PRIMARY_REF_NONE {
		initNonCoeffCdfs()
		u.setupPastIndependence()
	} else {
		u.loadCdfs(u.RefFrameIdx[u.PrimaryRefFrame])
		u.loadPrevious()
	}

	if useRefFrameMvs {
		u.motionFieldEstimation()
	}

	tileInfo, resultState := tileinfo.NewTileInfo(b, u.State.newTileInfoState())
	u.TileInfo = tileInfo
	u.State.update(resultState)

	u.quantizationParams(b)
	u.segmentationParams(b)
	u.deltaQParams(b)
	u.deltaLfParams(b)

	if u.PrimaryRefFrame == shared.PRIMARY_REF_NONE {

		u.initCoeffCdfs()
	} else {
		u.loadPreviousSegementIds()
	}

	u.CodedLossless = true

	SegQMLevel := [][]int{}

	for segmentId := 0; segmentId < shared.MAX_SEGMENTS; segmentId++ {
		qIndex := u.getQIndex(1, segmentId)
		u.LosslessArray[segmentId] = qIndex == 0 && u.DeltaQYDc == 0 && u.DeltaQUAc == 0 && u.DeltaQUDc == 0 && u.DeltaQVAc == 0 && u.DeltaQVDc == 0

		if !u.LosslessArray[segmentId] {
			u.CodedLossless = false
		}

		if u.UsingQMatrix {
			if u.LosslessArray[segmentId] {
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

	u.loopFilterParams()
	u.cdefParams()
	u.lrParams()
	u.readTxMode(b)
	u.frameReferenceMode(b)
	u.skipModeParams(b)

	if u.FrameIsIntra || errorResilientMode || !u.State.SequenceHeader.EnableWarpedMotion {
		u.AllowWarpedMotion = false
	} else {
		u.AllowWarpedMotion = util.Bool(b.F(1))
	}

	u.ReducedTxSet = util.Bool(b.F(1))

	u.globalMotionParams(b)
	u.filmGrainParams()
}

// mark_ref_frames( idLen)
func (u *UncompressedHeader) markRefFrames(idLen int) {
	diffLen := u.State.SequenceHeader.DeltaFrameIdLengthMinusTwo + 2

	for i := 0; i < shared.NUM_REF_FRAMES; i++ {
		if u.CurrentFrameId > (1 << diffLen) {
			if u.RefFrameId[i] > u.CurrentFrameId ||
				u.RefFrameId[i] < (u.CurrentFrameId-(1<<diffLen)) {
				u.RefValid[i] = 0
			}
		} else {
			if u.RefFrameId[i] > u.CurrentFrameId && u.RefFrameId[i] < ((1<<idLen)+u.CurrentFrameId-(1<<diffLen)) {
				u.RefValid[i] = 0
			}
		}
	}
}

func (u *UncompressedHeader) setFrameRefs() {
	panic("not implemented")
}

func (u *UncompressedHeader) frameSizeWithRefs() {
	panic("not implemented")
}

// frame_size()
func (u *UncompressedHeader) frameSize(b *bitstream.BitStream) {
	if u.FrameSizeOverrideFlag {
		n := u.State.SequenceHeader.FrameWidthBitsMinusOne + 1
		frameWidthMinusOne := b.F(n)

		n = u.State.SequenceHeader.FrameHeightBitsMinusOne + 1
		frameHeightMinusOne := b.F(n)

		u.FrameWidth = frameWidthMinusOne + 1
		u.FrameHeight = frameHeightMinusOne + 1
	} else {
		u.FrameWidth = u.State.SequenceHeader.MaxFrameWidthMinusOne + 1
		u.FrameHeight = u.State.SequenceHeader.MaxFrameHeightMinusOne + 1
	}

	u.superResParams(b)
	u.computeImageSize()
}

// superres_params()
func (u *UncompressedHeader) superResParams(b *bitstream.BitStream) {
	if u.State.SequenceHeader.EnableSuperRes {
		u.UseSuperRes = util.Bool(b.F(1))
	} else {
		u.UseSuperRes = false
	}

	if u.UseSuperRes {
		codedDenom := b.F(SUPERRES_DENOM_BITS)
		u.SuperResDenom = codedDenom + SUPERRES_DENOM_MIN
	} else {
		u.SuperResDenom = shared.SUPERRES_NUM
	}

	u.UpscaledWidth = u.FrameWidth
	u.FrameWidth = (u.UpscaledWidth*shared.SUPERRES_NUM + (u.SuperResDenom / 2)) / u.SuperResDenom
}

// compute_image_size()
func (u *UncompressedHeader) computeImageSize() {
	u.State.MiCols = 2 * ((u.FrameWidth + 7) >> 3)
	u.State.MiRows = 2 * ((u.FrameHeight + 7) >> 3)
}

// render_size()
func (u *UncompressedHeader) renderSize(b *bitstream.BitStream) {
	renderAndFramSizeDifferent := util.Bool(b.F(1))

	if renderAndFramSizeDifferent {
		renderWidthMinusOne := b.F(16)
		renderHeightMinusOne := b.F(16)

		u.RenderWidth = renderWidthMinusOne + 1
		u.RenderHeight = renderHeightMinusOne + 1
	} else {
		u.RenderWidth = u.UpscaledWidth
		u.RenderHeight = u.UpscaledHeight
	}
}

func (u *UncompressedHeader) readInterpolationFilter(b *bitstream.BitStream) {
	isFilterSwitchable := util.Bool(b.F(1))

	if isFilterSwitchable {
		u.InterpolationFilter = shared.SWITCHABLE
	} else {
		u.InterpolationFilter = b.F(2)
	}
}

// get_relative_dist()
func (u *UncompressedHeader) GetRelativeDist(a int, b int) int {
	if !u.EnableOrderHint {
		return 0
	}

	diff := a - b
	m := 1 << (u.State.SequenceHeader.OrderHintBits - 1)
	diff = (diff & (m - 1)) - (diff & m)

	return diff
}

// setup_past_independence()
func (u *UncompressedHeader) setupPastIndependence() {
	for i := 0; i < shared.MAX_SEGMENTS; i++ {
		for j := 0; j < shared.SEG_LVL_MAX; j++ {
			u.State.FeatureData[i][j] = 0
			u.State.FeatureEnabled[i][j] = 0
		}
	}

	u.State.PrevSegmentIds = make([][]int, u.State.MiRows)
	for i := range u.State.PrevSegmentIds {
		u.State.PrevSegmentIds[i] = make([]int, u.State.MiCols)
	}

	for row := 0; row < u.State.MiRows; row++ {
		for col := 0; col < u.State.MiCols; col++ {
			u.State.PrevSegmentIds[row][col] = 0
		}
	}

	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		u.State.GmType[ref] = shared.IDENTITY
	}

	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		for i := 0; i <= 5; i++ {
			if (i % 3) == 2 {
				u.State.PrevGmParams[ref][i] = 1 << shared.WARPEDMODEL_PREC_BITS
			} else {
				u.State.PrevGmParams[ref][i] = 0
			}
		}
	}

	u.State.LoopFilterDeltaEnabled = true
	u.State.LoopFilterRefDeltas[shared.INTRA_FRAME] = 1
	u.State.LoopFilterRefDeltas[shared.LAST_FRAME] = 0
	u.State.LoopFilterRefDeltas[shared.LAST2_FRAME] = 0
	u.State.LoopFilterRefDeltas[shared.LAST3_FRAME] = 0
	u.State.LoopFilterRefDeltas[shared.BWDREF_FRAME] = 0
	u.State.LoopFilterRefDeltas[shared.GOLDEN_FRAME] = -1
	u.State.LoopFilterRefDeltas[shared.ALTREF_FRAME] = -1
	u.State.LoopFilterRefDeltas[shared.ALTREF2_FRAME] = -1

	u.State.LoopFilterModeDeltas[0] = 0
	u.State.LoopFilterModeDeltas[1] = 0
}

func (u *UncompressedHeader) loadCdfs(a int) {
	panic("not implemented")
}

func (u *UncompressedHeader) loadPrevious() {
	panic("not implemented")
}

func (u *UncompressedHeader) motionFieldEstimation() {
	panic("not implemented")
}

func (u *UncompressedHeader) tileInfo() {
	panic("not implemented")
}

// quantization_params()
func (u *UncompressedHeader) quantizationParams(b *bitstream.BitStream) {
	u.BaseQIdx = b.F(8)

	u.DeltaQYDc = u.readDeltaQ(b)

	var diffUvDelta bool
	if u.State.SequenceHeader.ColorConfig.NumPlanes > 1 {
		if u.State.SequenceHeader.ColorConfig.SeparateUvDeltaQ {
			diffUvDelta = util.Bool(b.F(1))
		} else {
			diffUvDelta = false
		}

		u.DeltaQUDc = u.readDeltaQ(b)
		u.DeltaQUAc = u.readDeltaQ(b)

		if diffUvDelta {
			u.DeltaQVDc = u.readDeltaQ(b)
			u.DeltaQVAc = u.readDeltaQ(b)

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

	u.UsingQMatrix = util.Bool(b.F(1))
	if u.UsingQMatrix {
		u.Qmy = b.F(4)
		u.Qmu = b.F(4)

		if !u.State.SequenceHeader.ColorConfig.SeparateUvDeltaQ {
			u.Qmv = u.Qmu
		} else {
			u.Qmv = b.F(4)
		}
	}
}

// read_delta_q()
func (u *UncompressedHeader) readDeltaQ(b *bitstream.BitStream) int {
	deltaCoded := util.Bool(b.F(1))
	if deltaCoded {
		return b.Su(1 + 6)
	} else {
		return 0
	}
}

// segmentation_params
func (u *UncompressedHeader) segmentationParams(b *bitstream.BitStream) {
	u.SegmentationEnabled = b.F(1)
	if u.SegmentationEnabled == 1 {
		if u.PrimaryRefFrame == shared.PRIMARY_REF_NONE {
			u.SegmentationUpdateMap = 1
			u.SegmentationTemporalUpdate = 0
			u.SegmentationUpdateData = 1

		} else {
			u.SegmentationUpdateMap = b.F(1)
			if u.SegmentationUpdateMap == 1 {
				u.SegmentationTemporalUpdate = b.F(1)
			}
			u.SegmentationUpdateData = b.F(1)
		}

		if u.SegmentationUpdateData == 1 {
			for i := 0; i < shared.MAX_SEGMENTS; i++ {
				for j := 0; j < shared.SEG_LVL_MAX; i++ {
					featureValue := 0
					featureEnabled := b.F(1)
					u.State.FeatureEnabled[i][j] = featureEnabled
					clippedValue := 0

					if featureEnabled == 1 {
						bitsToRead := shared.Segmentation_Feature_Bits[j]
						limit := shared.Segmentation_Feature_Max[j]
						if shared.Segmentation_Feature_Signed[j] == 1 {
							featureValue = b.Su(1 + bitsToRead)
							clippedValue = util.Clip3(-limit, limit, featureValue)
						} else {
							featureValue = b.Su(bitsToRead)
							clippedValue = util.Clip3(0, limit, featureValue)

						}
					}
					u.State.FeatureData[i][j] = clippedValue
				}
			}
		}
	} else {
		for i := 0; i < shared.MAX_SEGMENTS; i++ {
			for j := 0; j < shared.SEG_LVL_MAX; i++ {
				u.State.FeatureEnabled[i][j] = 0
				u.State.FeatureData[i][j] = 0
			}

		}
	}
	u.SegIdPreSkip = 0
	u.LastActiveSegId = 0

	for i := 0; i < shared.MAX_SEGMENTS; i++ {
		for j := 0; j < shared.SEG_LVL_MAX; i++ {
			if u.State.FeatureEnabled[i][j] == 1 {
				u.LastActiveSegId = i

				if j >= shared.SEG_LVL_REF_FRAME {
					u.SegIdPreSkip = 1
				}
			}
		}

	}

}

// delta_q_parms()
func (u *UncompressedHeader) deltaQParams(b *bitstream.BitStream) {
	u.DeltaQRes = 0
	u.DeltaQPresent = false

	if u.BaseQIdx > 0 {
		u.DeltaQPresent = util.Bool(b.F(1))
	}

	if u.DeltaQPresent {
		u.DeltaQRes = b.F(2)
	}
}

// delta_lf_params()
func (u *UncompressedHeader) deltaLfParams(b *bitstream.BitStream) {
	u.DeltaLfPresent = false
	u.DeltaLfRes = 0
	u.DeltaLfMulti = 0

	if u.DeltaQPresent {
		if !u.AllowIntraBc {
			u.DeltaLfPresent = b.F(1) != 0
		}

		if u.DeltaLfPresent {
			u.DeltaLfRes = b.F(2)
			u.DeltaLfMulti = b.F(1)
		}
	}
}

func (u *UncompressedHeader) initCoeffCdfs() {
	panic("not implemented")
}

func (u *UncompressedHeader) loadPreviousSegementIds() {
	panic("not implemented")
}

func (u *UncompressedHeader) getQIndex(a int, b int) int {
	panic("not implemented")
}

func (u *UncompressedHeader) loopFilterParams() {
	panic("not implemented")
}

func (u *UncompressedHeader) cdefParams() {
	panic("not implemented")
}

func (u *UncompressedHeader) lrParams() {
	panic("not implemented")
}

// read_tx_mode()
func (u *UncompressedHeader) readTxMode(b *bitstream.BitStream) {
	if u.CodedLossless {
		u.TxMode = ONLY_4X4
	} else {
		txModeSelect := util.Bool(b.F(1))

		if txModeSelect {
			u.TxMode = shared.TX_MODE_SELECT
		} else {
			u.TxMode = shared.TX_MODE_LARGEST
		}
	}
}

// frame_reference_mode()
func (u *UncompressedHeader) frameReferenceMode(b *bitstream.BitStream) {
	if u.FrameIsIntra {
		u.ReferenceSelect = false
	} else {
		u.ReferenceSelect = util.Bool(b.F(1))
	}

}

// skip_mode_params()
func (u *UncompressedHeader) skipModeParams(b *bitstream.BitStream) {
	var skipModeAllowed = 0
	var forwardHint int
	var backwardHint int
	if u.FrameIsIntra || u.ReferenceSelect || u.EnableOrderHint {
		skipModeAllowed = 0
	} else {
		forwardIdx := -1
		backwardIdx := -1

		for i := 0; i < shared.REFS_PER_FRAME; i++ {
			refHint := u.RefOrderHint[u.RefFrameIdx[i]]

			if u.GetRelativeDist(refHint, u.OrderHint) < 0 {
				if forwardIdx < 0 || u.GetRelativeDist(refHint, forwardHint) > 0 {
					forwardIdx = i
					forwardHint = refHint
				}
			} else if u.GetRelativeDist(refHint, u.OrderHint) > 0 {
				if backwardIdx < 0 || u.GetRelativeDist(refHint, backwardHint) > 0 {
					backwardIdx = i
					backwardHint = refHint
				}
			}
		}

		if forwardIdx < 0 {
			skipModeAllowed = 0
		} else if backwardIdx >= 0 {
			skipModeAllowed = 1
			u.SkipModeFrame[0] = shared.LAST_FRAME + util.Min(forwardIdx, backwardIdx)
			u.SkipModeFrame[1] = shared.LAST_FRAME + util.Max(forwardIdx, backwardIdx)
		} else {
			secondForwardIdx := -1
			var secondForwardHint int
			for i := 0; i < shared.REFS_PER_FRAME; i++ {
				refHint := u.RefOrderHint[u.RefFrameIdx[i]]
				if u.GetRelativeDist(refHint, forwardHint) < 0 {
					if secondForwardIdx < 0 || u.GetRelativeDist(refHint, secondForwardHint) > 0 {
						secondForwardIdx = i
						secondForwardHint = refHint
					}
				}
			}

			if secondForwardIdx < 0 {
				skipModeAllowed = 0
			} else {
				skipModeAllowed = 1
				u.SkipModeFrame[0] = shared.LAST_FRAME + util.Min(forwardIdx, secondForwardIdx)
				u.SkipModeFrame[1] = shared.LAST_FRAME + util.Max(forwardIdx, secondForwardIdx)
			}
		}
	}

	if util.Bool(skipModeAllowed) {
		u.SkipModePresent = b.F(1)
	} else {
		u.SkipModePresent = 0
	}
}

func (u *UncompressedHeader) globalMotionParams(b *bitstream.BitStream) {
	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		u.State.GmType[ref] = shared.IDENTITY
		for i := 0; i < 6; i++ {
			if i%3 == 2 {
				u.GmParams[ref][i] = 1 << shared.WARPEDMODEL_PREC_BITS

			} else {
				u.GmParams[ref][i] = 0

			}
		}
	}

	if u.FrameIsIntra {
		return
	}

	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		var typ int
		isGlobal := b.F(1)
		if util.Bool(isGlobal) {
			isRotZoom := b.F(1)
			if util.Bool(isRotZoom) {
				typ = shared.ROTZOOM
			} else {

				isTranslation := b.F(1)
				if util.Bool(isTranslation) {
					typ = shared.TRANSLATION
				} else {
					typ = shared.AFFINE
				}
			}
		} else {
			typ = shared.IDENTITY
		}
		u.State.GmType[ref] = typ
	}

}

func (u *UncompressedHeader) readGlobalParam(typ int, ref int, idx int, b *bitstream.BitStream) {
	absBits := shared.GM_ABS_ALPHA_BITS
	precBits := shared.GM_ALPHA_PREC_BITS

	if idx < 2 {
		if typ == shared.TRANSLATION {
			absBits = shared.GM_ABS_TRANS_ONLY_BITS - util.Int(!u.AllowHighPrecisionMv)
			precBits = shared.GM_TRANS_ONLY_PREC_BITS - util.Int(!u.AllowHighPrecisionMv)
		} else {
			absBits = shared.GM_ABS_TRANS_BITS
			precBits = shared.GM_TRANS_PREC_BITS
		}
	}

	precDiff := shared.WARPEDMODEL_PREC_BITS - precBits

	var round int
	if idx%3 == 2 {
		round = 1 << shared.WARPEDMODEL_PREC_BITS
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
	r := (u.State.PrevGmParams[ref][idx] >> precDiff) - sub
	u.GmParams[ref][idx] = (u.decodeSignedSubexpWithRef(mx, r, b) << precDiff) + round
}

func (u *UncompressedHeader) decodeSignedSubexpWithRef(mx int, r int, b *bitstream.BitStream) int {
	v := u.decodeSubexp(mx, b)
	if (r << 1) <= mx {
		return util.InverseRecenter(r, v)
	} else {
		return mx - 1 - util.InverseRecenter(mx-1-r, v)
	}
}

func (u *UncompressedHeader) decodeSubexp(numSyms int, b *bitstream.BitStream) int {
	i := 0
	mk := 0
	k := 3
	var b2 int
	for {
		if util.Bool(i) {
			b2 = k + i - 1
		} else {
			b2 = k
		}

		a := 1 << b2
		if numSyms <= mk+3*a {
			subexpFinalBits := b.Ns(numSyms - mk)
			return subexpFinalBits + mk
		} else {
			subexmpMoreBits := b.F(1)
			if util.Bool(subexmpMoreBits) {
				i++
				mk += a
			} else {
				subexpBits := b.F(b2)
				return subexpBits + mk
			}
		}
	}
}

func (u *UncompressedHeader) filmGrainParams() {
	panic("not implemented")
}

// decode_frame_wrapup()
func (u *UncompressedHeader) DecodeFrameWrapup() {
	panic("not implemented")
}

// temporal_point_info()
func (u *UncompressedHeader) TemporalPointInfo(b *bitstream.BitStream) {
	n := u.State.SequenceHeader.DecoderModelInfo.FramePresentationTimeLengthMinusOne + 1
	u.FramePresentationTime = b.F(n)
}

// load_grain_params( idx )
func (u *UncompressedHeader) loadGrainParams(idx int) {
	panic("not implemented")
}
