package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/parser"
)

const FRAME_LF_COUNT = 4
const WIENER_COEFFS = 3

const PARTITION_NONE = 0
const PARTITION_HORZ = 1
const PARTITION_VERT = 2
const PARTITION_SPLIT = 3
const PARTITION_HORZ_A = 4
const PARTITION_HORZ_B = 5
const PARTITION_VERT_A = 6
const PARTITION_VERT_B = 7
const PARTITION_HORZ_4 = 8
const PARTITION_VERT_4 = 9

const MAX_SB_SIZE = 128
const MAX_FRAME_DISTANCE = 31

const INTRA_EDGE_TAPS = 5
const SUBPEL_BITS = 4
const SCALE_SUBPEL_BITS = 10
const SUBPEL_MASK = 15

const PALETTE_COLORS = 8
const PALETTE_NUM_NEIGHBORS = 3

const DELTA_Q_SMALL = 3
const DELTA_LF_SMALL = 3

const INTRA_FRAME = 0
const NONE = -1

const SINGLE_REFERENCE = 0
const COMPOUND_REFERENCE = 1

const COMPOUND_WEDGE = 0
const COMPOUND_DIFFWTD = 1
const COMPOUND_AVERAGE = 2
const COMPOUND_INTRA = 3
const COMPOUND_DISTANCE = 4

const UNIDIR_COMP_REFERENCE = 0
const BIDIR_COMP_REFERENCE = 1

var Palette_Color_Hash_Multipliers = []int{1, 2, 2}

var Mode_To_Angle = []int{0, 90, 180, 45, 135, 113, 157, 203, 67, 0, 0, 0, 0}

const ANGLE_STEP = 3

const RESTORE_NONE = 0
const RESTORE_WIENER = 1
const RESTORE_SGRPROJ = 2
const RESTORE_SWITCHABLE = 3

const MI_SIZE = 4

const SGRPROJ_PARAMS_BITS = 4
const SGRPROJ_BITS = 7
const SGRPROJ_PRJ_SUBEXP_K = 4

const DC_PRED = 0
const II_DC_PRED = 0
const II_V_PRED = 1
const II_H_PRED = 2
const II_SMOOTH_PRED = 3

const SIMPLE = 0
const OBMC = 1
const LOCALWARP = 2

const COMPUND_AVERAGE = 2

const MAX_REF_MV_STACK_SIZE = 8

const V_PRED = 1
const H_PRED = 2
const SMOOTH_PRED = 9
const SMOOTH_V_PRED = 10
const SMOOTH_H_PRED = 11

const D67_PRED = 8
const UV_CFL_PRED = 13

const MAX_ANGLE_DELTA = 3

const CFL_SIGN_ZERO = 0
const CFL_SIGN_NEG = 2

const INTRABC_DELAY_PIXELS = 256

const MV_INTRABC_CONTEXT = 1

const MV_JOINT_ZERO = 0
const MV_JOINT_HNZVZ = 1
const MV_JOINT_HZVNZ = 2
const MV_JOINT_HNZVNZ = 3

const CLASS0_SIZE = 2

const MV_CLASS_0 = 0
const MV_CLASS_1 = 1
const MV_CLASS_2 = 2
const MV_CLASS_3 = 3
const MV_CLASS_4 = 4
const MV_CLASS_5 = 5
const MV_CLASS_6 = 6
const MV_CLASS_7 = 7
const MV_CLASS_8 = 8
const MV_CLASS_9 = 9
const MV_CLASS_10 = 10

const LEAST_SQUARES_SAMPLES_MAX = 8

const REF_SCALE_SHIFT = 14

const INTRA_FILTER_SCALE_BITS = 4

const LS_MV_MAX = 256
const DIV_LUT_PREC_BITS = 13
const DIV_LUT_BITS = 8

