package shared

const BLOCK_INVALID = 3
const BLOCK_SIZES = 22
const BLOCK_4X4 = 0
const BLOCK_4X8 = 1
const BLOCK_8X4 = 2
const BLOCK_8X8 = 3
const BLOCK_8X16 = 4
const BLOCK_16X8 = 5
const BLOCK_16X16 = 6
const BLOCK_16X32 = 7
const BLOCK_32X16 = 8
const BLOCK_32X32 = 9
const BLOCK_32X64 = 10
const BLOCK_64X32 = 11
const BLOCK_64X64 = 12
const BLOCK_64X128 = 13
const BLOCK_128X64 = 14
const BLOCK_128X128 = 15
const BLOCK_4X16 = 16
const BLOCK_16X4 = 17
const BLOCK_8X32 = 18
const BLOCK_32X8 = 19
const BLOCK_16X64 = 20
const BLOCK_64X16 = 21

var Partition_Subsize = [][]int{
	{
		BLOCK_4X4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128X128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4X4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_128X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X128,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X4,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_32X8,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_64X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
	{
		BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_4X16,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_8X32,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_16X64,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
		BLOCK_INVALID, BLOCK_INVALID, BLOCK_INVALID,
	},
}

var Subsampled_Size = [][][]int{
	{{BLOCK_4X4, BLOCK_4X4}, {BLOCK_4X4, BLOCK_4X4}},
	{{BLOCK_4X8, BLOCK_4X4}, {BLOCK_INVALID, BLOCK_4X4}},
	{{BLOCK_8X4, BLOCK_INVALID}, {BLOCK_4X4, BLOCK_4X4}},
	{{BLOCK_8X8, BLOCK_8X4}, {BLOCK_4X8, BLOCK_4X4}},
	{{BLOCK_8X16, BLOCK_8X8}, {BLOCK_INVALID, BLOCK_4X8}},
	{{BLOCK_16X8, BLOCK_INVALID}, {BLOCK_8X8, BLOCK_8X4}},
	{{BLOCK_16X16, BLOCK_16X8}, {BLOCK_8X16, BLOCK_8X8}},
	{{BLOCK_16X32, BLOCK_16X16}, {BLOCK_INVALID, BLOCK_8X16}},
	{{BLOCK_32X16, BLOCK_INVALID}, {BLOCK_16X16, BLOCK_16X8}},
	{{BLOCK_32X32, BLOCK_32X16}, {BLOCK_16X32, BLOCK_16X16}},
	{{BLOCK_32X64, BLOCK_32X32}, {BLOCK_INVALID, BLOCK_16X32}},
	{{BLOCK_64X32, BLOCK_INVALID}, {BLOCK_32X32, BLOCK_32X16}},
	{{BLOCK_64X64, BLOCK_64X32}, {BLOCK_32X64, BLOCK_32X32}},
	{{BLOCK_64X128, BLOCK_64X64}, {BLOCK_INVALID, BLOCK_32X64}},
	{{BLOCK_128X64, BLOCK_INVALID}, {BLOCK_64X64, BLOCK_64X32}},
	{{BLOCK_128X128, BLOCK_128X64}, {BLOCK_64X128, BLOCK_64X64}},
	{{BLOCK_4X16, BLOCK_4X8}, {BLOCK_INVALID, BLOCK_4X8}},
	{{BLOCK_16X4, BLOCK_INVALID}, {BLOCK_8X4, BLOCK_8X4}},
	{{BLOCK_8X32, BLOCK_8X16}, {BLOCK_INVALID, BLOCK_4X16}},
	{{BLOCK_32X8, BLOCK_INVALID}, {BLOCK_16X8, BLOCK_16X4}},
	{{BLOCK_16X64, BLOCK_16X32}, {BLOCK_INVALID, BLOCK_8X32}},
	{{BLOCK_64X16, BLOCK_INVALID}, {BLOCK_32X16, BLOCK_32X8}},
}

const NONE = -1
const INTRA_FRAME = 0
const LAST_FRAME = 1
const LAST2_FRAME = 2
const LAST3_FRAME = 3
const GOLDEN_FRAME = 4
const BWDREF_FRAME = 5
const ALTREF2_FRAME = 6
const ALTREF_FRAME = 7

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

const SUPERRES_NUM = 8

const TX_MODE_LARGEST = 1
const TX_MODE_SELECT = 2

const MI_SIZE_LOG2 = 2

var MI_WIDTH_LOG2 = []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 0, 2, 1, 3, 2, 4}
var MI_HEIGHT_LOG2 = []int{0, 1, 0, 1, 2, 1, 2, 3, 2, 3, 4, 3, 4, 5, 4, 5, 2, 0, 3, 1, 4, 2}

