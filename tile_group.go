package main

const FRAME_LF_COUNT = 4
const WIENER_COEFFS = 3

const BLOCK_INVALID = 3
const BLOCK_4x4 = 0
const BLOCK_4x8 = 1
const BLOCK_8x4 = 2
const BLOCK_8x8 = 3
const BLOCK_8x16 = 4
const BLOCK_16x8 = 5
const BLOCK_16x16 = 6
const BLOCK_16x32 = 7
const BLOCK_32x16 = 8
const BLOCK_32x32 = 9
const BLOCK_32x64 = 10
const BLOCK_64x32 = 11
const BLOCK_64x64 = 12
const BLOCK_64x128 = 13
const BLOCK_128x64 = 14
const BLOCK_128x128 = 15
const BLOCK_4x16 = 16
const BLOCK_16x4 = 17
const BLOCK_8x32 = 18
const BLOCK_32x8 = 19
const BLOCK_16x64 = 20
const BLOCK_64x16 = 21
const PARTITION_NONE = 0
const PARTITION_HORZ = 1
const PARTITION_VERT = 2
const PARTITION_SPLIT = 3

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

const LAST_FRAME = 1
const LAST2_FRAME = 2
const LAST3_FRAME = 3
const GOLDEN_FRAME = 4
const BWDREF_FRAME = 5
const ALTREF2_FRAME = 6
const ALTREF_FRAME = 7

var Wedge_Bits = []int{0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 0, 0}
var Palette_Color_Hash_Multipliers = []int{1, 2, 2}

const MAX_VARTX_DEPTH = 2

const TX_4X4 = 0
const TX_8X8 = 1
const TX_16X16 = 2
const TX_32X32 = 3
const TX_64X64 = 4
const TX_4X8 = 5
const TX_8X4 = 6
const TX_8X16 = 7
const TX_16X8 = 8
const TX_16X32 = 9
const TX_32X16 = 10
const TX_32X64 = 11
const TX_64X32 = 12
const TX_4X16 = 13
const TX_16X4 = 14
const TX_8X32 = 15
const TX_32X8 = 16
const TX_16X64 = 17
const TX_64X16 = 18

var Max_Tx_Size_Rect = []int{
	TX_4X4, TX_4X8, TX_8X4, TX_8X8,
	TX_8X16, TX_16X8, TX_16X16, TX_16X32,
	TX_32X16, TX_32X32, TX_32X64, TX_64X32,
	TX_64X64, TX_64X64, TX_64X64, TX_64X64,
	TX_4X16, TX_16X4, TX_8X32, TX_32X8,
	TX_16X64, TX_64X16,
}

var Max_Tx_Depth = []int{
	0, 1, 1, 1,
	2, 2, 2, 3,
	3, 3, 4, 4,
	4, 4, 4, 4,
	2, 2, 3, 3,
	4, 4,
}

var Split_Tx_Size = []int{
	TX_4X4,
	TX_4X4,
	TX_8X8,
	TX_16X16,
	TX_32X32,
	TX_4X4,
	TX_4X4,
	TX_8X8,
	TX_8X8,
	TX_16X16,
	TX_16X16,
	TX_32X32,
	TX_32X32,
	TX_4X8,
	TX_8X4,
	TX_8X16,
	TX_16X8,
	TX_16X32,
	TX_32X16,
}

var Tx_Width = []int{4, 8, 16, 32, 64, 4, 8, 8, 16, 16, 32, 32, 64, 4, 16, 8, 32, 16, 64}
var Tx_Height = []int{4, 8, 16, 32, 64, 8, 4, 16, 8, 32, 16, 64, 32, 16, 4, 32, 8, 64, 16}

var Partition_Subsize = [][]int{
	{
		BLOCK_4x4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128x128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4x4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32x8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4x16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8x32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16x64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
}

var Sgrproj_Xqd_Mid = []int{-32, 31}
var Sgrproj_Xqd_Min = []int{-96, -32}
var Sgrproj_Xqd_Max = []int{31, 95}
var Wiener_Taps_Mid = []int{3, -7, 15}
var Wiener_Taps_Min = []int{-5, -23, -17}
var Wiener_Taps_Max = []int{10, 8, 46}
var Wiener_Taps_K = []int{1, 2, 3}
var SgrParams = [][]int{}

const RESTORE_NONE = 0
const RESTORE_WIENER = 1
const RESTORE_SGRPROJ = 2
const RESTORE_SWITCHABLE = 3

const MI_SIZE = 4

const SGRPROJ_PARAMS_BITS = 4
const SGRPROJ_BITS = 7
const SGRPROJ_PRJ_SUBEXP_K = 4

const DC_PRED = 0

const SIMPLE = 0
const OBMC = 1
const LOCALWARP = 2

const COMPUND_AVERAGE = 2

const NEARESTMV = 14
const NEARMV = 15
const GLOBALMV = 16
const NEWMV = 17
const NEAREST_NEARESTMV = 18
const NEAR_NEARMV = 19
const NEAREST_NEWMV = 20
const NEW_NEARESTMV = 21
const NEAR_NEWMV = 22
const NEW_NEARMV = 23
const GLOBAL_GLOBALMV = 24
const NEW_NEWMV = 25

const MAX_REF_MV_STACK_SIZE = 8

const V_PRED = 1
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
}

func NewTileGroup(p *Parser, sz int) TileGroup {
	t := TileGroup{}
	t.build(p, sz)
	return t
}

func (t *TileGroup) build(p *Parser, sz int) {
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

	}
}

