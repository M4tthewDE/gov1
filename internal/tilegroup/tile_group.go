package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

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
	Skips               [][]int
	SkipModes           [][]int
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
	MotionFieldMvs      [][][][]int
	Block_Width         []int
	Block_Height        []int
	IsInters            [][]int
	Mv                  [][]int
	Mvs                 [][][][]int
	FoundMatch          int
	TotalMatches        int
	CloseMatches        int
	RefStackMv          [][][]int
	RefIdCount          []int
	RefIdMvs            [][][]int
	RefDiffCount        []int
	RefDiffMvs          [][][]int
	RefFrameSignBias    []int
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
	LeftRefFrame        [2]int
	AboveRefFrame       [2]int
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
	TxSizes      [][]int

	AboveLevelContext  [][]int
	AboveDcContext     [][]int
	LeftLevelContext   [][]int
	LeftDcContext      [][]int
	LeftSegPredContext []int

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
	LocalWarpParams [6]int

	FrameStore [][][][]int
	Mask       [][]int

	FwdWeight int
	BckWeight int

	ZeroMvContext int
	DrlCtxStack   []int

	MaxLumaH int
	MaxLumaW int

	Quant             []int
	Dequant           [][]int
	TxType            int
	PlaneTxType       int
	TxTypes           [][]int
	T                 []int
	Residual          [][]int
	LoopFilterTxSizes [][][]int
}

func NewTileGroup(sz int, b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) TileGroup {
	t := TileGroup{}
	t.build(sz, b, state, uh, sh)
	return t
}

func (t *TileGroup) build(sz int, b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	NumTiles := state.TileCols * state.TileRows
	startbitPos := b.Position
	tileStartAndEndPresentFlag := false

	if NumTiles > 1 {
		tileStartAndEndPresentFlag = util.Bool(b.F(1))
	}

	var tgStart int
	var tgEnd int
	var tileBits int

	if NumTiles == 1 || !tileStartAndEndPresentFlag {
		tgStart = 0
		tgEnd = NumTiles - 1
	} else {
		tileBits = state.TileColsLog2 + state.TileRowsLog2
		tgStart = b.F(tileBits)
		tgEnd = b.F(tileBits)
	}

	b.ByteAlignment()
	endBitBos := b.Position
	headerBytes := (endBitBos - startbitPos) / 8
	sz -= headerBytes

	for state.TileNum = tgStart; state.TileNum <= tgEnd; state.TileNum++ {
		tileRow := state.TileNum / state.TileCols
		tileCol := state.TileNum % state.TileCols
		lastTile := state.TileNum == tgEnd

		var tileSize int
		if lastTile {
			tileSize = sz
		} else {
			tileSizeMinusOne := b.Le(state.TileSizeBytes)
			tileSize = tileSizeMinusOne + 1
			sz -= tileSize + state.TileSizeBytes
		}

		state.MiRowStart = state.MiRowStarts[tileRow]
		state.MiRowEnd = state.MiRowStarts[tileRow+1]
		state.MiColStart = state.MiColStarts[tileCol]
		state.MiColEnd = state.MiColStarts[tileCol+1]
		state.CurrentQIndex = uh.BaseQIdx
		t.initSymbol(tileSize, b, state)
		t.decodeTile(b, state, sh, uh)
		t.exitSymbol(b, state, uh)
	}

	if tgEnd == NumTiles-1 {
		if !uh.DisableFrameEndUpdateCdf {
			t.framEndUpdateCdf()
		}
		t.decodeFrameWrapup(state, uh)
		state.SeenFrameHeader = false
	}
}

// decode_frame_wrapup( )
func (t *TileGroup) decodeFrameWrapup(state *state.State, uh uncompressedheader.UncompressedHeader) {
	if !uh.ShowExistingFrame {
		if uh.LoopFilterLevel[0] != 0 || uh.LoopFilterLevel[1] != 0 {
			t.loopFilterProcess()
		}

		state.CdefFrame = t.cdefProcess()
		state.UpscaledCurrFrame = t.upscalingProcess()
		state.UpscaledCurrFrame = t.upscalingProcess()
		state.LrFrame = t.loopRestorationProcess()

		if uh.SegmentationEnabled && uh.SegmentationUpdateMap == 0 {
			for row := 0; row < state.MiRows; row++ {
				for col := 0; col < state.MiCols; col++ {
					t.SegmentIds[row][col] = state.PrevSegmentIds[row][col]
				}
			}
		}
	} else {
		if uh.FrameType == shared.KEY_FRAME {
			t.referenceFrameLoadingProcess()
		}
	}

	t.referenceFrameUpdateProcess()
	if uh.ShowFrame || uh.ShowExistingFrame {
		t.outputProcess()
	}
}