var NUM_4X4_BLOCKS_WIDE = []int{
	1, 1, 2, 2, 2, 4, 4, 4, 8, 8, 8,
	16, 16, 16, 32, 32, 1, 4, 2, 8, 4, 16,
}

var NUM_4X4_BLOCKS_HIGH = []int{
	1, 2, 1, 2, 4, 2, 4, 8, 4, 8, 16,
	8, 16, 32, 16, 32, 4, 1, 8, 2, 16, 4,
}

var BLOCK_WIDTH = []int{
	1 * 4, 1 * 4, 2 * 4, 2 * 4, 2 * 4, 4 * 4, 4 * 4, 4 * 4, 8 * 4, 8 * 4, 8 * 4,
	16 * 4, 16 * 4, 16 * 4, 32 * 4, 32 * 4, 1 * 4, 4 * 4, 2 * 4, 8 * 4, 4 * 4, 16 * 4,
}

var BLOCK_HEIGHT = []int{
	1 * 4, 2 * 4, 1 * 4, 2 * 4, 4 * 4, 2 * 4, 4 * 4, 8 * 4, 4 * 4, 8 * 4, 16 * 4,
	8 * 4, 16 * 4, 32 * 4, 16 * 4, 32 * 4, 4 * 4, 1 * 4, 8 * 4, 2 * 4, 16 * 4, 4 * 4,
}

const WARPEDMODEL_PREC_BITS = 16
const WARPEDMODEL_NONDIAGAFFINE_CLAMP = 1 << 13
const WARPEDMODEL_TRANS_CLAMP = 1 << 23
const WARPEDPIXEL_PREC_SHIFTS = 1 << 6
const WARPEDDIFF_PREC_BITS = 10
const WARP_PARAM_REDUCE_BITS = 6

const GM_TRANS_ONLY_PREC_BITS = 3
const GM_TRANS_PREC_BITS = 6
const GM_ABS_TRANS_ONLY_BITS = 9
const GM_ABS_ALPHA_BITS = 12
const GM_ABS_TRANS_BITS = 12
const GM_ALPHA_PREC_BITS = 15

const TRANSLATION = 1

const EIGHTTAP = 0
const EIGHTTAP_SMOOTH = 1
const EIGHTTAP_SHARP = 2
const BILINEAR = 3
const SWITCHABLE = 4

const NUM_REF_FRAMES = 8
const REFS_PER_FRAME = 7
const KEY_FRAME = 0
const PRIMARY_REF_NONE = 7
const MAX_SEGMENTS = 8

const SEG_LVL_ALT_Q = 0
const SEG_LVL_REF_FRAME = 5
const SEG_LVL_SKIP = 6
const SEG_LVL_GLOBALMV = 7
const SEG_LVL_MAX = 8

const MAX_LOOP_FILTER = 63

const IDENTITY = 0
const ROTZOOM = 2
const AFFINE = 3

var Segmentation_Feature_Bits = []int{8, 6, 6, 6, 6, 3, 0, 0}
var Segmentation_Feature_Signed = []int{1, 1, 1, 1, 1, 0, 0, 0}
var Segmentation_Feature_Max = []int{255, MAX_LOOP_FILTER, MAX_LOOP_FILTER, MAX_LOOP_FILTER, MAX_LOOP_FILTER, 7, 0, 0}

const RESTORE_NONE = 0
const RESTORE_WIENER = 1
const RESTORE_SGRPROJ = 2
const RESTORE_SWITCHABLE = 3

var REMAP_LR_TYPE = []int{
	RESTORE_NONE, RESTORE_SWITCHABLE, RESTORE_WIENER, RESTORE_SGRPROJ,
}

const RESTORATION_TILESIZE_MAX = 256

const MV_CONTEXTS = 2
const FRAME_LF_COUNT = 4
