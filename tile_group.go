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

const DELTA_Q_SMALL = 3
const DELTA_LF_SMALL = 3

const INTRA_FRAME = 0
const NONE = -1

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

const COMPUND_AVERAGE = 2

type TileGroup struct {
	LrType         [][][]int
	RefLrWiener    [][][]int
	LrWiener       [][][][][]int
	LrSgrSet       [][][]int
	RefSgrXqd      [][]int
	LrSgrXqd       [][][][]int
	HasChroma      bool
	SegmentId      int
	SegmentIds     [][]int
	Lossless       bool
	Skip           int
	YMode          int
	UVMode         int
	PalletteSizeY  int
	PalletteSizeUV int
	InterpFilter   []int
	NumMvFound     int
	NewMvCount     int
	GlobalMvs      [][]int
	Block_Width    []int
	Block_Height   []int
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
}

// mode_info()
func (t *TileGroup) modeInfo(p *Parser) {
	if p.uncompressedHeader.FrameIsIntra {
		t.intraFrameModeInfo(p)
	} else {
		t.interFrameModeInfo()
	}
}

// intra_frame_mode_info()
func (t *TileGroup) intraFrameModeInfo(p *Parser) {
	t.Skip = 0
	if p.uncompressedHeader.SegIdPreSkip == 1 {
		t.intraSegmentId(p)
	}

	skipMode := 0
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

	var useIntrabc int
	if p.uncompressedHeader.AllowIntraBc {
		useIntrabc = p.S()
	} else {
		useIntrabc = 0
	}

	if Bool(useIntrabc) {
		isInter := -1
		t.YMode = DC_PRED
		t.UVMode = DC_PRED
		motionMode := SIMPLE
		compoundType := COMPUND_AVERAGE
		t.PalletteSizeY = 0
		t.PalletteSizeUV = 0
		t.InterpFilter[0] = BILINEAR
		t.InterpFilter[1] = BILINEAR

		// NEXT:
		// GOOD LUCK
		t.findMvStack(0, p)
	}

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

	FoundMatch := 0
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