// frame_end_update_cdf( )
func (t *TileGroup) framEndUpdateCdf() {
	panic("not implemented")
}

// init_symbol( sz )
func (t *TileGroup) initSymbol(sz int, b *bitstream.BitStream, state *state.State) {
	numBits := util.Min(sz*8, 15)
	buf := b.F(numBits)
	paddedBuf := buf << (15 - numBits)

	state.SymbolValue = ((1 << 15) - 1) ^ paddedBuf
	state.SymbolRange = 1 << 15
	state.SymbolMaxBits = 8*sz - 15
	state.TileIntraFrameYModeCdf = uncompressedheader.DEFAULT_INTRA_FRAME_Y_MODE_CDF

	state.TileYModeCdf = state.YModeCdf
	state.TileUVModeCflNotAllowedCdf = state.UVModeCflNotAllowedCdf
	state.TileUVModeCflAllowedCdf = state.UVModeCflAllowedCdf
	state.TileAngleDeltaCdf = state.AngleDeltaCdf
	state.TileIntrabcCdf = state.IntrabcCdf
	state.TilePartitionW8Cdf = state.PartitionW8Cdf
	state.TilePartitionW16Cdf = state.PartitionW16Cdf
	state.TilePartitionW32Cdf = state.PartitionW32Cdf
	state.TilePartitionW64Cdf = state.PartitionW64Cdf
	state.TilePartitionW128Cdf = state.PartitionW128Cdf
	state.TileSegmentIdCdf = state.SegmentIdCdf
	state.TileSegmentIdPredictedCdf = state.SegmentIdPredictedCdf
	state.TileTx8x8Cdf = state.Tx8x8Cdf
	state.TileTx16x16Cdf = state.Tx16x16Cdf
	state.TileTx32x32Cdf = state.Tx32x32Cdf
	state.TileTx64x64Cdf = state.Tx64x64Cdf
	state.TileTxfmSplitCdf = state.TxfmSplitCdf
	state.TileFilterIntraModeCdf = state.FilterIntraModeCdf
	state.TileFilterIntraCdf = state.FilterIntraCdf
	state.TileInterpFilterCdf = state.InterpFilterCdf
	state.TileMotionModeCdf = state.MotionModeCdf
	state.TileNewMvCdf = state.NewMvCdf
	state.TileZeroMvCdf = state.ZeroMvCdf
	state.TileRefMvCdf = state.RefMvCdf
	state.TileCompoundModeCdf = state.CompoundModeCdf
	state.TileDrlModeCdf = state.DrlModeCdf
	state.TileIsInterCdf = state.IsInterCdf
	state.TileCompModeCdf = state.CompModeCdf
	state.TileSkipModeCdf = state.SkipModeCdf
	state.TileSkipCdf = state.SkipCdf
	state.TileCompRefCdf = state.CompRefCdf
	state.TileCompBwdRefCdf = state.CompBwdRefCdf
	state.TileSingleRefCdf = state.SingleRefCdf
	state.TileMvJointCdf = state.MvJointCdf
	state.TileMvClassCdf = state.MvClassCdf
	state.TileMvClass0BitCdf = state.MvClass0BitCdf
	state.TileMvFrCdf = state.MvFrCdf
	state.TileMvClass0FrCdf = state.MvClass0FrCdf
	state.TileMvClass0HpCdf = state.MvClass0HpCdf
	state.TileMvSignCdf = state.MvSignCdf
	state.TileMvBitCdf = state.MvBitCdf
	state.TileMvHpCdf = state.MvHpCdf
	state.TilePaletteYModeCdf = state.PaletteYModeCdf
	state.TilePaletteUVModeCdf = state.PaletteUVModeCdf
	state.TilePaletteUVSizeCdf = state.PaletteUVSizeCdf
	state.TilePaletteSize2YColorCdf = state.PaletteSize2YColorCdf
	state.TilePaletteSize2UVColorCdf = state.PaletteSize2UVColorCdf
	state.TilePaletteSize3YColorCdf = state.PaletteSize3YColorCdf
	state.TilePaletteSize3UVColorCdf = state.PaletteSize3UVColorCdf
	state.TilePaletteSize4YColorCdf = state.PaletteSize4YColorCdf
	state.TilePaletteSize4UVColorCdf = state.PaletteSize4UVColorCdf
	state.TilePaletteSize5YColorCdf = state.PaletteSize5YColorCdf
	state.TilePaletteSize5UVColorCdf = state.PaletteSize5UVColorCdf
	state.TilePaletteSize6YColorCdf = state.PaletteSize6YColorCdf
	state.TilePaletteSize6UVColorCdf = state.PaletteSize6UVColorCdf
	state.TilePaletteSize7YColorCdf = state.PaletteSize7YColorCdf
	state.TilePaletteSize7UVColorCdf = state.PaletteSize7UVColorCdf
	state.TilePaletteSize8YColorCdf = state.PaletteSize8YColorCdf
	state.TilePaletteSize8UVColorCdf = state.PaletteSize8UVColorCdf

	state.TileDeltaQCdf = state.DeltaQCdf
	state.TileDeltaLFCdf = state.DeltaLFCdf
	state.TileDeltaLFMultiCdf = state.DeltaLFMultiCdf
	state.TileIntraTxTypeSet1Cdf = state.IntraTxTypeSet1Cdf
	state.TileIntraTxTypeSet2Cdf = state.IntraTxTypeSet2Cdf
	state.TileInterTxTypeSet1Cdf = state.InterTxTypeSet1Cdf
	state.TileInterTxTypeSet2Cdf = state.InterTxTypeSet2Cdf
	state.TileInterTxTypeSet3Cdf = state.InterTxTypeSet3Cdf
	state.TileUseObmcCdf = state.UseObmcCdf
	state.TileInterIntraCdf = state.InterIntraCdf
	state.TileCompRefTypeCdf = state.CompRefTypeCdf
	state.TileCflSignCdf = state.CflSignCdf
	state.TileUniCompRefCdf = state.UniCompRefCdf
	state.TileWedgeInterIntraCdf = state.WedgeInterIntraCdf
	state.TileCompGroupIdxCdf = state.CompGroupIdxCdf
	state.TileCompoundIdxCdf = state.CompoundIdxCdf
	state.TileCompoundTypeCdf = state.CompoundTypeCdf
	state.TileInterIntraModeCdf = state.InterIntraModeCdf
	state.TileWedgeIndexCdf = state.WedgeIndexCdf
	state.TileCflAlphaCdf = state.CflAlphaCdf
	state.TileUseWienerCdf = state.UseWienerCdf
	state.TileUseSgrprojCdf = state.UseSgrprojCdf
	state.TileRestorationTypeCdf = state.RestorationTypeCdf

	state.TileTxbSkipCdf = state.TxbSkipCdf
	state.TileEobPt16Cdf = state.EobPt16Cdf
	state.TileEobPt32Cdf = state.EobPt32Cdf
	state.TileEobPt64Cdf = state.EobPt64Cdf
	state.TileEobPt128Cdf = state.EobPt128Cdf
	state.TileEobPt256Cdf = state.EobPt256Cdf
	state.TileEobPt512Cdf = state.EobPt512Cdf
	state.TileEobPt1024Cdf = state.EobPt1024Cdf
	state.TileEobExtraCdf = state.EobExtraCdf
	state.TileDcSignCdf = state.DcSignCdf
	state.TileCoeffBaseEobCdf = state.CoeffBaseEobCdf
	state.TileCoeffBaseCdf = state.CoeffBaseCdf
	state.TileCoeffBrCdf = state.CoeffBrCdf
}

