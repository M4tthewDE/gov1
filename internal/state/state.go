package state

import (
	"github.com/m4tthewde/gov1/internal/cdef"
	"github.com/m4tthewde/gov1/internal/shared"
)

type State struct {
	MiCols                 int
	MiRows                 int
	TileColsLog2           int
	TileRowsLog2           int
	MiRowStarts            []int
	MiColStarts            []int
	TileCols               int
	TileRows               int
	TileSizeBytes          int
	FeatureEnabled         [shared.MAX_SEGMENTS][shared.SEG_LVL_MAX]int
	FeatureData            [shared.MAX_SEGMENTS][shared.SEG_LVL_MAX]int
	GmType                 [shared.ALTREF_FRAME + 1]int
	PrevGmParams           [shared.ALTREF_FRAME + 1][6]int
	PrevSegmentIds         [][]int
	RefFrameType           [7]int
	CurrentQIndex          int
	OperatingPointIdc      int
	MiColStart             int
	MiColEnd               int
	MiRowStart             int
	MiRowEnd               int
	BlockWidth             []int
	BlockHeight            []int
	Num4x4BlocksWide       []int
	Num4x4BlocksHigh       []int
	SeenFrameHeader        bool
	RenderWidth            int
	RenderHeight           int
	UpscaledWidth          int
	UpscaledHeight         int
	TileNum                int
	MiCol                  int
	MiRow                  int
	MiSize                 int
	MiSizes                [][]int
	DeltaLF                []int
	RefLrWiener            [][][]int
	ReadDeltas             bool
	Cdef                   cdef.Cdef
	BlockDecoded           [][][]int
	FrameRestorationType   []int
	LoopRestorationSize    []int
	AvailU                 bool
	AvailL                 bool
	AvailUChroma           bool
	AvailLChroma           bool
	RefFrame               []int
	RefFrames              [][][]int
	RefFrameWidth          []int
	RefFrameHeight         []int
	CurrFrame              [][][]int
	SymbolMaxBits          int
	SymbolValue            int
	SymbolRange            int
	RefFrameTypes          [7]int
	Skip                   int
	IsInter                int
	CdefSize4              int
	CdefFrame              [][][]int
	UpscaledCurrFrame      [][][]int
	LrFrame                [][][]int
	TileIntraFrameYModeCdf [][][]int

	TxbSkipCdf      [][][]int
	EobPt16Cdf      [][][]int
	EobPt32Cdf      [][][]int
	EobPt64Cdf      [][][]int
	EobPt128Cdf     [][][]int
	EobPt256Cdf     [][][]int
	EobPt512Cdf     [][]int
	EobPt1024Cdf    [][]int
	EobExtraCdf     [][][][]int
	DcSignCdf       [][][]int
	CoeffBaseEobCdf [][][][]int
	CoeffBaseCdf    [][][][]int
	CoeffBrCdf      [][][][]int
}
