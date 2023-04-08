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
	RefFrameType      [shared.NUM_REF_FRAMES]int
	CurrentQIndex     int
	OperatingPointIdc int
	MiColStart        int
	MiColEnd          int
	MiRowStart        int
	MiRowEnd          int
	BlockWidth        [shared.MAX_TILE_ROWS]int
	BlockHeight       [shared.MAX_TILE_COLS]int

	SeenFrameHeader bool
	RenderWidth     int
	RenderHeight    int
	TileNum         int
	MiCol           int
	MiRow           int
	MiSize          int
	MiSizes         [shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS]int
	DeltaLF         [shared.FRAME_LF_COUNT]int
	DeltaLFs        [2][2][4]int
	RefLrWiener     [][][]int
	ReadDeltas      bool
	Cdef            cdef.Cdef

	BlockDecoded           [3][13][13]bool
	FrameRestorationType   []int
	LoopRestorationSize    [3]int
	AvailU                 bool
	AvailL                 bool
	AvailUChroma           bool
	AvailLChroma           bool
	RefFrame               [2]int
	RefFrames              [shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS][2]int
	RefFrameWidth          [shared.NUM_REF_FRAMES]int
	RefFrameHeight         [shared.NUM_REF_FRAMES]int
	RefRenderWidth         [shared.NUM_REF_FRAMES]int
	RefRenderHeight        [shared.NUM_REF_FRAMES]int
	RefMiCols              [shared.NUM_REF_FRAMES]int
	RefMiRows              [shared.NUM_REF_FRAMES]int
	RefSubsamplingX        [shared.NUM_REF_FRAMES]int
	RefSubsamplingY        [shared.NUM_REF_FRAMES]int
	RefBitDepth            [shared.NUM_REF_FRAMES]int
	RefUpscaledWidth       [8]int
	SavedOrderHints        [shared.NUM_REF_FRAMES][shared.REFS_PER_FRAME + shared.LAST_FRAME]int
	CurrFrame              [3][16][16]int
	SymbolMaxBits          int
	SymbolValue            int
	SymbolRange            int
	RefFrameTypes          [7]int
	Skip                   int
	IsInter                int
	CdefSize4              int
	CdefFrame              [3][16][16]int
	UpscaledCurrFrame      [3][9][9]int
	UpscaledCdefFrame      [3][9][9]int
	LrFrame                [3][9][9]int
	TileIntraFrameYModeCdf [][][]int
	MfMvs                  [shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS][2]int
	SavedMvs               [shared.NUM_REF_FRAMES][shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS][2]int
	SavedRefFrames         [shared.NUM_REF_FRAMES][shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS]int
	MfRefFrames            [shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS]int
	SavedGmParams          [shared.NUM_REF_FRAMES][shared.ALTREF_FRAME + 1][6]int
	SavedSegmentIds        [shared.NUM_REF_FRAMES][shared.MAX_TILE_ROWS][shared.MAX_TILE_COLS]int

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
	RefOrderHint           [shared.NUM_REF_FRAMES]int
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

	SavedYModeCdf               [][]int
	SavedUVModeCflNotAllowedCdf [][]int
	SavedUVModeCflAllowedCdf    [][]int
	SavedAngleDeltaCdf          [][]int
	SavedIntrabcCdf             []int
	SavedPartitionW8Cdf         [][]int
	SavedPartitionW16Cdf        [][]int
	SavedPartitionW32Cdf        [][]int
	SavedPartitionW64Cdf        [][]int
	SavedPartitionW128Cdf       [][]int
	SavedSegmentIdCdf           [][]int
	SavedSegmentIdPredictedCdf  [][]int
	SavedTx8x8Cdf               [][]int
	SavedTx16x16Cdf             [][]int
	SavedTx32x32Cdf             [][]int
	SavedTx64x64Cdf             [][]int
	SavedTxfmSplitCdf           [][]int
	SavedFilterIntraModeCdf     []int
	SavedFilterIntraCdf         [][]int
	SavedInterpFilterCdf        [][]int
	SavedMotionModeCdf          [][]int
	SavedNewMvCdf               [][]int
	SavedZeroMvCdf              [][]int
	SavedRefMvCdf               [][]int
	SavedCompoundModeCdf        [][]int
	SavedDrlModeCdf             [][]int
	SavedIsInterCdf             [][]int
	SavedCompModeCdf            [][]int
	SavedSkipModeCdf            [][]int
	SavedSkipCdf                [][]int
	SavedCompRefCdf             [][][]int
	SavedCompBwdRefCdf          [][][]int
	SavedSingleRefCdf           [][][]int
	SavedMvJointCdf             [shared.MV_CONTEXTS][]int
	SavedMvClassCdf             [shared.MV_CONTEXTS][][]int
	SavedMvClass0BitCdf         [shared.MV_CONTEXTS][2][]int
	SavedMvFrCdf                [shared.MV_CONTEXTS][][]int
	SavedMvClass0FrCdf          [shared.MV_CONTEXTS][][][]int
	SavedMvClass0HpCdf          [shared.MV_CONTEXTS][2][]int
	SavedMvSignCdf              [shared.MV_CONTEXTS][2][]int
	SavedMvBitCdf               [shared.MV_CONTEXTS][2][][]int
	SavedMvHpCdf                [shared.MV_CONTEXTS][2][]int
	SavedPaletteYModeCdf        [][][]int
	SavedPaletteUVModeCdf       [][]int
	SavedPaletteUVSizeCdf       [][]int
	SavedPaletteSize2YColorCdf  [][]int
	SavedPaletteSize2UVColorCdf [][]int
	SavedPaletteSize3YColorCdf  [][]int
	SavedPaletteSize3UVColorCdf [][]int
	SavedPaletteSize4YColorCdf  [][]int
	SavedPaletteSize4UVColorCdf [][]int
	SavedPaletteSize5YColorCdf  [][]int
	SavedPaletteSize5UVColorCdf [][]int
	SavedPaletteSize6YColorCdf  [][]int
	SavedPaletteSize6UVColorCdf [][]int
	SavedPaletteSize7YColorCdf  [][]int
	SavedPaletteSize7UVColorCdf [][]int
	SavedPaletteSize8YColorCdf  [][]int
	SavedPaletteSize8UVColorCdf [][]int

	SavedDeltaQCdf          []int
	SavedDeltaLFCdf         []int
	SavedDeltaLFMultiCdf    [shared.FRAME_LF_COUNT][]int
	SavedIntraTxTypeSet1Cdf [][][]int
	SavedIntraTxTypeSet2Cdf [][][]int
	SavedInterTxTypeSet1Cdf [][]int
	SavedInterTxTypeSet2Cdf []int
	SavedInterTxTypeSet3Cdf [][]int
	SavedUseObmcCdf         [][]int
	SavedInterIntraCdf      [][]int
	SavedCompRefTypeCdf     [][]int
	SavedCflSignCdf         []int
	SavedUniCompRefCdf      [][][]int
	SavedWedgeInterIntraCdf [][]int
	SavedCompGroupIdxCdf    [][]int
	SavedCompoundIdxCdf     [][]int
	SavedCompoundTypeCdf    [][]int
	SavedInterIntraModeCdf  [][]int
	SavedWedgeIndexCdf      [][]int
	SavedCflAlphaCdf        [][]int
	SavedUseWienerCdf       []int
	SavedUseSgrprojCdf      []int
	SavedRestorationTypeCdf []int

	SavedTxbSkipCdf      [][][]int
	SavedEobPt16Cdf      [][][]int
	SavedEobPt32Cdf      [][][]int
	SavedEobPt64Cdf      [][][]int
	SavedEobPt128Cdf     [][][]int
	SavedEobPt256Cdf     [][][]int
	SavedEobPt512Cdf     [][]int
	SavedEobPt1024Cdf    [][]int
	SavedEobExtraCdf     [][][][]int
	SavedDcSignCdf       [][][]int
	SavedCoeffBaseEobCdf [][][][]int
	SavedCoeffBaseCdf    [][][][]int
	SavedCoeffBrCdf      [][][][]int
	RefValid             [shared.NUM_REF_FRAMES]int

	Memory []MemoryArea
}

type MemoryArea struct {
	Ctx int

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
	RefOrderHint           [shared.NUM_REF_FRAMES]int
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
}