// 8.2.4 Exit process for symbol decoder
func (t *TileGroup) exitSymbol(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	if state.SymbolMaxBits < -14 {
		panic("Violating bitstream conformance!")
	}

	b.Position += util.Max(0, state.SymbolMaxBits)

	if !uh.DisableFrameEndUpdateCdf && state.TileNum == uh.TileInfo.ContextUpdateTileId {
		// TODO: whatever is supposed to happen here
	}
}

// clear_above_context()
func (t *TileGroup) clearAboveContext(state *state.State) {
	t.AboveLevelContext = make([][]int, 3)
	t.AboveDcContext = make([][]int, 3)
	t.AboveSegPredContext = make([]int, 3)

	for i := 0; i < 3; i++ {
		t.AboveLevelContext[i] = make([]int, state.MiCols)
		t.AboveDcContext[i] = make([]int, state.MiCols)
	}

	for i := 0; i < state.MiCols; i++ {
		for plane := 0; plane <= 2; plane++ {
			t.AboveLevelContext[plane][i] = 0
			t.AboveDcContext[plane][i] = 0
			t.AboveSegPredContext[i] = 0
		}
	}
}

// clear_left_context( )
func (t *TileGroup) clearLeftContext(state *state.State) {
	t.LeftLevelContext = make([][]int, 3)
	t.LeftDcContext = make([][]int, 3)
	t.LeftSegPredContext = make([]int, 3)

	for i := 0; i < 3; i++ {
		t.LeftLevelContext[i] = make([]int, state.MiRows)
		t.LeftDcContext[i] = make([]int, state.MiRows)
	}

	for i := 0; i < state.MiRows; i++ {
		for plane := 0; plane <= 2; plane++ {
			t.LeftLevelContext[plane][i] = 0
			t.LeftDcContext[plane][i] = 0
			t.LeftSegPredContext[i] = 0
		}
	}
}