type TileGroup struct {
	State State

	LrType              [][][]int
	RefLrWiener         [][][]int
	LrWiener            [][][][][]int
	LrSgrSet            [][][]int
	RefSgrXqd           [][]int
	LrSgrXqd            [][][][]int
	HasChroma           bool
	SegmentId           int
	SegmentIds          [][]int
	Lossless            bool
	Skip                int
	YMode               int
	YModes              [][]int
	UVMode              int
	UVModes             [][]int
	PaletteSizeY        int
	PaletteSizeUV       int
	InterpFilter        []int
	InterpFilters       [][][]int
	NumMvFound          int
	NewMvCount          int
	GlobalMvs           [][]int
	Block_Width         []int
	Block_Height        []int
	IsInters            [][]int
	Mv                  [][]int
	Mvs                 [][][][]int
	FoundMatch          int
	RefStackMv          [][][]int
	WeightStack         []int
	AngleDeltaY         int
	AngleDeltaUV        int
	CflAlphaU           int
	CflAlphaV           int
	useIntrabc          int
	PredMv              [][]int
	RefMvIdx            int
	MvCtx               int
	PaletteSizes        [][][]int
	PaletteColors       [][][][]int
	PaletteCache        []int
	PaletteColorsY      []int
	PaletteColorsU      []int
	PaletteColorsV      []int
	FilterIntraMode     int
	SkipMode            int
	IsInter             int
	MotionMode          int
	CompoundType        int
	LeftRefFrame        []int
	AboveRefFrame       []int
	LeftIntra           bool
	AboveIntra          bool
	LeftSingle          bool
	AboveSingle         bool
	AboveSegPredContext []int

	InterIntra      int
	InterIntraMode  int
	UseFilterIntra  int
	WedgeIndex      int
	WedgeSign       int
	WedgeInterIntra int

	NumSamples        int
	NumSamplesScanned int
	CandList          [][]int

	RefUpscaledWidth  []int
	RefUpscaledHeight []int
	MaskType          int

	ColorMapY        [][]int
	ColorMapUV       [][]int
	ColorOrder       []int
	ColorContextHash int

	InterTxSizes [][]int
	TxSize       int

	AboveLevelContext [][]int
	AboveDcContext    [][]int
	LeftLevelContext  [][]int
	LeftDcContext     [][]int

	CompGroupIdxs [][]int
	CompoundIdxs  [][]int
	CompGroupIdx  int
	CompoundIdx   int

	IsInterIntra bool

	AboveRow []int
	LeftCol  []int

	InterRound0    int
	InterRound1    int
	InterPostRound int

	LocalValid      bool
	LocalWarpParams []int

	FrameStore [][][][]int
	Mask       [][]int

	FwdWeight int
	BckWeight int
}

func NewTileGroup(p *parser.Parser, sz int) TileGroup {
	t := TileGroup{}
	t.build(p, sz)
	return t
}

func (t *TileGroup) build(p *parser.Parser, sz int) {
	NumTiles := p.TileCols * p.TileRows
	startbitPos := p.position
	tileStartAndEndPresentFlag := false

	if NumTiles > 1 {
		tileStartAndEndPresentFlag = p.f(1) != 0
	}

	var tgStart int
	var tgEnd int
	var tileBits int

	if NumTiles == 1 || !tileStartAndEndPresentFlag {
		tgStart = 0
		tgEnd = NumTiles - 1
	} else {
		tileBits = p.TileColsLog2 + p.TileRowsLog2
		tgStart = p.f(tileBits)
		tgEnd = p.f(tileBits)
	}

	p.byteAlignment()
	endBitBos := p.position
	headerBytes := (endBitBos - startbitPos) / 8
	sz -= headerBytes

	for p.TileNum = tgStart; p.TileNum <= tgEnd; p.TileNum++ {
		tileRow := p.TileNum / p.TileCols
		tileCol := p.TileNum % p.TileCols
		lastTile := p.TileNum == tgEnd

		var tileSize int
		if lastTile {
			tileSize = sz
		} else {
			tileSizeMinusOne := p.le(p.TileSizeBytes)
			tileSize = tileSizeMinusOne + 1
			sz -= tileSize + p.TileSizeBytes
		}

		p.MiRowStart = p.MiRowStarts[tileRow]
		p.MiRowEnd = p.MiRowStarts[tileRow+1]
		p.MiColStart = p.MiColStarts[tileCol]
		p.MiColEnd = p.MiColStarts[tileCol+1]
		p.CurrentQIndex = p.uncompressedHeader.BaseQIdx
		p.initSymbol(tileSize)
		t.decodeTile(p)
	}
}

// 8.2.4 Exit process for symbol decoder
func (t *TileGroup) exitSymbol(p *parser.Parser) {
	if p.SymbolMaxBits < -14 {
		panic("Violating bitstream conformance!")
	}

	p.position += Max(0, p.SymbolMaxBits)

	if !p.uncompressedHeader.DisableFrameEndUpdateCdf && p.TileNum == p.uncompressedHeader.TileInfo.ContextUpdateTileId {
		// TODO: whatever is supposed to happen ehre
	}
}

