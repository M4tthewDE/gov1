package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
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
	LocalWarpParams [6]int

	FrameStore [][][][]int
	Mask       [][]int

	FwdWeight int
	BckWeight int
}

func NewTileGroup(sz int, b *bitstream.BitStream, state State) TileGroup {
	t := TileGroup{}
	t.State = state
	t.build(sz, b)
	return t
}

func (t *TileGroup) build(sz int, b *bitstream.BitStream) {
	NumTiles := t.State.TileCols * t.State.TileRows
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
		tileBits = t.State.TileColsLog2 + t.State.TileRowsLog2
		tgStart = b.F(tileBits)
		tgEnd = b.F(tileBits)
	}

	b.ByteAlignment()
	endBitBos := b.Position
	headerBytes := (endBitBos - startbitPos) / 8
	sz -= headerBytes

	for t.State.TileNum = tgStart; t.State.TileNum <= tgEnd; t.State.TileNum++ {
		tileRow := t.State.TileNum / t.State.TileCols
		tileCol := t.State.TileNum % t.State.TileCols
		lastTile := t.State.TileNum == tgEnd

		var tileSize int
		if lastTile {
			tileSize = sz
		} else {
			tileSizeMinusOne := b.Le(t.State.TileSizeBytes)
			tileSize = tileSizeMinusOne + 1
			sz -= tileSize + t.State.TileSizeBytes
		}

		t.State.MiRowStart = t.State.MiRowStarts[tileRow]
		t.State.MiRowEnd = t.State.MiRowStarts[tileRow+1]
		t.State.MiColStart = t.State.MiColStarts[tileCol]
		t.State.MiColEnd = t.State.MiColStarts[tileCol+1]
		t.State.CurrentQIndex = t.State.UncompressedHeader.BaseQIdx
		t.initSymbol(tileSize)
		t.decodeTile(b)
	}
}

// init_symbol( sz )
func (t *TileGroup) initSymbol(sz int) {
	panic("not implemented init_symbol( sz )")
}

// 8.2.4 Exit process for symbol decoder
func (t *TileGroup) exitSymbol(b *bitstream.BitStream) {
	if t.State.SymbolMaxBits < -14 {
		panic("Violating bitstream conformance!")
	}

	b.Position += util.Max(0, t.State.SymbolMaxBits)

	if !t.State.UncompressedHeader.DisableFrameEndUpdateCdf && t.State.TileNum == t.State.UncompressedHeader.TileInfo.ContextUpdateTileId {
		// TODO: whatever is supposed to happen here
	}
}

// clear_above_context()
func (t *TileGroup) clearAboveContext() {
	panic("not implemented: clear_above_context()")
}

// clear_left_context( x )
func (t *TileGroup) clearLeftContext() {
	panic("not implemented: clear_left_context()")
}

// decode_tile()
func (t *TileGroup) decodeTile(b *bitstream.BitStream) {
	t.clearAboveContext()

	for i := 0; i < FRAME_LF_COUNT; i++ {
		t.State.DeltaLF[i] = 0
	}

	for plane := 0; plane < t.State.SequenceHeader.ColorConfig.NumPlanes; plane++ {
		for pass := 0; pass < 2; pass++ {
			t.RefSgrXqd[plane][pass] = Sgrproj_Xqd_Mid[pass]

			for i := 0; i < WIENER_COEFFS; i++ {
				t.RefLrWiener[plane][pass][i] = Wiener_Taps_Mid[i]
			}
		}

	}
	sbSize := shared.BLOCK_64X64
	if t.State.SequenceHeader.Use128x128SuperBlock {
		sbSize = shared.BLOCK_128X128
	}

	sbSize4 := t.State.Num4x4BlocksWide[sbSize]

	for r := t.State.MiRowStart; r < t.State.MiRowEnd; r += sbSize4 {
		t.clearLeftContext()

		for c := t.State.MiColStart; c < t.State.MiColEnd; c += sbSize4 {
			t.State.ReadDeltas = t.State.UncompressedHeader.DeltaQPresent
			t.State.Cdef.ClearCdef(r, c, t.State.SequenceHeader.Use128x128SuperBlock, t.State.CdefSize4)
			t.clearBlockDecodedFlags(r, c, sbSize)
			t.readLr(r, c, sbSize, b)
			t.decodePartition(r, c, sbSize, b)
		}
	}
}

// clear_block_decoded_flags( r, c, sbSize4 )
func (t *TileGroup) clearBlockDecodedFlags(r int, c int, sbSize4 int) {
	for plane := 0; plane < t.State.SequenceHeader.ColorConfig.NumPlanes; plane++ {
		subX := 0
		subY := 0
		if plane > 0 {
			if t.State.SequenceHeader.ColorConfig.SubsamplingX {
				subX = 1
			}
			if t.State.SequenceHeader.ColorConfig.SubsamplingY {
				subY = 1
			}
		}

		sbWidth4 := (t.State.MiColEnd - c) >> subX
		sbHeight4 := (t.State.MiRowEnd - r) >> subY

		for y := -1; y <= (sbSize4 >> subY); y++ {
			for x := -1; x <= (sbSize4 >> subX); x++ {

				if y < 0 && x < sbWidth4 {
					t.State.BlockDecoded[plane][y][x] = 1
				} else if x < 0 && y < sbHeight4 {
					t.State.BlockDecoded[plane][y][x] = 1
				} else {
					t.State.BlockDecoded[plane][y][x] = 0
				}
			}
		}
		lastElement := len(t.State.BlockDecoded[plane][sbSize4>>subY])
		t.State.BlockDecoded[plane][sbSize4>>subY][lastElement] = 0
	}

}