// decode_tile()
func (t *TileGroup) decodeTile(b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	t.clearAboveContext(state)

	for i := 0; i < shared.FRAME_LF_COUNT; i++ {
		state.DeltaLF[i] = 0
	}

	for plane := 0; plane < sh.ColorConfig.NumPlanes; plane++ {
		for pass := 0; pass < 2; pass++ {
			t.RefSgrXqd[plane][pass] = Sgrproj_Xqd_Mid[pass]

			for i := 0; i < WIENER_COEFFS; i++ {
				t.RefLrWiener[plane][pass][i] = Wiener_Taps_Mid[i]
			}
		}

	}
	sbSize := shared.BLOCK_64X64
	if sh.Use128x128SuperBlock {
		sbSize = shared.BLOCK_128X128
	}

	sbSize4 := state.Num4x4BlocksWide[sbSize]

	for r := state.MiRowStart; r < state.MiRowEnd; r += sbSize4 {
		t.clearLeftContext(state)

		for c := state.MiColStart; c < state.MiColEnd; c += sbSize4 {
			state.ReadDeltas = uh.DeltaQPresent
			state.Cdef.ClearCdef(r, c, sh.Use128x128SuperBlock, state.CdefSize4)
			t.clearBlockDecodedFlags(r, c, sbSize, state, sh)
			t.readLr(r, c, sbSize, b, state, uh, sh)
			t.decodePartition(r, c, sbSize, b, state, sh, uh)
		}
	}
}

// clear_block_decoded_flags( r, c, sbSize4 )
func (t *TileGroup) clearBlockDecodedFlags(r int, c int, sbSize4 int, state *state.State, sh sequenceheader.SequenceHeader) {
	for plane := 0; plane < sh.ColorConfig.NumPlanes; plane++ {
		subX := 0
		subY := 0
		if plane > 0 {
			if sh.ColorConfig.SubsamplingX {
				subX = 1
			}
			if sh.ColorConfig.SubsamplingY {
				subY = 1
			}
		}

		sbWidth4 := (state.MiColEnd - c) >> subX
		sbHeight4 := (state.MiRowEnd - r) >> subY

		for y := -1; y <= (sbSize4 >> subY); y++ {
			for x := -1; x <= (sbSize4 >> subX); x++ {

				if y < 0 && x < sbWidth4 {
					state.BlockDecoded[plane][y][x] = 1
				} else if x < 0 && y < sbHeight4 {
					state.BlockDecoded[plane][y][x] = 1
				} else {
					state.BlockDecoded[plane][y][x] = 0
				}
			}
		}
		lastElement := len(state.BlockDecoded[plane][sbSize4>>subY])
		state.BlockDecoded[plane][sbSize4>>subY][lastElement] = 0
	}

}
