package uncompressedheader

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
)

// init_non_coeff_cdfs( )
func initNonCoeffCdfs(state *state.State) {
	state.YModeCdf = DEFAULT_Y_MODE_CDF
	state.UVModeCflNotAllowedCdf = DEFAULT_UV_MODE_CFL_NOT_ALLOWED_CDF
	state.UVModeCflAllowedCdf = DEFAULT_UV_MODE_CFL_ALLOWED_CDF
	state.AngleDeltaCdf = DEFAULT_ANGLE_DELTA_CDF
	state.IntrabcCdf = DEFAULT_INTRABC_CDF
	state.PartitionW8Cdf = DEFAULT_PARTITION_W8_CDF
	state.PartitionW16Cdf = DEFAULT_PARTITION_W16_CDF
	state.PartitionW32Cdf = DEFAULT_PARTITION_W32_CDF
	state.PartitionW64Cdf = DEFAULT_PARTITION_W64_CDF
	state.PartitionW128Cdf = DEFAULT_PARTITION_W128_CDF
	state.SegmentIdCdf = DEFAULT_SEGMENT_ID_CDF
	state.SegmentIdPredictedCdf = DEFAULT_SEGMENT_ID_PREDICTED_CDF
	state.Tx8x8Cdf = DEFAULT_TX_8X8_CDF
	state.Tx16x16Cdf = DEFAULT_TX_16X16_CDF
	state.Tx32x32Cdf = DEFAULT_TX_32X32_CDF
	state.Tx64x64Cdf = DEFAULT_TX_64X64_CDF
	state.TxfmSplitCdf = DEFAULT_TXFM_SPLIT_CDF
	state.FilterIntraModeCdf = DEFAULT_FILTER_INTRA_MODE_CDF
	state.FilterIntraCdf = DEFAULT_FILTER_INTRA_CDF
	state.InterpFilterCdf = DEFAULT_INTERP_FILTER_CDF
	state.MotionModeCdf = DEFAULT_MOTION_MODE_CDF
	state.NewMvCdf = DEFAULT_NEW_MV_CDF
	state.ZeroMvCdf = DFEAULT_ZERO_MV_CDF
	state.CompoundModeCdf = DEFAULT_COMPOUND_MODE_CDF
	state.DrlModeCdf = DEFAULT_DRL_MODE_CDF
	state.IsInterCdf = DEFAULT_IS_INTER_CDF
	state.CompModeCdf = DEFAULT_COMP_MODE_CDF
	state.SkipModeCdf = DEFAULT_SKIP_MODE_CDF
	state.SkipCdf = DEFAULT_SKIP_CDF
	state.CompRefCdf = DEFAULT_COMP_REF_CDF
	state.CompBwdRefCdf = DEFAULT_COMP_BWD_REF_CDF
	state.SingleRefCdf = DEFAULT_SINGLE_REF_CDF

	for i := 0; i < shared.MV_CONTEXTS; i++ {
		state.MvJointCdf[i] = DEFAULT_MV_JOINT_CDF
		state.MvClassCdf[i] = DEFAULT_MV_CLASS_CDF
	}

	for i := 0; i < shared.MV_CONTEXTS; i++ {
		for comp := 0; comp <= 1; comp++ {
			state.MvClass0BitCdf[i][comp] = DEFAULT_MV_CLASS0_BIT_CDF
			state.MvFrCdf[i] = DEFAULT_MV_FR_CDF
			state.MvClass0FrCdf[i] = DEFAULT_MV_CLASS0_FR_CDF
			state.MvClass0HpCdf[i][comp] = DEFAULT_MV_CLASS0_HP_CDF
			state.MvSignCdf[i][comp] = DEFAULT_MV_SIGN_CDF
			state.MvBitCdf[i][comp] = DEFAULT_MV_BIT_CDF
			state.MvHpCdf[i][comp] = DEFAULT_MV_HP_CDF
		}
	}

	state.PaletteYModeCdf = DEFAULT_PALETTE_Y_MODE_CDF
	state.PaletteUVModeCdf = DEFAULT_PALETTE_UV_MODE_CDF
	state.PaletteUVSizeCdf = DEFAULT_PALETTE_UV_SIZE_CDF
	state.PaletteSize2YColorCdf = DEFAULT_PALETTE_SIZE_2_Y_COLOR_CDF
	state.PaletteSize2UVColorCdf = DEFAULT_PALETTE_SIZE_2_UV_COLOR_CDF
	state.PaletteSize3YColorCdf = DEFAULT_PALETTE_SIZE_3_Y_COLOR_CDF
	state.PaletteSize3UVColorCdf = DEFAULT_PALETTE_SIZE_3_UV_COLOR_CDF
	state.PaletteSize4YColorCdf = DEFAULT_PALETTE_SIZE_4_Y_COLOR_CDF
	state.PaletteSize4UVColorCdf = DEFAULT_PALETTE_SIZE_4_UV_COLOR_CDF
	state.PaletteSize5YColorCdf = DEFAULT_PALETTE_SIZE_5_Y_COLOR_CDF
	state.PaletteSize5UVColorCdf = DEFAULT_PALETTE_SIZE_5_UV_COLOR_CDF
	state.PaletteSize6YColorCdf = DEFAULT_PALETTE_SIZE_6_Y_COLOR_CDF
	state.PaletteSize6UVColorCdf = DEFAULT_PALETTE_SIZE_6_UV_COLOR_CDF
	state.PaletteSize7YColorCdf = DEFAULT_PALETTE_SIZE_7_Y_COLOR_CDF
	state.PaletteSize7UVColorCdf = DEFAULT_PALETTE_SIZE_7_UV_COLOR_CDF
	state.PaletteSize8YColorCdf = DEFAULT_PALETTE_SIZE_8_Y_COLOR_CDF
	state.PaletteSize8UVColorCdf = DEFAULT_PALETTE_SIZE_8_UV_COLOR_CDF

	state.DeltaQCdf = DEFAULT_DELTA_Q_CDF
	state.DeltaLFCdf = DEFAULT_DELTA_LF_CDF

	for i := 0; i < shared.FRAME_LF_COUNT; i++ {
		state.DeltaLFMultiCdf[i] = DEFAULT_DELTA_LF_CDF
	}

	state.IntraTxTypeSet1Cdf = DEFAULT_INTRA_TX_TYPE_SET1_CDF
	state.IntraTxTypeSet2Cdf = DEFAULT_INTRA_TX_TYPE_SET2_CDF
	state.InterTxTypeSet1Cdf = DEFAULT_INTER_TX_TYPE_SET1_CDF
	state.InterTxTypeSet2Cdf = DEFAULT_INTER_TX_TYPE_SET2_CDF
	state.InterTxTypeSet3Cdf = DEFAULT_INTER_TX_TYPE_SET3_CDF
	state.UseObmcCdf = DEFAULT_USE_OBMC_CDF
	state.InterIntraCdf = DEFAULT_INTER_INTRA_CDF
	state.CompRefTypeCdf = DEFAULT_COMP_REF_TYPE_CDF
	state.CflSignCdf = DEFAULT_CFL_SIGN_CDF
	state.UniCompRefCdf = DEFAULT_UNI_COMP_REF_CDF
	state.WedgeInterIntraCdf = DEFAULT_WEDGE_INTER_INTRA_CDF
	state.CompGroupIdxCdf = DEFAULT_COMP_GROUP_IDX_CDF
	state.CompoundIdxCdf = DEFAULT_COMPOUND_IDX_CDF
	state.CompoundTypeCdf = DEFAULT_COMPOUND_TYPE_CDF
	state.InterIntraModeCdf = DEFAULT_INTER_INTRA_MODE_CDF
	state.WedgeIndexCdf = DEFAULT_WEDGE_INDEX_CDF
	state.CflAlphaCdf = DEFAULT_CFL_ALPHA_CDF
	state.UseWienerCdf = DEFAULT_USE_WIENER_CDF
	state.UseSgrprojCdf = DEFAULT_USE_SGRPROJ_CDF
	state.RestorationTypeCdf = DEFAULT_RESTORATION_TYPE_CDF
}

// init_coeff_cdfs()
func initCoeffCdfs(state *state.State, uh UncompressedHeader) {
	var idx int
	if uh.BaseQIdx <= 20 {
		idx = 0
	} else if uh.BaseQIdx <= 60 {
		idx = 1
	} else if uh.BaseQIdx <= 120 {
		idx = 2
	} else {
		idx = 3
	}

	state.TxbSkipCdf = DEFAULT_TXB_SKIP_CDF[idx]
	state.EobPt16Cdf = DEFAULT_EOB_PT_16_CDF[idx]
	state.EobPt32Cdf = DEFAULT_EOB_PT_32_CDF[idx]
	state.EobPt64Cdf = DEFAULT_EOB_PT_64_CDF[idx]

	state.EobPt128Cdf = DEFAULT_EOB_PT_128_CDF[idx]
	state.EobPt256Cdf = DEFAULT_EOB_PT_256_CDF[idx]
	state.EobPt512Cdf = DEFAULT_EOB_PT_512_CDF[idx]
	state.EobPt1024Cdf = DEFAULT_EOB_PT_1024_CDF[idx]
	state.EobExtraCdf = DEFAULT_EOB_EXTRA_CDF[idx]
	state.DcSignCdf = DEFAULT_DC_SIGN_CDF[idx]
	state.CoeffBaseEobCdf = DEFAULT_COEFF_BASE_EOB_CDF[idx]
	state.CoeffBaseCdf = DEFAULT_COEFF_BASE_CDF[idx]
	state.CoeffBrCdf = DEFAULT_COEFF_BR_CDF[idx]
}

// save_cdfs( i )
func SaveCdfs(ctx int, s *state.State) {
	memoryArea := state.MemoryArea{}

	memoryArea.YModeCdf = s.YModeCdf
	memoryArea.UVModeCflNotAllowedCdf = s.UVModeCflNotAllowedCdf
	memoryArea.UVModeCflAllowedCdf = s.UVModeCflAllowedCdf
	memoryArea.AngleDeltaCdf = s.AngleDeltaCdf
	memoryArea.IntrabcCdf = s.IntrabcCdf
	memoryArea.PartitionW8Cdf = s.PartitionW8Cdf
	memoryArea.PartitionW16Cdf = s.PartitionW16Cdf
	memoryArea.PartitionW32Cdf = s.PartitionW32Cdf
	memoryArea.PartitionW64Cdf = s.PartitionW64Cdf
	memoryArea.PartitionW128Cdf = s.PartitionW128Cdf
	memoryArea.SegmentIdCdf = s.SegmentIdCdf
	memoryArea.SegmentIdPredictedCdf = s.SegmentIdPredictedCdf
	memoryArea.Tx8x8Cdf = s.Tx8x8Cdf
	memoryArea.Tx16x16Cdf = s.Tx16x16Cdf
	memoryArea.Tx32x32Cdf = s.Tx32x32Cdf
	memoryArea.Tx64x64Cdf = s.Tx64x64Cdf
	memoryArea.TxfmSplitCdf = s.TxfmSplitCdf
	memoryArea.FilterIntraModeCdf = s.FilterIntraModeCdf
	memoryArea.FilterIntraCdf = s.FilterIntraCdf
	memoryArea.InterpFilterCdf = s.InterpFilterCdf
	memoryArea.MotionModeCdf = s.MotionModeCdf
	memoryArea.NewMvCdf = s.NewMvCdf
	memoryArea.ZeroMvCdf = s.ZeroMvCdf
	memoryArea.RefMvCdf = s.RefMvCdf
	memoryArea.CompoundModeCdf = s.CompoundModeCdf
	memoryArea.DrlModeCdf = s.DrlModeCdf
	memoryArea.IsInterCdf = s.IsInterCdf
	memoryArea.CompModeCdf = s.CompModeCdf
	memoryArea.SkipModeCdf = s.SkipModeCdf
	memoryArea.SkipCdf = s.SkipCdf
	memoryArea.CompRefCdf = s.CompRefCdf
	memoryArea.CompBwdRefCdf = s.CompBwdRefCdf
	memoryArea.SingleRefCdf = s.SingleRefCdf
	memoryArea.MvJointCdf = s.MvJointCdf
	memoryArea.MvClassCdf = s.MvClassCdf
	memoryArea.MvClass0BitCdf = s.MvClass0BitCdf
	memoryArea.MvFrCdf = s.MvFrCdf
	memoryArea.MvClass0FrCdf = s.MvClass0FrCdf
	memoryArea.MvClass0HpCdf = s.MvClass0HpCdf
	memoryArea.MvSignCdf = s.MvSignCdf
	memoryArea.MvBitCdf = s.MvBitCdf
	memoryArea.MvHpCdf = s.MvHpCdf
	memoryArea.PaletteYModeCdf = s.PaletteYModeCdf
	memoryArea.PaletteUVModeCdf = s.PaletteUVModeCdf
	memoryArea.PaletteUVSizeCdf = s.PaletteUVSizeCdf
	memoryArea.PaletteSize2YColorCdf = s.PaletteSize2YColorCdf
	memoryArea.PaletteSize2UVColorCdf = s.PaletteSize2UVColorCdf
	memoryArea.PaletteSize3YColorCdf = s.PaletteSize3YColorCdf
	memoryArea.PaletteSize3UVColorCdf = s.PaletteSize3UVColorCdf
	memoryArea.PaletteSize4YColorCdf = s.PaletteSize4YColorCdf
	memoryArea.PaletteSize4UVColorCdf = s.PaletteSize4UVColorCdf
	memoryArea.PaletteSize5YColorCdf = s.PaletteSize5YColorCdf
	memoryArea.PaletteSize5UVColorCdf = s.PaletteSize5UVColorCdf
	memoryArea.PaletteSize6YColorCdf = s.PaletteSize6YColorCdf
	memoryArea.PaletteSize6UVColorCdf = s.PaletteSize6UVColorCdf
	memoryArea.PaletteSize7YColorCdf = s.PaletteSize7YColorCdf
	memoryArea.PaletteSize7UVColorCdf = s.PaletteSize7UVColorCdf
	memoryArea.PaletteSize8YColorCdf = s.PaletteSize8YColorCdf
	memoryArea.PaletteSize8UVColorCdf = s.PaletteSize8UVColorCdf

	memoryArea.DeltaQCdf = s.DeltaQCdf
	memoryArea.DeltaLFCdf = s.DeltaLFCdf
	memoryArea.DeltaLFMultiCdf = s.DeltaLFMultiCdf
	memoryArea.IntraTxTypeSet1Cdf = s.IntraTxTypeSet1Cdf
	memoryArea.IntraTxTypeSet2Cdf = s.IntraTxTypeSet2Cdf
	memoryArea.InterTxTypeSet1Cdf = s.InterTxTypeSet1Cdf
	memoryArea.InterTxTypeSet2Cdf = s.InterTxTypeSet2Cdf
	memoryArea.InterTxTypeSet3Cdf = s.InterTxTypeSet3Cdf
	memoryArea.UseObmcCdf = s.UseObmcCdf
	memoryArea.InterIntraCdf = s.InterIntraCdf
	memoryArea.CompRefTypeCdf = s.CompRefTypeCdf
	memoryArea.CflSignCdf = s.CflSignCdf
	memoryArea.UniCompRefCdf = s.UniCompRefCdf
	memoryArea.WedgeInterIntraCdf = s.WedgeInterIntraCdf
	memoryArea.CompGroupIdxCdf = s.CompGroupIdxCdf
	memoryArea.CompoundIdxCdf = s.CompoundIdxCdf
	memoryArea.CompoundTypeCdf = s.CompoundTypeCdf
	memoryArea.InterIntraModeCdf = s.InterIntraModeCdf
	memoryArea.WedgeIndexCdf = s.WedgeIndexCdf
	memoryArea.CflAlphaCdf = s.CflAlphaCdf
	memoryArea.UseWienerCdf = s.UseWienerCdf
	memoryArea.UseSgrprojCdf = s.UseSgrprojCdf
	memoryArea.RestorationTypeCdf = s.RestorationTypeCdf

	memoryArea.TxbSkipCdf = s.TxbSkipCdf
	memoryArea.EobPt16Cdf = s.EobPt16Cdf
	memoryArea.EobPt32Cdf = s.EobPt32Cdf
	memoryArea.EobPt64Cdf = s.EobPt64Cdf
	memoryArea.EobPt128Cdf = s.EobPt128Cdf
	memoryArea.EobPt256Cdf = s.EobPt256Cdf
	memoryArea.EobPt512Cdf = s.EobPt512Cdf
	memoryArea.EobPt1024Cdf = s.EobPt1024Cdf
	memoryArea.EobExtraCdf = s.EobExtraCdf
	memoryArea.DcSignCdf = s.DcSignCdf
	memoryArea.CoeffBaseEobCdf = s.CoeffBaseEobCdf
	memoryArea.CoeffBaseCdf = s.CoeffBaseCdf
	memoryArea.CoeffBrCdf = s.CoeffBrCdf

	s.Memory[ctx] = memoryArea
}

func LoadCdfs(ctx int, s *state.State) {
	memoryArea := s.Memory[ctx]

	s.YModeCdf = memoryArea.YModeCdf
	s.UVModeCflNotAllowedCdf = memoryArea.UVModeCflNotAllowedCdf
	s.UVModeCflAllowedCdf = memoryArea.UVModeCflAllowedCdf
	s.AngleDeltaCdf = memoryArea.AngleDeltaCdf
	s.IntrabcCdf = memoryArea.IntrabcCdf
	s.PartitionW8Cdf = memoryArea.PartitionW8Cdf
	s.PartitionW16Cdf = memoryArea.PartitionW16Cdf
	s.PartitionW32Cdf = memoryArea.PartitionW32Cdf
	s.PartitionW64Cdf = memoryArea.PartitionW64Cdf
	s.PartitionW128Cdf = memoryArea.PartitionW128Cdf
	s.SegmentIdCdf = memoryArea.SegmentIdCdf
	s.SegmentIdPredictedCdf = memoryArea.SegmentIdPredictedCdf
	s.Tx8x8Cdf = memoryArea.Tx8x8Cdf
	s.Tx16x16Cdf = memoryArea.Tx16x16Cdf
	s.Tx32x32Cdf = memoryArea.Tx32x32Cdf
	s.Tx64x64Cdf = memoryArea.Tx64x64Cdf
	s.TxfmSplitCdf = memoryArea.TxfmSplitCdf
	s.FilterIntraModeCdf = memoryArea.FilterIntraModeCdf
	s.FilterIntraCdf = memoryArea.FilterIntraCdf
	s.InterpFilterCdf = memoryArea.InterpFilterCdf
	s.MotionModeCdf = memoryArea.MotionModeCdf
	s.NewMvCdf = memoryArea.NewMvCdf
	s.ZeroMvCdf = memoryArea.ZeroMvCdf
	s.RefMvCdf = memoryArea.RefMvCdf
	s.CompoundModeCdf = memoryArea.CompoundModeCdf
	s.DrlModeCdf = memoryArea.DrlModeCdf
	s.IsInterCdf = memoryArea.IsInterCdf
	s.CompModeCdf = memoryArea.CompModeCdf
	s.SkipModeCdf = memoryArea.SkipModeCdf
	s.SkipCdf = memoryArea.SkipCdf
	s.CompRefCdf = memoryArea.CompRefCdf
	s.CompBwdRefCdf = memoryArea.CompBwdRefCdf
	s.SingleRefCdf = memoryArea.SingleRefCdf
	s.MvJointCdf = memoryArea.MvJointCdf
	s.MvClassCdf = memoryArea.MvClassCdf
	s.MvClass0BitCdf = memoryArea.MvClass0BitCdf
	s.MvFrCdf = memoryArea.MvFrCdf
	s.MvClass0FrCdf = memoryArea.MvClass0FrCdf
	s.MvClass0HpCdf = memoryArea.MvClass0HpCdf
	s.MvSignCdf = memoryArea.MvSignCdf
	s.MvBitCdf = memoryArea.MvBitCdf
	s.MvHpCdf = memoryArea.MvHpCdf
	s.PaletteYModeCdf = memoryArea.PaletteYModeCdf
	s.PaletteUVModeCdf = memoryArea.PaletteUVModeCdf
	s.PaletteUVSizeCdf = memoryArea.PaletteUVSizeCdf
	s.PaletteSize2YColorCdf = memoryArea.PaletteSize2YColorCdf
	s.PaletteSize2UVColorCdf = memoryArea.PaletteSize2UVColorCdf
	s.PaletteSize3YColorCdf = memoryArea.PaletteSize3YColorCdf
	s.PaletteSize3UVColorCdf = memoryArea.PaletteSize3UVColorCdf
	s.PaletteSize4YColorCdf = memoryArea.PaletteSize4YColorCdf
	s.PaletteSize4UVColorCdf = memoryArea.PaletteSize4UVColorCdf
	s.PaletteSize5YColorCdf = memoryArea.PaletteSize5YColorCdf
	s.PaletteSize5UVColorCdf = memoryArea.PaletteSize5UVColorCdf
	s.PaletteSize6YColorCdf = memoryArea.PaletteSize6YColorCdf
	s.PaletteSize6UVColorCdf = memoryArea.PaletteSize6UVColorCdf
	s.PaletteSize7YColorCdf = memoryArea.PaletteSize7YColorCdf
	s.PaletteSize7UVColorCdf = memoryArea.PaletteSize7UVColorCdf
	s.PaletteSize8YColorCdf = memoryArea.PaletteSize8YColorCdf
	s.PaletteSize8UVColorCdf = memoryArea.PaletteSize8UVColorCdf

	s.DeltaQCdf = memoryArea.DeltaQCdf
	s.DeltaLFCdf = memoryArea.DeltaLFCdf
	s.DeltaLFMultiCdf = memoryArea.DeltaLFMultiCdf
	s.IntraTxTypeSet1Cdf = memoryArea.IntraTxTypeSet1Cdf
	s.IntraTxTypeSet2Cdf = memoryArea.IntraTxTypeSet2Cdf
	s.InterTxTypeSet1Cdf = memoryArea.InterTxTypeSet1Cdf
	s.InterTxTypeSet2Cdf = memoryArea.InterTxTypeSet2Cdf
	s.InterTxTypeSet3Cdf = memoryArea.InterTxTypeSet3Cdf
	s.UseObmcCdf = memoryArea.UseObmcCdf
	s.InterIntraCdf = memoryArea.InterIntraCdf
	s.CompRefTypeCdf = memoryArea.CompRefTypeCdf
	s.CflSignCdf = memoryArea.CflSignCdf
	s.UniCompRefCdf = memoryArea.UniCompRefCdf
	s.WedgeInterIntraCdf = memoryArea.WedgeInterIntraCdf
	s.CompGroupIdxCdf = memoryArea.CompGroupIdxCdf
	s.CompoundIdxCdf = memoryArea.CompoundIdxCdf
	s.CompoundTypeCdf = memoryArea.CompoundTypeCdf
	s.InterIntraModeCdf = memoryArea.InterIntraModeCdf
	s.WedgeIndexCdf = memoryArea.WedgeIndexCdf
	s.CflAlphaCdf = memoryArea.CflAlphaCdf
	s.UseWienerCdf = memoryArea.UseWienerCdf
	s.UseSgrprojCdf = memoryArea.UseSgrprojCdf
	s.RestorationTypeCdf = memoryArea.RestorationTypeCdf

	s.TxbSkipCdf = memoryArea.TxbSkipCdf
	s.EobPt16Cdf = memoryArea.EobPt16Cdf
	s.EobPt32Cdf = memoryArea.EobPt32Cdf
	s.EobPt64Cdf = memoryArea.EobPt64Cdf
	s.EobPt128Cdf = memoryArea.EobPt128Cdf
	s.EobPt256Cdf = memoryArea.EobPt256Cdf
	s.EobPt512Cdf = memoryArea.EobPt512Cdf
	s.EobPt1024Cdf = memoryArea.EobPt1024Cdf
	s.EobExtraCdf = memoryArea.EobExtraCdf
	s.DcSignCdf = memoryArea.DcSignCdf
	s.CoeffBaseEobCdf = memoryArea.CoeffBaseEobCdf
	s.CoeffBaseCdf = memoryArea.CoeffBaseCdf
	s.CoeffBrCdf = memoryArea.CoeffBrCdf

	// TODO: set last entry of every array to 0
}

var DEFAULT_INTER_TX_TYPE_SET2_CDF = []int{
	770, 2421, 5225, 12907, 15819, 18927, 21561, 24089, 26595,
	28526, 30529, 32768, 0,
}

var DEFAULT_INTER_TX_TYPE_SET1_CDF = [][]int{
	{
		4458, 5560, 7695, 9709, 13330, 14789, 17537, 20266, 21504, 22848, 23934, 25474, 27727, 28915, 30631, 32768, 0},
	{
		1645, 2573, 4778, 5711, 7807, 8622, 10522, 15357, 17674, 20408, 22517, 25010, 27116, 28856, 30749, 32768, 0},
}

var DEFAULT_RESTORATION_TYPE_CDF = []int{
	9413, 22581, 32768, 0,
}

var DEFAULT_USE_SGRPROJ_CDF = []int{
	16855, 32768, 0,
}

var DEFAULT_USE_WIENER_CDF = []int{
	11570, 32768, 0,
}

var DEFAULT_CFL_ALPHA_CDF = [][]int{
	{
		7637, 20719, 31401, 32481, 32657, 32688, 32692, 32696, 32700, 32704, 32708, 32712, 32716, 32720, 32724, 32768, 0},
	{
		14365, 23603, 28135, 31168, 32167, 32395, 32487, 32573, 32620, 32647, 32668, 32672, 32676, 32680, 32684, 32768, 0},
	{
		11532, 22380, 28445, 31360, 32349, 32523, 32584, 32649, 32673, 32677, 32681, 32685, 32689, 32693, 32697, 32768, 0},
	{
		26990, 31402, 32282, 32571, 32692, 32696, 32700, 32704, 32708, 32712, 32716, 32720, 32724, 32728, 32732, 32768, 0},
	{
		17248, 26058, 28904, 30608, 31305, 31877, 32126, 32321, 32394, 32464, 32516, 32560, 32576, 32593, 32622, 32768, 0},
	{
		14738, 21678, 25779, 27901, 29024, 30302, 30980, 31843, 32144, 32413, 32520, 32594, 32622, 32656, 32660, 32768, 0},
}