// decode_tile()
func (t *TileGroup) decodeTile(p *parser.Parser) {
	p.clearAboveContext()

	for i := 0; i < FRAME_LF_COUNT; i++ {
		p.DeltaLF = SliceAssign(p.DeltaLF, i, 0)
	}

	for plane := 0; plane < p.sequenceHeader.ColorConfig.NumPlanes; plane++ {
		for pass := 0; pass < 2; pass++ {
			t.RefSgrXqd = SliceAssignNested(t.RefSgrXqd, plane, pass, Sgrproj_Xqd_Mid[pass])

			for i := 0; i < WIENER_COEFFS; i++ {
				t.RefLrWiener[plane][pass][i] = Wiener_Taps_Mid[i]
			}
		}

	}
	sbSize := BLOCK_64X64
	if p.sequenceHeader.Use128x128SuperBlock {
		sbSize = BLOCK_128X128
	}

	sbSize4 := p.Num4x4BlocksWide[sbSize]

	for r := p.MiRowStart; r < p.MiRowEnd; r += sbSize4 {
		p.clearLeftContext()

		for c := p.MiColStart; c < p.MiColEnd; c += sbSize4 {
			p.ReadDeltas = p.uncompressedHeader.DeltaQPresent
			p.Cdef.clear_cdef(r, c, p)
			t.clearBlockDecodedFlags(r, c, sbSize, p)
			t.readLr(r, c, sbSize, p)
			t.decodePartition(r, c, sbSize, p)
		}
	}
}

// inter_frame_mode_info()
func (t *TileGroup) interFrameModeInfo(p *parser.Parser) {
	t.useIntrabc = 0

	if p.AvailL {
		t.LeftRefFrame[0] = p.RefFrames[p.MiRow][p.MiCol-1][0]
		t.LeftRefFrame[1] = p.RefFrames[p.MiRow][p.MiCol-1][1]
	} else {
		t.LeftRefFrame[0] = INTRA_FRAME
		t.LeftRefFrame[1] = NONE
	}

	if p.AvailU {
		t.AboveRefFrame[0] = p.RefFrames[p.MiRow-1][p.MiCol][0]
		t.AboveRefFrame[1] = p.RefFrames[p.MiRow-1][p.MiCol][1]
	} else {
		t.AboveRefFrame[0] = INTRA_FRAME
		t.AboveRefFrame[1] = NONE
	}

	t.LeftIntra = t.LeftRefFrame[0] <= INTRA_FRAME
	t.AboveIntra = t.AboveRefFrame[0] <= INTRA_FRAME
	t.LeftSingle = t.LeftRefFrame[1] <= INTRA_FRAME
	t.AboveSingle = t.AboveRefFrame[1] <= INTRA_FRAME

	t.Skip = 0
	t.interSegmentId(1, p)
	t.readSkipMode(p)

	if Bool(t.SkipMode) {
		t.Skip = 1
	} else {
		t.readSkip(p)
	}

	if !Bool(p.uncompressedHeader.SegIdPreSkip) {
		t.interSegmentId(0, p)
	}

	t.Lossless = p.uncompressedHeader.LosslessArray[t.SegmentId]
	t.readCdef(p)
	t.readDeltaQIndex(p)
	t.readDeltaLf(p)
	p.ReadDeltas = false
	t.readIsInter(p)

	if Bool(t.IsInter) {
		t.interBlockModeInfo(p)
	} else {
		t.intraBlockModeInfo(p)
	}
}

// intra_block_mode_info()
func (t *TileGroup) intraBlockModeInfo(p *parser.Parser) {
	p.RefFrame[0] = INTRA_FRAME
	p.RefFrame[1] = NONE
	yMode := p.S()
	t.YMode = yMode
	t.intraAngleInfoY(p)

	if t.HasChroma {
		uvMode := p.S()
		t.UVMode = uvMode

		if t.UVMode == UV_CFL_PRED {
			t.readCflAlphas(p)
		}

		t.intraAngleInfoUv(p)
	}

	t.PaletteSizeY = 0
	t.PaletteSizeUV = 0
	if p.MiSize >= BLOCK_8X8 &&
		t.Block_Width[p.MiSize] <= 64 &&
		t.Block_Height[p.MiSize] <= 64 &&
		Bool(p.uncompressedHeader.AllowScreenContentTools) {
		t.paletteModeInfo(p)
	}

	t.filterIntraModeInfo(p)
}

