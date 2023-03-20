package uncompressedheader

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/tileinfo"
	"github.com/m4tthewde/gov1/internal/util"
)

const ONLY_4X4 = 0

const SUPERRES_DENOM_BITS = 3
const SUPERRES_DENOM_MIN = 9

const INTER_FRAME = 1
const SWITCH_FRAME = 3

const TOTAL_REFS_PER_FRAME = 8

type UncompressedHeader struct {
	ShowExistingFrame          bool
	ShowableFrame              bool
	ShowFrame                  bool
	FrameType                  int
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
	SegmentationEnabled        bool
	SegmentationTemporalUpdate int
	SegmentationUpdateMap      int
	SegmentationUpdateData     int
	LosslessArray              [shared.MAX_SEGMENTS]bool
	GmParams                   [shared.ALTREF_FRAME + 1][6]int
	ForceIntegerMv             bool
	AllowScreenContentTools    int
	RefOrderHint               []int
	RefFrameIdx                [7]int
	OrderHint                  int
	OrderHints                 []int
	SkipModeFrame              []int
	SkipModePresent            int
	SegQMLevel                 [3][shared.MAX_SEGMENTS]int
	DisableCdfUpdate           bool

	LoopFilterDeltaEnabled bool
	LoopFilterDeltaUpdate  int
	LoopFilterLevel        [4]int
	LoopFilterRefDeltas    [8]int
	LoopFilterModeDeltas   [2]int
	LoopFilterSharpness    int

	CdefBits          int
	CdefYPriStrength  []int
	CdefYSecStrength  []int
	CdefUVPriStrength []int
	CdefUVSecStrength []int
	CdefDampening     int

	FrameRestorationType [3]int
	UsesLr               bool
	UseRefFrameMvs       bool

	ApplyGrain            bool
	GrainSeed             int
	PointYValue           []int
	PointYScaling         []int
	ChromaScalingFromLuna bool
	PointCbValue          []int
	PointCbScaling        []int
	PointCrValue          []int
	PointCrScaling        []int
	GrainScalingMinus8    int
	ArCoeffLag            int
	ArCoeffsYPlus128      []int
	ArCoeffsCbPlus128     []int
	ArCoeffsCrPlus128     []int
	ArCoeffShiftMinus6    int
	GrainScaleShift       int
	CbMult                int
	CbLumaMult            int
	CbOffset              int
	CrMult                int
	CrLumaMult            int
	CrOffset              int
	OverlapFlag           bool
	ClipToRestrictedRange bool
}

func NewUncompressedHeader(h header.Header, sh sequenceheader.SequenceHeader, b *bitstream.BitStream, s *state.State) UncompressedHeader {
	u := UncompressedHeader{}

	u.build(h, sh, s, b)
	return u
}