var DEFAULT_WEDGE_INDEX_CDF = [][]int{
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2438, 4440, 6599, 8663, 11005, 12874, 15751, 18094, 20359, 22362, 24127, 25702, 27752, 29450, 31171, 32768, 0},
	{
		806, 3266, 6005, 6738, 7218, 7367, 7771, 14588, 16323, 17367, 18452, 19422, 22839, 26127, 29629, 32768, 0},
	{
		2779, 3738, 4683, 7213, 7775, 8017, 8655, 14357, 17939, 21332, 24520, 27470, 29456, 30529, 31656, 32768, 0},
	{
		1684, 3625, 5675, 7108, 9302, 11274, 14429, 17144, 19163, 20961, 22884, 24471, 26719, 28714, 30877, 32768, 0},
	{
		1142, 3491, 6277, 7314, 8089, 8355, 9023, 13624, 15369, 16730, 18114, 19313, 22521, 26012, 29550, 32768, 0},
	{
		2742, 4195, 5727, 8035, 8980, 9336, 10146, 14124, 17270, 20533, 23434, 25972, 27944, 29570, 31416, 32768, 0},
	{
		1727, 3948, 6101, 7796, 9841, 12344, 15766, 18944, 20638, 22038, 23963, 25311, 26988, 28766, 31012, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		154, 987, 1925, 2051, 2088, 2111, 2151, 23033, 23703, 24284, 24985, 25684, 27259, 28883, 30911, 32768, 0},
	{
		1135, 1322, 1493, 2635, 2696, 2737, 2770, 21016, 22935, 25057, 27251, 29173, 30089, 30960, 31933, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
	{
		2048, 4096, 6144, 8192, 10240, 12288, 14336, 16384, 18432, 20480, 22528, 24576, 26624, 28672, 30720, 32768, 0},
}

var DEFAULT_INTER_INTRA_MODE_CDF = [][]int{
	{
		1875, 11082, 27332, 32768, 0},
	{
		2473, 9996, 26388, 32768, 0},
	{
		4238, 11537, 25926, 32768, 0},
}

var DEFAULT_COMPOUND_TYPE_CDF = [][]int{
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		23431, 32768, 0},
	{
		13171, 32768, 0},
	{
		11470, 32768, 0},
	{
		9770, 32768, 0},
	{
		9100, 32768, 0},
	{
		8233, 32768, 0},
	{
		6172, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		11820, 32768, 0},
	{
		7701, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
}

var DEFAULT_COMPOUND_IDX_CDF = [][]int{
	{
		18244, 32768, 0},
	{
		12865, 32768, 0},
	{
		7053, 32768, 0},
	{
		13259, 32768, 0},
	{
		9334, 32768, 0},
	{
		4644, 32768, 0},
}

var DEFAULT_COMP_GROUP_IDX_CDF = [][]int{
	{
		26607, 32768, 0},
	{
		22891, 32768, 0},
	{
		18840, 32768, 0},
	{
		24594, 32768, 0},
	{
		19934, 32768, 0},
	{
		22674, 32768, 0},
}

var DEFAULT_WEDGE_INTER_INTRA_CDF = [][]int{
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		20036, 32768, 0},
	{
		24957, 32768, 0},
	{
		26704, 32768, 0},
	{
		27530, 32768, 0},
	{
		29564, 32768, 0},
	{
		29444, 32768, 0},
	{
		26872, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
}

var DEFAULT_COMP_REF_TYPE_CDF = [][]int{
	{
		1198, 32768, 0},
	{
		2070, 32768, 0},
	{
		9166, 32768, 0},
	{
		7499, 32768, 0},
	{
		22475, 32768, 0},
}

var DEFAULT_UNI_COMP_REF_CDF = [][][]int{
	{
		{
			5284, 32768, 0},
		{
			3865, 32768, 0},
		{
			3128, 32768, 0},
	},
	{
		{
			23152, 32768, 0},
		{
			14173, 32768, 0},
		{
			15270, 32768, 0},
	},
	{
		{
			31774, 32768, 0},
		{
			25120, 32768, 0},
		{
			26710, 32768, 0},
	},
}

var DEFAULT_CFL_SIGN_CDF = []int{
	1418, 2123, 13340, 18405, 26972, 28343, 32294, 32768, 0,
}

var Default_Comp_Ref_Type_Cdf = [][]int{
	{
		1198, 32768, 0},
	{
		2070, 32768, 0},
	{
		9166, 32768, 0},
	{
		7499, 32768, 0},
	{
		22475, 32768, 0},
}

var DEFAULT_INTER_INTRA_CDF = [][]int{
	{
		26887, 32768, 00},
	{
		27597, 32768, 00},
	{
		30237, 32768, 00},
}

var DEFAULT_USE_OBMC_CDF = [][]int{
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		10437, 32768, 0},
	{
		9371, 32768, 0},
	{
		9301, 32768, 0},
	{
		17432, 32768, 0},
	{
		14423, 32768, 0},
	{
		15142, 32768, 0},
	{
		25817, 32768, 0},
	{
		22823, 32768, 0},
	{
		22083, 32768, 0},
	{
		30128, 32768, 0},
	{
		31014, 32768, 0},
	{
		31560, 32768, 0},
	{
		32638, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		23664, 32768, 0},
	{
		20901, 32768, 0},
	{
		24008, 32768, 0},
	{
		26879, 32768, 0},
}

var DEFAULT_INTER_TX_TYPE_SET3_CDF = [][]int{
	{
		16384, 32768, 0},
	{
		4167, 32768, 0},
	{
		1998, 32768, 0},
	{
		748, 32768, 0},
}

var DEFAULT_INTRA_TX_TYPE_SET2_CDF = [][][]int{
	{
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
	},
	{
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
		{
			6554, 13107, 19661, 26214, 32768, 0},
	},
	{
		{
			1127, 12814, 22772, 27483, 32768, 0},
		{
			145, 6761, 11980, 26667, 32768, 0},
		{
			362, 5887, 11678, 16725, 32768, 0},
		{
			385, 15213, 18587, 30693, 32768, 0},
		{
			25, 2914, 23134, 27903, 32768, 0},
		{
			60, 4470, 11749, 23991, 32768, 0},
		{
			37, 3332, 14511, 21448, 32768, 0},
		{
			157, 6320, 13036, 17439, 32768, 0},
		{
			119, 6719, 12906, 29396, 32768, 0},
		{
			47, 5537, 12576, 21499, 32768, 0},
		{
			269, 6076, 11258, 23115, 32768, 0},
		{
			83, 5615, 12001, 17228, 32768, 0},
		{
			1968, 5556, 12023, 18547, 32768, 0},
	},
}

var DEFAULT_INTRA_TX_TYPE_SET1_CDF = [][][]int{
	{
		{
			1535, 8035, 9461, 12751, 23467, 27825, 32768, 0},
		{
			564, 3335, 9709, 10870, 18143, 28094, 32768, 0},
		{
			672, 3247, 3676, 11982, 19415, 23127, 32768, 0},
		{
			5279, 13885, 15487, 18044, 23527, 30252, 32768, 0},
		{
			4423, 6074, 7985, 10416, 25693, 29298, 32768, 0},
		{
			1486, 4241, 9460, 10662, 16456, 27694, 32768, 0},
		{
			439, 2838, 3522, 6737, 18058, 23754, 32768, 0},
		{
			1190, 4233, 4855, 11670, 20281, 24377, 32768, 0},
		{
			1045, 4312, 8647, 10159, 18644, 29335, 32768, 0},
		{
			202, 3734, 4747, 7298, 17127, 24016, 32768, 0},
		{
			447, 4312, 6819, 8884, 16010, 23858, 32768, 0},
		{
			277, 4369, 5255, 8905, 16465, 22271, 32768, 0},
		{
			3409, 5436, 10599, 15599, 19687, 24040, 32768, 0},
	},
	{
		{
			1870, 13742, 14530, 16498, 23770, 27698, 32768, 0},
		{
			326, 8796, 14632, 15079, 19272, 27486, 32768, 0},
		{
			484, 7576, 7712, 14443, 19159, 22591, 32768, 0},
		{
			1126, 15340, 15895, 17023, 20896, 30279, 32768, 0},
		{
			655, 4854, 5249, 5913, 22099, 27138, 32768, 0},
		{
			1299, 6458, 8885, 9290, 14851, 25497, 32768, 0},
		{
			311, 5295, 5552, 6885, 16107, 22672, 32768, 0},
		{
			883, 8059, 8270, 11258, 17289, 21549, 32768, 0},
		{
			741, 7580, 9318, 10345, 16688, 29046, 32768, 0},
		{
			110, 7406, 7915, 9195, 16041, 23329, 32768, 0},
		{
			363, 7974, 9357, 10673, 15629, 24474, 32768, 0},
		{
			153, 7647, 8112, 9936, 15307, 19996, 32768, 0},
		{
			3511, 6332, 11165, 15335, 19323, 23594, 32768, 0},
	},
}

var DEFAULT_DELTA_LF_CDF = []int{
	28160, 32120, 32677, 32768, 0,
}

var DEFAULT_DELTA_Q_CDF = []int{
	28160, 32120, 32677, 32768, 0,
}