// inter_block_mode_info()
func (t *TileGroup) interBlockModeInfo(p *parser.Parser) {
	t.PaletteSizeY = 0
	t.PaletteSizeUV = 0
	t.readRefFrames(p)

	isCompound := p.RefFrame[1] > INTRA_FRAME
	t.findMvStack(Int(isCompound), p)

	if Bool(t.SkipMode) {
		t.YMode = NEAREST_NEARESTMV
	} else if t.segFeatureActive(SEG_LVL_SKIP, p) || t.segFeatureActive(SEG_LVL_GLOBALMV, p) {
		t.YMode = GLOBALMV
	} else if isCompound {
		compoundMode := p.S()
		t.YMode = NEAREST_NEARESTMV + compoundMode
	} else {
		newMv := p.S()
		if newMv == 0 {
			t.YMode = NEWMV
		} else {
			zeroMv := p.S()
			if zeroMv == 0 {
				t.YMode = GLOBALMV
			} else {
				refMv := p.S()
				if refMv == 0 {
					t.YMode = NEARESTMV
				} else {
					t.YMode = NEARMV
				}
			}
		}
	}

	t.RefMvIdx = 0
	if t.YMode == NEWMV || t.YMode == NEW_NEWMV {
		for idx := 0; idx < 2; idx++ {
			if t.NumMvFound > idx+1 {
				drlMode := p.S()
				if drlMode == 0 {
					t.RefMvIdx = idx
					break
				}
				t.RefMvIdx = idx + 1
			}
		}
	} else if t.hasNearmv() {
		t.RefMvIdx = 1
		for idx := 1; idx < 3; idx++ {
			if t.NumMvFound > idx+1 {
				drlMode := p.S()
				if drlMode == 0 {
					t.RefMvIdx = idx
					break
				}
				t.RefMvIdx = idx + 1
			}
		}
	}

	t.assignMv(Int(isCompound), p)
	t.readInterIntraMode(isCompound, p)
	t.readMotionMode(isCompound, p)
	t.readCompoundType(isCompound, p)

	if p.uncompressedHeader.InterpolationFilter == SWITCHABLE {
		x := 1
		if p.sequenceHeader.EnableDualFilter {
			x = 2
		}
		for dir := 0; dir < x; dir++ {
			if t.needsInterpFilter(p) {
				t.InterpFilter[dir] = p.S()
			} else {
				t.InterpFilter[dir] = EIGHTTAP
			}
		}

		if !p.sequenceHeader.EnableDualFilter {
			t.InterpFilter[1] = t.InterpFilter[0]
		}
	} else {
		for dir := 0; dir < 2; dir++ {
			t.InterpFilter[dir] = p.uncompressedHeader.InterpolationFilter
		}
	}
}

// needs_interp_filter()
func (t *TileGroup) needsInterpFilter(p *parser.Parser) bool {
	large := Min(t.Block_Width[p.MiSize], t.Block_Height[p.MiSize]) >= 8

	if Bool(t.SkipMode) || t.MotionMode == LOCALWARP {
		return false
	} else if large && t.YMode == GLOBALMV {
		return p.GmType[p.RefFrame[0]] == TRANSLATION
	} else if large && t.YMode == GLOBAL_GLOBALMV {
		return p.GmType[p.RefFrame[0]] == TRANSLATION || p.GmType[1] == TRANSLATION
	} else {
		return true
	}
}

// read_compound_type( isCompound )
func (t *TileGroup) readCompoundType(isCompound bool, p *parser.Parser) {
	t.CompGroupIdx = 0
	t.CompoundIdx = 1
	if Bool(t.SkipMode) {
		t.CompoundType = COMPOUND_AVERAGE
		return
	}

	if isCompound {
		n := Wedge_Bits[p.MiSize]
		if p.sequenceHeader.EnableMaskedCompound {
			t.CompGroupIdx = p.S()
		}

		if t.CompGroupIdx == 0 {
			if p.sequenceHeader.EnableJntComp {
				t.CompoundIdx = p.S()
				if Bool(t.CompoundIdx) {
					t.CompoundType = COMPOUND_AVERAGE

				} else {
					t.CompoundType = COMPOUND_DISTANCE
				}
			} else {
				t.CompoundType = COMPOUND_AVERAGE
			}
		} else {
			if n == 0 {
				t.CompoundType = COMPOUND_DIFFWTD
			} else {
				t.CompoundType = p.S()
			}
		}

		if t.CompoundType == COMPOUND_WEDGE {
			t.WedgeIndex = p.S()
			t.WedgeIndex = p.L(1)
		} else if t.CompoundType == COMPOUND_DIFFWTD {
			t.MaskType = p.L(1)
		}
	} else {
		if Bool(t.InterIntra) {
			if Bool(t.WedgeInterIntra) {
				t.CompoundType = COMPOUND_WEDGE
			} else {
				t.CompoundType = COMPOUND_INTRA
			}
		} else {
			t.CompoundType = COMPOUND_AVERAGE
		}
	}

}