func (u *UncompressedHeader) build(h header.Header, sh sequenceheader.SequenceHeader, s *state.State, b *bitstream.BitStream) {
	var idLen int
	if sh.FrameIdNumbersPresent {
		idLen = sh.AdditionalFrameIdLengthMinusOne +
			sh.DeltaFrameIdLengthMinusTwo + 3
	}

	var errorResilientMode bool

	var refValid [shared.NUM_REF_FRAMES]int
	var refOrderHint [shared.NUM_REF_FRAMES]int
	var orderHints [shared.REFS_PER_FRAME]int

	bufferRemovalTime := []int{}

	allFrames := ((1 << shared.NUM_REF_FRAMES) - 1)
	if sh.ReducedStillPictureHeader {
		u.ShowExistingFrame = false
		u.FrameType = shared.KEY_FRAME
		u.FrameIsIntra = true

		u.ShowFrame = true
		u.ShowableFrame = false
	} else {
		showExistingFrame := util.Bool(b.F(1))

		if showExistingFrame {
			frameToShowMapIdx := b.F(3)

			if sh.DecoderModelInfoPresent && !sh.TimingInfo.EqualPictureInterval {
				u.TemporalPointInfo(sh, b)
			}

			u.RefreshImageFlags = 0
			if sh.FrameIdNumbersPresent {
				u.DisplayFrameId = b.F(idLen)
			}

			u.FrameType = s.RefFrameType[frameToShowMapIdx]

			if u.FrameType == shared.KEY_FRAME {
				u.RefreshImageFlags = allFrames
			}

			if sh.FilmGrainParamsPresent {
				u.loadGrainParams(frameToShowMapIdx)
			}

			return
		}

		u.FrameType = b.F(2)

		u.FrameIsIntra = (u.FrameType == 2 || u.FrameType == 0)

		u.ShowFrame = util.Bool(b.F(1))

		if u.ShowFrame && sh.DecoderModelInfoPresent && !sh.TimingInfo.EqualPictureInterval {
			u.TemporalPointInfo(sh, b)
		}

		if u.ShowFrame {
			u.ShowableFrame = u.FrameType != 0
		} else {
			u.ShowableFrame = util.Bool(b.F(1))
		}

		if u.FrameType == 3 || u.FrameType == 0 && u.ShowFrame {
			errorResilientMode = true
		} else {
			errorResilientMode = util.Bool(b.F(1))
		}
	}

	if u.FrameType == 0 && u.ShowFrame {
		for i := 0; i < shared.NUM_REF_FRAMES; i++ {
			refValid[i] = 0
			refOrderHint[i] = 0
		}

		for i := 0; i < shared.REFS_PER_FRAME; i++ {
			orderHints[shared.LAST_FRAME+1] = 0
		}
	}

	u.DisableCdfUpdate = util.Bool(b.F(1))

	if sh.SeqForceScreenContentTools == 2 {
		u.AllowScreenContentTools = b.F(1)
	} else {
		u.AllowScreenContentTools = 1
	}

	if util.Bool(u.AllowScreenContentTools) {
		if sh.SeqForceIntegerMv == 2 {
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

	if sh.FrameIdNumbersPresent {
		u.PrevFrameId = u.CurrentFrameId
		u.CurrentFrameId = b.F(idLen)
		u.markRefFrames(idLen, sh)
	} else {
		u.CurrentFrameId = 0
	}

	if u.FrameType == SWITCH_FRAME {
		u.FrameSizeOverrideFlag = true
	} else if sh.ReducedStillPictureHeader {
		u.FrameSizeOverrideFlag = false

	} else {
		u.FrameSizeOverrideFlag = util.Bool(b.F(1))
	}

	orderHint := b.F(sh.OrderHintBits)
	u.OrderHint = orderHint

	if u.FrameIsIntra || errorResilientMode {
		u.PrimaryRefFrame = shared.PRIMARY_REF_NONE
	} else {
		u.PrimaryRefFrame = b.F(3)
	}

	if sh.DecoderModelInfoPresent {
		bufferRemovalTimePresent := util.Bool(b.F(1))

		if bufferRemovalTimePresent {
			for opNum := 0; opNum <= sh.OperatingPointsCountMinusOne; opNum++ {
				if sh.DecoderModelPresentForThisOp[opNum] {
					opPtIdc := sh.OperatingPointIdc[opNum]
					inTemporalLayer := ((opPtIdc >> h.ExtensionHeader.TemporalID) & 1) != 0
					inSpatialLayer := ((opPtIdc >> h.ExtensionHeader.SpatialID) & 1) != 0

					if opPtIdc == 0 || (inTemporalLayer && inSpatialLayer) {
						n := sh.DecoderModelInfo.BufferRemovalTimeLengthMinusOne + 1
						bufferRemovalTime[opNum] = b.F(n)
					}
				}
			}
		}
	}

	RefValid := []int{}
	ref_order_hint := []int{}

	u.UseRefFrameMvs = false
	var refreshFrameFlags int

	if u.FrameType == 3 || u.FrameType == 0 || u.ShowFrame {
		refreshFrameFlags = allFrames
	} else {
		refreshFrameFlags = b.F(8)
	}

	if !u.FrameIsIntra || refreshFrameFlags != allFrames {
		if errorResilientMode && sh.EnableOrderHint {
			for i := 0; i < shared.NUM_REF_FRAMES; i++ {
				ref_order_hint[i] = b.F(sh.OrderHintBits)

				if ref_order_hint[i] != u.RefOrderHint[i] {
					RefValid[i] = 0
				}
			}
		}
	}

	expectedFrameId := []int{}

	if u.FrameIsIntra {
		u.frameSize(b, s, sh)
		u.renderSize(b)

		if util.Bool(u.AllowScreenContentTools) && u.UpscaledWidth == u.FrameWidth {
			u.AllowIntraBc = util.Bool(b.F(1))
		}
	} else {

		var frameRefsShortSignaling bool
		if !sh.EnableOrderHint {
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

			if sh.FrameIdNumbersPresent {
				n := sh.DeltaFrameIdLengthMinusTwo + 2
				deltaFrameIdMinusOne := b.F(n)
				DeltaFrameId := deltaFrameIdMinusOne + 1
				expectedFrameId[i] = (u.CurrentFrameId + (1 << idLen) - DeltaFrameId) % (1 << idLen)
			}
		}

		if u.FrameSizeOverrideFlag && !errorResilientMode {
			u.frameSizeWithRefs()
		} else {
			u.frameSize(b, s, sh)
			u.renderSize(b)

		}
		if u.ForceIntegerMv {
			u.AllowHighPrecisionMv = false
		} else {

			u.AllowHighPrecisionMv = util.Bool(b.F(1))
		}

		u.readInterpolationFilter(b)
		u.IsMotionModeSwitchable = util.Bool(b.F(1))

		if errorResilientMode || !sh.EnableRefFrameMvs {
			u.UseRefFrameMvs = false

		} else {
			u.UseRefFrameMvs = util.Bool(b.F(1))
		}

		RefFrameSignBias := []bool{}

		for i := 0; i < shared.REFS_PER_FRAME; i++ {
			refFrame := shared.LAST_FRAME + 1
			hint := u.RefOrderHint[u.RefFrameIdx[i]]
			u.OrderHints[refFrame] = hint
			if !sh.EnableOrderHint {
				RefFrameSignBias[refFrame] = false
			} else {
				RefFrameSignBias[refFrame] = u.GetRelativeDist(hint, u.OrderHint, sh) > 0
			}
		}
	}

	if sh.ReducedStillPictureHeader || u.DisableCdfUpdate {
		u.DisableFrameEndUpdateCdf = true
	} else {
		u.DisableFrameEndUpdateCdf = util.Bool(b.F(1))
	}

	if u.PrimaryRefFrame == shared.PRIMARY_REF_NONE {
		initNonCoeffCdfs(s)
		u.setupPastIndependence(s)
	} else {
		u.loadCdfs(u.RefFrameIdx[u.PrimaryRefFrame])
		u.loadPrevious()
	}

	if u.UseRefFrameMvs {
		u.motionFieldEstimation()
	}

	u.TileInfo = tileinfo.NewTileInfo(sh, s, b)

	u.quantizationParams(b, sh)
	u.segmentationParams(b, s)
	u.deltaQParams(b)
	u.deltaLfParams(b)

	if u.PrimaryRefFrame == shared.PRIMARY_REF_NONE {

		u.initCoeffCdfs(s)
	} else {
		u.loadPreviousSegementIds()
	}

	u.CodedLossless = true

	for segmentId := 0; segmentId < shared.MAX_SEGMENTS; segmentId++ {
		qIndex := u.GetQIndex(1, segmentId, s)
		u.LosslessArray[segmentId] = qIndex == 0 && u.DeltaQYDc == 0 && u.DeltaQUAc == 0 && u.DeltaQUDc == 0 && u.DeltaQVAc == 0 && u.DeltaQVDc == 0

		if !u.LosslessArray[segmentId] {
			u.CodedLossless = false
		}

		if u.UsingQMatrix {
			if u.LosslessArray[segmentId] {
				u.SegQMLevel[0][segmentId] = 15
				u.SegQMLevel[1][segmentId] = 15
				u.SegQMLevel[2][segmentId] = 15
			} else {
				u.SegQMLevel[0][segmentId] = u.Qmy
				u.SegQMLevel[1][segmentId] = u.Qmy
				u.SegQMLevel[2][segmentId] = u.Qmy
			}
		}
	}

	u.AllLossless = u.CodedLossless && (u.FrameWidth == u.UpscaledWidth)

	u.loopFilterParams(b, sh)
	u.cdefParams(b, sh)
	u.lrParams(b, s, sh)
	u.readTxMode(b)
	u.frameReferenceMode(b)
	u.skipModeParams(b, sh)

	if u.FrameIsIntra || errorResilientMode || !sh.EnableWarpedMotion {
		u.AllowWarpedMotion = false
	} else {
		u.AllowWarpedMotion = util.Bool(b.F(1))
	}

	u.ReducedTxSet = util.Bool(b.F(1))

	u.globalMotionParams(b, s)
	u.filmGrainParams(b, sh)
}

// mark_ref_frames( idLen)
func (u *UncompressedHeader) markRefFrames(idLen int, sh sequenceheader.SequenceHeader) {
	diffLen := sh.DeltaFrameIdLengthMinusTwo + 2

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
func (u *UncompressedHeader) frameSize(b *bitstream.BitStream, s *state.State, sh sequenceheader.SequenceHeader) {
	if u.FrameSizeOverrideFlag {
		n := sh.FrameWidthBitsMinusOne + 1
		frameWidthMinusOne := b.F(n)

		n = sh.FrameHeightBitsMinusOne + 1
		frameHeightMinusOne := b.F(n)

		u.FrameWidth = frameWidthMinusOne + 1
		u.FrameHeight = frameHeightMinusOne + 1
	} else {
		u.FrameWidth = sh.MaxFrameWidthMinusOne + 1
		u.FrameHeight = sh.MaxFrameHeightMinusOne + 1
	}

	u.superResParams(b, sh)
	u.computeImageSize(s)
}

// superres_params()
func (u *UncompressedHeader) superResParams(b *bitstream.BitStream, sh sequenceheader.SequenceHeader) {
	if sh.EnableSuperRes {
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
func (u *UncompressedHeader) computeImageSize(s *state.State) {
	s.MiCols = 2 * ((u.FrameWidth + 7) >> 3)
	s.MiRows = 2 * ((u.FrameHeight + 7) >> 3)
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
func (u *UncompressedHeader) GetRelativeDist(a int, b int, sh sequenceheader.SequenceHeader) int {
	if !u.EnableOrderHint {
		return 0
	}

	diff := a - b
	m := 1 << (sh.OrderHintBits - 1)
	diff = (diff & (m - 1)) - (diff & m)

	return diff
}

// setup_past_independence()
func (u *UncompressedHeader) setupPastIndependence(s *state.State) {
	for i := 0; i < shared.MAX_SEGMENTS; i++ {
		for j := 0; j < shared.SEG_LVL_MAX; j++ {
			s.FeatureData[i][j] = 0
			s.FeatureEnabled[i][j] = 0
		}
	}

	s.PrevSegmentIds = make([][]int, s.MiRows)
	for i := range s.PrevSegmentIds {
		s.PrevSegmentIds[i] = make([]int, s.MiCols)
	}

	for row := 0; row < s.MiRows; row++ {
		for col := 0; col < s.MiCols; col++ {
			s.PrevSegmentIds[row][col] = 0
		}
	}

	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		s.GmType[ref] = shared.IDENTITY
	}

	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		for i := 0; i <= 5; i++ {
			if (i % 3) == 2 {
				s.PrevGmParams[ref][i] = 1 << shared.WARPEDMODEL_PREC_BITS
			} else {
				s.PrevGmParams[ref][i] = 0
			}
		}
	}

	u.LoopFilterDeltaEnabled = true
	u.LoopFilterRefDeltas[shared.INTRA_FRAME] = 1
	u.LoopFilterRefDeltas[shared.LAST_FRAME] = 0
	u.LoopFilterRefDeltas[shared.LAST2_FRAME] = 0
	u.LoopFilterRefDeltas[shared.LAST3_FRAME] = 0
	u.LoopFilterRefDeltas[shared.BWDREF_FRAME] = 0
	u.LoopFilterRefDeltas[shared.GOLDEN_FRAME] = -1
	u.LoopFilterRefDeltas[shared.ALTREF_FRAME] = -1
	u.LoopFilterRefDeltas[shared.ALTREF2_FRAME] = -1

	u.LoopFilterModeDeltas[0] = 0
	u.LoopFilterModeDeltas[1] = 0
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
func (u *UncompressedHeader) quantizationParams(b *bitstream.BitStream, sh sequenceheader.SequenceHeader) {
	u.BaseQIdx = b.F(8)

	u.DeltaQYDc = u.readDeltaQ(b)

	var diffUvDelta bool
	if sh.ColorConfig.NumPlanes > 1 {
		if sh.ColorConfig.SeparateUvDeltaQ {
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

		if !sh.ColorConfig.SeparateUvDeltaQ {
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
func (u *UncompressedHeader) segmentationParams(b *bitstream.BitStream, s *state.State) {
	u.SegmentationEnabled = util.Bool(b.F(1))
	if u.SegmentationEnabled {
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
				for j := 0; j < shared.SEG_LVL_MAX; j++ {
					featureValue := 0
					featureEnabled := b.F(1)
					s.FeatureEnabled[i][j] = featureEnabled
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
					s.FeatureData[i][j] = clippedValue
				}
			}
		}
	} else {
		for i := 0; i < shared.MAX_SEGMENTS; i++ {
			for j := 0; j < shared.SEG_LVL_MAX; i++ {
				s.FeatureEnabled[i][j] = 0
				s.FeatureData[i][j] = 0
			}

		}
	}
	u.SegIdPreSkip = 0
	u.LastActiveSegId = 0

	for i := 0; i < shared.MAX_SEGMENTS; i++ {
		for j := 0; j < shared.SEG_LVL_MAX; j++ {
			if s.FeatureEnabled[i][j] == 1 {
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

func (u *UncompressedHeader) loadPreviousSegementIds() {
	panic("not implemented")
}

// get_qindex( ignoreDeltaQ, segmentId )
func (u *UncompressedHeader) GetQIndex(ignoreDeltaQ int, segmentId int, s *state.State) int {
	if u.segFeatureActiveIdx(segmentId, shared.SEG_LVL_ALT_Q, s) {
		data := s.FeatureData[segmentId][shared.SEG_LVL_ALT_Q]
		qindex := u.BaseQIdx + data

		if ignoreDeltaQ == 0 && u.DeltaQPresent {
			qindex = s.CurrentQIndex
		}

		return util.Clip3(0, 255, qindex)
	}

	if ignoreDeltaQ == 0 && u.DeltaQPresent {
		return s.CurrentQIndex
	}

	return u.BaseQIdx
}

// seg_feature_active_idx( idx, feature )
func (u *UncompressedHeader) segFeatureActiveIdx(idx int, feature int, s *state.State) bool {
	return u.SegmentationEnabled && util.Bool(s.FeatureEnabled[idx][feature])

}

// loop_filter_params()
func (u *UncompressedHeader) loopFilterParams(b *bitstream.BitStream, sh sequenceheader.SequenceHeader) {
	if u.CodedLossless || u.AllowIntraBc {
		u.LoopFilterLevel[0] = 0
		u.LoopFilterLevel[1] = 1
		u.LoopFilterRefDeltas[shared.INTRA_FRAME] = 1
		u.LoopFilterRefDeltas[shared.LAST_FRAME] = 0
		u.LoopFilterRefDeltas[shared.LAST2_FRAME] = 0
		u.LoopFilterRefDeltas[shared.LAST3_FRAME] = 0
		u.LoopFilterRefDeltas[shared.BWDREF_FRAME] = 0
		u.LoopFilterRefDeltas[shared.GOLDEN_FRAME] = -1
		u.LoopFilterRefDeltas[shared.ALTREF_FRAME] = -1
		u.LoopFilterRefDeltas[shared.ALTREF2_FRAME] = -1

		u.LoopFilterModeDeltas[0] = 0
		u.LoopFilterModeDeltas[1] = 0

		return
	}

	u.LoopFilterLevel[0] = b.F(6)
	u.LoopFilterLevel[1] = b.F(6)

	if sh.ColorConfig.NumPlanes > 1 {
		if util.Bool(u.LoopFilterLevel[0]) || util.Bool(u.LoopFilterLevel[1]) {
			u.LoopFilterLevel[2] = b.F(6)
			u.LoopFilterLevel[3] = b.F(6)
		}
	}

	u.LoopFilterSharpness = b.F(6)
	u.LoopFilterDeltaEnabled = util.Bool(b.F(1))

	if u.LoopFilterDeltaEnabled {
		u.LoopFilterDeltaUpdate = b.F(1)

		if u.LoopFilterDeltaUpdate == 1 {
			for i := 0; i < TOTAL_REFS_PER_FRAME; i++ {
				updateRefDelta := b.F(1)
				if updateRefDelta == 1 {
					u.LoopFilterRefDeltas[i] = b.Su(1 + 6)
				}
			}
			for i := 0; i < 2; i++ {
				updateModeDelta := b.F(1)
				if updateModeDelta == 1 {
					u.LoopFilterModeDeltas[i] = b.Su(1 + 6)
				}
			}
		}
	}
}

// cdef_params( )
func (u *UncompressedHeader) cdefParams(b *bitstream.BitStream, sh sequenceheader.SequenceHeader) {
	if u.CodedLossless || u.AllowIntraBc || !sh.EnableCdef {
		u.CdefYPriStrength = make([]int, 1)
		u.CdefYSecStrength = make([]int, 1)
		u.CdefUVPriStrength = make([]int, 1)
		u.CdefUVSecStrength = make([]int, 1)

		u.CdefBits = 0
		u.CdefYPriStrength[0] = 0
		u.CdefYSecStrength[0] = 0
		u.CdefUVPriStrength[0] = 0
		u.CdefUVSecStrength[0] = 0
		u.CdefDampening = 3
		return
	}

	cdefDampeningMinus3 := b.F(2)
	u.CdefDampening = cdefDampeningMinus3 + 3
	u.CdefBits = b.F(2)
	u.CdefYPriStrength = make([]int, (1 << u.CdefBits))
	u.CdefYSecStrength = make([]int, (1 << u.CdefBits))
	u.CdefUVPriStrength = make([]int, (1 << u.CdefBits))
	u.CdefUVSecStrength = make([]int, (1 << u.CdefBits))

	for i := 0; i < (1 << u.CdefBits); i++ {
		u.CdefYPriStrength[i] = b.F(4)
		u.CdefYSecStrength[i] = b.F(2)
		if u.CdefYSecStrength[i] == 3 {
			u.CdefYSecStrength[i] += 1
		}

		if sh.ColorConfig.NumPlanes > 1 {
			u.CdefUVPriStrength[i] = b.F(4)
			u.CdefUVSecStrength[i] = b.F(2)
			if u.CdefUVSecStrength[i] == 3 {
				u.CdefUVSecStrength[i] += 1
			}

		}
	}
}

// lr_params()
func (u *UncompressedHeader) lrParams(b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader) {
	if u.AllLossless || u.AllowIntraBc || !sh.EnableRestoration {
		u.FrameRestorationType[0] = shared.RESTORE_NONE
		u.FrameRestorationType[1] = shared.RESTORE_NONE
		u.FrameRestorationType[2] = shared.RESTORE_NONE
		u.UsesLr = false
		return
	}

	u.UsesLr = false
	usesChromaLr := false
	for i := 0; i < sh.ColorConfig.NumPlanes; i++ {
		lrType := b.F(2)
		u.FrameRestorationType[i] = shared.REMAP_LR_TYPE[lrType]
		if u.FrameRestorationType[i] != shared.RESTORE_NONE {
			u.UsesLr = true
			if i > 0 {
				usesChromaLr = true
			}
		}
	}

	var lrUnitShift int
	if u.UsesLr {
		if sh.Use128x128SuperBlock {
			lrUnitShift = b.F(1)
			lrUnitShift++
		} else {
			lrUnitShift = b.F(1)

			if util.Bool(lrUnitShift) {
				lrUnitExtraShift := b.F(1)
				lrUnitShift += lrUnitExtraShift
			}
		}
		state.LoopRestorationSize[0] = shared.RESTORATION_TILESIZE_MAX >> (2 - lrUnitShift)

		var lrUvShift int
		if sh.ColorConfig.SubsamplingX && sh.ColorConfig.SubsamplingY && usesChromaLr {
			lrUvShift = b.F(1)
		} else {
			lrUvShift = 0
		}

		state.LoopRestorationSize[1] = state.LoopRestorationSize[0] >> lrUvShift
		state.LoopRestorationSize[2] = state.LoopRestorationSize[0] >> lrUvShift
	}
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
func (u *UncompressedHeader) skipModeParams(b *bitstream.BitStream, sh sequenceheader.SequenceHeader) {
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

			if u.GetRelativeDist(refHint, u.OrderHint, sh) < 0 {
				if forwardIdx < 0 || u.GetRelativeDist(refHint, forwardHint, sh) > 0 {
					forwardIdx = i
					forwardHint = refHint
				}
			} else if u.GetRelativeDist(refHint, u.OrderHint, sh) > 0 {
				if backwardIdx < 0 || u.GetRelativeDist(refHint, backwardHint, sh) > 0 {
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
				if u.GetRelativeDist(refHint, forwardHint, sh) < 0 {
					if secondForwardIdx < 0 || u.GetRelativeDist(refHint, secondForwardHint, sh) > 0 {
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

func (u *UncompressedHeader) globalMotionParams(b *bitstream.BitStream, s *state.State) {
	for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
		s.GmType[ref] = shared.IDENTITY
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
		s.GmType[ref] = typ
	}

}

func (u *UncompressedHeader) readGlobalParam(typ int, ref int, idx int, b *bitstream.BitStream, s *state.State) {
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
	r := (s.PrevGmParams[ref][idx] >> precDiff) - sub
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

// film_grain_params()
func (u *UncompressedHeader) filmGrainParams(b *bitstream.BitStream, sh sequenceheader.SequenceHeader) {
	if !sh.FilmGrainParamsPresent || (!u.ShowFrame && !u.ShowableFrame) {
		u.resetGrainParams()
		return
	}

	u.ApplyGrain = util.Bool(b.F(1))
	if !u.ApplyGrain {
		u.resetGrainParams()
		return
	}

	u.GrainSeed = b.F(16)

	var updateGrain bool
	if u.FrameType == INTER_FRAME {
		updateGrain = util.Bool(b.F(1))
	} else {
		updateGrain = true
	}

	if !updateGrain {
		filmGrainParamsRefIdx := b.F(3)
		tempGrainSeed := u.GrainSeed
		u.loadGrainParams(filmGrainParamsRefIdx)
		u.GrainSeed = tempGrainSeed
		return
	}

	numYPoints := b.F(4)
	u.PointYValue = make([]int, numYPoints)
	u.PointYScaling = make([]int, numYPoints)
	for i := 0; i < numYPoints; i++ {
		u.PointYValue[i] = b.F(8)
		u.PointYScaling[i] = b.F(8)

	}

	if sh.ColorConfig.MonoChrome {
		u.ChromaScalingFromLuna = false
	} else {
		u.ChromaScalingFromLuna = util.Bool(b.F(1))
	}

	var numCbPoints int
	var numCrPoints int
	if sh.ColorConfig.MonoChrome || u.ChromaScalingFromLuna || (sh.ColorConfig.SubsamplingX && sh.ColorConfig.SubsamplingY && numYPoints == 0) {
		numCbPoints = 0
		numCrPoints = 0
	} else {
		numCbPoints = b.F(4)

		u.PointCbValue = make([]int, numCbPoints)
		u.PointCbScaling = make([]int, numCbPoints)
		for i := 0; i < numCbPoints; i++ {
			u.PointCbValue[i] = b.F(8)
			u.PointCbScaling[i] = b.F(8)
		}

		numCrPoints = b.F(4)

		u.PointCrValue = make([]int, numCrPoints)
		u.PointCrScaling = make([]int, numCrPoints)
		for i := 0; i < numCrPoints; i++ {
			u.PointCrValue[i] = b.F(8)
			u.PointCrScaling[i] = b.F(8)
		}
	}

	u.GrainScalingMinus8 = b.F(2)
	u.ArCoeffLag = b.F(2)

	numPosLuma := 2 * u.ArCoeffLag * (u.ArCoeffLag + 1)

	var numPosChroma int
	if util.Bool(numYPoints) {
		numPosChroma = numPosLuma + 1
		u.ArCoeffsYPlus128 = make([]int, numPosLuma)
		for i := 0; i < numPosLuma; i++ {
			u.ArCoeffsYPlus128[i] = b.F(8)
		}
	} else {
		numPosChroma = numPosLuma
	}

	if u.ChromaScalingFromLuna || util.Bool(numCbPoints) {
		u.ArCoeffsCbPlus128 = make([]int, numPosChroma)
		for i := 0; i < numPosChroma; i++ {
			u.ArCoeffsCbPlus128[i] = b.F(8)
		}
	}

	if u.ChromaScalingFromLuna || util.Bool(numCrPoints) {
		u.ArCoeffsCrPlus128 = make([]int, numPosChroma)
		for i := 0; i < numPosChroma; i++ {
			u.ArCoeffsCrPlus128[i] = b.F(8)
		}
	}

	u.ArCoeffShiftMinus6 = b.F(2)
	u.GrainScaleShift = b.F(2)

	if util.Bool(numCbPoints) {
		u.CbMult = b.F(8)
		u.CbLumaMult = b.F(8)
		u.CbOffset = b.F(9)
	}

	if util.Bool(numCrPoints) {
		u.CrMult = b.F(8)
		u.CrLumaMult = b.F(8)
		u.CrOffset = b.F(9)
	}

	u.OverlapFlag = util.Bool(b.F(1))
	u.ClipToRestrictedRange = util.Bool(b.F(1))
}

// rest_grain_params()
func (u *UncompressedHeader) resetGrainParams() {
	u.ApplyGrain = false
	u.GrainSeed = 0
	u.PointYValue = []int{}
	u.PointYScaling = []int{}
	u.ChromaScalingFromLuna = false
	u.PointCbValue = []int{}
	u.PointCbScaling = []int{}
	u.PointCrValue = []int{}
	u.PointCrScaling = []int{}
	u.GrainScalingMinus8 = 0
	u.ArCoeffLag = 0
	u.ArCoeffsYPlus128 = []int{}
	u.ArCoeffsCbPlus128 = []int{}
	u.ArCoeffsCrPlus128 = []int{}
	u.ArCoeffShiftMinus6 = 0
	u.GrainScaleShift = 0
	u.CbMult = 0
	u.CbLumaMult = 0
	u.CbOffset = 0
	u.CrMult = 0
	u.CrLumaMult = 0
	u.CrOffset = 0
	u.OverlapFlag = false
	u.ClipToRestrictedRange = false
}

// decode_frame_wrapup()
func (u *UncompressedHeader) DecodeFrameWrapup() {
	panic("not implemented")
}

// temporal_point_info()
func (u *UncompressedHeader) TemporalPointInfo(sh sequenceheader.SequenceHeader, b *bitstream.BitStream) {
	n := sh.DecoderModelInfo.FramePresentationTimeLengthMinusOne + 1
	u.FramePresentationTime = b.F(n)
}

// load_grain_params( idx )
func (u *UncompressedHeader) loadGrainParams(idx int) {
	panic("not implemented")
}