var DEFAULT_PALETTE_SIZE_8_UV_COLOR_CDF = [][]int{
	{
		21442, 23288, 24758, 26246, 27649, 28980, 30563, 32768, 0},
	{
		5863, 14933, 17552, 20668, 23683, 26411, 29273, 32768, 0},
	{
		3415, 25810, 26877, 27990, 29223, 30394, 31618, 32768, 0},
	{
		17965, 20084, 22232, 23974, 26274, 28402, 30390, 32768, 0},
	{
		31190, 31329, 31516, 31679, 31825, 32026, 32322, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_8_Y_COLOR_CDF = [][]int{
	{
		21689, 23883, 25163, 26352, 27506, 28827, 30195, 32768, 0},
	{
		6892, 15385, 17840, 21606, 24287, 26753, 29204, 32768, 0},
	{
		5651, 23182, 25042, 26518, 27982, 29392, 30900, 32768, 0},
	{
		19349, 22578, 24418, 25994, 27524, 29031, 30448, 32768, 0},
	{
		31028, 31270, 31504, 31705, 31927, 32153, 32392, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_7_UV_COLOR_CDF = [][]int{
	{
		21239, 23168, 25044, 26962, 28705, 30506, 32768, 0},
	{
		6545, 15012, 18004, 21817, 25503, 28701, 32768, 0},
	{
		3448, 26295, 27437, 28704, 30126, 31442, 32768, 0},
	{
		15889, 18323, 21704, 24698, 26976, 29690, 32768, 0},
	{
		30988, 31204, 31479, 31734, 31983, 32325, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_7_Y_COLOR_CDF = [][]int{
	{
		23105, 25199, 26464, 27684, 28931, 30318, 32768, 0},
	{
		6950, 15447, 18952, 22681, 25567, 28563, 32768, 0},
	{
		7560, 23474, 25490, 27203, 28921, 30708, 32768, 0},
	{
		18544, 22373, 24457, 26195, 28119, 30045, 32768, 0},
	{
		31198, 31451, 31670, 31882, 32123, 32391, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_6_UV_COLOR_CDF = [][]int{
	{
		22217, 24567, 26637, 28683, 30548, 32768, 0},
	{
		7307, 16406, 19636, 24632, 28424, 32768, 0},
	{
		4441, 25064, 26879, 28942, 30919, 32768, 0},
	{
		17210, 20528, 23319, 26750, 29582, 32768, 0},
	{
		30674, 30953, 31396, 31735, 32207, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_6_Y_COLOR_CDF = [][]int{
	{
		23132, 25407, 26970, 28435, 30073, 32768, 0},
	{
		7443, 17242, 20717, 24762, 27982, 32768, 0},
	{
		6300, 24862, 26944, 28784, 30671, 32768, 0},
	{
		18916, 22895, 25267, 27435, 29652, 32768, 0},
	{
		31270, 31550, 31808, 32059, 32353, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_5_UV_COLOR_CDF = [][]int{
	{
		24779, 26955, 28576, 30282, 32768, 0},
	{
		8669, 20364, 24073, 28093, 32768, 0},
	{
		4255, 27565, 29377, 31067, 32768, 0},
	{
		19864, 23674, 26716, 29530, 32768, 0},
	{
		31646, 31893, 32147, 32426, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_5_Y_COLOR_CDF = [][]int{
	{
		24779, 26955, 28576, 30282, 32768, 0},
	{
		8669, 20364, 24073, 28093, 32768, 0},
	{
		4255, 27565, 29377, 31067, 32768, 0},
	{
		19864, 23674, 26716, 29530, 32768, 0},
	{
		31646, 31893, 32147, 32426, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_4_UV_COLOR_CDF = [][]int{
	{
		24210, 27175, 29903, 32768, 0},
	{
		9888, 22386, 27214, 32768, 0},
	{
		5901, 26053, 29293, 32768, 0},
	{
		18318, 22152, 28333, 32768, 0},
	{
		30459, 31136, 31926, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_4_Y_COLOR_CDF = [][]int{
	{
		25572, 28046, 30045, 32768, 0},
	{
		9478, 21590, 27256, 32768, 0},
	{
		7248, 26837, 29824, 32768, 0},
	{
		19167, 24486, 28349, 32768, 0},
	{
		31400, 31825, 32250, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_3_UV_COLOR_CDF = [][]int{
	{
		25257, 29145, 32768, 0},
	{
		12287, 27293, 32768, 0},
	{
		7033, 27960, 32768, 0},
	{
		20145, 25405, 32768, 0},
	{
		30608, 31639, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_3_Y_COLOR_CDF = [][]int{
	{
		27877, 30490, 32768, 0},
	{
		11532, 25697, 32768, 0},
	{
		6544, 30234, 32768, 0},
	{
		23018, 28072, 32768, 0},
	{
		31915, 32385, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_2_UV_COLOR_CDF = [][]int{
	{
		29089, 32768, 0},
	{
		16384, 32768, 0},
	{
		8713, 32768, 0},
	{
		29257, 32768, 0},
	{
		31610, 32768, 0},
}

var DEFAULT_PALETTE_SIZE_2_Y_COLOR_CDF = [][]int{
	{
		28710, 32768, 0},
	{
		16384, 32768, 0},
	{
		10553, 32768, 0},
	{
		27036, 32768, 0},
	{
		31603, 32768, 0},
}

var DEFAULT_PALETTE_UV_SIZE_CDF = [][]int{
	{
		8713, 19979, 27128, 29609, 31331, 32272, 32768, 0},
	{
		5839, 15573, 23581, 26947, 29848, 31700, 32768, 0},
	{
		4426, 11260, 17999, 21483, 25863, 29430, 32768, 0},
	{
		3228, 9464, 14993, 18089, 22523, 27420, 32768, 0},
	{
		3768, 8886, 13091, 17852, 22495, 27207, 32768, 0},
	{
		2464, 8451, 12861, 21632, 25525, 28555, 32768, 0},
	{
		1269, 5435, 10433, 18963, 21700, 25865, 32768, 0},
}

var DEFAULT_PALETTE_Y_SIZE_CDF = [][]int{
	{
		7952, 13000, 18149, 21478, 25527, 29241, 32768, 0},
	{
		7139, 11421, 16195, 19544, 23666, 28073, 32768, 0},
	{
		7788, 12741, 17325, 20500, 24315, 28530, 32768, 0},
	{
		8271, 14064, 18246, 21564, 25071, 28533, 32768, 0},
	{
		12725, 19180, 21863, 24839, 27535, 30120, 32768, 0},
	{
		9711, 14888, 16923, 21052, 25661, 27875, 32768, 0},
	{
		14940, 20797, 21678, 24186, 27033, 28999, 32768, 0},
}

var DEFAULT_PALETTE_UV_MODE_CDF = [][]int{
	{
		32461, 32768, 0},
	{
		21488, 32768, 0},
}

var DEFAULT_PALETTE_Y_MODE_CDF = [][][]int{
	{
		{
			31676, 32768, 0},
		{
			3419, 32768, 0},
		{
			1261, 32768, 0},
	},
	{
		{
			31912, 32768, 0},
		{
			2859, 32768, 0},
		{
			980, 32768, 0},
	},
	{
		{
			31823, 32768, 0},
		{
			3400, 32768, 0},
		{
			781, 32768, 0},
	},
	{
		{
			32030, 32768, 0},
		{
			3561, 32768, 0},
		{
			904, 32768, 0},
	},
	{
		{
			32309, 32768, 0},
		{
			7337, 32768, 0},
		{
			1462, 32768, 0},
	},
	{
		{
			32265, 32768, 0},
		{
			4015, 32768, 0},
		{
			1521, 32768, 0},
	},
	{
		{
			32450, 32768, 0},
		{
			7946, 32768, 0},
		{
			129, 32768, 0},
	},
}

var DEFAULT_MV_HP_CDF = []int{128 * 128, 32768, 0}

var DEFAULT_MV_BIT_CDF = [][]int{
	{
		136 * 128, 32768, 0},
	{
		140 * 128, 32768, 0},
	{
		148 * 128, 32768, 0},
	{
		160 * 128, 32768, 0},
	{
		176 * 128, 32768, 0},
	{
		192 * 128, 32768, 0},
	{
		224 * 128, 32768, 0},
	{
		234 * 128, 32768, 0},
	{
		234 * 128, 32768, 0},
	{
		240 * 128, 32768, 0},
}

var DEFAULT_MV_SIGN_CDF = []int{128 * 128, 32768, 0}

var DEFAULT_MV_CLASS0_HP_CDF = []int{160 * 128, 32768, 0}

var DEFAULT_MV_CLASS0_FR_CDF = [][][]int{
	{
		{
			16384, 24576, 26624, 32768, 0},
		{
			12288, 21248, 24128, 32768, 0},
	},
	{
		{
			16384, 24576, 26624, 32768, 0},
		{
			12288, 21248, 24128, 32768, 0},
	},
}

var DEFAULT_MV_FR_CDF = [][]int{
	{
		8192, 17408, 21248, 32768, 0},
	{
		8192, 17408, 21248, 32768, 0},
}

var DEFAULT_MV_CLASS0_BIT_CDF = []int{
	216 * 128, 32768, 0,
}

var DEFAULT_MV_CLASS_CDF = [][]int{
	{
		28672, 30976, 31858, 32320, 32551, 32656, 32740, 32757, 32762, 32767, 32768, 0},
	{
		28672, 30976, 31858, 32320, 32551, 32656, 32740, 32757, 32762, 32767, 32768, 0},
}

var DEFAULT_MV_JOINT_CDF = []int{
	4096, 11264, 19328, 32768, 0,
}

var DEFAULT_SINGLE_REF_CDF = [][][]int{
	{
		{
			4897, 32768, 0},
		{
			1555, 32768, 0},
		{
			4236, 32768, 0},
		{
			8650, 32768, 0},
		{
			904, 32768, 0},
		{
			1444, 32768, 0},
	},
	{
		{
			16973, 32768, 0},
		{
			16751, 32768, 0},
		{
			19647, 32768, 0},
		{
			24773, 32768, 0},
		{
			11014, 32768, 0},
		{
			15087, 32768, 0},
	},
	{
		{
			29744, 32768, 0},
		{
			30279, 32768, 0},
		{
			31194, 32768, 0},
		{
			31895, 32768, 0},
		{
			26875, 32768, 0},
		{
			30304, 32768, 0},
	},
}

var DEFAULT_COMP_BWD_REF_CDF = [][][]int{
	{
		{
			2235, 32768, 0},
		{
			1423, 32768, 0},
	},
	{
		{
			17182, 32768, 0},
		{
			15175, 32768, 0},
	},
	{
		{
			30606, 32768, 0},
		{
			30489, 32768, 0},
	},
}

var DEFAULT_COMP_REF_CDF = [][][]int{
	{
		{
			4946, 32768, 0},
		{
			9468, 32768, 0},
		{
			1503, 32768, 0},
	},
	{
		{
			19891, 32768, 0},
		{
			22441, 32768, 0},
		{
			15160, 32768, 0},
	},
	{
		{
			30731, 32768, 0},
		{
			31059, 32768, 0},
		{
			27544, 32768, 0},
	},
}

var DEFAULT_SKIP_CDF = [][]int{
	{
		31671, 32768, 0},
	{
		16515, 32768, 0},
	{
		4576, 32768, 0},
}

var DEFAULT_SKIP_MODE_CDF = [][]int{
	{3262132621,
		32768, 00},
	{2070820708,
		32768, 00},
	{81278127,
		32768, 00},
}

var DEFAULT_COMP_MODE_CDF = [][]int{
	{
		26828, 32768, 0},
	{
		24035, 32768, 0},
	{
		12031, 32768, 0},
	{
		10640, 32768, 0},
	{
		2901, 32768, 0},
}

var DEFAULT_IS_INTER_CDF = [][]int{
	{
		806, 32768, 0},
	{
		16662, 32768, 0},
	{
		20186, 32768, 0},
	{
		26538, 32768, 0},
}

var DEFAULT_DRL_MODE_CDF = [][]int{
	{
		13104, 32768, 0},
	{
		24560, 32768, 0},
	{
		18945, 32768, 0},
}

var DEFAULT_COMPOUND_MODE_CDF = [][]int{
	{
		7760, 13823, 15808, 17641, 19156, 20666, 26891, 32768, 0},
	{
		10730, 19452, 21145, 22749, 24039, 25131, 28724, 32768, 0},
	{
		10664, 20221, 21588, 22906, 24295, 25387, 28436, 32768, 0},
	{
		13298, 16984, 20471, 24182, 25067, 25736, 26422, 32768, 0},
	{
		18904, 23325, 25242, 27432, 27898, 28258, 30758, 32768, 0},
	{
		10725, 17454, 20124, 22820, 24195, 25168, 26046, 32768, 0},
	{
		17125, 24273, 25814, 27492, 28214, 28704, 30592, 32768, 0},
	{
		13046, 23214, 24505, 25942, 27435, 28442, 29330, 32768, 0},
}

var DFEAULT_ZERO_MV_CDF = [][]int{
	{2175, 32768, 0},
	{1054, 32768, 0},
}

var DEFAULT_NEW_MV_CDF = [][]int{
	{
		24035, 32768, 0},
	{
		16630, 32768, 0},
	{
		15339, 32768, 0},
	{
		8386, 32768, 0},
	{
		12222, 32768, 0},
	{
		4676, 32768, 0},
}

var DEFAULT_MOTION_MODE_CDF = [][]int{
	{
		10923, 21845, 32768, 0},
	{
		10923, 21845, 32768, 0},
	{
		10923, 21845, 32768, 0},
	{
		7651, 24760, 32768, 0},
	{
		4738, 24765, 32768, 0},
	{
		5391, 25528, 32768, 0},
	{
		19419, 26810, 32768, 0},
	{
		5123, 23606, 32768, 0},
	{
		11606, 24308, 32768, 0},
	{
		26260, 29116, 32768, 0},
	{
		20360, 28062, 32768, 0},
	{
		21679, 26830, 32768, 0},
	{
		29516, 30701, 32768, 0},
	{
		28898, 30397, 32768, 0},
	{
		30878, 31335, 32768, 0},
	{
		32507, 32558, 32768, 0},
	{
		10923, 21845, 32768, 0},
	{
		10923, 21845, 32768, 0},
	{
		28799, 31390, 32768, 0},
	{
		26431, 30774, 32768, 0},
	{
		28973, 31594, 32768, 0},
	{
		29742, 31203, 32768, 0},
}

var DEFAULT_INTERP_FILTER_CDF = [][]int{
	{
		31935, 32720, 32768, 0},
	{
		5568, 32719, 32768, 0},
	{
		422, 2938, 32768, 0},
	{
		28244, 32608, 32768, 0},
	{
		31206, 31953, 32768, 0},
	{
		4862, 32121, 32768, 0},
	{
		770, 1152, 32768, 0},
	{
		20889, 25637, 32768, 0},
	{
		31910, 32724, 32768, 0},
	{
		4120, 32712, 32768, 0},
	{
		305, 2247, 32768, 0},
	{
		27403, 32636, 32768, 0},
	{
		31022, 32009, 32768, 0},
	{
		2963, 32093, 32768, 0},
	{
		601, 943, 32768, 0},
	{
		14969, 21398, 32768, 0},
}

var DEFAULT_FILTER_INTRA_MODE_CDF = []int{
	8949, 12776, 17211, 29558, 32768, 0,
}

var DEFAULT_FILTER_INTRA_CDF = [][]int{
	{
		4621, 32768, 0},
	{
		6743, 32768, 0},
	{
		5893, 32768, 0},
	{
		7866, 32768, 0},
	{
		12551, 32768, 0},
	{
		9394, 32768, 0},
	{
		12408, 32768, 0},
	{
		14301, 32768, 0},
	{
		12756, 32768, 0},
	{
		22343, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
	{
		12770, 32768, 0},
	{
		10368, 32768, 0},
	{
		20229, 32768, 0},
	{
		18101, 32768, 0},
	{
		16384, 32768, 0},
	{
		16384, 32768, 0},
}

var DEFAULT_TXFM_SPLIT_CDF = [][]int{
	{
		28581, 32768, 0},
	{
		23846, 32768, 0},
	{
		20847, 32768, 0},
	{
		24315, 32768, 0},
	{
		18196, 32768, 0},
	{
		12133, 32768, 0},
	{
		18791, 32768, 0},
	{
		10887, 32768, 0},
	{
		11005, 32768, 0},
	{
		27179, 32768, 0},
	{
		20004, 32768, 0},
	{
		11281, 32768, 0},
	{
		26549, 32768, 0},
	{
		19308, 32768, 0},
	{
		14224, 32768, 0},
	{
		28015, 32768, 0},
	{
		21546, 32768, 0},
	{
		14400, 32768, 0},
	{
		28165, 32768, 0},
	{
		22401, 32768, 0},
	{
		16088, 32768, 0},
}

var DEFAULT_TX_64X64_CDF = [][]int{
	{
		5782, 11475, 32768, 0},
	{
		5782, 11475, 32768, 0},
	{
		16803, 22759, 32768, 0},
}

var DEFAULT_TX_32X32_CDF = [][]int{
	{
		12986, 15180, 32768, 0},
	{
		12986, 15180, 32768, 0},
	{
		24302, 25602, 32768, 0},
}

var DEFAULT_TX_16X16_CDF = [][]int{
	{
		12272, 30172, 32768, 0},
	{
		12272, 30172, 32768, 0},
	{
		18677, 30848, 32768, 0},
}

var DEFAULT_TX_8X8_CDF = [][]int{
	{
		19968, 32768, 0},
	{
		19968, 32768, 0},
	{
		24320, 32768, 0},
}

var DEFAULT_SEGMENT_ID_PREDICTED_CDF = [][]int{
	{128 * 128, 32678, 0},
	{128 * 128, 32678, 0},
	{128 * 128, 32678, 0},
}

var DEFAULT_SEGMENT_ID_CDF = [][]int{
	{
		5622, 7893, 16093, 18233, 27809, 28373, 32533, 32768, 0},
	{
		14274, 18230, 22557, 24935, 29980, 30851, 32344, 32768, 0},
	{
		27527, 28487, 28723, 28890, 32397, 32647, 32679, 32768, 0},
}

var DEFAULT_PARTITION_W128_CDF = [][]int{
	{
		27899, 28219, 28529, 32484, 32539, 32619, 32639, 32768, 0},
	{
		6607, 6990, 8268, 32060, 32219, 32338, 32371, 32768, 0},
	{
		5429, 6676, 7122, 32027, 32227, 32531, 32582, 32768, 0},
	{
		711, 966, 1172, 32448, 32538, 32617, 32664, 32768, 0},
}
var DEFAULT_PARTITION_W64_CDF = [][]int{
	{
		20137, 21547, 23078, 29566, 29837, 30261, 30524, 30892, 31724, 32768, 0},
	{
		6732, 7490, 9497, 27944, 28250, 28515, 28969, 29630, 30104, 32768, 0},
	{
		5945, 7663, 8348, 28683, 29117, 29749, 30064, 30298, 32238, 32768, 0},
	{
		870, 1212, 1487, 31198, 31394, 31574, 31743, 31881, 32332, 32768, 0},
}

var DEFAULT_PARTITION_W32_CDF = [][]int{
	{
		18462, 20920, 23124, 27647, 28227, 29049, 29519, 30178, 31544, 32768, 0},
	{
		7689, 9060, 12056, 24992, 25660, 26182, 26951, 28041, 29052, 32768, 0},
	{
		6015, 9009, 10062, 24544, 25409, 26545, 27071, 27526, 32047, 32768, 0},
	{
		1394, 2208, 2796, 28614, 29061, 29466, 29840, 30185, 31899, 32768, 0},
}

var DEFAULT_PARTITION_W16_CDF = [][]int{
	{
		15597, 20929, 24571, 26706, 27664, 28821, 29601, 30571, 31902, 32768, 0},
	{
		7925, 11043, 16785, 22470, 23971, 25043, 26651, 28701, 29834, 32768, 0},
	{
		5414, 13269, 15111, 20488, 22360, 24500, 25537, 26336, 32117, 32768, 0},
	{
		2662, 6362, 8614, 20860, 23053, 24778, 26436, 27829, 31171, 32768, 0},
}

var DEFAULT_PARTITION_W8_CDF = [][]int{
	{
		19132, 25510, 30392, 32768, 0},
	{
		13928, 19855, 28540, 32768, 0},
	{
		12522, 23679, 28629, 32768, 0},
	{
		9896, 18783, 25853, 32768, 0},
}

var DEFAULT_INTRABC_CDF = []int{
	30531, 32768, 0,
}

var DEFAULT_ANGLE_DELTA_CDF = [][]int{
	{
		2180, 5032, 7567, 22776, 26989, 30217, 32768, 0},
	{
		2301, 5608, 8801, 23487, 26974, 30330, 32768, 0},
	{
		3780, 11018, 13699, 19354, 23083, 31286, 32768, 0},
	{
		4581, 11226, 15147, 17138, 21834, 28397, 32768, 0},
	{
		1737, 10927, 14509, 19588, 22745, 28823, 32768, 0},
	{
		2664, 10176, 12485, 17650, 21600, 30495, 32768, 0},
	{
		2240, 11096, 15453, 20341, 22561, 28917, 32768, 0},
	{
		3605, 10428, 12459, 17676, 21244, 30655, 32768, 0},
}

var DEFAULT_UV_MODE_CFL_ALLOWED_CDF = [][]int{
	{
		10407, 11208, 12900, 13181, 13823, 14175, 14899, 15656, 15986, 20086, 20995, 22455, 24212, 32768, 0},
	{
		4532, 19780, 20057, 20215, 20428, 21071, 21199, 21451, 22099, 24228, 24693, 27032, 29472, 32768, 0},
	{
		5273, 5379, 20177, 20270, 20385, 20439, 20949, 21695, 21774, 23138, 24256, 24703, 26679, 32768, 0},
	{
		6740, 7167, 7662, 14152, 14536, 14785, 15034, 16741, 18371, 21520, 22206, 23389, 24182, 32768, 0},
	{
		4987, 5368, 5928, 6068, 19114, 20315, 21857, 22253, 22411, 24911, 25380, 26027, 26376, 32768, 0},
	{
		5370, 6889, 7247, 7393, 9498, 21114, 21402, 21753, 21981, 24780, 25386, 26517, 27176, 32768, 0},
	{
		4816, 4961, 7204, 7326, 8765, 8930, 20169, 20682, 20803, 23188, 23763, 24455, 24940, 32768, 0},
	{
		6608, 6740, 8529, 9049, 9257, 9356, 9735, 18827, 19059, 22336, 23204, 23964, 24793, 32768, 0},
	{
		5998, 7419, 7781, 8933, 9255, 9549, 9753, 10417, 18898, 22494, 23139, 24764, 25989, 32768, 0},
	{
		10660, 11298, 12550, 12957, 13322, 13624, 14040, 15004, 15534, 20714, 21789, 23443, 24861, 32768, 0},
	{
		10522, 11530, 12552, 12963, 13378, 13779, 14245, 15235, 15902, 20102, 22696, 23774, 25838, 32768, 0},
	{
		10099, 10691, 12639, 13049, 13386, 13665, 14125, 15163, 15636, 19676, 20474, 23519, 25208, 32768, 0},
	{
		3144, 5087, 7382, 7504, 7593, 7690, 7801, 8064, 8232, 9248, 9875, 10521, 29048, 32768, 0},
}

var DEFAULT_Y_MODE_CDF = [][]int{
	{
		22801,
		23489,
		24293,
		24756,
		25601,
		26123,
		26606,
		27418,
		27945,
		29228,
		29685,
		30349,
		32768,
		0,
	},
	{
		18673,
		19845,
		22631,
		23318,
		23950,
		24649,
		25527,
		27364,
		28152,
		29701,
		29984,
		30852,
		32768,
		0,
	},
	{
		19770,
		20979,
		23396,
		23939,
		24241,
		24654,
		25136,
		27073,
		27830,
		29360,
		29730,
		30659,
		32768,
		0,
	},
	{
		20155,
		21301,
		22838,
		23178,
		23261,
		23533,
		23703,
		24804,
		25352,
		26575,
		27016,
		28049,
		32768,
		0,
	},
}

var DEFAULT_UV_MODE_CFL_NOT_ALLOWED_CDF = [][]int{
	{
		22631,
		24152,
		25378,
		25661,
		25986,
		26520,
		27055,
		27923,
		28244,
		30059,
		30941,
		31961,
		32768,
		0,
	},
	{
		9513,
		26881,
		26973,
		27046,
		27118,
		27664,
		27739,
		27824,
		28359,
		29505,
		29800,
		31796,
		32768,
		0,
	},
	{
		9845,
		9915,
		28663,
		28704,
		28757,
		28780,
		29198,
		29822,
		29854,
		30764,
		31777,
		32029,
		32768,
		0,
	},
	{
		13639,
		13897,
		14171,
		25331,
		25606,
		25727,
		25953,
		27148,
		28577,
		30612,
		31355,
		32493,
		32768,
		0,
	},
	{
		9764,
		9835,
		9930,
		9954,
		25386,
		27053,
		27958,
		28148,
		28243,
		31101,
		31744,
		32363,
		32768,
		0,
	},
	{
		11825,
		13589,
		13677,
		13720,
		15048,
		29213,
		29301,
		29458,
		29711,
		31161,
		31441,
		32550,
		32768,
		0,
	},
	{
		14175,
		14399,
		16608,
		16821,
		17718,
		17775,
		28551,
		30200,
		30245,
		31837,
		32342,
		32667,
		32768,
		0,
	},
	{
		12885,
		13038,
		14978,
		15590,
		15673,
		15748,
		16176,
		29128,
		29267,
		30643,
		31961,
		32461,
		32768,
		0,
	},
	{
		12026,
		13661,
		13874,
		15305,
		15490,
		15726,
		15995,
		16273,
		28443,
		30388,
		30767,
		32416,
		32768,
		0,
	},
	{
		19052,
		19840,
		20579,
		20916,
		21150,
		21467,
		21885,
		22719,
		23174,
		28861,
		30379,
		32175,
		32768,
		0,
	},
	{
		18627,
		19649,
		20974,
		21219,
		21492,
		21816,
		22199,
		23119,
		23527,
		27053,
		31397,
		32148,
		32768,
		0,
	},
	{
		17026,
		19004,
		19997,
		20339,
		20586,
		21103,
		21349,
		21907,
		22482,
		25896,
		26541,
		31819,
		32768,
		0,
	},
	{
		12124,
		13759,
		14959,
		14992,
		15007,
		15051,
		15078,
		15166,
		15255,
		15753,
		16039,
		16606,
		32768,
		0,
	},
}

var DEFAULT_INTRA_FRAME_Y_MODE_CDF = [][][]int{
	{
		{
			15588,
			17027,
			19338,
			20218,
			20682,
			21110,
			21825,
			23244,
			24189,
			28165,
			29093,
			30466,
			32768,
			0,
		},
		{
			12016,
			18066,
			19516,
			20303,
			20719,
			21444,
			21888,
			23032,
			24434,
			28658,
			30172,
			31409,
			32768,
			0,
		},
		{
			10052,
			10771,
			22296,
			22788,
			23055,
			23239,
			24133,
			25620,
			26160,
			29336,
			29929,
			31567,
			32768,
			0,
		},
		{
			14091,
			15406,
			16442,
			18808,
			19136,
			19546,
			19998,
			22096,
			24746,
			29585,
			30958,
			32462,
			32768,
			0,
		},
		{
			12122,
			13265,
			15603,
			16501,
			18609,
			20033,
			22391,
			25583,
			26437,
			30261,
			31073,
			32475,
			32768,
			0,
		},
	},
	{
		{
			10023,
			19585,
			20848,
			21440,
			21832,
			22760,
			23089,
			24023,
			25381,
			29014,
			30482,
			31436,
			32768,
			0,
		},
		{
			5983,
			24099,
			24560,
			24886,
			25066,
			25795,
			25913,
			26423,
			27610,
			29905,
			31276,
			31794,
			32768,
			0,
		},
		{
			7444,
			12781,
			20177,
			20728,
			21077,
			21607,
			22170,
			23405,
			24469,
			27915,
			29090,
			30492,
			32768,
			0,
		},
		{
			8537,
			14689,
			15432,
			17087,
			17408,
			18172,
			18408,
			19825,
			24649,
			29153,
			31096,
			32210,
			32768,
			0,
		},
		{
			7543,
			14231,
			15496,
			16195,
			17905,
			20717,
			21984,
			24516,
			26001,
			29675,
			30981,
			31994,
			32768,
			0,
		},
	},
	{
		{
			12613,
			13591,
			21383,
			22004,
			22312,
			22577,
			23401,
			25055,
			25729,
			29538,
			30305,
			32077,
			32768,
			0,
		},
		{
			9687,
			13470,
			18506,
			19230,
			19604,
			20147,
			20695,
			22062,
			23219,
			27743,
			29211,
			30907,
			32768,
			0,
		},
		{
			6183,
			6505,
			26024,
			26252,
			26366,
			26434,
			27082,
			28354,
			28555,
			30467,
			30794,
			32086,
			32768,
			0,
		},
		{
			10718,
			11734,
			14954,
			17224,
			17565,
			17924,
			18561,
			21523,
			23878,
			28975,
			30287,
			32252,
			32768,
			0,
		},
		{
			9194,
			9858,
			16501,
			17263,
			18424,
			19171,
			21563,
			25961,
			26561,
			30072,
			30737,
			32463,
			32768,
			0,
		},
	},
	{
		{
			12602,
			14399,
			15488,
			18381,
			18778,
			19315,
			19724,
			21419,
			25060,
			29696,
			30917,
			32409,
			32768,
			0,
		},
		{
			8203,
			13821,
			14524,
			17105,
			17439,
			18131,
			18404,
			19468,
			25225,
			29485,
			31158,
			32342,
			32768,
			0,
		},
		{
			8451,
			9731,
			15004,
			17643,
			18012,
			18425,
			19070,
			21538,
			24605,
			29118,
			30078,
			32018,
			32768,
			0,
		},
		{
			7714,
			9048,
			9516,
			16667,
			16817,
			16994,
			17153,
			18767,
			26743,
			30389,
			31536,
			32528,
			32768,
			0,
		},
		{
			8843,
			10280,
			11496,
			15317,
			16652,
			17943,
			19108,
			22718,
			25769,
			29953,
			30983,
			32485,
			32768,
			0,
		},
	},
	{
		{
			12578,
			13671,
			15979,
			16834,
			19075,
			20913,
			22989,
			25449,
			26219,
			30214,
			31150,
			32477,
			32768,
			0,
		},
		{
			9563,
			13626,
			15080,
			15892,
			17756,
			20863,
			22207,
			24236,
			25380,
			29653,
			31143,
			32277,
			32768,
			0,
		},
		{
			8356,
			8901,
			17616,
			18256,
			19350,
			20106,
			22598,
			25947,
			26466,
			29900,
			30523,
			32261,
			32768,
			0,
		},
		{
			10835,
			11815,
			13124,
			16042,
			17018,
			18039,
			18947,
			22753,
			24615,
			29489,
			30883,
			32482,
			32768,
			0,
		},
		{
			7618,
			8288,
			9859,
			10509,
			15386,
			18657,
			22903,
			28776,
			29180,
			31355,
			31802,
			32593,
			32768,
			0,
		},
	},
}

var DEFAULT_TXB_SKIP_CDF = [][][][]int{
	{
		{
			{
				31849,
				32768,
				0,
			},
			{
				5892,
				32768,
				0,
			},
			{
				12112,
				32768,
				0,
			},
			{
				21935,
				32768,
				0,
			},
			{
				20289,
				32768,
				0,
			},
			{
				27473,
				32768,
				0,
			},
			{
				32487,
				32768,
				0,
			},
			{
				7654,
				32768,
				0,
			},
			{
				19473,
				32768,
				0,
			},
			{
				29984,
				32768,
				0,
			},
			{
				9961,
				32768,
				0,
			},
			{
				30242,
				32768,
				0,
			},
			{
				32117,
				32768,
				0,
			},
		},
		{
			{
				31548,
				32768,
				0,
			},
			{
				1549,
				32768,
				0,
			},
			{
				10130,
				32768,
				0,
			},
			{
				16656,
				32768,
				0,
			},
			{
				18591,
				32768,
				0,
			},
			{
				26308,
				32768,
				0,
			},
			{
				32537,
				32768,
				0,
			},
			{
				5403,
				32768,
				0,
			},
			{
				18096,
				32768,
				0,
			},
			{
				30003,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				29957,
				32768,
				0,
			},
			{
				5391,
				32768,
				0,
			},
			{
				18039,
				32768,
				0,
			},
			{
				23566,
				32768,
				0,
			},
			{
				22431,
				32768,
				0,
			},
			{
				25822,
				32768,
				0,
			},
			{
				32197,
				32768,
				0,
			},
			{
				3778,
				32768,
				0,
			},
			{
				15336,
				32768,
				0,
			},
			{
				28981,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				17920,
				32768,
				0,
			},
			{
				1818,
				32768,
				0,
			},
			{
				7282,
				32768,
				0,
			},
			{
				25273,
				32768,
				0,
			},
			{
				10923,
				32768,
				0,
			},
			{
				31554,
				32768,
				0,
			},
			{
				32624,
				32768,
				0,
			},
			{
				1366,
				32768,
				0,
			},
			{
				15628,
				32768,
				0,
			},
			{
				30462,
				32768,
				0,
			},
			{
				146,
				32768,
				0,
			},
			{
				5132,
				32768,
				0,
			},
			{
				31657,
				32768,
				0,
			},
		},
		{
			{
				6308,
				32768,
				0,
			},
			{
				117,
				32768,
				0,
			},
			{
				1638,
				32768,
				0,
			},
			{
				2161,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				10923,
				32768,
				0,
			},
			{
				30247,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				30371,
				32768,
				0,
			},
			{
				7570,
				32768,
				0,
			},
			{
				13155,
				32768,
				0,
			},
			{
				20751,
				32768,
				0,
			},
			{
				20969,
				32768,
				0,
			},
			{
				27067,
				32768,
				0,
			},
			{
				32013,
				32768,
				0,
			},
			{
				5495,
				32768,
				0,
			},
			{
				17942,
				32768,
				0,
			},
			{
				28280,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				31782,
				32768,
				0,
			},
			{
				1836,
				32768,
				0,
			},
			{
				10689,
				32768,
				0,
			},
			{
				17604,
				32768,
				0,
			},
			{
				21622,
				32768,
				0,
			},
			{
				27518,
				32768,
				0,
			},
			{
				32399,
				32768,
				0,
			},
			{
				4419,
				32768,
				0,
			},
			{
				16294,
				32768,
				0,
			},
			{
				28345,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				31901,
				32768,
				0,
			},
			{
				10311,
				32768,
				0,
			},
			{
				18047,
				32768,
				0,
			},
			{
				24806,
				32768,
				0,
			},
			{
				23288,
				32768,
				0,
			},
			{
				27914,
				32768,
				0,
			},
			{
				32296,
				32768,
				0,
			},
			{
				4215,
				32768,
				0,
			},
			{
				15756,
				32768,
				0,
			},
			{
				28341,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				26726,
				32768,
				0,
			},
			{
				1045,
				32768,
				0,
			},
			{
				11703,
				32768,
				0,
			},
			{
				20590,
				32768,
				0,
			},
			{
				18554,
				32768,
				0,
			},
			{
				25970,
				32768,
				0,
			},
			{
				31938,
				32768,
				0,
			},
			{
				5583,
				32768,
				0,
			},
			{
				21313,
				32768,
				0,
			},
			{
				29390,
				32768,
				0,
			},
			{
				641,
				32768,
				0,
			},
			{
				22265,
				32768,
				0,
			},
			{
				31452,
				32768,
				0,
			},
		},
		{
			{
				26584,
				32768,
				0,
			},
			{
				188,
				32768,
				0,
			},
			{
				8847,
				32768,
				0,
			},
			{
				24519,
				32768,
				0,
			},
			{
				22938,
				32768,
				0,
			},
			{
				30583,
				32768,
				0,
			},
			{
				32608,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				29614,
				32768,
				0,
			},
			{
				9068,
				32768,
				0,
			},
			{
				12924,
				32768,
				0,
			},
			{
				19538,
				32768,
				0,
			},
			{
				17737,
				32768,
				0,
			},
			{
				24619,
				32768,
				0,
			},
			{
				30642,
				32768,
				0,
			},
			{
				4119,
				32768,
				0,
			},
			{
				16026,
				32768,
				0,
			},
			{
				25657,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				31957,
				32768,
				0,
			},
			{
				3230,
				32768,
				0,
			},
			{
				11153,
				32768,
				0,
			},
			{
				18123,
				32768,
				0,
			},
			{
				20143,
				32768,
				0,
			},
			{
				26536,
				32768,
				0,
			},
			{
				31986,
				32768,
				0,
			},
			{
				3050,
				32768,
				0,
			},
			{
				14603,
				32768,
				0,
			},
			{
				25155,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				32363,
				32768,
				0,
			},
			{
				10692,
				32768,
				0,
			},
			{
				19090,
				32768,
				0,
			},
			{
				24357,
				32768,
				0,
			},
			{
				24442,
				32768,
				0,
			},
			{
				28312,
				32768,
				0,
			},
			{
				32169,
				32768,
				0,
			},
			{
				3648,
				32768,
				0,
			},
			{
				15690,
				32768,
				0,
			},
			{
				26815,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				30669,
				32768,
				0,
			},
			{
				3832,
				32768,
				0,
			},
			{
				11663,
				32768,
				0,
			},
			{
				18889,
				32768,
				0,
			},
			{
				19782,
				32768,
				0,
			},
			{
				23313,
				32768,
				0,
			},
			{
				31330,
				32768,
				0,
			},
			{
				5124,
				32768,
				0,
			},
			{
				18719,
				32768,
				0,
			},
			{
				28468,
				32768,
				0,
			},
			{
				3082,
				32768,
				0,
			},
			{
				20982,
				32768,
				0,
			},
			{
				29443,
				32768,
				0,
			},
		},
		{
			{
				28573,
				32768,
				0,
			},
			{
				3183,
				32768,
				0,
			},
			{
				17802,
				32768,
				0,
			},
			{
				25977,
				32768,
				0,
			},
			{
				26677,
				32768,
				0,
			},
			{
				27832,
				32768,
				0,
			},
			{
				32387,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				26887,
				32768,
				0,
			},
			{
				6729,
				32768,
				0,
			},
			{
				10361,
				32768,
				0,
			},
			{
				17442,
				32768,
				0,
			},
			{
				15045,
				32768,
				0,
			},
			{
				22478,
				32768,
				0,
			},
			{
				29072,
				32768,
				0,
			},
			{
				2713,
				32768,
				0,
			},
			{
				11861,
				32768,
				0,
			},
			{
				20773,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				31903,
				32768,
				0,
			},
			{
				2044,
				32768,
				0,
			},
			{
				7528,
				32768,
				0,
			},
			{
				14618,
				32768,
				0,
			},
			{
				16182,
				32768,
				0,
			},
			{
				24168,
				32768,
				0,
			},
			{
				31037,
				32768,
				0,
			},
			{
				2786,
				32768,
				0,
			},
			{
				11194,
				32768,
				0,
			},
			{
				20155,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				32510,
				32768,
				0,
			},
			{
				8430,
				32768,
				0,
			},
			{
				17318,
				32768,
				0,
			},
			{
				24154,
				32768,
				0,
			},
			{
				23674,
				32768,
				0,
			},
			{
				28789,
				32768,
				0,
			},
			{
				32139,
				32768,
				0,
			},
			{
				3440,
				32768,
				0,
			},
			{
				13117,
				32768,
				0,
			},
			{
				22702,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
		{
			{
				31671,
				32768,
				0,
			},
			{
				2056,
				32768,
				0,
			},
			{
				11746,
				32768,
				0,
			},
			{
				16852,
				32768,
				0,
			},
			{
				18635,
				32768,
				0,
			},
			{
				24715,
				32768,
				0,
			},
			{
				31484,
				32768,
				0,
			},
			{
				4656,
				32768,
				0,
			},
			{
				16074,
				32768,
				0,
			},
			{
				24704,
				32768,
				0,
			},
			{
				1806,
				32768,
				0,
			},
			{
				14645,
				32768,
				0,
			},
			{
				25336,
				32768,
				0,
			},
		},
		{
			{
				31539,
				32768,
				0,
			},
			{
				8433,
				32768,
				0,
			},
			{
				20576,
				32768,
				0,
			},
			{
				27904,
				32768,
				0,
			},
			{
				27852,
				32768,
				0,
			},
			{
				30026,
				32768,
				0,
			},
			{
				32441,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
			{
				16384,
				32768,
				0,
			},
		},
	},
}

var DEFAULT_EOB_PT_16_CDF = [][][][]int{{
	{
		{
			840,
			1039,
			1980,
			4895,
			32768,
			0,
		},
		{
			370,
			671,
			1883,
			4471,
			32768,
			0,
		},
	},
	{
		{
			3247,
			4950,
			9688,
			14563,
			32768,
			0,
		},
		{
			1904,
			3354,
			7763,
			14647,
			32768,
			0,
		},
	},
},
	{
		{
			{
				2125,
				2551,
				5165,
				8946,
				32768,
				0,
			},
			{
				513,
				765,
				1859,
				6339,
				32768,
				0,
			},
		},
		{
			{
				7637,
				9498,
				14259,
				19108,
				32768,
				0,
			},
			{
				2497,
				4096,
				8866,
				16993,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				4016,
				4897,
				8881,
				14968,
				32768,
				0,
			},
			{
				716,
				1105,
				2646,
				10056,
				32768,
				0,
			},
		},
		{
			{
				11139,
				13270,
				18241,
				23566,
				32768,
				0,
			},
			{
				3192,
				5032,
				10297,
				19755,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				6708,
				8958,
				14746,
				22133,
				32768,
				0,
			},
			{
				1222,
				2074,
				4783,
				15410,
				32768,
				0,
			},
		},
		{
			{
				19575,
				21766,
				26044,
				29709,
				32768,
				0,
			},
			{
				7297,
				10767,
				19273,
				28194,
				32768,
				0,
			},
		},
	},
}

var DEFAULT_EOB_PT_32_CDF = [][][][]int{
	{
		{
			{
				400,
				520,
				977,
				2102,
				6542,
				32768,
				0,
			},
			{
				210,
				405,
				1315,
				3326,
				7537,
				32768,
				0,
			},
		},
		{
			{
				2636,
				4273,
				7588,
				11794,
				20401,
				32768,
				0,
			},
			{
				1786,
				3179,
				6902,
				11357,
				19054,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				989,
				1249,
				2019,
				4151,
				10785,
				32768,
				0,
			},
			{
				313,
				441,
				1099,
				2917,
				8562,
				32768,
				0,
			},
		},
		{
			{
				8394,
				10352,
				13932,
				18855,
				26014,
				32768,
				0,
			},
			{
				2578,
				4124,
				8181,
				13670,
				24234,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				2515,
				3003,
				4452,
				8162,
				16041,
				32768,
				0,
			},
			{
				574,
				821,
				1836,
				5089,
				13128,
				32768,
				0,
			},
		},
		{
			{
				13468,
				16303,
				20361,
				25105,
				29281,
				32768,
				0,
			},
			{
				3542,
				5502,
				10415,
				16760,
				25644,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				4617,
				5709,
				8446,
				13584,
				23135,
				32768,
				0,
			},
			{
				1156,
				1702,
				3675,
				9274,
				20539,
				32768,
				0,
			},
		},
		{
			{
				22086,
				24282,
				27010,
				29770,
				31743,
				32768,
				0,
			},
			{
				7699,
				10897,
				20891,
				26926,
				31628,
				32768,
				0,
			},
		},
	},
}

var DEFAULT_EOB_PT_64_CDF = [][][][]int{
	{
		{
			{
				329,
				498,
				1101,
				1784,
				3265,
				7758,
				32768,
				0,
			},
			{
				335,
				730,
				1459,
				5494,
				8755,
				12997,
				32768,
				0,
			},
		},
		{
			{
				3505,
				5304,
				10086,
				13814,
				17684,
				23370,
				32768,
				0,
			},
			{
				1563,
				2700,
				4876,
				10911,
				14706,
				22480,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				1260,
				1446,
				2253,
				3712,
				6652,
				13369,
				32768,
				0,
			},
			{
				401,
				605,
				1029,
				2563,
				5845,
				12626,
				32768,
				0,
			},
		},
		{
			{
				8609,
				10612,
				14624,
				18714,
				22614,
				29024,
				32768,
				0,
			},
			{
				1923,
				3127,
				5867,
				9703,
				14277,
				27100,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				2374,
				2772,
				4583,
				7276,
				12288,
				19706,
				32768,
				0,
			},
			{
				497,
				810,
				1315,
				3000,
				7004,
				15641,
				32768,
				0,
			},
		},
		{
			{
				15050,
				17126,
				21410,
				24886,
				28156,
				30726,
				32768,
				0,
			},
			{
				4034,
				6290,
				10235,
				14982,
				21214,
				28491,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				6307,
				7541,
				12060,
				16358,
				22553,
				27865,
				32768,
				0,
			},
			{
				1289,
				2320,
				3971,
				7926,
				14153,
				24291,
				32768,
				0,
			},
		},
		{
			{
				24212,
				25708,
				28268,
				30035,
				31307,
				32049,
				32768,
				0,
			},
			{
				8726,
				12378,
				19409,
				26450,
				30038,
				32462,
				32768,
				0,
			},
		},
	},
}

var DEFAULT_EOB_PT_128_CDF = [][][][]int{
	{
		{
			{
				219,
				482,
				1140,
				2091,
				3680,
				6028,
				12586,
				32768,
				0,
			},
			{
				371,
				699,
				1254,
				4830,
				9479,
				12562,
				17497,
				32768,
				0,
			},
		},
		{
			{
				5245,
				7456,
				12880,
				15852,
				20033,
				23932,
				27608,
				32768,
				0,
			},
			{
				2054,
				3472,
				5869,
				14232,
				18242,
				20590,
				26752,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				685,
				933,
				1488,
				2714,
				4766,
				8562,
				19254,
				32768,
				0,
			},
			{
				217,
				352,
				618,
				2303,
				5261,
				9969,
				17472,
				32768,
				0,
			},
		},
		{
			{
				8045,
				11200,
				15497,
				19595,
				23948,
				27408,
				30938,
				32768,
				0,
			},
			{
				2310,
				4160,
				7471,
				14997,
				17931,
				20768,
				30240,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				1366,
				1738,
				2527,
				5016,
				9355,
				15797,
				24643,
				32768,
				0,
			},
			{
				354,
				558,
				944,
				2760,
				7287,
				14037,
				21779,
				32768,
				0,
			},
		},
		{
			{
				13627,
				16246,
				20173,
				24429,
				27948,
				30415,
				31863,
				32768,
				0,
			},
			{
				6275,
				9889,
				14769,
				23164,
				27988,
				30493,
				32272,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				3472,
				4885,
				7489,
				12481,
				18517,
				24536,
				29635,
				32768,
				0,
			},
			{
				886,
				1731,
				3271,
				8469,
				15569,
				22126,
				28383,
				32768,
				0,
			},
		},
		{
			{
				24313,
				26062,
				28385,
				30107,
				31217,
				31898,
				32345,
				32768,
				0,
			},
			{
				9165,
				13282,
				21150,
				30286,
				31894,
				32571,
				32712,
				32768,
				0,
			},
		},
	},
}

var DEFAULT_EOB_PT_256_CDF = [][][][]int{
	{
		{
			{
				310,
				584,
				1887,
				3589,
				6168,
				8611,
				11352,
				15652,
				32768,
				0,
			},
			{
				998,
				1850,
				2998,
				5604,
				17341,
				19888,
				22899,
				25583,
				32768,
				0,
			},
		},
		{
			{
				2520,
				3240,
				5952,
				8870,
				12577,
				17558,
				19954,
				24168,
				32768,
				0,
			},
			{
				2203,
				4130,
				7435,
				10739,
				20652,
				23681,
				25609,
				27261,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				1448,
				2109,
				4151,
				6263,
				9329,
				13260,
				17944,
				23300,
				32768,
				0,
			},
			{
				399,
				1019,
				1749,
				3038,
				10444,
				15546,
				22739,
				27294,
				32768,
				0,
			},
		},
		{
			{
				6402,
				8148,
				12623,
				15072,
				18728,
				22847,
				26447,
				29377,
				32768,
				0,
			},
			{
				1674,
				3252,
				5734,
				10159,
				22397,
				23802,
				24821,
				30940,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				3089,
				3920,
				6038,
				9460,
				14266,
				19881,
				25766,
				29176,
				32768,
				0,
			},
			{
				1084,
				2358,
				3488,
				5122,
				11483,
				18103,
				26023,
				29799,
				32768,
				0,
			},
		},
		{
			{
				11514,
				13794,
				17480,
				20754,
				24361,
				27378,
				29492,
				31277,
				32768,
				0,
			},
			{
				6571,
				9610,
				15516,
				21826,
				29092,
				30829,
				31842,
				32708,
				32768,
				0,
			},
		},
	},
	{
		{
			{
				5348,
				7113,
				11820,
				15924,
				22106,
				26777,
				30334,
				31757,
				32768,
				0,
			},
			{
				2453,
				4474,
				6307,
				8777,
				16474,
				22975,
				29000,
				31547,
				32768,
				0,
			},
		},
		{
			{
				23110,
				24597,
				27140,
				28894,
				30167,
				30927,
				31392,
				32094,
				32768,
				0,
			},
			{
				9998,
				17661,
				25178,
				28097,
				31308,
				32038,
				32403,
				32695,
				32768,
				0,
			},
		},
	},
}

var DEFAULT_EOB_PT_512_CDF = [][][]int{
	{
		{
			641,
			983,
			3707,
			5430,
			10234,
			14958,
			18788,
			23412,
			26061,
			32768,
			0,
		},
		{
			5095,
			6446,
			9996,
			13354,
			16017,
			17986,
			20919,
			26129,
			29140,
			32768,
			0,
		},
	},
	{
		{
			1230,
			2278,
			5035,
			7776,
			11871,
			15346,
			19590,
			24584,
			28749,
			32768,
			0,
		},
		{
			7265,
			9979,
			15819,
			19250,
			21780,
			23846,
			26478,
			28396,
			31811,
			32768,
			0,
		},
	},
	{
		{
			2624,
			3936,
			6480,
			9686,
			13979,
			17726,
			23267,
			28410,
			31078,
			32768,
			0,
		},
		{
			12015,
			14769,
			19588,
			22052,
			24222,
			25812,
			27300,
			29219,
			32114,
			32768,
			0,
		},
	},
	{
		{
			5927,
			7809,
			10923,
			14597,
			19439,
			24135,
			28456,
			31142,
			32060,
			32768,
			0,
		},
		{
			21093,
			23043,
			25742,
			27658,
			29097,
			29716,
			30073,
			30820,
			31956,
			32768,
			0,
		},
	},
}

var DEFAULT_EOB_PT_1024_CDF = [][][]int{
	{
		{
			393,
			421,
			751,
			1623,
			3160,
			6352,
			13345,
			18047,
			22571,
			25830,
			32768,
			0,
		},
		{
			1865,
			1988,
			2930,
			4242,
			10533,
			16538,
			21354,
			27255,
			28546,
			31784,
			32768,
			0,
		},
	},
	{
		{
			696,
			948,
			3145,
			5702,
			9706,
			13217,
			17851,
			21856,
			25692,
			28034,
			32768,
			0,
		},
		{
			2672,
			3591,
			9330,
			17084,
			22725,
			24284,
			26527,
			28027,
			28377,
			30876,
			32768,
			0,
		},
	},
	{
		{
			2784,
			3831,
			7041,
			10521,
			14847,
			18844,
			23155,
			26682,
			29229,
			31045,
			32768,
			0,
		},
		{
			9577,
			12466,
			17739,
			20750,
			22061,
			23215,
			24601,
			25483,
			25843,
			32056,
			32768,
			0,
		},
	},
	{
		{
			6698,
			8334,
			11961,
			15762,
			20186,
			23862,
			27434,
			29326,
			31082,
			32050,
			32768,
			0,
		},
		{
			20569,
			22426,
			25569,
			26859,
			28053,
			28913,
			29486,
			29724,
			29807,
			32570,
			32768,
			0,
		},
	},
}

var DEFAULT_EOB_EXTRA_CDF = [][][][][]int{
	{
		{
			{
				{
					16961,
					32768,
					0,
				},
				{
					17223,
					32768,
					0,
				},
				{
					7621,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					19069,
					32768,
					0,
				},
				{
					22525,
					32768,
					0,
				},
				{
					13377,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					20401,
					32768,
					0,
				},
				{
					17025,
					32768,
					0,
				},
				{
					12845,
					32768,
					0,
				},
				{
					12873,
					32768,
					0,
				},
				{
					14094,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					20681,
					32768,
					0,
				},
				{
					20701,
					32768,
					0,
				},
				{
					15250,
					32768,
					0,
				},
				{
					15017,
					32768,
					0,
				},
				{
					14928,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					23905,
					32768,
					0,
				},
				{
					17194,
					32768,
					0,
				},
				{
					16170,
					32768,
					0,
				},
				{
					17695,
					32768,
					0,
				},
				{
					13826,
					32768,
					0,
				},
				{
					15810,
					32768,
					0,
				},
				{
					12036,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					23959,
					32768,
					0,
				},
				{
					20799,
					32768,
					0,
				},
				{
					19021,
					32768,
					0,
				},
				{
					16203,
					32768,
					0,
				},
				{
					17886,
					32768,
					0,
				},
				{
					14144,
					32768,
					0,
				},
				{
					12010,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					27399,
					32768,
					0,
				},
				{
					16327,
					32768,
					0,
				},
				{
					18071,
					32768,
					0,
				},
				{
					19584,
					32768,
					0,
				},
				{
					20721,
					32768,
					0,
				},
				{
					18432,
					32768,
					0,
				},
				{
					19560,
					32768,
					0,
				},
				{
					10150,
					32768,
					0,
				},
				{
					8805,
					32768,
					0,
				},
			},
			{
				{
					24932,
					32768,
					0,
				},
				{
					20833,
					32768,
					0,
				},
				{
					12027,
					32768,
					0,
				},
				{
					16670,
					32768,
					0,
				},
				{
					19914,
					32768,
					0,
				},
				{
					15106,
					32768,
					0,
				},
				{
					17662,
					32768,
					0,
				},
				{
					13783,
					32768,
					0,
				},
				{
					28756,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					23406,
					32768,
					0,
				},
				{
					21845,
					32768,
					0,
				},
				{
					18432,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					17096,
					32768,
					0,
				},
				{
					12561,
					32768,
					0,
				},
				{
					17320,
					32768,
					0,
				},
				{
					22395,
					32768,
					0,
				},
				{
					21370,
					32768,
					0,
				},
			},
			{
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
	},
	{
		{
			{
				{
					17471,
					32768,
					0,
				},
				{
					20223,
					32768,
					0,
				},
				{
					11357,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					20335,
					32768,
					0,
				},
				{
					21667,
					32768,
					0,
				},
				{
					14818,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					20430,
					32768,
					0,
				},
				{
					20662,
					32768,
					0,
				},
				{
					15367,
					32768,
					0,
				},
				{
					16970,
					32768,
					0,
				},
				{
					14657,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					22117,
					32768,
					0,
				},
				{
					22028,
					32768,
					0,
				},
				{
					18650,
					32768,
					0,
				},
				{
					16042,
					32768,
					0,
				},
				{
					15885,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					22409,
					32768,
					0,
				},
				{
					21012,
					32768,
					0,
				},
				{
					15650,
					32768,
					0,
				},
				{
					17395,
					32768,
					0,
				},
				{
					15469,
					32768,
					0,
				},
				{
					20205,
					32768,
					0,
				},
				{
					19511,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					24220,
					32768,
					0,
				},
				{
					22480,
					32768,
					0,
				},
				{
					17737,
					32768,
					0,
				},
				{
					18916,
					32768,
					0,
				},
				{
					19268,
					32768,
					0,
				},
				{
					18412,
					32768,
					0,
				},
				{
					18844,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					25991,
					32768,
					0,
				},
				{
					20314,
					32768,
					0,
				},
				{
					17731,
					32768,
					0,
				},
				{
					19678,
					32768,
					0,
				},
				{
					18649,
					32768,
					0,
				},
				{
					17307,
					32768,
					0,
				},
				{
					21798,
					32768,
					0,
				},
				{
					17549,
					32768,
					0,
				},
				{
					15630,
					32768,
					0,
				},
			},
			{
				{
					26585,
					32768,
					0,
				},
				{
					21469,
					32768,
					0,
				},
				{
					20432,
					32768,
					0,
				},
				{
					17735,
					32768,
					0,
				},
				{
					19280,
					32768,
					0,
				},
				{
					15235,
					32768,
					0,
				},
				{
					20297,
					32768,
					0,
				},
				{
					22471,
					32768,
					0,
				},
				{
					28997,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					26605,
					32768,
					0,
				},
				{
					11304,
					32768,
					0,
				},
				{
					16726,
					32768,
					0,
				},
				{
					16560,
					32768,
					0,
				},
				{
					20866,
					32768,
					0,
				},
				{
					23524,
					32768,
					0,
				},
				{
					19878,
					32768,
					0,
				},
				{
					13469,
					32768,
					0,
				},
				{
					23084,
					32768,
					0,
				},
			},
			{
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
	},
	{
		{
			{
				{
					18983,
					32768,
					0,
				},
				{
					20512,
					32768,
					0,
				},
				{
					14885,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					20090,
					32768,
					0,
				},
				{
					19444,
					32768,
					0,
				},
				{
					17286,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					19139,
					32768,
					0,
				},
				{
					21487,
					32768,
					0,
				},
				{
					18959,
					32768,
					0,
				},
				{
					20910,
					32768,
					0,
				},
				{
					19089,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					20536,
					32768,
					0,
				},
				{
					20664,
					32768,
					0,
				},
				{
					20625,
					32768,
					0,
				},
				{
					19123,
					32768,
					0,
				},
				{
					14862,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					19833,
					32768,
					0,
				},
				{
					21502,
					32768,
					0,
				},
				{
					17485,
					32768,
					0,
				},
				{
					20267,
					32768,
					0,
				},
				{
					18353,
					32768,
					0,
				},
				{
					23329,
					32768,
					0,
				},
				{
					21478,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					22041,
					32768,
					0,
				},
				{
					23434,
					32768,
					0,
				},
				{
					20001,
					32768,
					0,
				},
				{
					20554,
					32768,
					0,
				},
				{
					20951,
					32768,
					0,
				},
				{
					20145,
					32768,
					0,
				},
				{
					15562,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					23312,
					32768,
					0,
				},
				{
					21607,
					32768,
					0,
				},
				{
					16526,
					32768,
					0,
				},
				{
					18957,
					32768,
					0,
				},
				{
					18034,
					32768,
					0,
				},
				{
					18934,
					32768,
					0,
				},
				{
					24247,
					32768,
					0,
				},
				{
					16921,
					32768,
					0,
				},
				{
					17080,
					32768,
					0,
				},
			},
			{
				{
					26579,
					32768,
					0,
				},
				{
					24910,
					32768,
					0,
				},
				{
					18637,
					32768,
					0,
				},
				{
					19800,
					32768,
					0,
				},
				{
					20388,
					32768,
					0,
				},
				{
					9887,
					32768,
					0,
				},
				{
					15642,
					32768,
					0,
				},
				{
					30198,
					32768,
					0,
				},
				{
					24721,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					26998,
					32768,
					0,
				},
				{
					16737,
					32768,
					0,
				},
				{
					17838,
					32768,
					0,
				},
				{
					18922,
					32768,
					0,
				},
				{
					19515,
					32768,
					0,
				},
				{
					18636,
					32768,
					0,
				},
				{
					17333,
					32768,
					0,
				},
				{
					15776,
					32768,
					0,
				},
				{
					22658,
					32768,
					0,
				},
			},
			{
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
	},
	{
		{
			{
				{
					20177,
					32768,
					0,
				},
				{
					20789,
					32768,
					0,
				},
				{
					20262,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					21416,
					32768,
					0,
				},
				{
					20855,
					32768,
					0,
				},
				{
					23410,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					20238,
					32768,
					0,
				},
				{
					21057,
					32768,
					0,
				},
				{
					19159,
					32768,
					0,
				},
				{
					22337,
					32768,
					0,
				},
				{
					20159,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					20125,
					32768,
					0,
				},
				{
					20559,
					32768,
					0,
				},
				{
					21707,
					32768,
					0,
				},
				{
					22296,
					32768,
					0,
				},
				{
					17333,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					19941,
					32768,
					0,
				},
				{
					20527,
					32768,
					0,
				},
				{
					21470,
					32768,
					0,
				},
				{
					22487,
					32768,
					0,
				},
				{
					19558,
					32768,
					0,
				},
				{
					22354,
					32768,
					0,
				},
				{
					20331,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
			{
				{
					22752,
					32768,
					0,
				},
				{
					25006,
					32768,
					0,
				},
				{
					22075,
					32768,
					0,
				},
				{
					21576,
					32768,
					0,
				},
				{
					17740,
					32768,
					0,
				},
				{
					21690,
					32768,
					0,
				},
				{
					19211,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					21442,
					32768,
					0,
				},
				{
					22358,
					32768,
					0,
				},
				{
					18503,
					32768,
					0,
				},
				{
					20291,
					32768,
					0,
				},
				{
					19945,
					32768,
					0,
				},
				{
					21294,
					32768,
					0,
				},
				{
					21178,
					32768,
					0,
				},
				{
					19400,
					32768,
					0,
				},
				{
					10556,
					32768,
					0,
				},
			},
			{
				{
					24648,
					32768,
					0,
				},
				{
					24949,
					32768,
					0,
				},
				{
					20708,
					32768,
					0,
				},
				{
					23905,
					32768,
					0,
				},
				{
					20501,
					32768,
					0,
				},
				{
					9558,
					32768,
					0,
				},
				{
					9423,
					32768,
					0,
				},
				{
					30365,
					32768,
					0,
				},
				{
					19253,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					26064,
					32768,
					0,
				},
				{
					22098,
					32768,
					0,
				},
				{
					19613,
					32768,
					0,
				},
				{
					20525,
					32768,
					0,
				},
				{
					17595,
					32768,
					0,
				},
				{
					16618,
					32768,
					0,
				},
				{
					20497,
					32768,
					0,
				},
				{
					18989,
					32768,
					0,
				},
				{
					15513,
					32768,
					0,
				},
			},
			{
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
				{
					16384,
					32768,
					0,
				},
			},
		},
	},
}

var DEFAULT_DC_SIGN_CDF = [][][][]int{
	{
		{
			{128 * 125, 32768, 0},
			{128 * 102, 32768, 0},
			{128 * 147, 32768, 0},
		},
		{
			{128 * 119, 32768, 0},
			{128 * 101, 32768, 0},
			{128 * 135, 32768, 0},
		},
	},
	{
		{
			{128 * 125, 32768, 0},
			{128 * 102, 32768, 0},
			{128 * 147, 32768, 0},
		},
		{
			{128 * 119, 32768, 0},
			{128 * 101, 32768, 0},
			{128 * 135, 32768, 0},
		},
	},
	{
		{
			{128 * 125, 32768, 0},
			{128 * 102, 32768, 0},
			{128 * 147, 32768, 0},
		},
		{
			{128 * 119, 32768, 0},
			{128 * 101, 32768, 0},
			{128 * 135, 32768, 0},
		},
	},
	{
		{
			{128 * 125, 32768, 0},
			{128 * 102, 32768, 0},
			{128 * 147, 32768, 0},
		},
		{
			{128 * 119, 32768, 0},
			{128 * 101, 32768, 0},
			{128 * 135, 32768, 0},
		},
	},
}

var DEFAULT_COEFF_BASE_EOB_CDF = [][][][][]int{
	{
		{
			{
				{
					17837,
					29055,
					32768,
					0,
				},
				{
					29600,
					31446,
					32768,
					0,
				},
				{
					30844,
					31878,
					32768,
					0,
				},
				{
					24926,
					28948,
					32768,
					0,
				},
			},
			{
				{
					21365,
					30026,
					32768,
					0,
				},
				{
					30512,
					32423,
					32768,
					0,
				},
				{
					31658,
					32621,
					32768,
					0,
				},
				{
					29630,
					31881,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					5717,
					26477,
					32768,
					0,
				},
				{
					30491,
					31703,
					32768,
					0,
				},
				{
					31550,
					32158,
					32768,
					0,
				},
				{
					29648,
					31491,
					32768,
					0,
				},
			},
			{
				{
					12608,
					27820,
					32768,
					0,
				},
				{
					30680,
					32225,
					32768,
					0,
				},
				{
					30809,
					32335,
					32768,
					0,
				},
				{
					31299,
					32423,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					1786,
					12612,
					32768,
					0,
				},
				{
					30663,
					31625,
					32768,
					0,
				},
				{
					32339,
					32468,
					32768,
					0,
				},
				{
					31148,
					31833,
					32768,
					0,
				},
			},
			{
				{
					18857,
					23865,
					32768,
					0,
				},
				{
					31428,
					32428,
					32768,
					0,
				},
				{
					31744,
					32373,
					32768,
					0,
				},
				{
					31775,
					32526,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					1787,
					2532,
					32768,
					0,
				},
				{
					30832,
					31662,
					32768,
					0,
				},
				{
					31824,
					32682,
					32768,
					0,
				},
				{
					32133,
					32569,
					32768,
					0,
				},
			},
			{
				{
					13751,
					22235,
					32768,
					0,
				},
				{
					32089,
					32409,
					32768,
					0,
				},
				{
					27084,
					27920,
					32768,
					0,
				},
				{
					29291,
					32594,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					1725,
					3449,
					32768,
					0,
				},
				{
					31102,
					31935,
					32768,
					0,
				},
				{
					32457,
					32613,
					32768,
					0,
				},
				{
					32412,
					32649,
					32768,
					0,
				},
			},
			{
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
			},
		},
	},
	{
		{
			{
				{
					17560,
					29888,
					32768,
					0,
				},
				{
					29671,
					31549,
					32768,
					0,
				},
				{
					31007,
					32056,
					32768,
					0,
				},
				{
					27286,
					30006,
					32768,
					0,
				},
			},
			{
				{
					26594,
					31212,
					32768,
					0,
				},
				{
					31208,
					32582,
					32768,
					0,
				},
				{
					31835,
					32637,
					32768,
					0,
				},
				{
					30595,
					32206,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					15239,
					29932,
					32768,
					0,
				},
				{
					31315,
					32095,
					32768,
					0,
				},
				{
					32130,
					32434,
					32768,
					0,
				},
				{
					30864,
					31996,
					32768,
					0,
				},
			},
			{
				{
					26279,
					30968,
					32768,
					0,
				},
				{
					31142,
					32495,
					32768,
					0,
				},
				{
					31713,
					32540,
					32768,
					0,
				},
				{
					31929,
					32594,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					2644,
					25198,
					32768,
					0,
				},
				{
					32038,
					32451,
					32768,
					0,
				},
				{
					32639,
					32695,
					32768,
					0,
				},
				{
					32166,
					32518,
					32768,
					0,
				},
			},
			{
				{
					17187,
					27668,
					32768,
					0,
				},
				{
					31714,
					32550,
					32768,
					0,
				},
				{
					32283,
					32678,
					32768,
					0,
				},
				{
					31930,
					32563,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					1044,
					2257,
					32768,
					0,
				},
				{
					30755,
					31923,
					32768,
					0,
				},
				{
					32208,
					32693,
					32768,
					0,
				},
				{
					32244,
					32615,
					32768,
					0,
				},
			},
			{
				{
					21317,
					26207,
					32768,
					0,
				},
				{
					29133,
					30868,
					32768,
					0,
				},
				{
					29311,
					31231,
					32768,
					0,
				},
				{
					29657,
					31087,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					478,
					1834,
					32768,
					0,
				},
				{
					31005,
					31987,
					32768,
					0,
				},
				{
					32317,
					32724,
					32768,
					0,
				},
				{
					30865,
					32648,
					32768,
					0,
				},
			},
			{
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
			},
		},
	},
	{
		{
			{
				{
					20092,
					30774,
					32768,
					0,
				},
				{
					30695,
					32020,
					32768,
					0,
				},
				{
					31131,
					32103,
					32768,
					0,
				},
				{
					28666,
					30870,
					32768,
					0,
				},
			},
			{
				{
					27258,
					31095,
					32768,
					0,
				},
				{
					31804,
					32623,
					32768,
					0,
				},
				{
					31763,
					32528,
					32768,
					0,
				},
				{
					31438,
					32506,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					18049,
					30489,
					32768,
					0,
				},
				{
					31706,
					32286,
					32768,
					0,
				},
				{
					32163,
					32473,
					32768,
					0,
				},
				{
					31550,
					32184,
					32768,
					0,
				},
			},
			{
				{
					27116,
					30842,
					32768,
					0,
				},
				{
					31971,
					32598,
					32768,
					0,
				},
				{
					32088,
					32576,
					32768,
					0,
				},
				{
					32067,
					32664,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					12854,
					29093,
					32768,
					0,
				},
				{
					32272,
					32558,
					32768,
					0,
				},
				{
					32667,
					32729,
					32768,
					0,
				},
				{
					32306,
					32585,
					32768,
					0,
				},
			},
			{
				{
					25476,
					30366,
					32768,
					0,
				},
				{
					32169,
					32687,
					32768,
					0,
				},
				{
					32479,
					32689,
					32768,
					0,
				},
				{
					31673,
					32634,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					2809,
					19301,
					32768,
					0,
				},
				{
					32205,
					32622,
					32768,
					0,
				},
				{
					32338,
					32730,
					32768,
					0,
				},
				{
					31786,
					32616,
					32768,
					0,
				},
			},
			{
				{
					22737,
					29105,
					32768,
					0,
				},
				{
					30810,
					32362,
					32768,
					0,
				},
				{
					30014,
					32627,
					32768,
					0,
				},
				{
					30528,
					32574,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					935,
					3382,
					32768,
					0,
				},
				{
					30789,
					31909,
					32768,
					0,
				},
				{
					32466,
					32756,
					32768,
					0,
				},
				{
					30860,
					32513,
					32768,
					0,
				},
			},
			{
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
			},
		},
	},
	{
		{
			{
				{
					22497,
					31198,
					32768,
					0,
				},
				{
					31715,
					32495,
					32768,
					0,
				},
				{
					31606,
					32337,
					32768,
					0,
				},
				{
					30388,
					31990,
					32768,
					0,
				},
			},
			{
				{
					27877,
					31584,
					32768,
					0,
				},
				{
					32170,
					32728,
					32768,
					0,
				},
				{
					32155,
					32688,
					32768,
					0,
				},
				{
					32219,
					32702,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					21457,
					31043,
					32768,
					0,
				},
				{
					31951,
					32483,
					32768,
					0,
				},
				{
					32153,
					32562,
					32768,
					0,
				},
				{
					31473,
					32215,
					32768,
					0,
				},
			},
			{
				{
					27558,
					31151,
					32768,
					0,
				},
				{
					32020,
					32640,
					32768,
					0,
				},
				{
					32097,
					32575,
					32768,
					0,
				},
				{
					32242,
					32719,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					19980,
					30591,
					32768,
					0,
				},
				{
					32219,
					32597,
					32768,
					0,
				},
				{
					32581,
					32706,
					32768,
					0,
				},
				{
					31803,
					32287,
					32768,
					0,
				},
			},
			{
				{
					26473,
					30507,
					32768,
					0,
				},
				{
					32431,
					32723,
					32768,
					0,
				},
				{
					32196,
					32611,
					32768,
					0,
				},
				{
					31588,
					32528,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					24647,
					30463,
					32768,
					0,
				},
				{
					32412,
					32695,
					32768,
					0,
				},
				{
					32468,
					32720,
					32768,
					0,
				},
				{
					31269,
					32523,
					32768,
					0,
				},
			},
			{
				{
					28482,
					31505,
					32768,
					0,
				},
				{
					32152,
					32701,
					32768,
					0,
				},
				{
					31732,
					32598,
					32768,
					0,
				},
				{
					31767,
					32712,
					32768,
					0,
				},
			},
		},
		{
			{
				{
					12358,
					24977,
					32768,
					0,
				},
				{
					31331,
					32385,
					32768,
					0,
				},
				{
					32634,
					32756,
					32768,
					0,
				},
				{
					30411,
					32548,
					32768,
					0,
				},
			},
			{
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
				{
					10923,
					21845,
					32768,
					0,
				},
			},
		},
	},
}

var DEFAULT_COEFF_BASE_CDF = [][][][][]int{
	{
		{
			{
				{
					4034, 8930, 12727, 32768, 0},
				{
					18082, 29741, 31877, 32768, 0},
				{
					12596, 26124, 30493, 32768, 0},
				{
					9446, 21118, 27005, 32768, 0},
				{
					6308, 15141, 21279, 32768, 0},
				{
					2463, 6357, 9783, 32768, 0},
				{
					20667, 30546, 31929, 32768, 0},
				{
					13043, 26123, 30134, 32768, 0},
				{
					8151, 18757, 24778, 32768, 0},
				{
					5255, 12839, 18632, 32768, 0},
				{
					2820, 7206, 11161, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					15736, 27553, 30604, 32768, 0},
				{
					11210, 23794, 28787, 32768, 0},
				{
					5947, 13874, 19701, 32768, 0},
				{
					4215, 9323, 13891, 32768, 0},
				{
					2833, 6462, 10059, 32768, 0},
				{
					19605, 30393, 31582, 32768, 0},
				{
					13523, 26252, 30248, 32768, 0},
				{
					8446, 18622, 24512, 32768, 0},
				{
					3818, 10343, 15974, 32768, 0},
				{
					1481, 4117, 6796, 32768, 0},
				{
					22649, 31302, 32190, 32768, 0},
				{
					14829, 27127, 30449, 32768, 0},
				{
					8313, 17702, 23304, 32768, 0},
				{
					3022, 8301, 12786, 32768, 0},
				{
					1536, 4412, 7184, 32768, 0},
				{
					22354, 29774, 31372, 32768, 0},
				{
					14723, 25472, 29214, 32768, 0},
				{
					6673, 13745, 18662, 32768, 0},
				{
					2068, 5766, 9322, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					6302, 16444, 21761, 32768, 0},
				{
					23040, 31538, 32475, 32768, 0},
				{
					15196, 28452, 31496, 32768, 0},
				{
					10020, 22946, 28514, 32768, 0},
				{
					6533, 16862, 23501, 32768, 0},
				{
					3538, 9816, 15076, 32768, 0},
				{
					24444, 31875, 32525, 32768, 0},
				{
					15881, 28924, 31635, 32768, 0},
				{
					9922, 22873, 28466, 32768, 0},
				{
					6527, 16966, 23691, 32768, 0},
				{
					4114, 11303, 17220, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					20201, 30770, 32209, 32768, 0},
				{
					14754, 28071, 31258, 32768, 0},
				{
					8378, 20186, 26517, 32768, 0},
				{
					5916, 15299, 21978, 32768, 0},
				{
					4268, 11583, 17901, 32768, 0},
				{
					24361, 32025, 32581, 32768, 0},
				{
					18673, 30105, 31943, 32768, 0},
				{
					10196, 22244, 27576, 32768, 0},
				{
					5495, 14349, 20417, 32768, 0},
				{
					2676, 7415, 11498, 32768, 0},
				{
					24678, 31958, 32585, 32768, 0},
				{
					18629, 29906, 31831, 32768, 0},
				{
					9364, 20724, 26315, 32768, 0},
				{
					4641, 12318, 18094, 32768, 0},
				{
					2758, 7387, 11579, 32768, 0},
				{
					25433, 31842, 32469, 32768, 0},
				{
					18795, 29289, 31411, 32768, 0},
				{
					7644, 17584, 23592, 32768, 0},
				{
					3408, 9014, 15047, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					4536, 10072, 14001, 32768, 0},
				{
					25459, 31416, 32206, 32768, 0},
				{
					16605, 28048, 30818, 32768, 0},
				{
					11008, 22857, 27719, 32768, 0},
				{
					6915, 16268, 22315, 32768, 0},
				{
					2625, 6812, 10537, 32768, 0},
				{
					24257, 31788, 32499, 32768, 0},
				{
					16880, 29454, 31879, 32768, 0},
				{
					11958, 25054, 29778, 32768, 0},
				{
					7916, 18718, 25084, 32768, 0},
				{
					3383, 8777, 13446, 32768, 0},
				{
					22720, 31603, 32393, 32768, 0},
				{
					14960, 28125, 31335, 32768, 0},
				{
					9731, 22210, 27928, 32768, 0},
				{
					6304, 15832, 22277, 32768, 0},
				{
					2910, 7818, 12166, 32768, 0},
				{
					20375, 30627, 32131, 32768, 0},
				{
					13904, 27284, 30887, 32768, 0},
				{
					9368, 21558, 27144, 32768, 0},
				{
					5937, 14966, 21119, 32768, 0},
				{
					2667, 7225, 11319, 32768, 0},
				{
					23970, 31470, 32378, 32768, 0},
				{
					17173, 29734, 32018, 32768, 0},
				{
					12795, 25441, 29965, 32768, 0},
				{
					8981, 19680, 25893, 32768, 0},
				{
					4728, 11372, 16902, 32768, 0},
				{
					24287, 31797, 32439, 32768, 0},
				{
					16703, 29145, 31696, 32768, 0},
				{
					10833, 23554, 28725, 32768, 0},
				{
					6468, 16566, 23057, 32768, 0},
				{
					2415, 6562, 10278, 32768, 0},
				{
					26610, 32395, 32659, 32768, 0},
				{
					18590, 30498, 32117, 32768, 0},
				{
					12420, 25756, 29950, 32768, 0},
				{
					7639, 18746, 24710, 32768, 0},
				{
					3001, 8086, 12347, 32768, 0},
				{
					25076, 32064, 32580, 32768, 0},
				{
					17946, 30128, 32028, 32768, 0},
				{
					12024, 24985, 29378, 32768, 0},
				{
					7517, 18390, 24304, 32768, 0},
				{
					3243, 8781, 13331, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					6037, 16771, 21957, 32768, 0},
				{
					24774, 31704, 32426, 32768, 0},
				{
					16830, 28589, 31056, 32768, 0},
				{
					10602, 22828, 27760, 32768, 0},
				{
					6733, 16829, 23071, 32768, 0},
				{
					3250, 8914, 13556, 32768, 0},
				{
					25582, 32220, 32668, 32768, 0},
				{
					18659, 30342, 32223, 32768, 0},
				{
					12546, 26149, 30515, 32768, 0},
				{
					8420, 20451, 26801, 32768, 0},
				{
					4636, 12420, 18344, 32768, 0},
				{
					27581, 32362, 32639, 32768, 0},
				{
					18987, 30083, 31978, 32768, 0},
				{
					11327, 24248, 29084, 32768, 0},
				{
					7264, 17719, 24120, 32768, 0},
				{
					3995, 10768, 16169, 32768, 0},
				{
					25893, 31831, 32487, 32768, 0},
				{
					16577, 28587, 31379, 32768, 0},
				{
					10189, 22748, 28182, 32768, 0},
				{
					6832, 17094, 23556, 32768, 0},
				{
					3708, 10110, 15334, 32768, 0},
				{
					25904, 32282, 32656, 32768, 0},
				{
					19721, 30792, 32276, 32768, 0},
				{
					12819, 26243, 30411, 32768, 0},
				{
					8572, 20614, 26891, 32768, 0},
				{
					5364, 14059, 20467, 32768, 0},
				{
					26580, 32438, 32677, 32768, 0},
				{
					20852, 31225, 32340, 32768, 0},
				{
					12435, 25700, 29967, 32768, 0},
				{
					8691, 20825, 26976, 32768, 0},
				{
					4446, 12209, 17269, 32768, 0},
				{
					27350, 32429, 32696, 32768, 0},
				{
					21372, 30977, 32272, 32768, 0},
				{
					12673, 25270, 29853, 32768, 0},
				{
					9208, 20925, 26640, 32768, 0},
				{
					5018, 13351, 18732, 32768, 0},
				{
					27351, 32479, 32713, 32768, 0},
				{
					21398, 31209, 32387, 32768, 0},
				{
					12162, 25047, 29842, 32768, 0},
				{
					7896, 18691, 25319, 32768, 0},
				{
					4670, 12882, 18881, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					5487, 10460, 13708, 32768, 0},
				{
					21597, 28303, 30674, 32768, 0},
				{
					11037, 21953, 26476, 32768, 0},
				{
					8147, 17962, 22952, 32768, 0},
				{
					5242, 13061, 18532, 32768, 0},
				{
					1889, 5208, 8182, 32768, 0},
				{
					26774, 32133, 32590, 32768, 0},
				{
					17844, 29564, 31767, 32768, 0},
				{
					11690, 24438, 29171, 32768, 0},
				{
					7542, 18215, 24459, 32768, 0},
				{
					2993, 8050, 12319, 32768, 0},
				{
					28023, 32328, 32591, 32768, 0},
				{
					18651, 30126, 31954, 32768, 0},
				{
					12164, 25146, 29589, 32768, 0},
				{
					7762, 18530, 24771, 32768, 0},
				{
					3492, 9183, 13920, 32768, 0},
				{
					27591, 32008, 32491, 32768, 0},
				{
					17149, 28853, 31510, 32768, 0},
				{
					11485, 24003, 28860, 32768, 0},
				{
					7697, 18086, 24210, 32768, 0},
				{
					3075, 7999, 12218, 32768, 0},
				{
					28268, 32482, 32654, 32768, 0},
				{
					19631, 31051, 32404, 32768, 0},
				{
					13860, 27260, 31020, 32768, 0},
				{
					9605, 21613, 27594, 32768, 0},
				{
					4876, 12162, 17908, 32768, 0},
				{
					27248, 32316, 32576, 32768, 0},
				{
					18955, 30457, 32075, 32768, 0},
				{
					11824, 23997, 28795, 32768, 0},
				{
					7346, 18196, 24647, 32768, 0},
				{
					3403, 9247, 14111, 32768, 0},
				{
					29711, 32655, 32735, 32768, 0},
				{
					21169, 31394, 32417, 32768, 0},
				{
					13487, 27198, 30957, 32768, 0},
				{
					8828, 21683, 27614, 32768, 0},
				{
					4270, 11451, 17038, 32768, 0},
				{
					28708, 32578, 32731, 32768, 0},
				{
					20120, 31241, 32482, 32768, 0},
				{
					13692, 27550, 31321, 32768, 0},
				{
					9418, 22514, 28439, 32768, 0},
				{
					4999, 13283, 19462, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					5673, 14302, 19711, 32768, 0},
				{
					26251, 30701, 31834, 32768, 0},
				{
					12782, 23783, 27803, 32768, 0},
				{
					9127, 20657, 25808, 32768, 0},
				{
					6368, 16208, 21462, 32768, 0},
				{
					2465, 7177, 10822, 32768, 0},
				{
					29961, 32563, 32719, 32768, 0},
				{
					18318, 29891, 31949, 32768, 0},
				{
					11361, 24514, 29357, 32768, 0},
				{
					7900, 19603, 25607, 32768, 0},
				{
					4002, 10590, 15546, 32768, 0},
				{
					29637, 32310, 32595, 32768, 0},
				{
					18296, 29913, 31809, 32768, 0},
				{
					10144, 21515, 26871, 32768, 0},
				{
					5358, 14322, 20394, 32768, 0},
				{
					3067, 8362, 13346, 32768, 0},
				{
					28652, 32470, 32676, 32768, 0},
				{
					17538, 30771, 32209, 32768, 0},
				{
					13924, 26882, 30494, 32768, 0},
				{
					10496, 22837, 27869, 32768, 0},
				{
					7236, 16396, 21621, 32768, 0},
				{
					30743, 32687, 32746, 32768, 0},
				{
					23006, 31676, 32489, 32768, 0},
				{
					14494, 27828, 31120, 32768, 0},
				{
					10174, 22801, 28352, 32768, 0},
				{
					6242, 15281, 21043, 32768, 0},
				{
					25817, 32243, 32720, 32768, 0},
				{
					18618, 31367, 32325, 32768, 0},
				{
					13997, 28318, 31878, 32768, 0},
				{
					12255, 26534, 31383, 32768, 0},
				{
					9561, 21588, 28450, 32768, 0},
				{
					28188, 32635, 32724, 32768, 0},
				{
					22060, 32365, 32728, 32768, 0},
				{
					18102, 30690, 32528, 32768, 0},
				{
					14196, 28864, 31999, 32768, 0},
				{
					12262, 25792, 30865, 32768, 0},
				{
					24176, 32109, 32628, 32768, 0},
				{
					18280, 29681, 31963, 32768, 0},
				{
					10205, 23703, 29664, 32768, 0},
				{
					7889, 20025, 27676, 32768, 0},
				{
					6060, 16743, 23970, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					5141, 7096, 8260, 32768, 0},
				{
					27186, 29022, 29789, 32768, 0},
				{
					6668, 12568, 15682, 32768, 0},
				{
					2172, 6181, 8638, 32768, 0},
				{
					1126, 3379, 4531, 32768, 0},
				{
					443, 1361, 2254, 32768, 0},
				{
					26083, 31153, 32436, 32768, 0},
				{
					13486, 24603, 28483, 32768, 0},
				{
					6508, 14840, 19910, 32768, 0},
				{
					3386, 8800, 13286, 32768, 0},
				{
					1530, 4322, 7054, 32768, 0},
				{
					29639, 32080, 32548, 32768, 0},
				{
					15897, 27552, 30290, 32768, 0},
				{
					8588, 20047, 25383, 32768, 0},
				{
					4889, 13339, 19269, 32768, 0},
				{
					2240, 6871, 10498, 32768, 0},
				{
					28165, 32197, 32517, 32768, 0},
				{
					20735, 30427, 31568, 32768, 0},
				{
					14325, 24671, 27692, 32768, 0},
				{
					5119, 12554, 17805, 32768, 0},
				{
					1810, 5441, 8261, 32768, 0},
				{
					31212, 32724, 32748, 32768, 0},
				{
					23352, 31766, 32545, 32768, 0},
				{
					14669, 27570, 31059, 32768, 0},
				{
					8492, 20894, 27272, 32768, 0},
				{
					3644, 10194, 15204, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					2461, 7013, 9371, 32768, 0},
				{
					24749, 29600, 30986, 32768, 0},
				{
					9466, 19037, 22417, 32768, 0},
				{
					3584, 9280, 14400, 32768, 0},
				{
					1505, 3929, 5433, 32768, 0},
				{
					677, 1500, 2736, 32768, 0},
				{
					23987, 30702, 32117, 32768, 0},
				{
					13554, 24571, 29263, 32768, 0},
				{
					6211, 14556, 21155, 32768, 0},
				{
					3135, 10972, 15625, 32768, 0},
				{
					2435, 7127, 11427, 32768, 0},
				{
					31300, 32532, 32550, 32768, 0},
				{
					14757, 30365, 31954, 32768, 0},
				{
					4405, 11612, 18553, 32768, 0},
				{
					580, 4132, 7322, 32768, 0},
				{
					1695, 10169, 14124, 32768, 0},
				{
					30008, 32282, 32591, 32768, 0},
				{
					19244, 30108, 31748, 32768, 0},
				{
					11180, 24158, 29555, 32768, 0},
				{
					5650, 14972, 19209, 32768, 0},
				{
					2114, 5109, 8456, 32768, 0},
				{
					31856, 32716, 32748, 32768, 0},
				{
					23012, 31664, 32572, 32768, 0},
				{
					13694, 26656, 30636, 32768, 0},
				{
					8142, 19508, 26093, 32768, 0},
				{
					4253, 10955, 16724, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					601, 983, 1311, 32768, 0},
				{
					18725, 23406, 28087, 32768, 0},
				{
					5461, 8192, 10923, 32768, 0},
				{
					3781, 15124, 21425, 32768, 0},
				{
					2587, 7761, 12072, 32768, 0},
				{
					106, 458, 810, 32768, 0},
				{
					22282, 29710, 31894, 32768, 0},
				{
					8508, 20926, 25984, 32768, 0},
				{
					3726, 12713, 18083, 32768, 0},
				{
					1620, 7112, 10893, 32768, 0},
				{
					729, 2236, 3495, 32768, 0},
				{
					30163, 32474, 32684, 32768, 0},
				{
					18304, 30464, 32000, 32768, 0},
				{
					11443, 26526, 29647, 32768, 0},
				{
					6007, 15292, 21299, 32768, 0},
				{
					2234, 6703, 8937, 32768, 0},
				{
					30954, 32177, 32571, 32768, 0},
				{
					17363, 29562, 31076, 32768, 0},
				{
					9686, 22464, 27410, 32768, 0},
				{
					8192, 16384, 21390, 32768, 0},
				{
					1755, 8046, 11264, 32768, 0},
				{
					31168, 32734, 32748, 32768, 0},
				{
					22486, 31441, 32471, 32768, 0},
				{
					12833, 25627, 29738, 32768, 0},
				{
					6980, 17379, 23122, 32768, 0},
				{
					3111, 8887, 13479, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
	{
		{
			{
				{
					6041, 11854, 15927, 32768, 0},
				{
					20326, 30905, 32251, 32768, 0},
				{
					14164, 26831, 30725, 32768, 0},
				{
					9760, 20647, 26585, 32768, 0},
				{
					6416, 14953, 21219, 32768, 0},
				{
					2966, 7151, 10891, 32768, 0},
				{
					23567, 31374, 32254, 32768, 0},
				{
					14978, 27416, 30946, 32768, 0},
				{
					9434, 20225, 26254, 32768, 0},
				{
					6658, 14558, 20535, 32768, 0},
				{
					3916, 8677, 12989, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					18088, 29545, 31587, 32768, 0},
				{
					13062, 25843, 30073, 32768, 0},
				{
					8940, 16827, 22251, 32768, 0},
				{
					7654, 13220, 17973, 32768, 0},
				{
					5733, 10316, 14456, 32768, 0},
				{
					22879, 31388, 32114, 32768, 0},
				{
					15215, 27993, 30955, 32768, 0},
				{
					9397, 19445, 24978, 32768, 0},
				{
					3442, 9813, 15344, 32768, 0},
				{
					1368, 3936, 6532, 32768, 0},
				{
					25494, 32033, 32406, 32768, 0},
				{
					16772, 27963, 30718, 32768, 0},
				{
					9419, 18165, 23260, 32768, 0},
				{
					2677, 7501, 11797, 32768, 0},
				{
					1516, 4344, 7170, 32768, 0},
				{
					26556, 31454, 32101, 32768, 0},
				{
					17128, 27035, 30108, 32768, 0},
				{
					8324, 15344, 20249, 32768, 0},
				{
					1903, 5696, 9469, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8455, 19003, 24368, 32768, 0},
				{
					23563, 32021, 32604, 32768, 0},
				{
					16237, 29446, 31935, 32768, 0},
				{
					10724, 23999, 29358, 32768, 0},
				{
					6725, 17528, 24416, 32768, 0},
				{
					3927, 10927, 16825, 32768, 0},
				{
					26313, 32288, 32634, 32768, 0},
				{
					17430, 30095, 32095, 32768, 0},
				{
					11116, 24606, 29679, 32768, 0},
				{
					7195, 18384, 25269, 32768, 0},
				{
					4726, 12852, 19315, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					22822, 31648, 32483, 32768, 0},
				{
					16724, 29633, 31929, 32768, 0},
				{
					10261, 23033, 28725, 32768, 0},
				{
					7029, 17840, 24528, 32768, 0},
				{
					4867, 13886, 21502, 32768, 0},
				{
					25298, 31892, 32491, 32768, 0},
				{
					17809, 29330, 31512, 32768, 0},
				{
					9668, 21329, 26579, 32768, 0},
				{
					4774, 12956, 18976, 32768, 0},
				{
					2322, 7030, 11540, 32768, 0},
				{
					25472, 31920, 32543, 32768, 0},
				{
					17957, 29387, 31632, 32768, 0},
				{
					9196, 20593, 26400, 32768, 0},
				{
					4680, 12705, 19202, 32768, 0},
				{
					2917, 8456, 13436, 32768, 0},
				{
					26471, 32059, 32574, 32768, 0},
				{
					18458, 29783, 31909, 32768, 0},
				{
					8400, 19464, 25956, 32768, 0},
				{
					3812, 10973, 17206, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					6779, 13743, 17678, 32768, 0},
				{
					24806, 31797, 32457, 32768, 0},
				{
					17616, 29047, 31372, 32768, 0},
				{
					11063, 23175, 28003, 32768, 0},
				{
					6521, 16110, 22324, 32768, 0},
				{
					2764, 7504, 11654, 32768, 0},
				{
					25266, 32367, 32637, 32768, 0},
				{
					19054, 30553, 32175, 32768, 0},
				{
					12139, 25212, 29807, 32768, 0},
				{
					7311, 18162, 24704, 32768, 0},
				{
					3397, 9164, 14074, 32768, 0},
				{
					25988, 32208, 32522, 32768, 0},
				{
					16253, 28912, 31526, 32768, 0},
				{
					9151, 21387, 27372, 32768, 0},
				{
					5688, 14915, 21496, 32768, 0},
				{
					2717, 7627, 12004, 32768, 0},
				{
					23144, 31855, 32443, 32768, 0},
				{
					16070, 28491, 31325, 32768, 0},
				{
					8702, 20467, 26517, 32768, 0},
				{
					5243, 13956, 20367, 32768, 0},
				{
					2621, 7335, 11567, 32768, 0},
				{
					26636, 32340, 32630, 32768, 0},
				{
					19990, 31050, 32341, 32768, 0},
				{
					13243, 26105, 30315, 32768, 0},
				{
					8588, 19521, 25918, 32768, 0},
				{
					4717, 11585, 17304, 32768, 0},
				{
					25844, 32292, 32582, 32768, 0},
				{
					19090, 30635, 32097, 32768, 0},
				{
					11963, 24546, 28939, 32768, 0},
				{
					6218, 16087, 22354, 32768, 0},
				{
					2340, 6608, 10426, 32768, 0},
				{
					28046, 32576, 32694, 32768, 0},
				{
					21178, 31313, 32296, 32768, 0},
				{
					13486, 26184, 29870, 32768, 0},
				{
					7149, 17871, 23723, 32768, 0},
				{
					2833, 7958, 12259, 32768, 0},
				{
					27710, 32528, 32686, 32768, 0},
				{
					20674, 31076, 32268, 32768, 0},
				{
					12413, 24955, 29243, 32768, 0},
				{
					6676, 16927, 23097, 32768, 0},
				{
					2966, 8333, 12919, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8639, 19339, 24429, 32768, 0},
				{
					24404, 31837, 32525, 32768, 0},
				{
					16997, 29425, 31784, 32768, 0},
				{
					11253, 24234, 29149, 32768, 0},
				{
					6751, 17394, 24028, 32768, 0},
				{
					3490, 9830, 15191, 32768, 0},
				{
					26283, 32471, 32714, 32768, 0},
				{
					19599, 31168, 32442, 32768, 0},
				{
					13146, 26954, 30893, 32768, 0},
				{
					8214, 20588, 26890, 32768, 0},
				{
					4699, 13081, 19300, 32768, 0},
				{
					28212, 32458, 32669, 32768, 0},
				{
					18594, 30316, 32100, 32768, 0},
				{
					11219, 24408, 29234, 32768, 0},
				{
					6865, 17656, 24149, 32768, 0},
				{
					3678, 10362, 16006, 32768, 0},
				{
					25825, 32136, 32616, 32768, 0},
				{
					17313, 29853, 32021, 32768, 0},
				{
					11197, 24471, 29472, 32768, 0},
				{
					6947, 17781, 24405, 32768, 0},
				{
					3768, 10660, 16261, 32768, 0},
				{
					27352, 32500, 32706, 32768, 0},
				{
					20850, 31468, 32469, 32768, 0},
				{
					14021, 27707, 31133, 32768, 0},
				{
					8964, 21748, 27838, 32768, 0},
				{
					5437, 14665, 21187, 32768, 0},
				{
					26304, 32492, 32698, 32768, 0},
				{
					20409, 31380, 32385, 32768, 0},
				{
					13682, 27222, 30632, 32768, 0},
				{
					8974, 21236, 26685, 32768, 0},
				{
					4234, 11665, 16934, 32768, 0},
				{
					26273, 32357, 32711, 32768, 0},
				{
					20672, 31242, 32441, 32768, 0},
				{
					14172, 27254, 30902, 32768, 0},
				{
					9870, 21898, 27275, 32768, 0},
				{
					5164, 13506, 19270, 32768, 0},
				{
					26725, 32459, 32728, 32768, 0},
				{
					20991, 31442, 32527, 32768, 0},
				{
					13071, 26434, 30811, 32768, 0},
				{
					8184, 20090, 26742, 32768, 0},
				{
					4803, 13255, 19895, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					7555, 14942, 18501, 32768, 0},
				{
					24410, 31178, 32287, 32768, 0},
				{
					14394, 26738, 30253, 32768, 0},
				{
					8413, 19554, 25195, 32768, 0},
				{
					4766, 12924, 18785, 32768, 0},
				{
					2029, 5806, 9207, 32768, 0},
				{
					26776, 32364, 32663, 32768, 0},
				{
					18732, 29967, 31931, 32768, 0},
				{
					11005, 23786, 28852, 32768, 0},
				{
					6466, 16909, 23510, 32768, 0},
				{
					3044, 8638, 13419, 32768, 0},
				{
					29208, 32582, 32704, 32768, 0},
				{
					20068, 30857, 32208, 32768, 0},
				{
					12003, 25085, 29595, 32768, 0},
				{
					6947, 17750, 24189, 32768, 0},
				{
					3245, 9103, 14007, 32768, 0},
				{
					27359, 32465, 32669, 32768, 0},
				{
					19421, 30614, 32174, 32768, 0},
				{
					11915, 25010, 29579, 32768, 0},
				{
					6950, 17676, 24074, 32768, 0},
				{
					3007, 8473, 13096, 32768, 0},
				{
					29002, 32676, 32735, 32768, 0},
				{
					22102, 31849, 32576, 32768, 0},
				{
					14408, 28009, 31405, 32768, 0},
				{
					9027, 21679, 27931, 32768, 0},
				{
					4694, 12678, 18748, 32768, 0},
				{
					28216, 32528, 32682, 32768, 0},
				{
					20849, 31264, 32318, 32768, 0},
				{
					12756, 25815, 29751, 32768, 0},
				{
					7565, 18801, 24923, 32768, 0},
				{
					3509, 9533, 14477, 32768, 0},
				{
					30133, 32687, 32739, 32768, 0},
				{
					23063, 31910, 32515, 32768, 0},
				{
					14588, 28051, 31132, 32768, 0},
				{
					9085, 21649, 27457, 32768, 0},
				{
					4261, 11654, 17264, 32768, 0},
				{
					29518, 32691, 32748, 32768, 0},
				{
					22451, 31959, 32613, 32768, 0},
				{
					14864, 28722, 31700, 32768, 0},
				{
					9695, 22964, 28716, 32768, 0},
				{
					4932, 13358, 19502, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					6465, 16958, 21688, 32768, 0},
				{
					25199, 31514, 32360, 32768, 0},
				{
					14774, 27149, 30607, 32768, 0},
				{
					9257, 21438, 26972, 32768, 0},
				{
					5723, 15183, 21882, 32768, 0},
				{
					3150, 8879, 13731, 32768, 0},
				{
					26989, 32262, 32682, 32768, 0},
				{
					17396, 29937, 32085, 32768, 0},
				{
					11387, 24901, 29784, 32768, 0},
				{
					7289, 18821, 25548, 32768, 0},
				{
					3734, 10577, 16086, 32768, 0},
				{
					29728, 32501, 32695, 32768, 0},
				{
					17431, 29701, 31903, 32768, 0},
				{
					9921, 22826, 28300, 32768, 0},
				{
					5896, 15434, 22068, 32768, 0},
				{
					3430, 9646, 14757, 32768, 0},
				{
					28614, 32511, 32705, 32768, 0},
				{
					19364, 30638, 32263, 32768, 0},
				{
					13129, 26254, 30402, 32768, 0},
				{
					8754, 20484, 26440, 32768, 0},
				{
					4378, 11607, 17110, 32768, 0},
				{
					30292, 32671, 32744, 32768, 0},
				{
					21780, 31603, 32501, 32768, 0},
				{
					14314, 27829, 31291, 32768, 0},
				{
					9611, 22327, 28263, 32768, 0},
				{
					4890, 13087, 19065, 32768, 0},
				{
					25862, 32567, 32733, 32768, 0},
				{
					20794, 32050, 32567, 32768, 0},
				{
					17243, 30625, 32254, 32768, 0},
				{
					13283, 27628, 31474, 32768, 0},
				{
					9669, 22532, 28918, 32768, 0},
				{
					27435, 32697, 32748, 32768, 0},
				{
					24922, 32390, 32714, 32768, 0},
				{
					21449, 31504, 32536, 32768, 0},
				{
					16392, 29729, 31832, 32768, 0},
				{
					11692, 24884, 29076, 32768, 0},
				{
					24193, 32290, 32735, 32768, 0},
				{
					18909, 31104, 32563, 32768, 0},
				{
					12236, 26841, 31403, 32768, 0},
				{
					8171, 21840, 29082, 32768, 0},
				{
					7224, 17280, 25275, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					3078, 6839, 9890, 32768, 0},
				{
					13837, 20450, 24479, 32768, 0},
				{
					5914, 14222, 19328, 32768, 0},
				{
					3866, 10267, 14762, 32768, 0},
				{
					2612, 7208, 11042, 32768, 0},
				{
					1067, 2991, 4776, 32768, 0},
				{
					25817, 31646, 32529, 32768, 0},
				{
					13708, 26338, 30385, 32768, 0},
				{
					7328, 18585, 24870, 32768, 0},
				{
					4691, 13080, 19276, 32768, 0},
				{
					1825, 5253, 8352, 32768, 0},
				{
					29386, 32315, 32624, 32768, 0},
				{
					17160, 29001, 31360, 32768, 0},
				{
					9602, 21862, 27396, 32768, 0},
				{
					5915, 15772, 22148, 32768, 0},
				{
					2786, 7779, 12047, 32768, 0},
				{
					29246, 32450, 32663, 32768, 0},
				{
					18696, 29929, 31818, 32768, 0},
				{
					10510, 23369, 28560, 32768, 0},
				{
					6229, 16499, 23125, 32768, 0},
				{
					2608, 7448, 11705, 32768, 0},
				{
					30753, 32710, 32748, 32768, 0},
				{
					21638, 31487, 32503, 32768, 0},
				{
					12937, 26854, 30870, 32768, 0},
				{
					8182, 20596, 26970, 32768, 0},
				{
					3637, 10269, 15497, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					5244, 12150, 16906, 32768, 0},
				{
					20486, 26858, 29701, 32768, 0},
				{
					7756, 18317, 23735, 32768, 0},
				{
					3452, 9256, 13146, 32768, 0},
				{
					2020, 5206, 8229, 32768, 0},
				{
					1801, 4993, 7903, 32768, 0},
				{
					27051, 31858, 32531, 32768, 0},
				{
					15988, 27531, 30619, 32768, 0},
				{
					9188, 21484, 26719, 32768, 0},
				{
					6273, 17186, 23800, 32768, 0},
				{
					3108, 9355, 14764, 32768, 0},
				{
					31076, 32520, 32680, 32768, 0},
				{
					18119, 30037, 31850, 32768, 0},
				{
					10244, 22969, 27472, 32768, 0},
				{
					4692, 14077, 19273, 32768, 0},
				{
					3694, 11677, 17556, 32768, 0},
				{
					30060, 32581, 32720, 32768, 0},
				{
					21011, 30775, 32120, 32768, 0},
				{
					11931, 24820, 29289, 32768, 0},
				{
					7119, 17662, 24356, 32768, 0},
				{
					3833, 10706, 16304, 32768, 0},
				{
					31954, 32731, 32748, 32768, 0},
				{
					23913, 31724, 32489, 32768, 0},
				{
					15520, 28060, 31286, 32768, 0},
				{
					11517, 23008, 28571, 32768, 0},
				{
					6193, 14508, 20629, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					1035, 2807, 4156, 32768, 0},
				{
					13162, 18138, 20939, 32768, 0},
				{
					2696, 6633, 8755, 32768, 0},
				{
					1373, 4161, 6853, 32768, 0},
				{
					1099, 2746, 4716, 32768, 0},
				{
					340, 1021, 1599, 32768, 0},
				{
					22826, 30419, 32135, 32768, 0},
				{
					10395, 21762, 26942, 32768, 0},
				{
					4726, 12407, 17361, 32768, 0},
				{
					2447, 7080, 10593, 32768, 0},
				{
					1227, 3717, 6011, 32768, 0},
				{
					28156, 31424, 31934, 32768, 0},
				{
					16915, 27754, 30373, 32768, 0},
				{
					9148, 20990, 26431, 32768, 0},
				{
					5950, 15515, 21148, 32768, 0},
				{
					2492, 7327, 11526, 32768, 0},
				{
					30602, 32477, 32670, 32768, 0},
				{
					20026, 29955, 31568, 32768, 0},
				{
					11220, 23628, 28105, 32768, 0},
				{
					6652, 17019, 22973, 32768, 0},
				{
					3064, 8536, 13043, 32768, 0},
				{
					31769, 32724, 32748, 32768, 0},
				{
					22230, 30887, 32373, 32768, 0},
				{
					12234, 25079, 29731, 32768, 0},
				{
					7326, 18816, 25353, 32768, 0},
				{
					3933, 10907, 16616, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
	{
		{
			{
				{
					8896, 16227, 20630, 32768, 0},
				{
					23629, 31782, 32527, 32768, 0},
				{
					15173, 27755, 31321, 32768, 0},
				{
					10158, 21233, 27382, 32768, 0},
				{
					6420, 14857, 21558, 32768, 0},
				{
					3269, 8155, 12646, 32768, 0},
				{
					24835, 32009, 32496, 32768, 0},
				{
					16509, 28421, 31579, 32768, 0},
				{
					10957, 21514, 27418, 32768, 0},
				{
					7881, 15930, 22096, 32768, 0},
				{
					5388, 10960, 15918, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					20745, 30773, 32093, 32768, 0},
				{
					15200, 27221, 30861, 32768, 0},
				{
					13032, 20873, 25667, 32768, 0},
				{
					12285, 18663, 23494, 32768, 0},
				{
					11563, 17481, 21489, 32768, 0},
				{
					26260, 31982, 32320, 32768, 0},
				{
					15397, 28083, 31100, 32768, 0},
				{
					9742, 19217, 24824, 32768, 0},
				{
					3261, 9629, 15362, 32768, 0},
				{
					1480, 4322, 7499, 32768, 0},
				{
					27599, 32256, 32460, 32768, 0},
				{
					16857, 27659, 30774, 32768, 0},
				{
					9551, 18290, 23748, 32768, 0},
				{
					3052, 8933, 14103, 32768, 0},
				{
					2021, 5910, 9787, 32768, 0},
				{
					29005, 32015, 32392, 32768, 0},
				{
					17677, 27694, 30863, 32768, 0},
				{
					9204, 17356, 23219, 32768, 0},
				{
					2403, 7516, 12814, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					10808, 22056, 26896, 32768, 0},
				{
					25739, 32313, 32676, 32768, 0},
				{
					17288, 30203, 32221, 32768, 0},
				{
					11359, 24878, 29896, 32768, 0},
				{
					6949, 17767, 24893, 32768, 0},
				{
					4287, 11796, 18071, 32768, 0},
				{
					27880, 32521, 32705, 32768, 0},
				{
					19038, 31004, 32414, 32768, 0},
				{
					12564, 26345, 30768, 32768, 0},
				{
					8269, 19947, 26779, 32768, 0},
				{
					5674, 14657, 21674, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					25742, 32319, 32671, 32768, 0},
				{
					19557, 31164, 32454, 32768, 0},
				{
					13381, 26381, 30755, 32768, 0},
				{
					10101, 21466, 26722, 32768, 0},
				{
					9209, 19650, 26825, 32768, 0},
				{
					27107, 31917, 32432, 32768, 0},
				{
					18056, 28893, 31203, 32768, 0},
				{
					10200, 21434, 26764, 32768, 0},
				{
					4660, 12913, 19502, 32768, 0},
				{
					2368, 6930, 12504, 32768, 0},
				{
					26960, 32158, 32613, 32768, 0},
				{
					18628, 30005, 32031, 32768, 0},
				{
					10233, 22442, 28232, 32768, 0},
				{
					5471, 14630, 21516, 32768, 0},
				{
					3235, 10767, 17109, 32768, 0},
				{
					27696, 32440, 32692, 32768, 0},
				{
					20032, 31167, 32438, 32768, 0},
				{
					8700, 21341, 28442, 32768, 0},
				{
					5662, 14831, 21795, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					9704, 17294, 21132, 32768, 0},
				{
					26762, 32278, 32633, 32768, 0},
				{
					18382, 29620, 31819, 32768, 0},
				{
					10891, 23475, 28723, 32768, 0},
				{
					6358, 16583, 23309, 32768, 0},
				{
					3248, 9118, 14141, 32768, 0},
				{
					27204, 32573, 32699, 32768, 0},
				{
					19818, 30824, 32329, 32768, 0},
				{
					11772, 25120, 30041, 32768, 0},
				{
					6995, 18033, 25039, 32768, 0},
				{
					3752, 10442, 16098, 32768, 0},
				{
					27222, 32256, 32559, 32768, 0},
				{
					15356, 28399, 31475, 32768, 0},
				{
					8821, 20635, 27057, 32768, 0},
				{
					5511, 14404, 21239, 32768, 0},
				{
					2935, 8222, 13051, 32768, 0},
				{
					24875, 32120, 32529, 32768, 0},
				{
					15233, 28265, 31445, 32768, 0},
				{
					8605, 20570, 26932, 32768, 0},
				{
					5431, 14413, 21196, 32768, 0},
				{
					2994, 8341, 13223, 32768, 0},
				{
					28201, 32604, 32700, 32768, 0},
				{
					21041, 31446, 32456, 32768, 0},
				{
					13221, 26213, 30475, 32768, 0},
				{
					8255, 19385, 26037, 32768, 0},
				{
					4930, 12585, 18830, 32768, 0},
				{
					28768, 32448, 32627, 32768, 0},
				{
					19705, 30561, 32021, 32768, 0},
				{
					11572, 23589, 28220, 32768, 0},
				{
					5532, 15034, 21446, 32768, 0},
				{
					2460, 7150, 11456, 32768, 0},
				{
					29874, 32619, 32699, 32768, 0},
				{
					21621, 31071, 32201, 32768, 0},
				{
					12511, 24747, 28992, 32768, 0},
				{
					6281, 16395, 22748, 32768, 0},
				{
					3246, 9278, 14497, 32768, 0},
				{
					29715, 32625, 32712, 32768, 0},
				{
					20958, 31011, 32283, 32768, 0},
				{
					11233, 23671, 28806, 32768, 0},
				{
					6012, 16128, 22868, 32768, 0},
				{
					3427, 9851, 15414, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					11016, 22111, 26794, 32768, 0},
				{
					25946, 32357, 32677, 32768, 0},
				{
					17890, 30452, 32252, 32768, 0},
				{
					11678, 25142, 29816, 32768, 0},
				{
					6720, 17534, 24584, 32768, 0},
				{
					4230, 11665, 17820, 32768, 0},
				{
					28400, 32623, 32747, 32768, 0},
				{
					21164, 31668, 32575, 32768, 0},
				{
					13572, 27388, 31182, 32768, 0},
				{
					8234, 20750, 27358, 32768, 0},
				{
					5065, 14055, 20897, 32768, 0},
				{
					28981, 32547, 32705, 32768, 0},
				{
					18681, 30543, 32239, 32768, 0},
				{
					10919, 24075, 29286, 32768, 0},
				{
					6431, 17199, 24077, 32768, 0},
				{
					3819, 10464, 16618, 32768, 0},
				{
					26870, 32467, 32693, 32768, 0},
				{
					19041, 30831, 32347, 32768, 0},
				{
					11794, 25211, 30016, 32768, 0},
				{
					6888, 18019, 24970, 32768, 0},
				{
					4370, 12363, 18992, 32768, 0},
				{
					29578, 32670, 32744, 32768, 0},
				{
					23159, 32007, 32613, 32768, 0},
				{
					15315, 28669, 31676, 32768, 0},
				{
					9298, 22607, 28782, 32768, 0},
				{
					6144, 15913, 22968, 32768, 0},
				{
					28110, 32499, 32669, 32768, 0},
				{
					21574, 30937, 32015, 32768, 0},
				{
					12759, 24818, 28727, 32768, 0},
				{
					6545, 16761, 23042, 32768, 0},
				{
					3649, 10597, 16833, 32768, 0},
				{
					28163, 32552, 32728, 32768, 0},
				{
					22101, 31469, 32464, 32768, 0},
				{
					13160, 25472, 30143, 32768, 0},
				{
					7303, 18684, 25468, 32768, 0},
				{
					5241, 13975, 20955, 32768, 0},
				{
					28400, 32631, 32744, 32768, 0},
				{
					22104, 31793, 32603, 32768, 0},
				{
					13557, 26571, 30846, 32768, 0},
				{
					7749, 19861, 26675, 32768, 0},
				{
					4873, 14030, 21234, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					9800, 17635, 21073, 32768, 0},
				{
					26153, 31885, 32527, 32768, 0},
				{
					15038, 27852, 31006, 32768, 0},
				{
					8718, 20564, 26486, 32768, 0},
				{
					5128, 14076, 20514, 32768, 0},
				{
					2636, 7566, 11925, 32768, 0},
				{
					27551, 32504, 32701, 32768, 0},
				{
					18310, 30054, 32100, 32768, 0},
				{
					10211, 23420, 29082, 32768, 0},
				{
					6222, 16876, 23916, 32768, 0},
				{
					3462, 9954, 15498, 32768, 0},
				{
					29991, 32633, 32721, 32768, 0},
				{
					19883, 30751, 32201, 32768, 0},
				{
					11141, 24184, 29285, 32768, 0},
				{
					6420, 16940, 23774, 32768, 0},
				{
					3392, 9753, 15118, 32768, 0},
				{
					28465, 32616, 32712, 32768, 0},
				{
					19850, 30702, 32244, 32768, 0},
				{
					10983, 24024, 29223, 32768, 0},
				{
					6294, 16770, 23582, 32768, 0},
				{
					3244, 9283, 14509, 32768, 0},
				{
					30023, 32717, 32748, 32768, 0},
				{
					22940, 32032, 32626, 32768, 0},
				{
					14282, 27928, 31473, 32768, 0},
				{
					8562, 21327, 27914, 32768, 0},
				{
					4846, 13393, 19919, 32768, 0},
				{
					29981, 32590, 32695, 32768, 0},
				{
					20465, 30963, 32166, 32768, 0},
				{
					11479, 23579, 28195, 32768, 0},
				{
					5916, 15648, 22073, 32768, 0},
				{
					3031, 8605, 13398, 32768, 0},
				{
					31146, 32691, 32739, 32768, 0},
				{
					23106, 31724, 32444, 32768, 0},
				{
					13783, 26738, 30439, 32768, 0},
				{
					7852, 19468, 25807, 32768, 0},
				{
					3860, 11124, 16853, 32768, 0},
				{
					31014, 32724, 32748, 32768, 0},
				{
					23629, 32109, 32628, 32768, 0},
				{
					14747, 28115, 31403, 32768, 0},
				{
					8545, 21242, 27478, 32768, 0},
				{
					4574, 12781, 19067, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					9185, 19694, 24688, 32768, 0},
				{
					26081, 31985, 32621, 32768, 0},
				{
					16015, 29000, 31787, 32768, 0},
				{
					10542, 23690, 29206, 32768, 0},
				{
					6732, 17945, 24677, 32768, 0},
				{
					3916, 11039, 16722, 32768, 0},
				{
					28224, 32566, 32744, 32768, 0},
				{
					19100, 31138, 32485, 32768, 0},
				{
					12528, 26620, 30879, 32768, 0},
				{
					7741, 20277, 26885, 32768, 0},
				{
					4566, 12845, 18990, 32768, 0},
				{
					29933, 32593, 32718, 32768, 0},
				{
					17670, 30333, 32155, 32768, 0},
				{
					10385, 23600, 28909, 32768, 0},
				{
					6243, 16236, 22407, 32768, 0},
				{
					3976, 10389, 16017, 32768, 0},
				{
					28377, 32561, 32738, 32768, 0},
				{
					19366, 31175, 32482, 32768, 0},
				{
					13327, 27175, 31094, 32768, 0},
				{
					8258, 20769, 27143, 32768, 0},
				{
					4703, 13198, 19527, 32768, 0},
				{
					31086, 32706, 32748, 32768, 0},
				{
					22853, 31902, 32583, 32768, 0},
				{
					14759, 28186, 31419, 32768, 0},
				{
					9284, 22382, 28348, 32768, 0},
				{
					5585, 15192, 21868, 32768, 0},
				{
					28291, 32652, 32746, 32768, 0},
				{
					19849, 32107, 32571, 32768, 0},
				{
					14834, 26818, 29214, 32768, 0},
				{
					10306, 22594, 28672, 32768, 0},
				{
					6615, 17384, 23384, 32768, 0},
				{
					28947, 32604, 32745, 32768, 0},
				{
					25625, 32289, 32646, 32768, 0},
				{
					18758, 28672, 31403, 32768, 0},
				{
					10017, 23430, 28523, 32768, 0},
				{
					6862, 15269, 22131, 32768, 0},
				{
					23933, 32509, 32739, 32768, 0},
				{
					19927, 31495, 32631, 32768, 0},
				{
					11903, 26023, 30621, 32768, 0},
				{
					7026, 20094, 27252, 32768, 0},
				{
					5998, 18106, 24437, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					4456, 11274, 15533, 32768, 0},
				{
					21219, 29079, 31616, 32768, 0},
				{
					11173, 23774, 28567, 32768, 0},
				{
					7282, 18293, 24263, 32768, 0},
				{
					4890, 13286, 19115, 32768, 0},
				{
					1890, 5508, 8659, 32768, 0},
				{
					26651, 32136, 32647, 32768, 0},
				{
					14630, 28254, 31455, 32768, 0},
				{
					8716, 21287, 27395, 32768, 0},
				{
					5615, 15331, 22008, 32768, 0},
				{
					2675, 7700, 12150, 32768, 0},
				{
					29954, 32526, 32690, 32768, 0},
				{
					16126, 28982, 31633, 32768, 0},
				{
					9030, 21361, 27352, 32768, 0},
				{
					5411, 14793, 21271, 32768, 0},
				{
					2943, 8422, 13163, 32768, 0},
				{
					29539, 32601, 32730, 32768, 0},
				{
					18125, 30385, 32201, 32768, 0},
				{
					10422, 24090, 29468, 32768, 0},
				{
					6468, 17487, 24438, 32768, 0},
				{
					2970, 8653, 13531, 32768, 0},
				{
					30912, 32715, 32748, 32768, 0},
				{
					20666, 31373, 32497, 32768, 0},
				{
					12509, 26640, 30917, 32768, 0},
				{
					8058, 20629, 27290, 32768, 0},
				{
					4231, 12006, 18052, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					10202, 20633, 25484, 32768, 0},
				{
					27336, 31445, 32352, 32768, 0},
				{
					12420, 24384, 28552, 32768, 0},
				{
					7648, 18115, 23856, 32768, 0},
				{
					5662, 14341, 19902, 32768, 0},
				{
					3611, 10328, 15390, 32768, 0},
				{
					30945, 32616, 32736, 32768, 0},
				{
					18682, 30505, 32253, 32768, 0},
				{
					11513, 25336, 30203, 32768, 0},
				{
					7449, 19452, 26148, 32768, 0},
				{
					4482, 13051, 18886, 32768, 0},
				{
					32022, 32690, 32747, 32768, 0},
				{
					18578, 30501, 32146, 32768, 0},
				{
					11249, 23368, 28631, 32768, 0},
				{
					5645, 16958, 22158, 32768, 0},
				{
					5009, 11444, 16637, 32768, 0},
				{
					31357, 32710, 32748, 32768, 0},
				{
					21552, 31494, 32504, 32768, 0},
				{
					13891, 27677, 31340, 32768, 0},
				{
					9051, 22098, 28172, 32768, 0},
				{
					5190, 13377, 19486, 32768, 0},
				{
					32364, 32740, 32748, 32768, 0},
				{
					24839, 31907, 32551, 32768, 0},
				{
					17160, 28779, 31696, 32768, 0},
				{
					12452, 24137, 29602, 32768, 0},
				{
					6165, 15389, 22477, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					2575, 7281, 11077, 32768, 0},
				{
					14002, 20866, 25402, 32768, 0},
				{
					6343, 15056, 19658, 32768, 0},
				{
					4474, 11858, 17041, 32768, 0},
				{
					2865, 8299, 12534, 32768, 0},
				{
					1344, 3949, 6391, 32768, 0},
				{
					24720, 31239, 32459, 32768, 0},
				{
					12585, 25356, 29968, 32768, 0},
				{
					7181, 18246, 24444, 32768, 0},
				{
					5025, 13667, 19885, 32768, 0},
				{
					2521, 7304, 11605, 32768, 0},
				{
					29908, 32252, 32584, 32768, 0},
				{
					17421, 29156, 31575, 32768, 0},
				{
					9889, 22188, 27782, 32768, 0},
				{
					5878, 15647, 22123, 32768, 0},
				{
					2814, 8665, 13323, 32768, 0},
				{
					30183, 32568, 32713, 32768, 0},
				{
					18528, 30195, 32049, 32768, 0},
				{
					10982, 24606, 29657, 32768, 0},
				{
					6957, 18165, 25231, 32768, 0},
				{
					3508, 10118, 15468, 32768, 0},
				{
					31761, 32736, 32748, 32768, 0},
				{
					21041, 31328, 32546, 32768, 0},
				{
					12568, 26732, 31166, 32768, 0},
				{
					8052, 20720, 27733, 32768, 0},
				{
					4336, 12192, 18396, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
	{
		{
			{
				{
					7062, 16472, 22319, 32768, 0},
				{
					24538, 32261, 32674, 32768, 0},
				{
					13675, 28041, 31779, 32768, 0},
				{
					8590, 20674, 27631, 32768, 0},
				{
					5685, 14675, 22013, 32768, 0},
				{
					3655, 9898, 15731, 32768, 0},
				{
					26493, 32418, 32658, 32768, 0},
				{
					16376, 29342, 32090, 32768, 0},
				{
					10594, 22649, 28970, 32768, 0},
				{
					8176, 17170, 24303, 32768, 0},
				{
					5605, 12694, 19139, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					23888, 31902, 32542, 32768, 0},
				{
					18612, 29687, 31987, 32768, 0},
				{
					16245, 24852, 29249, 32768, 0},
				{
					15765, 22608, 27559, 32768, 0},
				{
					19895, 24699, 27510, 32768, 0},
				{
					28401, 32212, 32457, 32768, 0},
				{
					15274, 27825, 30980, 32768, 0},
				{
					9364, 18128, 24332, 32768, 0},
				{
					2283, 8193, 15082, 32768, 0},
				{
					1228, 3972, 7881, 32768, 0},
				{
					29455, 32469, 32620, 32768, 0},
				{
					17981, 28245, 31388, 32768, 0},
				{
					10921, 20098, 26240, 32768, 0},
				{
					3743, 11829, 18657, 32768, 0},
				{
					2374, 9593, 15715, 32768, 0},
				{
					31068, 32466, 32635, 32768, 0},
				{
					20321, 29572, 31971, 32768, 0},
				{
					10771, 20255, 27119, 32768, 0},
				{
					2795, 10410, 17361, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					9320, 22102, 27840, 32768, 0},
				{
					2795, 10410, 17361, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					9320, 22102, 27840, 32768, 0},
				{
					27057, 32464, 32724, 32768, 0},
				{
					16331, 30268, 32309, 32768, 0},
				{
					10319, 23935, 29720, 32768, 0},
				{
					6189, 16448, 24106, 32768, 0},
				{
					3589, 10884, 18808, 32768, 0},
				{
					29026, 32624, 32748, 32768, 0},
				{
					19226, 31507, 32587, 32768, 0},
				{
					12692, 26921, 31203, 32768, 0},
				{
					7049, 19532, 27635, 32768, 0},
				{
					7727, 15669, 23252, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					28056, 32625, 32748, 32768, 0},
				{
					22383, 32075, 32669, 32768, 0},
				{
					15417, 27098, 31749, 32768, 0},
				{
					18127, 26493, 27190, 32768, 0},
				{
					5461, 16384, 21845, 32768, 0},
				{
					27982, 32091, 32584, 32768, 0},
				{
					19045, 29868, 31972, 32768, 0},
				{
					10397, 22266, 27932, 32768, 0},
				{
					5990, 13697, 21500, 32768, 0},
				{
					1792, 6912, 15104, 32768, 0},
				{
					28198, 32501, 32718, 32768, 0},
				{
					21534, 31521, 32569, 32768, 0},
				{
					11109, 25217, 30017, 32768, 0},
				{
					5671, 15124, 26151, 32768, 0},
				{
					4681, 14043, 18725, 32768, 0},
				{
					28688, 32580, 32741, 32768, 0},
				{
					22576, 32079, 32661, 32768, 0},
				{
					10627, 22141, 28340, 32768, 0},
				{
					9362, 14043, 28087, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					7754, 16948, 22142, 32768, 0},
				{
					25670, 32330, 32691, 32768, 0},
				{
					15663, 29225, 31994, 32768, 0},
				{
					9878, 23288, 29158, 32768, 0},
				{
					6419, 17088, 24336, 32768, 0},
				{
					3859, 11003, 17039, 32768, 0},
				{
					27562, 32595, 32725, 32768, 0},
				{
					17575, 30588, 32399, 32768, 0},
				{
					10819, 24838, 30309, 32768, 0},
				{
					7124, 18686, 25916, 32768, 0},
				{
					4479, 12688, 19340, 32768, 0},
				{
					28385, 32476, 32673, 32768, 0},
				{
					15306, 29005, 31938, 32768, 0},
				{
					8937, 21615, 28322, 32768, 0},
				{
					5982, 15603, 22786, 32768, 0},
				{
					3620, 10267, 16136, 32768, 0},
				{
					27280, 32464, 32667, 32768, 0},
				{
					15607, 29160, 32004, 32768, 0},
				{
					9091, 22135, 28740, 32768, 0},
				{
					6232, 16632, 24020, 32768, 0},
				{
					4047, 11377, 17672, 32768, 0},
				{
					29220, 32630, 32718, 32768, 0},
				{
					19650, 31220, 32462, 32768, 0},
				{
					13050, 26312, 30827, 32768, 0},
				{
					9228, 20870, 27468, 32768, 0},
				{
					6146, 15149, 21971, 32768, 0},
				{
					30169, 32481, 32623, 32768, 0},
				{
					17212, 29311, 31554, 32768, 0},
				{
					9911, 21311, 26882, 32768, 0},
				{
					4487, 13314, 20372, 32768, 0},
				{
					2570, 7772, 12889, 32768, 0},
				{
					30924, 32613, 32708, 32768, 0},
				{
					19490, 30206, 32107, 32768, 0},
				{
					11232, 23998, 29276, 32768, 0},
				{
					6769, 17955, 25035, 32768, 0},
				{
					4398, 12623, 19214, 32768, 0},
				{
					30609, 32627, 32722, 32768, 0},
				{
					19370, 30582, 32287, 32768, 0},
				{
					10457, 23619, 29409, 32768, 0},
				{
					6443, 17637, 24834, 32768, 0},
				{
					4645, 13236, 20106, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8626, 20271, 26216, 32768, 0},
				{
					26707, 32406, 32711, 32768, 0},
				{
					16999, 30329, 32286, 32768, 0},
				{
					11445, 25123, 30286, 32768, 0},
				{
					6411, 18828, 25601, 32768, 0},
				{
					6801, 12458, 20248, 32768, 0},
				{
					29918, 32682, 32748, 32768, 0},
				{
					20649, 31739, 32618, 32768, 0},
				{
					12879, 27773, 31581, 32768, 0},
				{
					7896, 21751, 28244, 32768, 0},
				{
					5260, 14870, 23698, 32768, 0},
				{
					29252, 32593, 32731, 32768, 0},
				{
					17072, 30460, 32294, 32768, 0},
				{
					10653, 24143, 29365, 32768, 0},
				{
					6536, 17490, 23983, 32768, 0},
				{
					4929, 13170, 20085, 32768, 0},
				{
					28137, 32518, 32715, 32768, 0},
				{
					18171, 30784, 32407, 32768, 0},
				{
					11437, 25436, 30459, 32768, 0},
				{
					7252, 18534, 26176, 32768, 0},
				{
					4126, 13353, 20978, 32768, 0},
				{
					31162, 32726, 32748, 32768, 0},
				{
					23017, 32222, 32701, 32768, 0},
				{
					15629, 29233, 32046, 32768, 0},
				{
					9387, 22621, 29480, 32768, 0},
				{
					6922, 17616, 25010, 32768, 0},
				{
					28838, 32265, 32614, 32768, 0},
				{
					19701, 30206, 31920, 32768, 0},
				{
					11214, 22410, 27933, 32768, 0},
				{
					5320, 14177, 23034, 32768, 0},
				{
					5049, 12881, 17827, 32768, 0},
				{
					27484, 32471, 32734, 32768, 0},
				{
					21076, 31526, 32561, 32768, 0},
				{
					12707, 26303, 31211, 32768, 0},
				{
					8169, 21722, 28219, 32768, 0},
				{
					6045, 19406, 27042, 32768, 0},
				{
					27753, 32572, 32745, 32768, 0},
				{
					20832, 31878, 32653, 32768, 0},
				{
					13250, 27356, 31674, 32768, 0},
				{
					7718, 21508, 29858, 32768, 0},
				{
					7209, 18350, 25559, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					7876, 16901, 21741, 32768, 0},
				{
					24001, 31898, 32625, 32768, 0},
				{
					14529, 27959, 31451, 32768, 0},
				{
					8273, 20818, 27258, 32768, 0},
				{
					5278, 14673, 21510, 32768, 0},
				{
					2983, 8843, 14039, 32768, 0},
				{
					28016, 32574, 32732, 32768, 0},
				{
					17471, 30306, 32301, 32768, 0},
				{
					10224, 24063, 29728, 32768, 0},
				{
					6602, 17954, 25052, 32768, 0},
				{
					4002, 11585, 17759, 32768, 0},
				{
					30190, 32634, 32739, 32768, 0},
				{
					17497, 30282, 32270, 32768, 0},
				{
					10229, 23729, 29538, 32768, 0},
				{
					6344, 17211, 24440, 32768, 0},
				{
					3849, 11189, 17108, 32768, 0},
				{
					28570, 32583, 32726, 32768, 0},
				{
					17521, 30161, 32238, 32768, 0},
				{
					10153, 23565, 29378, 32768, 0},
				{
					6455, 17341, 24443, 32768, 0},
				{
					3907, 11042, 17024, 32768, 0},
				{
					30689, 32715, 32748, 32768, 0},
				{
					21546, 31840, 32610, 32768, 0},
				{
					13547, 27581, 31459, 32768, 0},
				{
					8912, 21757, 28309, 32768, 0},
				{
					5548, 15080, 22046, 32768, 0},
				{
					30783, 32540, 32685, 32768, 0},
				{
					17540, 29528, 31668, 32768, 0},
				{
					10160, 21468, 26783, 32768, 0},
				{
					4724, 13393, 20054, 32768, 0},
				{
					2702, 8174, 13102, 32768, 0},
				{
					31648, 32686, 32742, 32768, 0},
				{
					20954, 31094, 32337, 32768, 0},
				{
					12420, 25698, 30179, 32768, 0},
				{
					7304, 19320, 26248, 32768, 0},
				{
					4366, 12261, 18864, 32768, 0},
				{
					31581, 32723, 32748, 32768, 0},
				{
					21373, 31586, 32525, 32768, 0},
				{
					12744, 26625, 30885, 32768, 0},
				{
					7431, 20322, 26950, 32768, 0},
				{
					4692, 13323, 20111, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					7833, 18369, 24095, 32768, 0},
				{
					26650, 32273, 32702, 32768, 0},
				{
					16371, 29961, 32191, 32768, 0},
				{
					11055, 24082, 29629, 32768, 0},
				{
					6892, 18644, 25400, 32768, 0},
				{
					5006, 13057, 19240, 32768, 0},
				{
					29834, 32666, 32748, 32768, 0},
				{
					19577, 31335, 32570, 32768, 0},
				{
					12253, 26509, 31122, 32768, 0},
				{
					7991, 20772, 27711, 32768, 0},
				{
					5677, 15910, 23059, 32768, 0},
				{
					30109, 32532, 32720, 32768, 0},
				{
					16747, 30166, 32252, 32768, 0},
				{
					10134, 23542, 29184, 32768, 0},
				{
					5791, 16176, 23556, 32768, 0},
				{
					4362, 10414, 17284, 32768, 0},
				{
					29492, 32626, 32748, 32768, 0},
				{
					19894, 31402, 32525, 32768, 0},
				{
					12942, 27071, 30869, 32768, 0},
				{
					8346, 21216, 27405, 32768, 0},
				{
					6572, 17087, 23859, 32768, 0},
				{
					32035, 32735, 32748, 32768, 0},
				{
					22957, 31838, 32618, 32768, 0},
				{
					14724, 28572, 31772, 32768, 0},
				{
					10364, 23999, 29553, 32768, 0},
				{
					7004, 18433, 25655, 32768, 0},
				{
					27528, 32277, 32681, 32768, 0},
				{
					16959, 31171, 32096, 32768, 0},
				{
					10486, 23593, 27962, 32768, 0},
				{
					8192, 16384, 23211, 32768, 0},
				{
					8937, 17873, 20852, 32768, 0},
				{
					27715, 32002, 32615, 32768, 0},
				{
					15073, 29491, 31676, 32768, 0},
				{
					11264, 24576, 28672, 32768, 0},
				{
					2341, 18725, 23406, 32768, 0},
				{
					7282, 18204, 25486, 32768, 0},
				{
					28547, 32213, 32657, 32768, 0},
				{
					20788, 29773, 32239, 32768, 0},
				{
					6780, 21469, 30508, 32768, 0},
				{
					5958, 14895, 23831, 32768, 0},
				{
					16384, 21845, 27307, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					5992, 14304, 19765, 32768, 0},
				{
					22612, 31238, 32456, 32768, 0},
				{
					13456, 27162, 31087, 32768, 0},
				{
					8001, 20062, 26504, 32768, 0},
				{
					5168, 14105, 20764, 32768, 0},
				{
					2632, 7771, 12385, 32768, 0},
				{
					27034, 32344, 32709, 32768, 0},
				{
					15850, 29415, 31997, 32768, 0},
				{
					9494, 22776, 28841, 32768, 0},
				{
					6151, 16830, 23969, 32768, 0},
				{
					3461, 10039, 15722, 32768, 0},
				{
					30134, 32569, 32731, 32768, 0},
				{
					15638, 29422, 31945, 32768, 0},
				{
					9150, 21865, 28218, 32768, 0},
				{
					5647, 15719, 22676, 32768, 0},
				{
					3402, 9772, 15477, 32768, 0},
				{
					28530, 32586, 32735, 32768, 0},
				{
					17139, 30298, 32292, 32768, 0},
				{
					10200, 24039, 29685, 32768, 0},
				{
					6419, 17674, 24786, 32768, 0},
				{
					3544, 10225, 15824, 32768, 0},
				{
					31333, 32726, 32748, 32768, 0},
				{
					20618, 31487, 32544, 32768, 0},
				{
					12901, 27217, 31232, 32768, 0},
				{
					8624, 21734, 28171, 32768, 0},
				{
					5104, 14191, 20748, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					11206, 21090, 26561, 32768, 0},
				{
					28759, 32279, 32671, 32768, 0},
				{
					14171, 27952, 31569, 32768, 0},
				{
					9743, 22907, 29141, 32768, 0},
				{
					6871, 17886, 24868, 32768, 0},
				{
					4960, 13152, 19315, 32768, 0},
				{
					31077, 32661, 32748, 32768, 0},
				{
					19400, 31195, 32515, 32768, 0},
				{
					12752, 26858, 31040, 32768, 0},
				{
					8370, 22098, 28591, 32768, 0},
				{
					5457, 15373, 22298, 32768, 0},
				{
					31697, 32706, 32748, 32768, 0},
				{
					17860, 30657, 32333, 32768, 0},
				{
					12510, 24812, 29261, 32768, 0},
				{
					6180, 19124, 24722, 32768, 0},
				{
					5041, 13548, 17959, 32768, 0},
				{
					31552, 32716, 32748, 32768, 0},
				{
					21908, 31769, 32623, 32768, 0},
				{
					14470, 28201, 31565, 32768, 0},
				{
					9493, 22982, 28608, 32768, 0},
				{
					6858, 17240, 24137, 32768, 0},
				{
					32543, 32752, 32756, 32768, 0},
				{
					24286, 32097, 32666, 32768, 0},
				{
					15958, 29217, 32024, 32768, 0},
				{
					10207, 24234, 29958, 32768, 0},
				{
					6929, 18305, 25652, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
		{
			{
				{
					4137, 10847, 15682, 32768, 0},
				{
					17824, 27001, 30058, 32768, 0},
				{
					10204, 22796, 28291, 32768, 0},
				{
					6076, 15935, 22125, 32768, 0},
				{
					3852, 10937, 16816, 32768, 0},
				{
					2252, 6324, 10131, 32768, 0},
				{
					25840, 32016, 32662, 32768, 0},
				{
					15109, 28268, 31531, 32768, 0},
				{
					9385, 22231, 28340, 32768, 0},
				{
					6082, 16672, 23479, 32768, 0},
				{
					3318, 9427, 14681, 32768, 0},
				{
					30594, 32574, 32718, 32768, 0},
				{
					16836, 29552, 31859, 32768, 0},
				{
					9556, 22542, 28356, 32768, 0},
				{
					6305, 16725, 23540, 32768, 0},
				{
					3376, 9895, 15184, 32768, 0},
				{
					29383, 32617, 32745, 32768, 0},
				{
					18891, 30809, 32401, 32768, 0},
				{
					11688, 25942, 30687, 32768, 0},
				{
					7468, 19469, 26651, 32768, 0},
				{
					3909, 11358, 17012, 32768, 0},
				{
					31564, 32736, 32748, 32768, 0},
				{
					20906, 31611, 32600, 32768, 0},
				{
					13191, 27621, 31537, 32768, 0},
				{
					8768, 22029, 28676, 32768, 0},
				{
					5079, 14109, 20906, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
}

var DEFAULT_COEFF_BR_CDF = [][][][][]int{
	{
		{
			{
				{
					14298, 20718, 24174, 32768, 0},
				{
					12536, 19601, 23789, 32768, 0},
				{
					8712, 15051, 19503, 32768, 0},
				{
					6170, 11327, 15434, 32768, 0},
				{
					4742, 8926, 12538, 32768, 0},
				{
					3803, 7317, 10546, 32768, 0},
				{
					1696, 3317, 4871, 32768, 0},
				{
					14392, 19951, 22756, 32768, 0},
				{
					15978, 23218, 26818, 32768, 0},
				{
					12187, 19474, 23889, 32768, 0},
				{
					9176, 15640, 20259, 32768, 0},
				{
					7068, 12655, 17028, 32768, 0},
				{
					5656, 10442, 14472, 32768, 0},
				{
					2580, 4992, 7244, 32768, 0},
				{
					12136, 18049, 21426, 32768, 0},
				{
					13784, 20721, 24481, 32768, 0},
				{
					10836, 17621, 21900, 32768, 0},
				{
					8372, 14444, 18847, 32768, 0},
				{
					6523, 11779, 16000, 32768, 0},
				{
					5337, 9898, 13760, 32768, 0},
				{
					3034, 5860, 8462, 32768, 0},
			},
			{
				{
					15967, 22905, 26286, 32768, 0},
				{
					13534, 20654, 24579, 32768, 0},
				{
					9504, 16092, 20535, 32768, 0},
				{
					6975, 12568, 16903, 32768, 0},
				{
					5364, 10091, 14020, 32768, 0},
				{
					4357, 8370, 11857, 32768, 0},
				{
					2506, 4934, 7218, 32768, 0},
				{
					23032, 28815, 30936, 32768, 0},
				{
					19540, 26704, 29719, 32768, 0},
				{
					15158, 22969, 27097, 32768, 0},
				{
					11408, 18865, 23650, 32768, 0},
				{
					8885, 15448, 20250, 32768, 0},
				{
					7108, 12853, 17416, 32768, 0},
				{
					4231, 8041, 11480, 32768, 0},
				{
					19823, 26490, 29156, 32768, 0},
				{
					18890, 25929, 28932, 32768, 0},
				{
					15660, 23491, 27433, 32768, 0},
				{
					12147, 19776, 24488, 32768, 0},
				{
					9728, 16774, 21649, 32768, 0},
				{
					7919, 14277, 19066, 32768, 0},
				{
					5440, 10170, 14185, 32768, 0},
			},
		},
		{
			{
				{
					14406, 20862, 24414, 32768, 0},
				{
					11824, 18907, 23109, 32768, 0},
				{
					8257, 14393, 18803, 32768, 0},
				{
					5860, 10747, 14778, 32768, 0},
				{
					4475, 8486, 11984, 32768, 0},
				{
					3606, 6954, 10043, 32768, 0},
				{
					1736, 3410, 5048, 32768, 0},
				{
					14430, 20046, 22882, 32768, 0},
				{
					15593, 22899, 26709, 32768, 0},
				{
					12102, 19368, 23811, 32768, 0},
				{
					9059, 15584, 20262, 32768, 0},
				{
					6999, 12603, 17048, 32768, 0},
				{
					5684, 10497, 14553, 32768, 0},
				{
					2822, 5438, 7862, 32768, 0},
				{
					15785, 21585, 24359, 32768, 0},
				{
					18347, 25229, 28266, 32768, 0},
				{
					14974, 22487, 26389, 32768, 0},
				{
					11423, 18681, 23271, 32768, 0},
				{
					8863, 15350, 20008, 32768, 0},
				{
					7153, 12852, 17278, 32768, 0},
				{
					3707, 7036, 9982, 32768, 0},
			},
			{
				{
					15460, 21696, 25469, 32768, 0},
				{
					12170, 19249, 23191, 32768, 0},
				{
					8723, 15027, 19332, 32768, 0},
				{
					6428, 11704, 15874, 32768, 0},
				{
					4922, 9292, 13052, 32768, 0},
				{
					4139, 7695, 11010, 32768, 0},
				{
					2291, 4508, 6598, 32768, 0},
				{
					19856, 26920, 29828, 32768, 0},
				{
					17923, 25289, 28792, 32768, 0},
				{
					14278, 21968, 26297, 32768, 0},
				{
					10910, 18136, 22950, 32768, 0},
				{
					8423, 14815, 19627, 32768, 0},
				{
					6771, 12283, 16774, 32768, 0},
				{
					4074, 7750, 11081, 32768, 0},
				{
					19852, 26074, 28672, 32768, 0},
				{
					19371, 26110, 28989, 32768, 0},
				{
					16265, 23873, 27663, 32768, 0},
				{
					12758, 20378, 24952, 32768, 0},
				{
					10095, 17098, 21961, 32768, 0},
				{
					8250, 14628, 19451, 32768, 0},
				{
					5205, 9745, 13622, 32768, 0},
			},
		},
		{
			{
				{
					10563, 16233, 19763, 32768, 0},
				{
					9794, 16022, 19804, 32768, 0},
				{
					6750, 11945, 15759, 32768, 0},
				{
					4963, 9186, 12752, 32768, 0},
				{
					3845, 7435, 10627, 32768, 0},
				{
					3051, 6085, 8834, 32768, 0},
				{
					1311, 2596, 3830, 32768, 0},
				{
					11246, 16404, 19689, 32768, 0},
				{
					12315, 18911, 22731, 32768, 0},
				{
					10557, 17095, 21289, 32768, 0},
				{
					8136, 14006, 18249, 32768, 0},
				{
					6348, 11474, 15565, 32768, 0},
				{
					5196, 9655, 13400, 32768, 0},
				{
					2349, 4526, 6587, 32768, 0},
				{
					13337, 18730, 21569, 32768, 0},
				{
					19306, 26071, 28882, 32768, 0},
				{
					15952, 23540, 27254, 32768, 0},
				{
					12409, 19934, 24430, 32768, 0},
				{
					9760, 16706, 21389, 32768, 0},
				{
					8004, 14220, 18818, 32768, 0},
				{
					4138, 7794, 10961, 32768, 0},
			},
			{
				{
					10870, 16684, 20949, 32768, 0},
				{
					9664, 15230, 18680, 32768, 0},
				{
					6886, 12109, 15408, 32768, 0},
				{
					4825, 8900, 12305, 32768, 0},
				{
					3630, 7162, 10314, 32768, 0},
				{
					3036, 6429, 9387, 32768, 0},
				{
					1671, 3296, 4940, 32768, 0},
				{
					13819, 19159, 23026, 32768, 0},
				{
					11984, 19108, 23120, 32768, 0},
				{
					10690, 17210, 21663, 32768, 0},
				{
					7984, 14154, 18333, 32768, 0},
				{
					6868, 12294, 16124, 32768, 0},
				{
					5274, 8994, 12868, 32768, 0},
				{
					2988, 5771, 8424, 32768, 0},
				{
					19736, 26647, 29141, 32768, 0},
				{
					18933, 26070, 28984, 32768, 0},
				{
					15779, 23048, 27200, 32768, 0},
				{
					12638, 20061, 24532, 32768, 0},
				{
					10692, 17545, 22220, 32768, 0},
				{
					9217, 15251, 20054, 32768, 0},
				{
					5078, 9284, 12594, 32768, 0},
			},
		},
		{
			{
				{
					2331, 3662, 5244, 32768, 0},
				{
					2891, 4771, 6145, 32768, 0},
				{
					4598, 7623, 9729, 32768, 0},
				{
					3520, 6845, 9199, 32768, 0},
				{
					3417, 6119, 9324, 32768, 0},
				{
					2601, 5412, 7385, 32768, 0},
				{
					600, 1173, 1744, 32768, 0},
				{
					7672, 13286, 17469, 32768, 0},
				{
					4232, 7792, 10793, 32768, 0},
				{
					2915, 5317, 7397, 32768, 0},
				{
					2318, 4356, 6152, 32768, 0},
				{
					2127, 4000, 5554, 32768, 0},
				{
					1850, 3478, 5275, 32768, 0},
				{
					977, 1933, 2843, 32768, 0},
				{
					18280, 24387, 27989, 32768, 0},
				{
					15852, 22671, 26185, 32768, 0},
				{
					13845, 20951, 24789, 32768, 0},
				{
					11055, 17966, 22129, 32768, 0},
				{
					9138, 15422, 19801, 32768, 0},
				{
					7454, 13145, 17456, 32768, 0},
				{
					3370, 6393, 9013, 32768, 0},
			},
			{
				{
					5842, 9229, 10838, 32768, 0},
				{
					2313, 3491, 4276, 32768, 0},
				{
					2998, 6104, 7496, 32768, 0},
				{
					2420, 7447, 9868, 32768, 0},
				{
					3034, 8495, 10923, 32768, 0},
				{
					4076, 8937, 10975, 32768, 0},
				{
					1086, 2370, 3299, 32768, 0},
				{
					9714, 17254, 20444, 32768, 0},
				{
					8543, 13698, 17123, 32768, 0},
				{
					4918, 9007, 11910, 32768, 0},
				{
					4129, 7532, 10553, 32768, 0},
				{
					2364, 5533, 8058, 32768, 0},
				{
					1834, 3546, 5563, 32768, 0},
				{
					1473, 2908, 4133, 32768, 0},
				{
					15405, 21193, 25619, 32768, 0},
				{
					15691, 21952, 26561, 32768, 0},
				{
					12962, 19194, 24165, 32768, 0},
				{
					10272, 17855, 22129, 32768, 0},
				{
					8588, 15270, 20718, 32768, 0},
				{
					8682, 14669, 19500, 32768, 0},
				{
					4870, 9636, 13205, 32768, 0},
			},
		},
		{
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
	{
		{
			{
				{
					14995, 21341, 24749, 32768, 0},
				{
					13158, 20289, 24601, 32768, 0},
				{
					8941, 15326, 19876, 32768, 0},
				{
					6297, 11541, 15807, 32768, 0},
				{
					4817, 9029, 12776, 32768, 0},
				{
					3731, 7273, 10627, 32768, 0},
				{
					1847, 3617, 5354, 32768, 0},
				{
					14472, 19659, 22343, 32768, 0},
				{
					16806, 24162, 27533, 32768, 0},
				{
					12900, 20404, 24713, 32768, 0},
				{
					9411, 16112, 20797, 32768, 0},
				{
					7056, 12697, 17148, 32768, 0},
				{
					5544, 10339, 14460, 32768, 0},
				{
					2954, 5704, 8319, 32768, 0},
				{
					12464, 18071, 21354, 32768, 0},
				{
					15482, 22528, 26034, 32768, 0},
				{
					12070, 19269, 23624, 32768, 0},
				{
					8953, 15406, 20106, 32768, 0},
				{
					7027, 12730, 17220, 32768, 0},
				{
					5887, 10913, 15140, 32768, 0},
				{
					3793, 7278, 10447, 32768, 0},
			},
			{
				{
					15571, 22232, 25749, 32768, 0},
				{
					14506, 21575, 25374, 32768, 0},
				{
					10189, 17089, 21569, 32768, 0},
				{
					7316, 13301, 17915, 32768, 0},
				{
					5783, 10912, 15190, 32768, 0},
				{
					4760, 9155, 13088, 32768, 0},
				{
					2993, 5966, 8774, 32768, 0},
				{
					23424, 28903, 30778, 32768, 0},
				{
					20775, 27666, 30290, 32768, 0},
				{
					16474, 24410, 28299, 32768, 0},
				{
					12471, 20180, 24987, 32768, 0},
				{
					9410, 16487, 21439, 32768, 0},
				{
					7536, 13614, 18529, 32768, 0},
				{
					5048, 9586, 13549, 32768, 0},
				{
					21090, 27290, 29756, 32768, 0},
				{
					20796, 27402, 30026, 32768, 0},
				{
					17819, 25485, 28969, 32768, 0},
				{
					13860, 21909, 26462, 32768, 0},
				{
					11002, 18494, 23529, 32768, 0},
				{
					8953, 15929, 20897, 32768, 0},
				{
					6448, 11918, 16454, 32768, 0},
			},
		},
		{
			{
				{
					15999, 22208, 25449, 32768, 0},
				{
					13050, 19988, 24122, 32768, 0},
				{
					8594, 14864, 19378, 32768, 0},
				{
					6033, 11079, 15238, 32768, 0},
				{
					4554, 8683, 12347, 32768, 0},
				{
					3672, 7139, 10337, 32768, 0},
				{
					1900, 3771, 5576, 32768, 0},
				{
					15788, 21340, 23949, 32768, 0},
				{
					16825, 24235, 27758, 32768, 0},
				{
					12873, 20402, 24810, 32768, 0},
				{
					9590, 16363, 21094, 32768, 0},
				{
					7352, 13209, 17733, 32768, 0},
				{
					5960, 10989, 15184, 32768, 0},
				{
					3232, 6234, 9007, 32768, 0},
				{
					15761, 20716, 23224, 32768, 0},
				{
					19318, 25989, 28759, 32768, 0},
				{
					15529, 23094, 26929, 32768, 0},
				{
					11662, 18989, 23641, 32768, 0},
				{
					8955, 15568, 20366, 32768, 0},
				{
					7281, 13106, 17708, 32768, 0},
				{
					4248, 8059, 11440, 32768, 0},
			},
			{
				{
					14899, 21217, 24503, 32768, 0},
				{
					13519, 20283, 24047, 32768, 0},
				{
					9429, 15966, 20365, 32768, 0},
				{
					6700, 12355, 16652, 32768, 0},
				{
					5088, 9704, 13716, 32768, 0},
				{
					4243, 8154, 11731, 32768, 0},
				{
					2702, 5364, 7861, 32768, 0},
				{
					22745, 28388, 30454, 32768, 0},
				{
					20235, 27146, 29922, 32768, 0},
				{
					15896, 23715, 27637, 32768, 0},
				{
					11840, 19350, 24131, 32768, 0},
				{
					9122, 15932, 20880, 32768, 0},
				{
					7488, 13581, 18362, 32768, 0},
				{
					5114, 9568, 13370, 32768, 0},
				{
					20845, 26553, 28932, 32768, 0},
				{
					20981, 27372, 29884, 32768, 0},
				{
					17781, 25335, 28785, 32768, 0},
				{
					13760, 21708, 26297, 32768, 0},
				{
					10975, 18415, 23365, 32768, 0},
				{
					9045, 15789, 20686, 32768, 0},
				{
					6130, 11199, 15423, 32768, 0},
			},
		},
		{
			{
				{
					13549, 19724, 23158, 32768, 0},
				{
					11844, 18382, 22246, 32768, 0},
				{
					7919, 13619, 17773, 32768, 0},
				{
					5486, 10143, 13946, 32768, 0},
				{
					4166, 7983, 11324, 32768, 0},
				{
					3364, 6506, 9427, 32768, 0},
				{
					1598, 3160, 4674, 32768, 0},
				{
					15281, 20979, 23781, 32768, 0},
				{
					14939, 22119, 25952, 32768, 0},
				{
					11363, 18407, 22812, 32768, 0},
				{
					8609, 14857, 19370, 32768, 0},
				{
					6737, 12184, 16480, 32768, 0},
				{
					5506, 10263, 14262, 32768, 0},
				{
					2990, 5786, 8380, 32768, 0},
				{
					20249, 25253, 27417, 32768, 0},
				{
					21070, 27518, 30001, 32768, 0},
				{
					16854, 24469, 28074, 32768, 0},
				{
					12864, 20486, 25000, 32768, 0},
				{
					9962, 16978, 21778, 32768, 0},
				{
					8074, 14338, 19048, 32768, 0},
				{
					4494, 8479, 11906, 32768, 0},
			},
			{
				{
					13960, 19617, 22829, 32768, 0},
				{
					11150, 17341, 21228, 32768, 0},
				{
					7150, 12964, 17190, 32768, 0},
				{
					5331, 10002, 13867, 32768, 0},
				{
					4167, 7744, 11057, 32768, 0},
				{
					3480, 6629, 9646, 32768, 0},
				{
					1883, 3784, 5686, 32768, 0},
				{
					18752, 25660, 28912, 32768, 0},
				{
					16968, 24586, 28030, 32768, 0},
				{
					13520, 21055, 25313, 32768, 0},
				{
					10453, 17626, 22280, 32768, 0},
				{
					8386, 14505, 19116, 32768, 0},
				{
					6742, 12595, 17008, 32768, 0},
				{
					4273, 8140, 11499, 32768, 0},
				{
					22120, 27827, 30233, 32768, 0},
				{
					20563, 27358, 29895, 32768, 0},
				{
					17076, 24644, 28153, 32768, 0},
				{
					13362, 20942, 25309, 32768, 0},
				{
					10794, 17965, 22695, 32768, 0},
				{
					9014, 15652, 20319, 32768, 0},
				{
					5708, 10512, 14497, 32768, 0},
			},
		},
		{
			{
				{
					5705, 10930, 15725, 32768, 0},
				{
					7946, 12765, 16115, 32768, 0},
				{
					6801, 12123, 16226, 32768, 0},
				{
					5462, 10135, 14200, 32768, 0},
				{
					4189, 8011, 11507, 32768, 0},
				{
					3191, 6229, 9408, 32768, 0},
				{
					1057, 2137, 3212, 32768, 0},
				{
					10018, 17067, 21491, 32768, 0},
				{
					7380, 12582, 16453, 32768, 0},
				{
					6068, 10845, 14339, 32768, 0},
				{
					5098, 9198, 12555, 32768, 0},
				{
					4312, 8010, 11119, 32768, 0},
				{
					3700, 6966, 9781, 32768, 0},
				{
					1693, 3326, 4887, 32768, 0},
				{
					18757, 24930, 27774, 32768, 0},
				{
					17648, 24596, 27817, 32768, 0},
				{
					14707, 22052, 26026, 32768, 0},
				{
					11720, 18852, 23292, 32768, 0},
				{
					9357, 15952, 20525, 32768, 0},
				{
					7810, 13753, 18210, 32768, 0},
				{
					3879, 7333, 10328, 32768, 0},
			},
			{
				{
					8278, 13242, 15922, 32768, 0},
				{
					10547, 15867, 18919, 32768, 0},
				{
					9106, 15842, 20609, 32768, 0},
				{
					6833, 13007, 17218, 32768, 0},
				{
					4811, 9712, 13923, 32768, 0},
				{
					3985, 7352, 11128, 32768, 0},
				{
					1688, 3458, 5262, 32768, 0},
				{
					12951, 21861, 26510, 32768, 0},
				{
					9788, 16044, 20276, 32768, 0},
				{
					6309, 11244, 14870, 32768, 0},
				{
					5183, 9349, 12566, 32768, 0},
				{
					4389, 8229, 11492, 32768, 0},
				{
					3633, 6945, 10620, 32768, 0},
				{
					3600, 6847, 9907, 32768, 0},
				{
					21748, 28137, 30255, 32768, 0},
				{
					19436, 26581, 29560, 32768, 0},
				{
					16359, 24201, 27953, 32768, 0},
				{
					13961, 21693, 25871, 32768, 0},
				{
					11544, 18686, 23322, 32768, 0},
				{
					9372, 16462, 20952, 32768, 0},
				{
					6138, 11210, 15390, 32768, 0},
			},
		},
		{
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
	{
		{
			{
				{
					16138, 22223, 25509, 32768, 0},
				{
					15347, 22430, 26332, 32768, 0},
				{
					9614, 16736, 21332, 32768, 0},
				{
					6600, 12275, 16907, 32768, 0},
				{
					4811, 9424, 13547, 32768, 0},
				{
					3748, 7809, 11420, 32768, 0},
				{
					2254, 4587, 6890, 32768, 0},
				{
					15196, 20284, 23177, 32768, 0},
				{
					18317, 25469, 28451, 32768, 0},
				{
					13918, 21651, 25842, 32768, 0},
				{
					10052, 17150, 21995, 32768, 0},
				{
					7499, 13630, 18587, 32768, 0},
				{
					6158, 11417, 16003, 32768, 0},
				{
					4014, 7785, 11252, 32768, 0},
				{
					15048, 21067, 24384, 32768, 0},
				{
					18202, 25346, 28553, 32768, 0},
				{
					14302, 22019, 26356, 32768, 0},
				{
					10839, 18139, 23166, 32768, 0},
				{
					8715, 15744, 20806, 32768, 0},
				{
					7536, 13576, 18544, 32768, 0},
				{
					5413, 10335, 14498, 32768, 0},
			},
			{
				{
					17394, 24501, 27895, 32768, 0},
				{
					15889, 23420, 27185, 32768, 0},
				{
					11561, 19133, 23870, 32768, 0},
				{
					8285, 14812, 19844, 32768, 0},
				{
					6496, 12043, 16550, 32768, 0},
				{
					4771, 9574, 13677, 32768, 0},
				{
					3603, 6830, 10144, 32768, 0},
				{
					21656, 27704, 30200, 32768, 0},
				{
					21324, 27915, 30511, 32768, 0},
				{
					17327, 25336, 28997, 32768, 0},
				{
					13417, 21381, 26033, 32768, 0},
				{
					10132, 17425, 22338, 32768, 0},
				{
					8580, 15016, 19633, 32768, 0},
				{
					5694, 11477, 16411, 32768, 0},
				{
					24116, 29780, 31450, 32768, 0},
				{
					23853, 29695, 31591, 32768, 0},
				{
					20085, 27614, 30428, 32768, 0},
				{
					15326, 24335, 28575, 32768, 0},
				{
					11814, 19472, 24810, 32768, 0},
				{
					10221, 18611, 24767, 32768, 0},
				{
					7689, 14558, 20321, 32768, 0},
			},
		},
		{
			{
				{
					16214, 22380, 25770, 32768, 0},
				{
					14213, 21304, 25295, 32768, 0},
				{
					9213, 15823, 20455, 32768, 0},
				{
					6395, 11758, 16139, 32768, 0},
				{
					4779, 9187, 13066, 32768, 0},
				{
					3821, 7501, 10953, 32768, 0},
				{
					2293, 4567, 6795, 32768, 0},
				{
					15859, 21283, 23820, 32768, 0},
				{
					18404, 25602, 28726, 32768, 0},
				{
					14325, 21980, 26206, 32768, 0},
				{
					10669, 17937, 22720, 32768, 0},
				{
					8297, 14642, 19447, 32768, 0},
				{
					6746, 12389, 16893, 32768, 0},
				{
					4324, 8251, 11770, 32768, 0},
				{
					16532, 21631, 24475, 32768, 0},
				{
					20667, 27150, 29668, 32768, 0},
				{
					16728, 24510, 28175, 32768, 0},
				{
					12861, 20645, 25332, 32768, 0},
				{
					10076, 17361, 22417, 32768, 0},
				{
					8395, 14940, 19963, 32768, 0},
				{
					5731, 10683, 14912, 32768, 0},
			},
			{
				{
					14433, 21155, 24938, 32768, 0},
				{
					14658, 21716, 25545, 32768, 0},
				{
					9923, 16824, 21557, 32768, 0},
				{
					6982, 13052, 17721, 32768, 0},
				{
					5419, 10503, 15050, 32768, 0},
				{
					4852, 9162, 13014, 32768, 0},
				{
					3271, 6395, 9630, 32768, 0},
				{
					22210, 27833, 30109, 32768, 0},
				{
					20750, 27368, 29821, 32768, 0},
				{
					16894, 24828, 28573, 32768, 0},
				{
					13247, 21276, 25757, 32768, 0},
				{
					10038, 17265, 22563, 32768, 0},
				{
					8587, 14947, 20327, 32768, 0},
				{
					5645, 11371, 15252, 32768, 0},
				{
					22027, 27526, 29714, 32768, 0},
				{
					23098, 29146, 31221, 32768, 0},
				{
					19886, 27341, 30272, 32768, 0},
				{
					15609, 23747, 28046, 32768, 0},
				{
					11993, 20065, 24939, 32768, 0},
				{
					9637, 18267, 23671, 32768, 0},
				{
					7625, 13801, 19144, 32768, 0},
			},
		},
		{
			{
				{
					14438, 20798, 24089, 32768, 0},
				{
					12621, 19203, 23097, 32768, 0},
				{
					8177, 14125, 18402, 32768, 0},
				{
					5674, 10501, 14456, 32768, 0},
				{
					4236, 8239, 11733, 32768, 0},
				{
					3447, 6750, 9806, 32768, 0},
				{
					1986, 3950, 5864, 32768, 0},
				{
					16208, 22099, 24930, 32768, 0},
				{
					16537, 24025, 27585, 32768, 0},
				{
					12780, 20381, 24867, 32768, 0},
				{
					9767, 16612, 21416, 32768, 0},
				{
					7686, 13738, 18398, 32768, 0},
				{
					6333, 11614, 15964, 32768, 0},
				{
					3941, 7571, 10836, 32768, 0},
				{
					22819, 27422, 29202, 32768, 0},
				{
					22224, 28514, 30721, 32768, 0},
				{
					17660, 25433, 28913, 32768, 0},
				{
					13574, 21482, 26002, 32768, 0},
				{
					10629, 17977, 22938, 32768, 0},
				{
					8612, 15298, 20265, 32768, 0},
				{
					5607, 10491, 14596, 32768, 0},
			},
			{
				{
					13569, 19800, 23206, 32768, 0},
				{
					13128, 19924, 23869, 32768, 0},
				{
					8329, 14841, 19403, 32768, 0},
				{
					6130, 10976, 15057, 32768, 0},
				{
					4682, 8839, 12518, 32768, 0},
				{
					3656, 7409, 10588, 32768, 0},
				{
					2577, 5099, 7412, 32768, 0},
				{
					22427, 28684, 30585, 32768, 0},
				{
					20913, 27750, 30139, 32768, 0},
				{
					15840, 24109, 27834, 32768, 0},
				{
					12308, 20029, 24569, 32768, 0},
				{
					10216, 16785, 21458, 32768, 0},
				{
					8309, 14203, 19113, 32768, 0},
				{
					6043, 11168, 15307, 32768, 0},
				{
					23166, 28901, 30998, 32768, 0},
				{
					21899, 28405, 30751, 32768, 0},
				{
					18413, 26091, 29443, 32768, 0},
				{
					15233, 23114, 27352, 32768, 0},
				{
					12683, 20472, 25288, 32768, 0},
				{
					10702, 18259, 23409, 32768, 0},
				{
					8125, 14464, 19226, 32768, 0},
			},
		},
		{
			{
				{
					9040, 14786, 18360, 32768, 0},
				{
					9979, 15718, 19415, 32768, 0},
				{
					7913, 13918, 18311, 32768, 0},
				{
					5859, 10889, 15184, 32768, 0},
				{
					4593, 8677, 12510, 32768, 0},
				{
					3820, 7396, 10791, 32768, 0},
				{
					1730, 3471, 5192, 32768, 0},
				{
					11803, 18365, 22709, 32768, 0},
				{
					11419, 18058, 22225, 32768, 0},
				{
					9418, 15774, 20243, 32768, 0},
				{
					7539, 13325, 17657, 32768, 0},
				{
					6233, 11317, 15384, 32768, 0},
				{
					5137, 9656, 13545, 32768, 0},
				{
					2977, 5774, 8349, 32768, 0},
				{
					21207, 27246, 29640, 32768, 0},
				{
					19547, 26578, 29497, 32768, 0},
				{
					16169, 23871, 27690, 32768, 0},
				{
					12820, 20458, 25018, 32768, 0},
				{
					10224, 17332, 22214, 32768, 0},
				{
					8526, 15048, 19884, 32768, 0},
				{
					5037, 9410, 13118, 32768, 0},
			},
			{
				{
					12339, 17329, 20140, 32768, 0},
				{
					13505, 19895, 23225, 32768, 0},
				{
					9847, 16944, 21564, 32768, 0},
				{
					7280, 13256, 18348, 32768, 0},
				{
					4712, 10009, 14454, 32768, 0},
				{
					4361, 7914, 12477, 32768, 0},
				{
					2870, 5628, 7995, 32768, 0},
				{
					20061, 25504, 28526, 32768, 0},
				{
					15235, 22878, 26145, 32768, 0},
				{
					12985, 19958, 24155, 32768, 0},
				{
					9782, 16641, 21403, 32768, 0},
				{
					9456, 16360, 20760, 32768, 0},
				{
					6855, 12940, 18557, 32768, 0},
				{
					5661, 10564, 15002, 32768, 0},
				{
					25656, 30602, 31894, 32768, 0},
				{
					22570, 29107, 31092, 32768, 0},
				{
					18917, 26423, 29541, 32768, 0},
				{
					15940, 23649, 27754, 32768, 0},
				{
					12803, 20581, 25219, 32768, 0},
				{
					11082, 18695, 23376, 32768, 0},
				{
					7939, 14373, 19005, 32768, 0},
			},
		},
		{
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
	{
		{
			{
				{
					18315, 24289, 27551, 32768, 0},
				{
					16854, 24068, 27835, 32768, 0},
				{
					10140, 17927, 23173, 32768, 0},
				{
					6722, 12982, 18267, 32768, 0},
				{
					4661, 9826, 14706, 32768, 0},
				{
					3832, 8165, 12294, 32768, 0},
				{
					2795, 6098, 9245, 32768, 0},
				{
					17145, 23326, 26672, 32768, 0},
				{
					20733, 27680, 30308, 32768, 0},
				{
					16032, 24461, 28546, 32768, 0},
				{
					11653, 20093, 25081, 32768, 0},
				{
					9290, 16429, 22086, 32768, 0},
				{
					7796, 14598, 19982, 32768, 0},
				{
					6502, 12378, 17441, 32768, 0},
				{
					21681, 27732, 30320, 32768, 0},
				{
					22389, 29044, 31261, 32768, 0},
				{
					19027, 26731, 30087, 32768, 0},
				{
					14739, 23755, 28624, 32768, 0},
				{
					11358, 20778, 25511, 32768, 0},
				{
					10995, 18073, 24190, 32768, 0},
				{
					9162, 14990, 20617, 32768, 0},
			},
			{
				{
					21425, 27952, 30388, 32768, 0},
				{
					18062, 25838, 29034, 32768, 0},
				{
					11956, 19881, 24808, 32768, 0},
				{
					7718, 15000, 20980, 32768, 0},
				{
					5702, 11254, 16143, 32768, 0},
				{
					4898, 9088, 16864, 32768, 0},
				{
					3679, 6776, 11907, 32768, 0},
				{
					23294, 30160, 31663, 32768, 0},
				{
					24397, 29896, 31836, 32768, 0},
				{
					19245, 27128, 30593, 32768, 0},
				{
					13202, 19825, 26404, 32768, 0},
				{
					11578, 19297, 23957, 32768, 0},
				{
					8073, 13297, 21370, 32768, 0},
				{
					5461, 10923, 19745, 32768, 0},
				{
					27367, 30521, 31934, 32768, 0},
				{
					24904, 30671, 31940, 32768, 0},
				{
					23075, 28460, 31299, 32768, 0},
				{
					14400, 23658, 30417, 32768, 0},
				{
					13885, 23882, 28325, 32768, 0},
				{
					14746, 22938, 27853, 32768, 0},
				{
					5461, 16384, 27307, 32768, 0},
			},
		},
		{
			{
				{
					18274, 24813, 27890, 32768, 0},
				{
					15537, 23149, 27003, 32768, 0},
				{
					9449, 16740, 21827, 32768, 0},
				{
					6700, 12498, 17261, 32768, 0},
				{
					4988, 9866, 14198, 32768, 0},
				{
					4236, 8147, 11902, 32768, 0},
				{
					2867, 5860, 8654, 32768, 0},
				{
					17124, 23171, 26101, 32768, 0},
				{
					20396, 27477, 30148, 32768, 0},
				{
					16573, 24629, 28492, 32768, 0},
				{
					12749, 20846, 25674, 32768, 0},
				{
					10233, 17878, 22818, 32768, 0},
				{
					8525, 15332, 20363, 32768, 0},
				{
					6283, 11632, 16255, 32768, 0},
				{
					20466, 26511, 29286, 32768, 0},
				{
					23059, 29174, 31191, 32768, 0},
				{
					19481, 27263, 30241, 32768, 0},
				{
					15458, 23631, 28137, 32768, 0},
				{
					12416, 20608, 25693, 32768, 0},
				{
					10261, 18011, 23261, 32768, 0},
				{
					8016, 14655, 19666, 32768, 0},
			},
			{
				{
					17616, 24586, 28112, 32768, 0},
				{
					15809, 23299, 27155, 32768, 0},
				{
					10767, 18890, 23793, 32768, 0},
				{
					7727, 14255, 18865, 32768, 0},
				{
					6129, 11926, 16882, 32768, 0},
				{
					4482, 9704, 14861, 32768, 0},
				{
					3277, 7452, 11522, 32768, 0},
				{
					22956, 28551, 30730, 32768, 0},
				{
					22724, 28937, 30961, 32768, 0},
				{
					18467, 26324, 29580, 32768, 0},
				{
					13234, 20713, 25649, 32768, 0},
				{
					11181, 17592, 22481, 32768, 0},
				{
					8291, 18358, 24576, 32768, 0},
				{
					7568, 11881, 14984, 32768, 0},
				{
					24948, 29001, 31147, 32768, 0},
				{
					25674, 30619, 32151, 32768, 0},
				{
					20841, 26793, 29603, 32768, 0},
				{
					14669, 24356, 28666, 32768, 0},
				{
					11334, 23593, 28219, 32768, 0},
				{
					8922, 14762, 22873, 32768, 0},
				{
					8301, 13544, 20535, 32768, 0},
			},
		},
		{
			{
				{
					17113, 23733, 27081, 32768, 0},
				{
					14139, 21406, 25452, 32768, 0},
				{
					8552, 15002, 19776, 32768, 0},
				{
					5871, 11120, 15378, 32768, 0},
				{
					4455, 8616, 12253, 32768, 0},
				{
					3469, 6910, 10386, 32768, 0},
				{
					2255, 4553, 6782, 32768, 0},
				{
					18224, 24376, 27053, 32768, 0},
				{
					19290, 26710, 29614, 32768, 0},
				{
					14936, 22991, 27184, 32768, 0},
				{
					11238, 18951, 23762, 32768, 0},
				{
					8786, 15617, 20588, 32768, 0},
				{
					7317, 13228, 18003, 32768, 0},
				{
					5101, 9512, 13493, 32768, 0},
				{
					22639, 28222, 30210, 32768, 0},
				{
					23216, 29331, 31307, 32768, 0},
				{
					19075, 26762, 29895, 32768, 0},
				{
					15014, 23113, 27457, 32768, 0},
				{
					11938, 19857, 24752, 32768, 0},
				{
					9942, 17280, 22282, 32768, 0},
				{
					7167, 13144, 17752, 32768, 0},
			},
			{
				{
					15820, 22738, 26488, 32768, 0},
				{
					13530, 20885, 25216, 32768, 0},
				{
					8395, 15530, 20452, 32768, 0},
				{
					6574, 12321, 16380, 32768, 0},
				{
					5353, 10419, 14568, 32768, 0},
				{
					4613, 8446, 12381, 32768, 0},
				{
					3440, 7158, 9903, 32768, 0},
				{
					24247, 29051, 31224, 32768, 0},
				{
					22118, 28058, 30369, 32768, 0},
				{
					16498, 24768, 28389, 32768, 0},
				{
					12920, 21175, 26137, 32768, 0},
				{
					10730, 18619, 25352, 32768, 0},
				{
					10187, 16279, 22791, 32768, 0},
				{
					9310, 14631, 22127, 32768, 0},
				{
					24970, 30558, 32057, 32768, 0},
				{
					24801, 29942, 31698, 32768, 0},
				{
					22432, 28453, 30855, 32768, 0},
				{
					19054, 25680, 29580, 32768, 0},
				{
					14392, 23036, 28109, 32768, 0},
				{
					12495, 20947, 26650, 32768, 0},
				{
					12442, 20326, 26214, 32768, 0},
			},
		},
		{
			{
				{
					12162, 18785, 22648, 32768, 0},
				{
					12749, 19697, 23806, 32768, 0},
				{
					8580, 15297, 20346, 32768, 0},
				{
					6169, 11749, 16543, 32768, 0},
				{
					4836, 9391, 13448, 32768, 0},
				{
					3821, 7711, 11613, 32768, 0},
				{
					2228, 4601, 7070, 32768, 0},
				{
					16319, 24725, 28280, 32768, 0},
				{
					15698, 23277, 27168, 32768, 0},
				{
					12726, 20368, 25047, 32768, 0},
				{
					9912, 17015, 21976, 32768, 0},
				{
					7888, 14220, 19179, 32768, 0},
				{
					6777, 12284, 17018, 32768, 0},
				{
					4492, 8590, 12252, 32768, 0},
				{
					23249, 28904, 30947, 32768, 0},
				{
					21050, 27908, 30512, 32768, 0},
				{
					17440, 25340, 28949, 32768, 0},
				{
					14059, 22018, 26541, 32768, 0},
				{
					11288, 18903, 23898, 32768, 0},
				{
					9411, 16342, 21428, 32768, 0},
				{
					6278, 11588, 15944, 32768, 0},
			},
			{
				{
					13981, 20067, 23226, 32768, 0},
				{
					16922, 23580, 26783, 32768, 0},
				{
					11005, 19039, 24487, 32768, 0},
				{
					7389, 14218, 19798, 32768, 0},
				{
					5598, 11505, 17206, 32768, 0},
				{
					6090, 11213, 15659, 32768, 0},
				{
					3820, 7371, 10119, 32768, 0},
				{
					21082, 26925, 29675, 32768, 0},
				{
					21262, 28627, 31128, 32768, 0},
				{
					18392, 26454, 30437, 32768, 0},
				{
					14870, 22910, 27096, 32768, 0},
				{
					12620, 19484, 24908, 32768, 0},
				{
					9290, 16553, 22802, 32768, 0},
				{
					6668, 14288, 20004, 32768, 0},
				{
					27704, 31055, 31949, 32768, 0},
				{
					24709, 29978, 31788, 32768, 0},
				{
					21668, 29264, 31657, 32768, 0},
				{
					18295, 26968, 30074, 32768, 0},
				{
					16399, 24422, 29313, 32768, 0},
				{
					14347, 23026, 28104, 32768, 0},
				{
					12370, 19806, 24477, 32768, 0},
			},
		},
		{
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
			{
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
				{
					8192, 16384, 24576, 32768, 0},
			},
		},
	},
}