// read_motion_mode( isCompound )
func (t *TileGroup) readMotionMode(isCompound bool, p *parser.Parser) {
	if Bool(t.SkipMode) {
		t.MotionMode = SIMPLE
		return
	}

	if !p.uncompressedHeader.IsMotionModeSwitchable {
		t.MotionMode = SIMPLE
		return
	}

	if Min(t.Block_Width[p.MiSize], t.Block_Height[p.MiSize]) < 8 {
		t.MotionMode = SIMPLE
		return
	}

	if !p.uncompressedHeader.ForceIntegerMv && (t.YMode == GLOBALMV || t.YMode == GLOBAL_GLOBALMV) {
		if p.GmType[p.RefFrame[0]] > TRANSLATION {
			t.MotionMode = SIMPLE
			return
		}
	}

	t.findWarpSamples(p)
	if p.uncompressedHeader.ForceIntegerMv || t.NumSamples == 0 || !p.uncompressedHeader.AllowWarpedMotion || t.isScaled(p.RefFrame[0], p) {
		useObmc := p.S()
		if Bool(useObmc) {
			t.MotionMode = OBMC

		} else {
			t.MotionMode = SIMPLE
		}
	} else {
		t.MotionMode = p.S()
	}
}

// is_scaled( refFrame )
func (t *TileGroup) isScaled(refFrame int) bool {
	refIdx := p.uncompressedHeader.ref_frame_idx[refFrame-LAST_FRAME]
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameWidth / 2)) / p.uncompressedHeader.FrameWidth
	yScale := ((t.RefUpscaledHeight[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameHeight / 2)) / p.uncompressedHeader.FrameHeight
	noScale := 1 << REF_SCALE_SHIFT

	return xScale != noScale || yScale != noScale
}

// find_warp_samples() 7.10.4.
func (t *TileGroup) findWarpSamples(p *parser.Parser) {
	t.NumSamples = 0
	t.NumSamplesScanned = 0

	w4 := p.Num4x4BlocksWide[p.MiSize]
	h4 := p.Num4x4BlocksHigh[p.MiSize]

	doTopLeft := 1
	doTopRight := 1

	if p.AvailU {
		srcSize := p.MiSizes[p.MiRow-1][p.MiCol]
		srcW := p.Num4x4BlocksWide[srcSize]

		if w4 <= srcW {
			colOffset := -(p.MiCol & (srcW - 1))
			if colOffset < 0 {
				doTopLeft = 0
			}
			if colOffset+srcW > w4 {
				doTopRight = 0
			}
			t.addSample(-1, 0, p)
		} else {
			var miStep int
			for i := 0; i < Min(w4, p.MiCols-p.MiCol); i += miStep {
				srcSize = p.MiSizes[p.MiRow-1][p.MiCol+i]
				srcW = p.Num4x4BlocksWide[srcSize]
				miStep = Min(w4, srcW)
				t.addSample(-1, i, p)
			}
		}
	}
	if p.AvailL {
		srcSize := p.MiSizes[p.MiRow][p.MiCol-1]
		srcH := p.Num4x4BlocksHigh[srcSize]

		if h4 <= srcH {
			rowOffset := -(p.MiRow & (srcH - 1))
			if rowOffset < 0 {
				doTopLeft = 0
			}
			t.addSample(0, -1, p)
		} else {
			var miStep int
			for i := 0; i < Min(h4, p.MiRows-p.MiRow); i += miStep {
				srcSize = p.MiSizes[p.MiRow+i][p.MiCol-1]
				srcH = p.Num4x4BlocksHigh[srcSize]
				miStep = Min(h4, srcH)
				t.addSample(i, -1, p)
			}
		}
	}

	if Bool(doTopLeft) {
		t.addSample(-1, -1, p)
	}

	if Bool(doTopRight) {
		if Max(w4, h4) <= 16 {
			t.addSample(-1, w4, p)
		}
	}

	if t.NumSamples == 0 && t.NumSamplesScanned > 0 {
		t.NumSamples = 1
	}

}

