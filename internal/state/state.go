package state

import (
	"github.com/m4tthewde/gov1/internal/cdef"
	"github.com/m4tthewde/gov1/internal/shared"
)

type State struct {
	MiCols            int
	MiRows            int
	TileColsLog2      int
	TileRowsLog2      int
	MiRowStarts       []int
	MiColStarts       []int
	TileCols          int
	TileRows          int
	TileSizeBytes     int
	FeatureEnabled    [shared.MAX_SEGMENTS][shared.SEG_LVL_MAX]int
	FeatureData       [shared.MAX_SEGMENTS][shared.SEG_LVL_MAX]int
	GmType            [shared.ALTREF_FRAME + 1]int
	PrevGmParams      [shared.ALTREF_FRAME + 1][6]int
	PrevSegmentIds    [][]int
	RefFrameType      [7]int
	CurrentQIndex     int
	OperatingPointIdc int
	MiColStart        int
	MiColEnd          int
	MiRowStart        int
	MiRowEnd          int
	BlockWidth        []int
	BlockHeight       []int

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
	DeltaLF                [shared.FRAME_LF_COUNT]int
	DeltaLFs               [][][]int
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
	RefFrame               [2]int
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

	YModeCdf               [][]int
	UVModeCflNotAllowedCdf [][]int
	UVModeCflAllowedCdf    [][]int
	AngleDeltaCdf          [][]int
	IntrabcCdf             []int
	PartitionW8Cdf         [][]int
	PartitionW16Cdf        [][]int
	PartitionW32Cdf        [][]int
	PartitionW64Cdf        [][]int
	PartitionW128Cdf       [][]int
	SegmentIdCdf           [][]int
	SegmentIdPredictedCdf  [][]int
	Tx8x8Cdf               [][]int
	Tx16x16Cdf             [][]int
	Tx32x32Cdf             [][]int
	Tx64x64Cdf             [][]int
	TxfmSplitCdf           [][]int
	FilterIntraModeCdf     []int
	FilterIntraCdf         [][]int
	InterpFilterCdf        [][]int
	MotionModeCdf          [][]int
	NewMvCdf               [][]int
	ZeroMvCdf              [][]int
	RefMvCdf               [][]int
	CompoundModeCdf        [][]int
	DrlModeCdf             [][]int
	IsInterCdf             [][]int
	CompModeCdf            [][]int
	SkipModeCdf            [][]int
	SkipCdf                [][]int
	CompRefCdf             [][][]int
	CompBwdRefCdf          [][][]int
	SingleRefCdf           [][][]int
	MvJointCdf             [shared.MV_CONTEXTS][]int
	MvClassCdf             [shared.MV_CONTEXTS][][]int
	MvClass0BitCdf         [shared.MV_CONTEXTS][2][]int
	MvFrCdf                [shared.MV_CONTEXTS][][]int
	MvClass0FrCdf          [shared.MV_CONTEXTS][][][]int
	MvClass0HpCdf          [shared.MV_CONTEXTS][2][]int
	MvSignCdf              [shared.MV_CONTEXTS][2][]int
	MvBitCdf               [shared.MV_CONTEXTS][2][][]int
	MvHpCdf                [shared.MV_CONTEXTS][2][]int
	PaletteYModeCdf        [][][]int
	PaletteUVModeCdf       [][]int
	PaletteUVSizeCdf       [][]int
	PaletteSize2YColorCdf  [][]int
	PaletteSize2UVColorCdf [][]int
	PaletteSize3YColorCdf  [][]int
	PaletteSize3UVColorCdf [][]int
	PaletteSize4YColorCdf  [][]int
	PaletteSize4UVColorCdf [][]int
	PaletteSize5YColorCdf  [][]int
	PaletteSize5UVColorCdf [][]int
	PaletteSize6YColorCdf  [][]int
	PaletteSize6UVColorCdf [][]int
	PaletteSize7YColorCdf  [][]int
	PaletteSize7UVColorCdf [][]int
	PaletteSize8YColorCdf  [][]int
	PaletteSize8UVColorCdf [][]int

	DeltaQCdf          []int
	DeltaLFCdf         []int
	DeltaLFMultiCdf    [shared.FRAME_LF_COUNT][]int
	IntraTxTypeSet1Cdf [][][]int
	IntraTxTypeSet2Cdf [][][]int
	InterTxTypeSet1Cdf [][]int
	InterTxTypeSet2Cdf []int
	InterTxTypeSet3Cdf [][]int
	UseObmcCdf         [][]int
	InterIntraCdf      [][]int
	CompRefTypeCdf     [][]int
	CflSignCdf         []int
	UniCompRefCdf      [][][]int
	WedgeInterIntraCdf [][]int
	CompGroupIdxCdf    [][]int
	CompoundIdxCdf     [][]int
	CompoundTypeCdf    [][]int
	InterIntraModeCdf  [][]int
	WedgeIndexCdf      [][]int
	CflAlphaCdf        [][]int
	UseWienerCdf       []int
	UseSgrprojCdf      []int
	RestorationTypeCdf []int

	TileYModeCdf               [][]int
	TileUVModeCflNotAllowedCdf [][]int
	TileUVModeCflAllowedCdf    [][]int
	TileAngleDeltaCdf          [][]int
	TileIntrabcCdf             []int
	TilePartitionW8Cdf         [][]int
	TilePartitionW16Cdf        [][]int
	TilePartitionW32Cdf        [][]int
	TilePartitionW64Cdf        [][]int
	TilePartitionW128Cdf       [][]int
	TileSegmentIdCdf           [][]int
	TileSegmentIdPredictedCdf  [][]int
	TileTx8x8Cdf               [][]int
	TileTx16x16Cdf             [][]int
	TileTx32x32Cdf             [][]int
	TileTx64x64Cdf             [][]int
	TileTxfmSplitCdf           [][]int
	TileFilterIntraModeCdf     []int
	TileFilterIntraCdf         [][]int
	TileInterpFilterCdf        [][]int
	TileMotionModeCdf          [][]int
	TileNewMvCdf               [][]int
	TileZeroMvCdf              [][]int
	TileRefMvCdf               [][]int
	TileCompoundModeCdf        [][]int
	TileDrlModeCdf             [][]int
	TileIsInterCdf             [][]int
	TileCompModeCdf            [][]int
	TileSkipModeCdf            [][]int
	TileSkipCdf                [][]int
	TileCompRefCdf             [][][]int
	TileCompBwdRefCdf          [][][]int
	TileSingleRefCdf           [][][]int
	TileMvJointCdf             [shared.MV_CONTEXTS][]int
	TileMvClassCdf             [shared.MV_CONTEXTS][][]int
	TileMvClass0BitCdf         [shared.MV_CONTEXTS][2][]int
	TileMvFrCdf                [shared.MV_CONTEXTS][][]int
	TileMvClass0FrCdf          [shared.MV_CONTEXTS][][][]int
	TileMvClass0HpCdf          [shared.MV_CONTEXTS][2][]int
	TileMvSignCdf              [shared.MV_CONTEXTS][2][]int
	TileMvBitCdf               [shared.MV_CONTEXTS][2][][]int
	TileMvHpCdf                [shared.MV_CONTEXTS][2][]int
	TilePaletteYModeCdf        [][][]int
	TilePaletteUVModeCdf       [][]int
	TilePaletteUVSizeCdf       [][]int
	TilePaletteSize2YColorCdf  [][]int
	TilePaletteSize2UVColorCdf [][]int
	TilePaletteSize3YColorCdf  [][]int
	TilePaletteSize3UVColorCdf [][]int
	TilePaletteSize4YColorCdf  [][]int
	TilePaletteSize4UVColorCdf [][]int
	TilePaletteSize5YColorCdf  [][]int
	TilePaletteSize5UVColorCdf [][]int
	TilePaletteSize6YColorCdf  [][]int
	TilePaletteSize6UVColorCdf [][]int
	TilePaletteSize7YColorCdf  [][]int
	TilePaletteSize7UVColorCdf [][]int
	TilePaletteSize8YColorCdf  [][]int
	TilePaletteSize8UVColorCdf [][]int

	TileDeltaQCdf          []int
	TileDeltaLFCdf         []int
	TileDeltaLFMultiCdf    [shared.FRAME_LF_COUNT][]int
	TileIntraTxTypeSet1Cdf [][][]int
	TileIntraTxTypeSet2Cdf [][][]int
	TileInterTxTypeSet1Cdf [][]int
	TileInterTxTypeSet2Cdf []int
	TileInterTxTypeSet3Cdf [][]int
	TileUseObmcCdf         [][]int
	TileInterIntraCdf      [][]int
	TileCompRefTypeCdf     [][]int
	TileCflSignCdf         []int
	TileUniCompRefCdf      [][][]int
	TileWedgeInterIntraCdf [][]int
	TileCompGroupIdxCdf    [][]int
	TileCompoundIdxCdf     [][]int
	TileCompoundTypeCdf    [][]int
	TileInterIntraModeCdf  [][]int
	TileWedgeIndexCdf      [][]int
	TileCflAlphaCdf        [][]int
	TileUseWienerCdf       []int
	TileUseSgrprojCdf      []int
	TileRestorationTypeCdf []int

	TileTxbSkipCdf      [][][]int
	TileEobPt16Cdf      [][][]int
	TileEobPt32Cdf      [][][]int
	TileEobPt64Cdf      [][][]int
	TileEobPt128Cdf     [][][]int
	TileEobPt256Cdf     [][][]int
	TileEobPt512Cdf     [][]int
	TileEobPt1024Cdf    [][]int
	TileEobExtraCdf     [][][][]int
	TileDcSignCdf       [][][]int
	TileCoeffBaseEobCdf [][][][]int
	TileCoeffBaseCdf    [][][][]int
	TileCoeffBrCdf      [][][][]int
}