// decode_tile()
func (t *TileGroup) decodeTile(p *Parser) {
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
	sbSize := BLOCK_64x64
	if p.sequenceHeader.Use128x128SuperBlock {
		sbSize = BLOCK_128x128
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

// decode_partition(r, c, bSize)
func (t *TileGroup) decodePartition(r int, c int, bSize int, p *Parser) {
	if r >= p.MiRows || c >= p.MiCols {
		return
	}

	p.AvailU = p.isInside(r-1, c)
	p.AvailL = p.isInside(r, c-1)
	num4x4 := p.Num4x4BlocksWide[bSize]
	halfBlock4x4 := num4x4 >> 1
	quarterBlock4x4 := halfBlock4x4 >> 1
	hasRows := (r + halfBlock4x4) < p.MiRows
	hasCols := (c + halfBlock4x4) < p.MiCols

	var partition int
	if bSize < BLOCK_8x8 {
		partition = PARTITION_NONE
	} else if hasRows && hasCols {
		partition = p.S()
	} else if hasCols {
		splitOrHorz := p.S() != 0
		if splitOrHorz {
			partition = PARTITION_SPLIT
		} else {
			partition = PARTITION_HORZ
		}
	} else if hasRows {
		splitOrVert := p.S() != 0
		if splitOrVert {
			partition = PARTITION_SPLIT
		} else {
			partition = PARTITION_VERT
		}

	} else {
		partition = PARTITION_SPLIT
	}

	subSize := Partition_Subsize[partition][bSize]
	splitSize := Partition_Subsize[PARTITION_SPLIT][bSize]
	if partition == PARTITION_NONE {
		t.decodeBlock(r, c, subSize, p)
	}
}

// decode_block( r, c, subSize)
func (t *TileGroup) decodeBlock(r int, c int, subSize int, p *Parser) {
	p.MiRow = r
	p.MiCol = c
	p.MiSize = subSize
	bw4 := p.Num4x4BlocksWide[subSize]
	bh4 := p.Num4x4BlocksHigh[subSize]

	if bh4 == 1 && p.sequenceHeader.ColorConfig.SubsamplingY && (p.MiRow&1) == 0 {
		t.HasChroma = false
	} else if bw4 == 1 && p.sequenceHeader.ColorConfig.SubsamplingX && (p.MiCol&1) == 0 {
		t.HasChroma = false
	} else {
		t.HasChroma = p.sequenceHeader.ColorConfig.NumPlanes > 1
	}

	p.AvailU = p.isInside(r-1, c)
	p.AvailL = p.isInside(r, c-1)
	availUChroma := p.AvailU
	availLChroma := p.AvailL

	if t.HasChroma {
		if p.sequenceHeader.ColorConfig.SubsamplingY && bh4 == 1 {
			availUChroma = p.isInside(r-2, c)
		}
		if p.sequenceHeader.ColorConfig.SubsamplingX && bw4 == 1 {
			availLChroma = p.isInside(r, c-2)
		}
	} else {
		availUChroma = false
		availLChroma = false
	}

	t.modeInfo(p)
	t.paletteTokens(p)
	t.readBlockTxSize(p)

	if Bool(t.Skip) {
		t.resetBlockContext(bw4, bh4, p)
	}

	isCompound := p.RefFrame[1] > INTRA_FRAME

	for y := 0; y < bh4; y++ {
		for x := 0; x < bw4; x++ {
			t.YModes[r+y][c+x] = t.YMode

			if p.RefFrame[0] == INTRA_FRAME && t.HasChroma {
				t.UVModes[r+y][c+x] = t.UVMode
			}

			for refList := 0; refList < 2; refList++ {
				p.RefFrames[r+y][c+x][refList] = p.RefFrame[refList]
			}

			if Bool(t.IsInter) {
				if !Bool(t.useIntrabc) {
					t.CompGroupIdxs[r+y][c+x] = t.CompGroupIdx
					t.CompoundIdxs[r+y][c+x] = t.CompoundIdx
				}
				for dir := 0; dir < 2; dir++ {
					t.InterpFilters[r+y][c+x][dir] = t.InterpFilter[dir]
				}
				for refList := 0; refList < 1+Int(isCompound); refList++ {
					t.Mvs[r+y][c+x][refList] = t.Mv[refList]
				}
			}
		}
	}
}

// reset_block_context( bw4, bh4 )
func (t *TileGroup) resetBlockContext(bw4 int, bh4 int, p *Parser) {
	for plane := 0; plane < 1+2*Int(t.HasChroma); plane++ {
		subX := 0
		subY := 0
		if plane > 0 {
			subX = Int(p.sequenceHeader.ColorConfig.SubsamplingX)
			subY = Int(p.sequenceHeader.ColorConfig.SubsamplingY)
		}

		for i := p.MiCol >> subX; i < ((p.MiCol + bw4) >> subX); i++ {
			t.AboveLevelContext[plane][i] = 0
			t.AboveDcContext[plane][i] = 0
		}

		for i := p.MiRow >> subY; i < ((p.MiRow + bh4) >> subY); i++ {
			t.LeftLevelContext[plane][i] = 0
			t.LeftDcContext[plane][i] = 0
		}
	}

}

// read_block_tx_size()
func (t *TileGroup) readBlockTxSize(p *Parser) {
	bw4 := p.Num4x4BlocksWide[p.MiSize]
	bh4 := p.Num4x4BlocksHigh[p.MiSize]

	if p.uncompressedHeader.TxMode == TX_MODE_SELECT &&
		p.MiSize > BLOCK_4x4 &&
		Bool(t.IsInter) &&
		!Bool(t.Skip) &&
		!t.Lossless {
		maxTxSz := Max_Tx_Size_Rect[p.MiSize]
		txW4 := Tx_Width[maxTxSz] / MI_SIZE
		txH4 := Tx_Height[maxTxSz] / MI_SIZE

		for row := p.MiRow; row < p.MiRow+bh4; row += txH4 {
			for col := p.MiCol; col < p.MiCol+bw4; col += txW4 {
				t.readVarTxSize(row, col, maxTxSz, 0, p)
			}
		}
	} else {
		t.readTxSize(!Bool(t.Skip) || Bool(t.IsInter), p)
		for row := p.MiRow; row < p.MiRow+bh4; row++ {
			for col := p.MiCol; col < p.MiCol+bw4; col++ {
				t.InterTxSizes[row][col] = t.TxSize
			}
		}
	}

}

// read_tx_size( allowSelect )
func (t *TileGroup) readTxSize(allowSelect bool, p *Parser) {
	if t.Lossless {
		t.TxSize = TX_4X4
		return
	}

	maxRectTxSize := Max_Tx_Size_Rect[p.MiSize]
	// TODO: what is this for?
	//maxTxDepth := Max_Tx_Depth[p.MiSize]
	t.TxSize = maxRectTxSize

	if p.MiSize > BLOCK_4x4 && allowSelect && p.uncompressedHeader.TxMode == TX_MODE_SELECT {
		txDepth := p.S()
		for i := 0; i < txDepth; i++ {
			t.TxSize = Split_Tx_Size[t.TxSize]
		}
	}
}

// read_var_tx_size( row, col, txSz, depth )
func (t *TileGroup) readVarTxSize(row int, col int, txSz int, depth int, p *Parser) {
	if row >= p.MiRows || col >= p.MiCols {
		return
	}

	var txfmSplit int
	if txSz == TX_4X4 || depth == MAX_VARTX_DEPTH {
		txfmSplit = 0
	} else {
		txfmSplit = p.S()
	}

	w4 := Tx_Width[txSz] / MI_SIZE
	h4 := Tx_Height[txSz] / MI_SIZE

	if Bool(txfmSplit) {
		subTxSz := Split_Tx_Size[txSz]
		stepW := Tx_Width[subTxSz] / MI_SIZE
		stepH := Tx_Height[subTxSz] / MI_SIZE

		for i := 0; i < h4; i += stepH {
			for j := 0; j < w4; j += stepW {
				t.readVarTxSize(row+i, col+j, subTxSz, depth+1, p)
			}
		}
	} else {
		for i := 0; i < h4; i++ {
			for j := 0; j < w4; j++ {
				t.InterTxSizes[row+i][col+i] = txSz
			}
		}
		t.TxSize = txSz
	}
}

// palette_tokens()
func (t *TileGroup) paletteTokens(p *Parser) {
	blockHeight := t.Block_Height[p.MiSize]
	blockWidth := t.Block_Width[p.MiSize]
	onscreenHeight := Min(blockHeight, (p.MiRows-p.MiRow)*MI_SIZE)
	onscreenWidth := Min(blockWidth, (p.MiCols-p.MiCol)*MI_SIZE)

	if Bool(t.PaletteSizeY) {
		colorIndexMapY := p.NS(t.PaletteSizeY)
		t.ColorMapY[0][0] = colorIndexMapY

		for i := 1; i < onscreenHeight+onscreenWidth-1; i++ {
			for j := Min(i, onscreenWidth-1); j >= Max(0, i-onscreenHeight+1); j-- {
				t.getPaletteColorContext(t.ColorMapY, (i - j), j, t.PaletteSizeY, p)
				paletteColorIdxY := p.S()
				t.ColorMapY[i-j][j] = t.ColorOrder[paletteColorIdxY]
			}
		}
		for i := 0; i < onscreenHeight; i++ {
			for j := onscreenWidth; j < blockWidth; j++ {
				t.ColorMapY[i][j] = t.ColorMapY[i][onscreenWidth-1]
			}
		}
		for i := onscreenHeight; i < blockHeight; i++ {
			for j := 0; j < blockWidth; j++ {
				t.ColorMapY[i][j] = t.ColorMapY[onscreenHeight-1][j]
			}
		}
	}

	if Bool(t.PaletteSizeUV) {
		colorIndexMapUv := p.NS(t.PaletteSizeUV)
		t.ColorMapUV[0][0] = colorIndexMapUv
		blockHeight = blockHeight >> Int(p.sequenceHeader.ColorConfig.SubsamplingY)
		blockWidth = blockWidth >> Int(p.sequenceHeader.ColorConfig.SubsamplingX)
		onscreenHeight = onscreenHeight >> Int(p.sequenceHeader.ColorConfig.SubsamplingX)
		onscreenWidth = onscreenWidth >> Int(p.sequenceHeader.ColorConfig.SubsamplingX)

		if blockWidth < 4 {
			blockWidth += 2
			onscreenWidth += 2
		}

		if blockHeight < 4 {
			blockHeight += 2
			onscreenHeight += 2
		}
		for i := 1; i < onscreenHeight+onscreenWidth-1; i++ {
			for j := Min(i, onscreenWidth-1); j >= Max(0, i-onscreenHeight+1); j-- {
				t.getPaletteColorContext(t.ColorMapUV, (i - j), j, t.PaletteSizeUV, p)
				paletteColorIdxUv := p.S()
				t.ColorMapUV[i-j][j] = t.ColorOrder[paletteColorIdxUv]
			}
		}
		for i := 0; i < onscreenHeight; i++ {
			for j := onscreenWidth; j < blockWidth; j++ {
				t.ColorMapUV[i][j] = t.ColorMapUV[i][onscreenWidth-1]
			}
		}
		for i := onscreenHeight; i < blockHeight; i++ {
			for j := 0; j < blockWidth; j++ {
				t.ColorMapUV[i][j] = t.ColorMapUV[onscreenHeight-1][j]
			}
		}
	}
}

// get_palette_color_context( colorMap, r, c, n )
func (t *TileGroup) getPaletteColorContext(colorMap [][]int, r int, c int, n int, p *Parser) {
	var scores []int
	for i := 0; i < PALETTE_COLORS; i++ {
		scores[i] = 0
		t.ColorOrder[i] = i
	}

	var neighbor int
	if c > 0 {
		neighbor = colorMap[r][c-1]
		scores[neighbor] += 2
	}

	if r > 0 && c > 0 {
		neighbor = colorMap[r-1][c-1]
		scores[neighbor] += 1
	}
	if r > 0 {
		neighbor = colorMap[r-1][c]
		scores[neighbor] += 1
	}

	for i := 0; i < PALETTE_NUM_NEIGHBORS; i++ {
		maxScore := scores[i]
		maIdx := i
		for j := i + 1; j < n; j++ {
			if scores[j] > maxScore {
				maxScore = scores[j]
				maIdx = j
			}
		}
		if maIdx != i {
			maxScore = scores[maIdx]
			maxColorOrder := t.ColorOrder[maIdx]
			for k := maIdx; k > i; k-- {
				scores[k] = scores[k-1]
				t.ColorOrder[k] = t.ColorOrder[k-1]
			}
			scores[i] = maxScore
			t.ColorOrder[i] = maxColorOrder
		}
	}

	t.ColorContextHash = 0
	for i := 0; i < PALETTE_NUM_NEIGHBORS; i++ {
		t.ColorContextHash += scores[i] * Palette_Color_Hash_Multipliers[i]
	}
}

// mode_info()
func (t *TileGroup) modeInfo(p *Parser) {
	if p.uncompressedHeader.FrameIsIntra {
		t.intraFrameModeInfo(p)
	} else {
		t.interFrameModeInfo(p)
	}
}

// inter_frame_mode_info()
func (t *TileGroup) interFrameModeInfo(p *Parser) {
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
func (t *TileGroup) intraBlockModeInfo(p *Parser) {
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
	if p.MiSize >= BLOCK_8x8 &&
		t.Block_Width[p.MiSize] <= 64 &&
		t.Block_Height[p.MiSize] <= 64 &&
		Bool(p.uncompressedHeader.AllowScreenContentTools) {
		t.paletteModeInfo(p)
	}

	t.filterIntraModeInfo(p)
}

// inter_block_mode_info()
func (t *TileGroup) interBlockModeInfo(p *Parser) {
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
func (t *TileGroup) needsInterpFilter(p *Parser) bool {
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
func (t *TileGroup) readCompoundType(isCompound bool, p *Parser) {
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
func (t *TileGroup) readMotionMode(isCompound bool, p *Parser) {
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
func (t *TileGroup) isScaled(refFrame int, p *Parser) bool {
	refIdx := p.uncompressedHeader.ref_frame_idx[refFrame-LAST_FRAME]
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameWidth / 2)) / p.uncompressedHeader.FrameWidth
	yScale := ((t.RefUpscaledHeight[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameHeight / 2)) / p.uncompressedHeader.FrameHeight
	noScale := 1 << REF_SCALE_SHIFT

	return xScale != noScale || yScale != noScale
}

// find_warp_samples() 7.10.4.
func (t *TileGroup) findWarpSamples(p *Parser) {
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
func (t *TileGroup) addSample(deltaRow int, deltaCol int, p *Parser) {
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
func (t *TileGroup) readInterIntraMode(isCompound bool, p *Parser) {
	if Bool(t.SkipMode) && p.sequenceHeader.EnableInterIntraCompound && !isCompound && p.MiSize > +BLOCK_8x8 && p.MiSize <= BLOCK_32x32 {
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
func (t *TileGroup) readRefFrames(p *Parser) {
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
func (t *TileGroup) readIsInter(p *Parser) {
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
func (t *TileGroup) interSegmentId(preSkip int, p *Parser) {
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
func (t *TileGroup) readSkipMode(p *Parser) {
	if t.segFeatureActive(SEG_LVL_SKIP, p) || t.segFeatureActive(SEG_LVL_REF_FRAME, p) || t.segFeatureActive(SEG_LVL_GLOBALMV, p) || !Bool(p.uncompressedHeader.SkipModePresent) || t.Block_Width[p.MiSize] < 8 || t.Block_Height[p.MiSize] < 8 {
		t.SkipMode = 0
	} else {
		t.SkipMode = p.S()
	}
}

// get_segment_id( )
func (t *TileGroup) getSegmentId(p *Parser) int {
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

// intra_frame_mode_info()
func (t *TileGroup) intraFrameModeInfo(p *Parser) {
	t.Skip = 0
	if p.uncompressedHeader.SegIdPreSkip == 1 {
		t.intraSegmentId(p)
	}

	t.SkipMode = 0
	t.readSkip(p)

	if !Bool(p.uncompressedHeader.SegIdPreSkip) {
		t.intraSegmentId(p)
	}
	t.readCdef(p)
	t.readDeltaQIndex(p)
	t.readDeltaLf(p)

	p.ReadDeltas = false
	p.RefFrame[0] = INTRA_FRAME
	p.RefFrame[0] = NONE

	if p.uncompressedHeader.AllowIntraBc {
		t.useIntrabc = p.S()
	} else {
		t.useIntrabc = 0
	}

	if Bool(t.useIntrabc) {
		t.IsInter = -1
		t.YMode = DC_PRED
		t.UVMode = DC_PRED
		t.MotionMode = SIMPLE
		t.CompoundType = COMPUND_AVERAGE
		t.PaletteSizeY = 0
		t.PaletteSizeUV = 0
		t.InterpFilter[0] = BILINEAR
		t.InterpFilter[1] = BILINEAR
		t.findMvStack(0, p)
		t.assignMv(0, p)
	} else {
		t.IsInter = 0
		intraFrameYMode := p.S()
		t.YMode = intraFrameYMode
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

		if p.MiSize >= BLOCK_8x8 && t.Block_Width[p.MiSize] <= 64 && t.Block_Height[p.MiSize] <= 64 && Bool(p.uncompressedHeader.AllowScreenContentTools) {
			t.paletteModeInfo(p)
		}
		t.filterIntraModeInfo(p)
	}
}

// filter_intra_mode_info()
func (t *TileGroup) filterIntraModeInfo(p *Parser) {
	useFilterIntra := false
	if p.sequenceHeader.EnableFilterIntra && t.YMode == DC_PRED && t.PaletteSizeY == 0 && Max(t.Block_Width[p.MiSize], t.Block_Height[p.MiSize]) <= 32 {
		useFilterIntra = Bool(p.S())

		if useFilterIntra {
			t.FilterIntraMode = p.S()
		}
	}
}

// palette_mode_info()
func (t *TileGroup) paletteModeInfo(p *Parser) {
	// TODO: this is used for initilization of has_palette_y I think
	//bSizeCtx := Mi_Width_Log2[p.MiSize] + Mi_Height_Log2[p.MiSize] - 2

	if t.YMode == DC_PRED {
		hasPaletteY := p.S()

		if Bool(hasPaletteY) {
			paletteSizeYMinus2 := p.S()
			t.PaletteSizeY = paletteSizeYMinus2 + 2
			cacheN := t.getPaletteCache(0, p)
			idx := 0

			for i := 0; i < cacheN && idx < t.PaletteSizeY; i++ {
				usePaletteColorCacheY := p.L(1)

				if Bool(usePaletteColorCacheY) {
					t.PaletteColorsY[idx] = t.PaletteCache[i]
					idx++
				}
			}

			if idx < t.PaletteSizeY {
				t.PaletteColorsY[idx] = p.L(p.sequenceHeader.ColorConfig.BitDepth)
				idx++
			}

			var paletteBits int
			if idx < t.PaletteSizeY {
				minBits := p.sequenceHeader.ColorConfig.BitDepth - 1
				paletteNumExtraBitsY := p.L(2)
				paletteBits = minBits + paletteNumExtraBitsY
			}

			for idx < t.PaletteSizeY {
				paletteDeltaY := p.L(paletteBits)
				paletteDeltaY++
				t.PaletteColorsY[idx] = Clip1(t.PaletteColorsY[idx-1]+paletteDeltaY, p)
				rangE := (1 << p.sequenceHeader.ColorConfig.BitDepth) - t.PaletteColorsY[idx] - 1
				paletteBits = Min(paletteBits, CeilLog2(rangE))
				idx++
			}
			t.PaletteColorsY = Sort(t.PaletteColorsY, 0, t.PaletteSizeY-1)
		}
	}

	if t.HasChroma && t.UVMode == DC_PRED {
		hasPaletteUv := p.S()
		if Bool(hasPaletteUv) {
			paletteSizeUvMinus2 := p.S()
			t.PaletteSizeUV = paletteSizeUvMinus2 + 2
			cacheN := t.getPaletteCache(1, p)
			idx := 0

			for i := 0; i < cacheN && idx < t.PaletteSizeUV; i++ {
				usePaletteColorCacheU := p.L(1)

				if Bool(usePaletteColorCacheU) {
					t.PaletteColorsY[idx] = t.PaletteCache[i]
					idx++
				}
			}

			if idx < t.PaletteSizeUV {
				t.PaletteColorsU[idx] = p.L(p.sequenceHeader.ColorConfig.BitDepth)
				idx++
			}

			var paletteBits int
			if idx < t.PaletteSizeUV {
				minBits := p.sequenceHeader.ColorConfig.BitDepth - 3
				paletteNumExtraBitsU := p.L(2)
				paletteBits = minBits + paletteNumExtraBitsU
			}

			for idx < t.PaletteSizeUV {
				paletteDeltaU := p.L(paletteBits)
				t.PaletteColorsU[idx] = Clip1(t.PaletteColorsU[idx-1]+paletteDeltaU, p)
				rangE := (1 << p.sequenceHeader.ColorConfig.BitDepth) - t.PaletteColorsU[idx] - 1
				paletteBits = Min(paletteBits, CeilLog2(rangE))
				idx++
			}
			t.PaletteColorsU = Sort(t.PaletteColorsU, 0, t.PaletteSizeUV-1)

			deltaEncodePaletteColorsv := p.L(1)

			if Bool(deltaEncodePaletteColorsv) {
				minBits := p.sequenceHeader.ColorConfig.BitDepth - 4
				maxVal := 1 << p.sequenceHeader.ColorConfig.BitDepth
				paletteNumExtraBitsv := p.L(2)
				paletteBits = minBits + paletteNumExtraBitsv
				t.PaletteColorsV[0] = p.L(p.sequenceHeader.ColorConfig.BitDepth)

				for idx := 1; idx < t.PaletteSizeUV; idx++ {
					paletteDeltaV := p.L(paletteBits)
					if Bool(paletteDeltaV) {
						paletteDeltaSignBitV := p.L(1)
						if Bool(paletteDeltaSignBitV) {
							paletteDeltaV = -paletteDeltaV
						}
					}

					val := t.PaletteColorsV[idx-1] + paletteDeltaV
					if val < 0 {
						val += maxVal
					}
					if val >= maxVal {
						val -= maxVal
					}
					t.PaletteColorsV[idx] = Clip1(val, p)
				}
			} else {
				for idx := 0; idx < t.PaletteSizeUV; idx++ {
					t.PaletteColorsV[idx] = p.L(p.sequenceHeader.ColorConfig.BitDepth)
				}
			}
		}
	}
}

// get_palette_cache( plane )
func (t *TileGroup) getPaletteCache(plane int, p *Parser) int {
	aboveN := 0

	if Bool((p.MiRow * MI_SIZE) % 64) {
		aboveN = t.PaletteSizes[plane][p.MiRow-1][p.MiCol]
	}

	leftN := 0
	if p.AvailL {
		leftN = t.PaletteSizes[plane][p.MiRow][p.MiCol-1]
	}

	aboveIdx := 0
	leftIdx := 0
	n := 0

	for aboveIdx < aboveN && leftIdx < leftN {
		aboveC := t.PaletteColors[plane][p.MiRow-1][p.MiCol][aboveIdx]
		leftC := t.PaletteColors[plane][p.MiRow][p.MiCol-1][leftIdx]

		if leftC < aboveC {
			if n == 0 || leftC != t.PaletteCache[n-1] {
				t.PaletteCache[n] = leftC
				n++
			}
			leftIdx++
		} else {
			if n == 0 || aboveC != t.PaletteCache[n-1] {
				t.PaletteCache[n] = aboveC
				n++
			}
			aboveIdx++
			if leftC == aboveC {
				leftIdx++
			}
		}
	}

	for aboveIdx < aboveN {
		val := t.PaletteColors[plane][p.MiRow-1][p.MiCol][aboveIdx]
		aboveIdx++
		if n == 0 || val != t.PaletteCache[n-1] {
			t.PaletteCache[n] = val
			n++
		}
	}
	return n
}

// intra_angle_info_uv()
func (t *TileGroup) intraAngleInfoUv(p *Parser) {
	t.AngleDeltaUV = 0
	if p.MiSize >= BLOCK_8x8 {
		if t.isDirectionalMode(t.UVMode) {
			angleDeltaUv := p.S()
			t.AngleDeltaUV = angleDeltaUv - MAX_ANGLE_DELTA
		}
	}
}

// assign_mv( isCompound )
func (t *TileGroup) assignMv(isCompound int, p *Parser) {
	for i := 0; i < 1+isCompound; i++ {
		var compMode int
		if Bool(t.useIntrabc) {
			compMode = NEWMV
		} else {
			compMode = t.getMode(i)
		}

		if Bool(t.useIntrabc) {
			t.PredMv[0] = t.RefStackMv[0][0]
			if t.PredMv[0][0] == 0 && t.PredMv[0][1] == 0 {
				t.PredMv[0] = t.RefStackMv[1][0]
			}
			if t.PredMv[0][0] == 0 && t.PredMv[0][1] == 0 {
				var sbSize int
				if p.sequenceHeader.Use128x128SuperBlock {
					sbSize = BLOCK_128x128
				} else {
					sbSize = BLOCK_64x64
				}
				sbSize4 := p.Num4x4BlocksHigh[sbSize]

				if p.MiRow-sbSize4 < p.MiRowStart {
					t.PredMv[0][0] = 0
					t.PredMv[0][1] = -(sbSize4*MI_SIZE + INTRABC_DELAY_PIXELS) * 8
				} else {
					t.PredMv[0][0] = -(sbSize4 * MI_SIZE * 8)
					t.PredMv[0][0] = 1
				}
			}

		} else if compMode == GLOBALMV {
			t.PredMv[i] = t.GlobalMvs[i]
		} else {
			var pos int
			if compMode == NEARESTMV {
				pos = 0
			} else {
				pos = t.RefMvIdx
			}

			if compMode == NEWMV && t.NumMvFound <= 1 {
				pos = 0
			}

			t.PredMv[i] = t.RefStackMv[pos][i]
		}

		if compMode == NEWMV {
			t.readMv(i, p)
		} else {
			t.Mv[i] = t.PredMv[i]
		}
	}
}

// read_mv( ref )
func (t *TileGroup) readMv(ref int, p *Parser) {
	var diffMv []int
	diffMv[0] = 0
	diffMv[1] = 0

	if Bool(t.useIntrabc) {
		t.MvCtx = MV_INTRABC_CONTEXT
	} else {
		t.MvCtx = 0
	}

	mvJoint := p.S()

	if mvJoint == MV_JOINT_HZVNZ || mvJoint == MV_JOINT_HNZVNZ {
		diffMv[0] = t.readMvComponent(0, p)
	}

	if mvJoint == MV_JOINT_HNZVZ || mvJoint == MV_JOINT_HNZVNZ {
		diffMv[1] = t.readMvComponent(1, p)
	}

	t.Mv[ref][0] = t.PredMv[ref][0] + diffMv[0]
	t.Mv[ref][1] = t.PredMv[ref][1] + diffMv[1]
}

// read_mv_component( comp )
func (t *TileGroup) readMvComponent(comp int, p *Parser) int {
	mvSign := p.S()
	mvClass := p.S()

	var mag int
	if mvClass == MV_CLASS_0 {
		mvClass0Bit := p.S()

		var mvClass0Fr int
		if p.uncompressedHeader.ForceIntegerMv {
			mvClass0Fr = 3
		} else {
			mvClass0Fr = p.S()
		}

		var mvClass0Hp int
		if p.uncompressedHeader.AllowHighPrecisionMv {
			mvClass0Hp = p.S()
		} else {
			mvClass0Hp = 1
		}

		mag = ((mvClass0Bit << 3) | (mvClass0Fr << 1) | mvClass0Hp) + 1
	} else {
		d := 0
		for i := 0; i < mvClass; i++ {
			mvBit := p.S()
			d |= mvBit << 1
		}

		mag = CLASS0_SIZE << (mvClass + 2)

		var mvFr int
		var mvHp int
		if p.uncompressedHeader.ForceIntegerMv {
			mvFr = 3
		} else {
			mvFr = p.S()
		}

		if p.uncompressedHeader.AllowHighPrecisionMv {
			mvHp = p.S()
		} else {
			mvHp = 1
		}

		mag += ((d << 3) | (mvFr << 1) | mvHp) + 1
	}

	if Bool(mvSign) {
		return -mag
	} else {
		return mag
	}
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

// read_cfl_alphas()
func (t *TileGroup) readCflAlphas(p *Parser) {
	cflAlphaSigns := p.S()
	signU := (cflAlphaSigns + 1) / 3
	signV := (cflAlphaSigns + 1) % 3

	if signU != CFL_SIGN_ZERO {
		cflAlphaU := p.S()
		t.CflAlphaU = 1 + cflAlphaU
		if signU == CFL_SIGN_NEG {
			t.CflAlphaU = -t.CflAlphaU
		}
	} else {
		t.CflAlphaU = 0
	}

	if signV != CFL_SIGN_ZERO {
		cflAlphaV := p.S()
		t.CflAlphaV = 1 + cflAlphaV
		if signV == CFL_SIGN_NEG {
			t.CflAlphaV = -t.CflAlphaV
		}
	} else {
		t.CflAlphaV = 0
	}

}

// intra_angle_info_y()
func (t *TileGroup) intraAngleInfoY(p *Parser) {
	t.AngleDeltaY = 0

	if p.MiSize >= BLOCK_8x8 {

		if t.isDirectionalMode(t.YMode) {
			angleDeltaY := p.S()
			t.AngleDeltaY = angleDeltaY - MAX_ANGLE_DELTA
		}
	}
}

// is_directional_mode( mode )
func (t *TileGroup) isDirectionalMode(mode int) bool {
	return (mode >= V_PRED) && (mode <= D67_PRED)
}

// 7.10.2. Find MV stack process
// find_mv_stack( isCompound )
func (t *TileGroup) findMvStack(isCompound int, p *Parser) {
	// 1.
	t.NumMvFound = 0

	// 2.
	t.NewMvCount = 0

	// 3.
	t.GlobalMvs[0] = t.setupGlobalMvProcess(0, p)

	// 4.
	if Bool(isCompound) {
		t.GlobalMvs[1] = t.setupGlobalMvProcess(1, p)
	}

	// 5.
	t.FoundMatch = 0

	// 6.
	t.scanRowProcess(-1, isCompound, p)
}

func (t *TileGroup) scanRowProcess(deltaRow int, isCompound int, p *Parser) {
	bw4 := p.Num4x4BlocksWide[p.MiSize]
	end4 := Min(Min(bw4, p.MiCols-p.MiCol), 16)
	deltaCol := 0
	useStep16 := bw4 >= 16

	if Abs(deltaRow) > 1 {
		deltaRow += p.MiRow & 1
		deltaCol = 1 - (p.MiCol & 1)
	}

	i := 0

	for i < end4 {
		mvRow := p.MiRow + deltaRow
		mvCol := p.MiCol + deltaCol + i

		if !p.isInside(mvRow, mvCol) {
			break
		}

		len := Min(bw4, p.Num4x4BlocksWide[p.MiSizes[mvRow][mvCol]])
		if Abs(deltaRow) > 1 {
			len = Max(2, len)
		}
		if useStep16 {
			len = Max(4, len)
		}
		weight := len * 2
		t.addRefMvCandidate(mvRow, mvCol, isCompound, weight, p)
		i += len
	}
}

// 7.10.2.7. Add reference motion vector process
func (t *TileGroup) addRefMvCandidate(mvRow int, mvCol int, isCompound int, weight int, p *Parser) {
	if t.IsInters[mvRow][mvCol] == 0 {
		return
	}

	// TODO: not sure if this loop is correct here
	for candList := 0; candList < 2; candList++ {
		if isCompound == 0 {
			if p.RefFrames[mvRow][mvCol][candList] == p.RefFrame[0] {
				t.searchStackProcess(mvRow, mvCol, candList, weight, p)
			}

		} else {
			if p.RefFrames[mvRow][mvCol][0] == p.RefFrame[0] && p.RefFrames[mvRow][mvCol][1] == p.RefFrame[1] {
				t.compoundSearchStackProcess(mvRow, mvCol, weight, p)
			}
		}
	}
}

// 7.10.2.8. Search stack process
func (t *TileGroup) searchStackProcess(mvRow int, mvCol int, candList int, weight int, p *Parser) {
	candMode := t.YModes[mvRow][mvCol]
	candSize := p.MiSizes[mvRow][mvCol]
	large := Min(t.Block_Width[candSize], t.Block_Height[candSize]) >= 8

	var candMv []int
	if (candMode == GLOBALMV || candMode == GLOBAL_GLOBALMV) && (p.GmType[p.RefFrame[0]] > TRANSLATION) && large {
		candMv = t.GlobalMvs[0]
	} else {
		candMv = t.Mvs[mvRow][mvCol][candList]
	}

	candMv = t.lowerPrecisionProcess(candMv, p)
	if HasNewmv(candMode) {
		t.NewMvCount += 1
	}

	t.FoundMatch = 1

	for idx := 0; idx < t.NumMvFound; idx++ {
		if Equals(candMv, t.RefStackMv[idx][0]) {
			t.WeightStack[idx] += weight
			return
		}
	}

	if t.NumMvFound < MAX_REF_MV_STACK_SIZE {
		t.RefStackMv[t.NumMvFound][0] = candMv
		t.WeightStack[t.NumMvFound] = weight
		t.NumMvFound += 1
	}
}

// 7.10.2.9. Compound search stack process
func (t *TileGroup) compoundSearchStackProcess(mvRow int, mvCol int, weight int, p *Parser) {
}

// 7.10.2.1 Setup global MV process
func (t *TileGroup) setupGlobalMvProcess(refList int, p *Parser) []int {
	ref := p.RefFrame[refList]

	var typ int
	if ref != INTRA_FRAME {
		typ = p.GmType[ref]
	}

	bw := t.Block_Width[p.MiSize]
	bh := t.Block_Height[p.MiSize]

	var xc int
	var yc int
	mv := []int{}
	if ref == INTRA_FRAME || typ == IDENTITY {
		mv[0] = 0
		mv[1] = 0
	} else if typ == TRANSLATION {
		mv[0] = p.uncompressedHeader.GmParams[ref][0] >> (WARPEDMODEL_PREC_BITS - 3)
		mv[1] = p.uncompressedHeader.GmParams[ref][1] >> (WARPEDMODEL_PREC_BITS - 3)
	} else {
		x := p.MiCol*MI_SIZE + bw/2 - 1
		y := p.MiRow*MI_SIZE + bh/2 - 1

		xc = (p.uncompressedHeader.GmParams[ref][2]-(1<<WARPEDMODEL_PREC_BITS))*x + p.uncompressedHeader.GmParams[ref][3]*y + p.uncompressedHeader.GmParams[ref][0]
		yc = p.uncompressedHeader.GmParams[ref][4]*x + (p.uncompressedHeader.GmParams[ref][5]-(1<<WARPEDMODEL_PREC_BITS))*y + p.uncompressedHeader.GmParams[ref][1]

		if p.uncompressedHeader.AllowHighPrecisionMv {
			mv[0] = Round2Signed(yc, WARPEDMODEL_PREC_BITS-3)
			mv[1] = Round2Signed(xc, WARPEDMODEL_PREC_BITS-3)
		} else {
			mv[0] = Round2Signed(yc, WARPEDMODEL_PREC_BITS-2) * 2
			mv[1] = Round2Signed(xc, WARPEDMODEL_PREC_BITS-2) * 2
		}
	}
	mv = t.lowerPrecisionProcess(mv, p)

	return mv
}

// 7.10.2.10. Lower precision process
func (t *TileGroup) lowerPrecisionProcess(candMv []int, p *Parser) []int {
	if p.uncompressedHeader.AllowHighPrecisionMv {
		return candMv
	}

	for i := 0; i < 2; i++ {
		if p.uncompressedHeader.ForceIntegerMv {
			a := Abs(candMv[i])
			aInt := (a + 3) >> 3

			if candMv[i] > 0 {
				candMv[i] = aInt << 3
			} else {
				candMv[i] = -(aInt << 3)
			}
		} else {
			if Bool(candMv[i] & 1) {
				if candMv[i] > 0 {
					// TODO: does this work?/
					candMv[i]--
				} else {
					candMv[i]++
				}
			}
		}
	}

	return candMv
}

// read_delta_lf()
func (t *TileGroup) readDeltaLf(p *Parser) {
	var sbSize int
	if p.sequenceHeader.Use128x128SuperBlock {
		sbSize = BLOCK_128x128
	} else {
		sbSize = BLOCK_64x64
	}

	if p.MiSize == sbSize && Bool(t.Skip) {
		return
	}

	if p.ReadDeltas && p.uncompressedHeader.DeltaLfPresent {
		frameLfCount := 1

		if Bool(p.uncompressedHeader.DeltaLfMulti) {
			if p.sequenceHeader.ColorConfig.NumPlanes > 1 {
				frameLfCount = FRAME_LF_COUNT
			} else {
				frameLfCount = FRAME_LF_COUNT - 2
			}
		}

		for i := 0; i < frameLfCount; i++ {
			var deltaLfAbs int
			delta_lf_abs := p.S()

			if delta_lf_abs == DELTA_LF_SMALL {
				deltaLfRemBits := p.L(3)
				n := deltaLfRemBits + 1
				deltaLfAbsBits := p.L(n)
				deltaLfAbs = deltaLfAbsBits + (1 << n) + 1
			} else {
				deltaLfAbs = delta_lf_abs
			}

			var reducedDeltaLfLevel int
			if Bool(deltaLfAbs) {
				deltaLfSignBit := p.L(1)
				if Bool(deltaLfSignBit) {
					reducedDeltaLfLevel = -deltaLfAbs
				} else {
					reducedDeltaLfLevel = deltaLfAbs

				}

				p.DeltaLF[i] = Clip3(-MAX_LOOP_FILTER, MAX_LOOP_FILTER, p.DeltaLF[i]+(reducedDeltaLfLevel<<p.uncompressedHeader.DeltaLfRes))
			}
		}
	}
}

// read_delta_qindex()
func (t *TileGroup) readDeltaQIndex(p *Parser) {
	var sbSize int
	if p.sequenceHeader.Use128x128SuperBlock {
		sbSize = BLOCK_128x128
	} else {
		sbSize = BLOCK_64x64
	}

	if p.MiSize == sbSize && Bool(t.Skip) {
		return
	}

	if p.ReadDeltas {
		deltaQAbs := p.S()
		if deltaQAbs == DELTA_Q_SMALL {
			deltaQRemBits := p.L(3)
			deltaQRemBits++
			deltaQAbsBits := p.L(deltaQRemBits)
			deltaQAbs = deltaQAbsBits + (1 << deltaQRemBits) + 1
		}

		if Bool(deltaQAbs) {
			deltaQSignBit := p.L(1)
			var reducedDeltaQIndex int
			if Bool(deltaQSignBit) {
				reducedDeltaQIndex = -deltaQAbs
			} else {
				reducedDeltaQIndex = deltaQAbs
			}

			p.CurrentQIndex = Clip3(1, 255, p.CurrentQIndex+(reducedDeltaQIndex<<p.uncompressedHeader.DeltaQRes))

		}
	}
}

// read_cdef()
func (t *TileGroup) readCdef(p *Parser) {
	if Bool(t.Skip) || p.uncompressedHeader.CodedLossless || !p.sequenceHeader.EnableCdef || p.uncompressedHeader.AllowIntraBc {
		return
	}

	cdefSize4 := p.Num4x4BlocksWide[BLOCK_64x64]
	cdefMask4 := ^(cdefSize4 - 1)
	r := p.MiRow & cdefMask4
	c := p.MiCol & cdefMask4

	if p.Cdef.CdefIdx[r][c] == -1 {
		p.Cdef.CdefIdx[r][c] = p.L(p.Cdef.CdefBits)
		w4 := p.Num4x4BlocksWide[p.MiSize]
		h4 := p.Num4x4BlocksHigh[p.MiSize]

		for i := r; i < r+h4; i += cdefSize4 {
			for j := c; i < c+w4; i += cdefSize4 {
				p.Cdef.CdefIdx[i][j] = p.Cdef.CdefIdx[r][c]
			}

		}
	}
}

// read_skip()
func (t *TileGroup) readSkip(p *Parser) {
	if (p.uncompressedHeader.SegIdPreSkip == 1) && t.segFeatureActive(SEG_LVL_SKIP, p) {
		t.Skip = 1
	} else {
		t.Skip = p.S()
	}
}

// seg_feature_active( feature )
func (t *TileGroup) segFeatureActive(feature int, p *Parser) bool {
	return t.segFeatureActiveIdx(t.SegmentId, feature, p)
}

// seg_feature_active_idx( idx, feature )
func (t *TileGroup) segFeatureActiveIdx(idx int, feature int, p *Parser) bool {
	return (p.uncompressedHeader.SegmentationEnabled == 1) && (p.FeatureEnabled[idx][feature] == 1)
}

// intra_segment_id()
func (t *TileGroup) intraSegmentId(p *Parser) {
	if p.uncompressedHeader.SegmentationEnabled == 1 {
		t.readSegmentId(p)
	} else {
		t.SegmentId = 0
	}

	t.Lossless = p.uncompressedHeader.LosslessArray[t.SegmentId]
}

// read_segment_id()
func (t *TileGroup) readSegmentId(p *Parser) {
	var prevU int
	var prevL int
	var prevUL int
	var pred int
	if p.AvailU && p.AvailL {
		prevUL = t.SegmentIds[p.MiRow-1][p.MiCol-1]
	} else {
		prevUL = -1
	}

	if p.AvailU {
		prevU = t.SegmentIds[p.MiRow-1][p.MiCol]
	} else {
		prevU = -1
	}

	if p.AvailL {
		prevL = t.SegmentIds[p.MiRow][p.MiCol-1]
	} else {
		prevL = -1
	}

	if prevU == -1 {
		if prevL == -1 {
			pred = 0
		} else {
			pred = prevL
		}
	} else if prevL == -1 {
		pred = prevU
	} else {
		if prevUL == prevU {
			pred = prevU
		} else {
			pred = prevL
		}
	}

	if t.Skip == 1 {
		t.SegmentId = pred
	} else {
		t.SegmentId = p.S()
		t.SegmentId = NegDeinterleave(t.SegmentId, pred, p.uncompressedHeader.LastActiveSegId+1)
	}
}

// clear_block_decoded_flags( r, c, sbSize4 )
func (t *TileGroup) clearBlockDecodedFlags(r int, c int, sbSize4 int, p *Parser) {
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

// read_lr( r, c, bSize )
func (t *TileGroup) readLr(r int, c int, bSize int, p *Parser) {
	if p.uncompressedHeader.AllowIntraBc {
		return
	}

	w := p.Num4x4BlocksWide[bSize]
	h := p.Num4x4BlocksHigh[bSize]

	for plane := 0; plane < p.sequenceHeader.ColorConfig.NumPlanes; plane++ {
		if p.FrameRestorationType[plane] != RESTORE_NONE {
			// FIXME: lots of creative freedom here, dangerous!
			subX := 0
			subY := 0

			if p.sequenceHeader.ColorConfig.SubsamplingX {
				subX = 1
			}

			if p.sequenceHeader.ColorConfig.SubsamplingY {
				subY = 1
			}

			unitSize := p.LoopRestorationSize[plane]
			unitRows := countUnitsInFrame(unitSize, Round2(p.uncompressedHeader.FrameHeight, subY))
			unitCols := countUnitsInFrame(unitSize, Round2(p.upscaledWidth, subX))
			unitRowStart := (r*(MI_SIZE>>subY) + unitSize - 1) / unitSize
			unitRowEnd := Min(unitRows, ((r+h)*(MI_SIZE>>subY)+unitSize-1)/unitSize)

			var numerator int
			var denominator int
			if p.uncompressedHeader.UseSuperRes {
				numerator = (MI_SIZE >> subX) * p.uncompressedHeader.SuperResDenom
				denominator = unitSize * SUPERRES_NUM
			} else {
				numerator = MI_SIZE >> subX
				denominator = unitSize
			}
			unitColStart := (c*numerator + denominator - 1) / denominator
			unitColEnd := Min(unitCols, ((c+w)*numerator+denominator-1)/denominator)

			for unitRow := unitRowStart; unitRow < unitRowEnd; unitRow++ {
				for unitCol := unitColStart; unitCol < unitColEnd; unitCol++ {
					t.readLrUnit(plane, unitRow, unitCol, p)
				}
			}
		}
	}
}

// read_lr_unit(plane, unitRow, unitCol)
func (t *TileGroup) readLrUnit(plane int, unitRow int, unitCol int, p *Parser) {
	var restorationType int
	if p.FrameRestorationType[plane] == RESTORE_WIENER {
		useWiener := p.S()
		restorationType = RESTORE_NONE
		if useWiener == 1 {
			restorationType = RESTORE_WIENER
		}
	} else if p.FrameRestorationType[plane] == RESTORE_SGRPROJ {
		useSgrproj := p.S()
		restorationType = RESTORE_NONE
		if useSgrproj == 1 {
			restorationType = RESTORE_SGRPROJ
		}
	} else {
		restorationType = p.S()
	}

	t.LrType[plane][unitRow][unitCol] = restorationType

	if restorationType == RESTORE_WIENER {
		for pass := 0; pass < 2; pass++ {
			var firstCoeff int
			if plane == 1 {
				firstCoeff = 1
				t.LrWiener[plane][unitRow][unitCol][pass][0] = 0
			} else {
				firstCoeff = 0
			}
			for j := firstCoeff; j < 3; j++ {
				min := Wiener_Taps_Min[j]
				max := Wiener_Taps_Max[j]
				k := Wiener_Taps_K[j]
				v := t.decodeSignedSubexpWithRefBool(min, max+1, k, t.RefLrWiener[plane][pass][j], p)
				t.LrWiener[plane][unitRow][unitCol][pass][j] = v
				t.RefLrWiener[plane][pass][j] = v
			}
		}
	} else if restorationType == RESTORE_SGRPROJ {
		lrSgrSet := p.L(SGRPROJ_PARAMS_BITS)
		t.LrSgrSet[plane][unitRow][unitCol] = lrSgrSet

		for i := 0; i < 2; i++ {
			radius := SgrParams[lrSgrSet][i*2]
			min := Sgrproj_Xqd_Min[i]
			max := Sgrproj_Xqd_Max[i]

			var v int
			if radius != 0 {
				v = t.decodeSignedSubexpWithRefBool(min, max+1, SGRPROJ_PRJ_SUBEXP_K, t.RefSgrXqd[plane][i], p)
			} else {
				v = 0
				if i == 1 {
					v = Clip3(min, max, (1<<SGRPROJ_BITS)-t.RefSgrXqd[plane][0])
				}
			}

			t.LrSgrXqd[plane][unitRow][unitCol][i] = v
			t.RefSgrXqd[plane][i] = v
		}
	}

}

func (t *TileGroup) decodeSignedSubexpWithRefBool(low int, high int, k int, r int, p *Parser) int {
	x := t.decodeUnsignedSubexpWithRefBool(high-low, k, r-low, p)
	return x + low

}

func (t *TileGroup) decodeUnsignedSubexpWithRefBool(mx int, k int, r int, p *Parser) int {
	v := t.decodeSubexpBool(mx, k, p)
	if (r << 1) <= mx {
		return InverseRecenter(r, v)
	} else {
		return mx - 1 - InverseRecenter(mx-1-r, v)
	}
}

func (t *TileGroup) decodeSubexpBool(numSyms int, k int, p *Parser) int {
	i := 0
	mk := 0
	for {
		b2 := k
		if i == 1 {
			b2 = k + i - 1
		}

		a := 1 << b2

		if numSyms <= -mk+3*a {
			subexpUnifBools := p.L(1)
			return subexpUnifBools + mk
		} else {
			subexpMoreBools := p.L(1) != 0
			if subexpMoreBools {
				i++
				mk += a
			} else {
				subexpBools := p.L(b2)
				return subexpBools + mk
			}
		}
	}
}

func countUnitsInFrame(unitSize int, frameSize int) int {
	return Max((frameSize+(unitSize>>1))/unitSize, 1)
}