// add_sample 7.10.4.2.
func (t *TileGroup) addSample(deltaRow int, deltaCol int, p *parser.Parser) {
	if t.NumSamplesScanned >= LEAST_SQUARES_SAMPLES_MAX {
		return
	}

	mvRow := p.MiRow + deltaRow
	mvCol := p.MiCol + deltaCol

	if !p.isInside(mvRow, mvCol) {
		return
	}

	// TODO: how do we know if something has not been writte to?
	if p.RefFrames[mvRow][mvCol][0] == 0 {
		return
	}

	if p.RefFrames[mvRow][mvCol][0] != p.RefFrame[0] {
		return
	}

	if p.RefFrames[mvRow][mvCol][1] != NONE {
		return
	}

	candSz := p.MiSizes[mvRow][mvCol]
	candW4 := p.Num4x4BlocksWide[candSz]
	candH4 := p.Num4x4BlocksHigh[candSz]
	candRow := mvRow & ^(candH4 - 1)
	candCol := mvCol & ^(candW4 - 1)
	midY := candRow*4 + candH4*2 - 1
	midX := candCol*4 + candW4*2 - 1
	threshold := Clip3(16, 112, Max(t.Block_Width[p.MiSize], t.Block_Height[p.MiSize]))
	mvDiffRow := Abs(t.Mvs[candRow][candCol][0][0] - t.Mv[0][0])
	mvDiffCol := Abs(t.Mvs[candRow][candCol][0][1] - t.Mv[0][1])
	valid := (mvDiffRow + mvDiffCol) <= threshold

	var cand []int
	cand[0] = midY * 8
	cand[1] = midX * 8
	cand[2] = midY*8 + t.Mvs[candRow][candCol][0][0]
	cand[3] = midX*8 + t.Mvs[candRow][candCol][0][1]

	t.NumSamplesScanned++
	if valid && t.NumSamplesScanned > 1 {
		return
	}

	for j := 0; j < 4; j++ {
		t.CandList[t.NumSamples][j] = cand[j]
	}

	if valid {
		t.NumSamples++
	}
}

// read_interintra_mode( isCompound )
func (t *TileGroup) readInterIntraMode(isCompound bool, p *parser.Parser) {
	if Bool(t.SkipMode) && p.sequenceHeader.EnableInterIntraCompound && !isCompound && p.MiSize > +BLOCK_8X8 && p.MiSize <= BLOCK_32X32 {
		t.InterIntra = p.S()

		if Bool(t.InterIntra) {
			t.InterIntraMode = p.S()
			p.RefFrame[1] = INTRA_FRAME
			t.AngleDeltaY = 0
			t.AngleDeltaUV = 0
			t.UseFilterIntra = 0
			t.WedgeInterIntra = p.S()
			if Bool(t.WedgeInterIntra) {
				t.WedgeIndex = p.S()
				t.WedgeSign = 0
			}
		}
	} else {
		t.InterIntra = 0
	}
}

// has_nearmv()
func (t *TileGroup) hasNearmv() bool {
	return t.YMode == NEARMV || t.YMode == NEAR_NEARMV || t.YMode == NEAR_NEWMV || t.YMode == NEW_NEARMV
}

// read_ref_frames()
func (t *TileGroup) readRefFrames(p *parser.Parser) {
	if Bool(t.SkipMode) {
		p.RefFrame[0] = p.uncompressedHeader.SkipModeFrame[0]
		p.RefFrame[1] = p.uncompressedHeader.SkipModeFrame[1]
	} else if t.segFeatureActive(SEG_LVL_REF_FRAME, p) {
		p.RefFrame[0] = p.FeatureData[t.SegmentId][SEG_LVL_REF_FRAME]
		p.RefFrame[1] = NONE
	} else {
		bw4 := p.Num4x4BlocksWide[p.MiSize]
		bh4 := p.Num4x4BlocksHigh[p.MiSize]

		var compMode int
		if p.uncompressedHeader.ReferenceSelect && Min(bw4, bh4) >= 2 {
			compMode = p.S()
		} else {
			compMode = SINGLE_REFERENCE
		}

		if compMode == COMPOUND_REFERENCE {
			compRefType := p.S()
			if compRefType == UNIDIR_COMP_REFERENCE {
				uniCompRef := p.S()
				if Bool(uniCompRef) {
					p.RefFrame[0] = BWDREF_FRAME
					p.RefFrame[1] = ALTREF_FRAME
				} else {
					uniCompRefP1 := p.S()
					if Bool(uniCompRefP1) {
						uniCompRefP2 := p.S()

						if Bool(uniCompRefP2) {
							p.RefFrame[0] = LAST_FRAME
							p.RefFrame[1] = GOLDEN_FRAME
						} else {
							p.RefFrame[0] = LAST_FRAME
							p.RefFrame[1] = LAST3_FRAME
						}
					} else {
						p.RefFrame[0] = LAST_FRAME
						p.RefFrame[1] = LAST2_FRAME

					}
				}
			} else {
				compRef := p.S()
				if compRef == 0 {
					compRefP1 := p.S()

					if Bool(compRefP1) {
						p.RefFrame[0] = LAST2_FRAME
					} else {
						p.RefFrame[0] = LAST_FRAME

					}
				} else {
					compRefP2 := p.S()

					if Bool(compRefP2) {
						p.RefFrame[0] = GOLDEN_FRAME
					} else {
						p.RefFrame[0] = LAST3_FRAME

					}

				}

				compBwdref := p.S()
				if compBwdref == 0 {
					compBwdrefP1 := p.S()

					if Bool(compBwdrefP1) {
						p.RefFrame[1] = ALTREF2_FRAME
					} else {
						p.RefFrame[1] = BWDREF_FRAME

					}
				} else {
					p.RefFrame[1] = ALTREF_FRAME
				}
			}
		} else {
			singleRefP1 := p.S()
			if Bool(singleRefP1) {
				singleRefP2 := p.S()
				if singleRefP2 == 0 {
					singleRefP6 := p.S()
					if Bool(singleRefP6) {
						p.RefFrame[0] = ALTREF2_FRAME
					} else {
						p.RefFrame[0] = BWDREF_FRAME

					}
				} else {
					p.RefFrame[0] = ALTREF_FRAME
				}
			} else {
				singleRefP3 := p.S()
				if Bool(singleRefP3) {
					singleRefP5 := p.S()
					if Bool(singleRefP5) {
						p.RefFrame[0] = GOLDEN_FRAME
					} else {
						p.RefFrame[0] = LAST3_FRAME
					}
				} else {
					singleRefP4 := p.S()
					if Bool(singleRefP4) {
						p.RefFrame[0] = LAST2_FRAME
					} else {
						p.RefFrame[0] = LAST_FRAME
					}
				}
			}
			p.RefFrame[1] = NONE
		}
	}
}

// read_is_inter()
func (t *TileGroup) readIsInter(p *parser.Parser) {
	if Bool(t.SkipMode) {
		t.IsInter = 1
	} else if t.segFeatureActive(SEG_LVL_REF_FRAME, p) {
		t.IsInter = Int(p.FeatureData[t.SegmentId][SEG_LVL_REF_FRAME] != INTRA_FRAME)
	} else if t.segFeatureActive(SEG_LVL_GLOBALMV, p) {
		t.IsInter = 0
	} else {
		t.IsInter = p.S()
	}
}

// inter_segment_id( preSkip )
func (t *TileGroup) interSegmentId(preSkip int, p *parser.Parser) {
	if Bool(p.uncompressedHeader.SegmentationEnabled) {
		predictedSegmentId := t.getSegmentId(p)

		if Bool(p.uncompressedHeader.SegmentationUpdateMap) {
			if Bool(preSkip) && !Bool(p.uncompressedHeader.SegIdPreSkip) {
				t.SegmentId = 0
				return
			}
			if !Bool(preSkip) {
				if Bool(t.Skip) {
					segIdPredicted := 0

					for i := 0; i < p.Num4x4BlocksWide[p.MiSize]; i++ {
						t.AboveSegPredContext[p.MiCol+i] = segIdPredicted
					}
					for i := 0; i < p.Num4x4BlocksHigh[p.MiSize]; i++ {
						t.AboveSegPredContext[p.MiRow+i] = segIdPredicted
					}
					t.readSegmentId(p)
					return
				}
			}

			if p.uncompressedHeader.SegmentationTemporalUpdate == 1 {
				segIdPredicted := p.S()
				if Bool(segIdPredicted) {
					t.SegmentId = predictedSegmentId
				} else {
					t.readSegmentId(p)
				}

				for i := 0; i < p.Num4x4BlocksWide[p.MiSize]; i++ {
					t.AboveSegPredContext[p.MiCol+i] = segIdPredicted
				}
				for i := 0; i < p.Num4x4BlocksHigh[p.MiSize]; i++ {
					t.AboveSegPredContext[p.MiRow+i] = segIdPredicted
				}

			} else {
				t.readSegmentId(p)
			}
		} else {
			t.SegmentId = predictedSegmentId
		}
	} else {
		t.SegmentId = 0
	}
}

// read_skip_mode()
func (t *TileGroup) readSkipMode(p *parser.Parser) {
	if t.segFeatureActive(SEG_LVL_SKIP, p) || t.segFeatureActive(SEG_LVL_REF_FRAME, p) || t.segFeatureActive(SEG_LVL_GLOBALMV, p) || !Bool(p.uncompressedHeader.SkipModePresent) || t.Block_Width[p.MiSize] < 8 || t.Block_Height[p.MiSize] < 8 {
		t.SkipMode = 0
	} else {
		t.SkipMode = p.S()
	}
}

// get_segment_id( )
func (t *TileGroup) getSegmentId(p *parser.Parser) int {
	bw4 := p.Num4x4BlocksWide[p.MiSize]
	bh4 := p.Num4x4BlocksHigh[p.MiSize]
	xMis := Min(p.MiCols-p.MiCol, bw4)
	yMis := Min(p.MiRows-p.MiRow, bh4)
	seg := 7

	for y := 0; y < yMis; y++ {
		for x := 0; x < xMis; x++ {
			seg = Min(seg, p.PrevSegmentIds[p.MiRow+y][p.MiCol+x])
		}
	}

	return seg
}

// get_mode( refList )
func (t *TileGroup) getMode(refList int) int {
	var compMode int
	if refList == 0 {
		if t.YMode < NEAREST_NEARESTMV {
			compMode = t.YMode
		} else if t.YMode == NEW_NEWMV || t.YMode == NEW_NEARESTMV || t.YMode == NEW_NEARMV {
			compMode = NEWMV
		} else if t.YMode == NEAREST_NEARESTMV || t.YMode == NEAREST_NEWMV {
			compMode = NEARESTMV
		} else if t.YMode == NEAR_NEARMV || t.YMode == NEAR_NEWMV {
			compMode = NEARMV
		} else {
			compMode = GLOBALMV
		}
	} else {
		if t.YMode == NEW_NEWMV || t.YMode == NEAREST_NEWMV || t.YMode == NEAR_NEWMV {
			compMode = NEWMV
		} else if t.YMode == NEAREST_NEARESTMV || t.YMode == NEW_NEARESTMV {
			compMode = NEARMV
		} else if t.YMode == NEAR_NEARMV || t.YMode == NEW_NEARMV {
			compMode = NEARMV
		} else {
			compMode = GLOBALMV
		}
	}

	return compMode
}

// is_directional_mode( mode )
func (t *TileGroup) isDirectionalMode(mode int) bool {
	return (mode >= V_PRED) && (mode <= D67_PRED)
}

// clear_block_decoded_flags( r, c, sbSize4 )
func (t *TileGroup) clearBlockDecodedFlags(r int, c int, sbSize4 int, p *parser.Parser) {
	for plane := 0; plane < p.sequenceHeader.ColorConfig.NumPlanes; plane++ {
		subX := 0
		subY := 0
		if plane > 0 {
			if p.sequenceHeader.ColorConfig.SubsamplingX {
				subX = 1
			}
			if p.sequenceHeader.ColorConfig.SubsamplingY {
				subY = 1
			}
		}

		sbWidth4 := (p.MiColEnd - c) >> subX
		sbHeight4 := (p.MiRowEnd - r) >> subY

		for y := -1; y <= (sbSize4 >> subY); y++ {
			for x := -1; x <= (sbSize4 >> subX); x++ {

				if y < 0 && x < sbWidth4 {
					p.BlockDecoded[plane][y][x] = 1
				} else if x < 0 && y < sbHeight4 {
					p.BlockDecoded[plane][y][x] = 1
				} else {
					p.BlockDecoded[plane][y][x] = 0
				}
			}
		}
		lastElement := len(p.BlockDecoded[plane][sbSize4>>subY])
		p.BlockDecoded[plane][sbSize4>>subY][lastElement] = 0
	}

}
