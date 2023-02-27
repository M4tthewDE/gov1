package main

const FRAME_LF_COUNT = 4
const WIENER_COEFFS = 3

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
const PARTITION_NONE = 0
const PARTITION_HORZ = 1
const PARTITION_VERT = 2
const PARTITION_SPLIT = 3

const MAX_SB_SIZE = 128
const MAX_FRAME_DISTANCE = 31

const INTRA_EDGE_TAPS = 5
const SUBPEL_BITS = 4
const SCALE_SUBPEL_BITS = 10
const SUBPEL_MASK = 15

var Subpel_Filters = [][][]int{
	{
		{0, 0, 0, 128, 0, 0, 0, 0}, {0, 2, -6, 126, 8, -2, 0, 0},
		{0, 2, -10, 122, 18, -4, 0, 0}, {0, 2, -12, 116, 28, -8, 2, 0},
		{0, 2, -14, 110, 38, -10, 2, 0}, {0, 2, -14, 102, 48, -12, 2, 0},
		{0, 2, -16, 94, 58, -12, 2, 0}, {0, 2, -14, 84, 66, -12, 2, 0},
		{0, 2, -14, 76, 76, -14, 2, 0}, {0, 2, -12, 66, 84, -14, 2, 0},
		{0, 2, -12, 58, 94, -16, 2, 0}, {0, 2, -12, 48, 102, -14, 2, 0},
		{0, 2, -10, 38, 110, -14, 2, 0}, {0, 2, -8, 28, 116, -12, 2, 0},
		{0, 0, -4, 18, 122, -10, 2, 0}, {0, 0, -2, 8, 126, -6, 2, 0},
	},
	{
		{0, 0, 0, 128, 0, 0, 0, 0}, {0, 2, 28, 62, 34, 2, 0, 0},
		{0, 0, 26, 62, 36, 4, 0, 0}, {0, 0, 22, 62, 40, 4, 0, 0},
		{0, 0, 20, 60, 42, 6, 0, 0}, {0, 0, 18, 58, 44, 8, 0, 0},
		{0, 0, 16, 56, 46, 10, 0, 0}, {0, -2, 16, 54, 48, 12, 0, 0},
		{0, -2, 14, 52, 52, 14, -2, 0}, {0, 0, 12, 48, 54, 16, -2, 0},
		{0, 0, 10, 46, 56, 16, 0, 0}, {0, 0, 8, 44, 58, 18, 0, 0},
		{0, 0, 6, 42, 60, 20, 0, 0}, {0, 0, 4, 40, 62, 22, 0, 0},
		{0, 0, 4, 36, 62, 26, 0, 0}, {0, 0, 2, 34, 62, 28, 2, 0},
	},
	{
		{0, 0, 0, 128, 0, 0, 0, 0}, {-2, 2, -6, 126, 8, -2, 2, 0},
		{-2, 6, -12, 124, 16, -6, 4, -2}, {-2, 8, -18, 120, 26, -10, 6, -2},
		{-4, 10, -22, 116, 38, -14, 6, -2}, {-4, 10, -22, 108, 48, -18, 8, -2},
		{-4, 10, -24, 100, 60, -20, 8, -2}, {-4, 10, -24, 90, 70, -22, 10, -2},
		{-4, 12, -24, 80, 80, -24, 12, -4}, {-2, 10, -22, 70, 90, -24, 10, -4},
		{-2, 8, -20, 60, 100, -24, 10, -4}, {-2, 8, -18, 48, 108, -22, 10, -4},
		{-2, 6, -14, 38, 116, -22, 10, -4}, {-2, 6, -10, 26, 120, -18, 8, -2},
		{-2, 4, -6, 16, 124, -12, 6, -2}, {0, 2, -2, 8, 126, -6, 2, -2},
	},
	{
		{0, 0, 0, 128, 0, 0, 0, 0}, {0, 0, 0, 120, 8, 0, 0, 0},
		{0, 0, 0, 112, 16, 0, 0, 0}, {0, 0, 0, 104, 24, 0, 0, 0},
		{0, 0, 0, 96, 32, 0, 0, 0}, {0, 0, 0, 88, 40, 0, 0, 0},
		{0, 0, 0, 80, 48, 0, 0, 0}, {0, 0, 0, 72, 56, 0, 0, 0},
		{0, 0, 0, 64, 64, 0, 0, 0}, {0, 0, 0, 56, 72, 0, 0, 0},
		{0, 0, 0, 48, 80, 0, 0, 0}, {0, 0, 0, 40, 88, 0, 0, 0},
		{0, 0, 0, 32, 96, 0, 0, 0}, {0, 0, 0, 24, 104, 0, 0, 0},
		{0, 0, 0, 16, 112, 0, 0, 0}, {0, 0, 0, 8, 120, 0, 0, 0},
	},
	{
		{0, 0, 0, 128, 0, 0, 0, 0}, {0, 0, -4, 126, 8, -2, 0, 0},
		{0, 0, -8, 122, 18, -4, 0, 0}, {0, 0, -10, 116, 28, -6, 0, 0},
		{0, 0, -12, 110, 38, -8, 0, 0}, {0, 0, -12, 102, 48, -10, 0, 0},
		{0, 0, -14, 94, 58, -10, 0, 0}, {0, 0, -12, 84, 66, -10, 0, 0},
		{0, 0, -12, 76, 76, -12, 0, 0}, {0, 0, -10, 66, 84, -12, 0, 0},
		{0, 0, -10, 58, 94, -14, 0, 0}, {0, 0, -10, 48, 102, -12, 0, 0},
		{0, 0, -8, 38, 110, -12, 0, 0}, {0, 0, -6, 28, 116, -10, 0, 0},
		{0, 0, -4, 18, 122, -8, 0, 0}, {0, 0, -2, 8, 126, -4, 0, 0},
	},
	{
		{0, 0, 0, 128, 0, 0, 0, 0}, {0, 0, 30, 62, 34, 2, 0, 0},
		{0, 0, 26, 62, 36, 4, 0, 0}, {0, 0, 22, 62, 40, 4, 0, 0},
		{0, 0, 20, 60, 42, 6, 0, 0}, {0, 0, 18, 58, 44, 8, 0, 0},
		{0, 0, 16, 56, 46, 10, 0, 0}, {0, 0, 14, 54, 48, 12, 0, 0},
		{0, 0, 12, 52, 52, 12, 0, 0}, {0, 0, 12, 48, 54, 14, 0, 0},
		{0, 0, 10, 46, 56, 16, 0, 0}, {0, 0, 8, 44, 58, 18, 0, 0},
		{0, 0, 6, 42, 60, 20, 0, 0}, {0, 0, 4, 40, 62, 22, 0, 0},
		{0, 0, 4, 36, 62, 26, 0, 0}, {0, 0, 2, 34, 62, 30, 0, 0},
	},
}

var Intra_Edge_Kernel = [][]int{
	{0, 4, 8, 4, 0},
	{0, 5, 6, 5, 0},
	{2, 4, 4, 4, 2},
}

var Dr_Intra_Derivative = []int{
	0, 0, 0, 1023, 0, 0, 547, 0, 0, 372, 0, 0, 0, 0,
	273, 0, 0, 215, 0, 0, 178, 0, 0, 151, 0, 0, 132, 0, 0,
	116, 0, 0, 102, 0, 0, 0, 90, 0, 0, 80, 0, 0, 71, 0, 0,
	64, 0, 0, 57, 0, 0, 51, 0, 0, 45, 0, 0, 0, 40, 0, 0,
	35, 0, 0, 31, 0, 0, 27, 0, 0, 23, 0, 0, 19, 0, 0,
	15, 0, 0, 0, 0, 11, 0, 0, 7, 0, 0, 3, 0, 0,
}

var Ii_Weights_1d = []int{
	60, 58, 56, 54, 52, 50, 48, 47, 45, 44, 42, 41, 39, 38, 37, 35, 34, 33, 32,
	31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 22, 21, 20, 19, 19, 18, 18, 17, 16,
	16, 15, 15, 14, 14, 13, 13, 12, 12, 12, 11, 11, 10, 10, 10, 9, 9, 9, 8,
	8, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 4, 4,
	4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var Sm_Weights_Tx_4x4 = []int{255, 149, 85, 64}
var Sm_Weights_Tx_8x8 = []int{255, 197, 146, 105, 73, 50, 37, 32}
var Sm_Weights_Tx_16x16 = []int{255, 225, 196, 170, 145, 123, 102, 84, 68, 54, 43, 33, 26, 20, 17, 16}
var Sm_Weights_Tx_32x32 = []int{255, 240, 225, 210, 196, 182, 169, 157, 145, 133, 122, 111, 101, 92,
	83, 74,
	66, 59, 52, 45, 39, 34, 29, 25, 21, 17, 14, 12, 10, 9,
	8, 8}
var Sm_Weights_Tx_64x64 = []int{255, 248, 240, 233, 225, 218, 210, 203, 196, 189, 182, 176, 169, 163,
	156,
	150, 144, 138, 133, 127, 121, 116, 111, 106, 101, 96, 91, 86, 82, 77,
	73, 69,
	65, 61, 57, 54, 50, 47, 44, 41, 38, 35, 32, 29, 27, 25, 22, 20, 18, 16,
	15,
	13, 12, 10, 9, 8, 7, 6, 6, 5, 5, 4, 4, 4}

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

var Sgrproj_Xqd_Mid = []int{-32, 31}
var Sgrproj_Xqd_Min = []int{-96, -32}
var Sgrproj_Xqd_Max = []int{31, 95}
var Wiener_Taps_Mid = []int{3, -7, 15}
var Wiener_Taps_Min = []int{-5, -23, -17}
var Wiener_Taps_Max = []int{10, 8, 46}
var Wiener_Taps_K = []int{1, 2, 3}
var SgrParams = [][]int{}

var Intra_Filter_Taps = [][][]int{
	{
		{-6, 10, 0, 0, 0, 12, 0},
		{-5, 2, 10, 0, 0, 9, 0},
		{-3, 1, 1, 10, 0, 7, 0},
		{-3, 1, 1, 2, 10, 5, 0},
		{-4, 6, 0, 0, 0, 2, 12},
		{-3, 2, 6, 0, 0, 2, 9},
		{-3, 2, 2, 6, 0, 2, 7},
		{-3, 1, 2, 2, 6, 3, 5},
	},
	{
		{-10, 16, 0, 0, 0, 10, 0},
		{-6, 0, 16, 0, 0, 6, 0},
		{-4, 0, 0, 16, 0, 4, 0},
		{-2, 0, 0, 0, 16, 2, 0},
		{-10, 16, 0, 0, 0, 0, 10},
		{-6, 0, 16, 0, 0, 0, 6},
		{-4, 0, 0, 16, 0, 0, 4},
		{-2, 0, 0, 0, 16, 0, 2},
	},
	{
		{-8, 8, 0, 0, 0, 16, 0},
		{-8, 0, 8, 0, 0, 16, 0},
		{-8, 0, 0, 8, 0, 16, 0},
		{-8, 0, 0, 0, 8, 16, 0},
		{-4, 4, 0, 0, 0, 0, 16},
		{-4, 0, 4, 0, 0, 0, 16},
		{-4, 0, 0, 4, 0, 0, 16},
		{-4, 0, 0, 0, 4, 0, 16},
	},
	{
		{-2, 8, 0, 0, 0, 10, 0},
		{-1, 3, 8, 0, 0, 6, 0},
		{-1, 2, 3, 8, 0, 4, 0},
		{0, 1, 2, 3, 8, 2, 0},
		{-1, 4, 0, 0, 0, 3, 10},
		{-1, 3, 4, 0, 0, 4, 6},
		{-1, 2, 3, 4, 0, 4, 4},
		{-1, 2, 2, 3, 4, 3, 3},
	},
	{
		{-12, 14, 0, 0, 0, 14, 0},
		{-10, 0, 14, 0, 0, 12, 0},
		{-9, 0, 0, 14, 0, 11, 0},
		{-8, 0, 0, 0, 14, 10, 0},
		{-10, 12, 0, 0, 0, 0, 14},
		{-9, 1, 12, 0, 0, 0, 12},
		{-8, 0, 0, 12, 0, 1, 11},
		{-7, 0, 0, 1, 12, 1, 9},
	},
}

var Mode_To_Angle = []int{0, 90, 180, 45, 135, 113, 157, 203, 67, 0, 0, 0, 0}

var Div_Lut = []int{
	16384, 16320, 16257, 16194, 16132, 16070, 16009, 15948, 15888, 15828, 15768,
	15709, 15650, 15592, 15534, 15477, 15420, 15364, 15308, 15252, 15197, 15142,
	15087, 15033, 14980, 14926, 14873, 14821, 14769, 14717, 14665, 14614, 14564,
	14513, 14463, 14413, 14364, 14315, 14266, 14218, 14170, 14122, 14075, 14028,
	13981, 13935, 13888, 13843, 13797, 13752, 13707, 13662, 13618, 13574, 13530,
	13487, 13443, 13400, 13358, 13315, 13273, 13231, 13190, 13148, 13107, 13066,
	13026, 12985, 12945, 12906, 12866, 12827, 12788, 12749, 12710, 12672, 12633,
	12596, 12558, 12520, 12483, 12446, 12409, 12373, 12336, 12300, 12264, 12228,
	12193, 12157, 12122, 12087, 12053, 12018, 11984, 11950, 11916, 11882, 11848,
	11815, 11782, 11749, 11716, 11683, 11651, 11619, 11586, 11555, 11523, 11491,
	11460, 11429, 11398, 11367, 11336, 11305, 11275, 11245, 11215, 11185, 11155,
	11125, 11096, 11067, 11038, 11009, 10980, 10951, 10923, 10894, 10866, 10838,
	10810, 10782, 10755, 10727, 10700, 10673, 10645, 10618, 10592, 10565, 10538,
	10512, 10486, 10460, 10434, 10408, 10382, 10356, 10331, 10305, 10280, 10255,
	10230, 10205, 10180, 10156, 10131, 10107, 10082, 10058, 10034, 10010, 99869,
	9963, 9939, 9916, 9892, 9869, 9846, 9823, 9800, 9777, 9754, 9732,
	9709, 9687, 9664, 9642, 9620, 9598, 9576, 9554, 9533, 9511, 9489,
	9468, 9447, 9425, 9404, 9383, 9362, 9341, 9321, 9300, 9279, 9259,
	9239, 9218, 9198, 9178, 9158, 9138, 9118, 9098, 9079, 9059, 9039,
	9020, 9001, 8981, 8962, 8943, 8924, 8905, 8886, 8867, 8849, 8830,
	8812, 8793, 8775, 8756, 8738, 8720, 8702, 8684, 8666, 8648, 8630,
	8613, 8595, 8577, 8560, 8542, 8525, 8508, 8490, 8473, 8456, 8439,
	8422, 8405, 8389, 8372, 8355, 8339, 8322, 8306, 8289, 8273, 8257,
	8240, 8224, 8208, 8192,
}

var Warped_Filters = [][]int{
	{0, 0, 127, 1, 0, 0, 0, 0}, {0, -1, 127, 2, 0, 0, 0, 0},
	{1, -3, 127, 4, -1, 0, 0, 0}, {1, -4, 126, 6, -2, 1, 0, 0},
	{1, -5, 126, 8, -3, 1, 0, 0}, {1, -6, 125, 11, -4, 1, 0, 0},
	{1, -7, 124, 13, -4, 1, 0, 0}, {2, -8, 123, 15, -5, 1, 0, 0},
	{2, -9, 122, 18, -6, 1, 0, 0}, {2, -10, 121, 20, -6, 1, 0, 0},
	{2, -11, 120, 22, -7, 2, 0, 0}, {2, -12, 119, 25, -8, 2, 0, 0},
	{3, -13, 117, 27, -8, 2, 0, 0}, {3, -13, 116, 29, -9, 2, 0, 0},
	{3, -14, 114, 32, -10, 3, 0, 0}, {3, -15, 113, 35, -10, 2, 0, 0},
	{3, -15, 111, 37, -11, 3, 0, 0}, {3, -16, 109, 40, -11, 3, 0, 0},
	{3, -16, 108, 42, -12, 3, 0, 0}, {4, -17, 106, 45, -13, 3, 0, 0},
	{4, -17, 104, 47, -13, 3, 0, 0}, {4, -17, 102, 50, -14, 3, 0, 0},
	{4, -17, 100, 52, -14, 3, 0, 0}, {4, -18, 98, 55, -15, 4, 0, 0},
	{4, -18, 96, 58, -15, 3, 0, 0}, {4, -18, 94, 60, -16, 4, 0, 0},
	{4, -18, 91, 63, -16, 4, 0, 0}, {4, -18, 89, 65, -16, 4, 0, 0},
	{4, -18, 87, 68, -17, 4, 0, 0}, {4, -18, 85, 70, -17, 4, 0, 0},
	{4, -18, 82, 73, -17, 4, 0, 0}, {4, -18, 80, 75, -17, 4, 0, 0},
	{4, -18, 78, 78, -18, 4, 0, 0}, {4, -17, 75, 80, -18, 4, 0, 0},
	{4, -17, 73, 82, -18, 4, 0, 0}, {4, -17, 70, 85, -18, 4, 0, 0},
	{4, -17, 68, 87, -18, 4, 0, 0}, {4, -16, 65, 89, -18, 4, 0, 0},
	{4, -16, 63, 91, -18, 4, 0, 0}, {4, -16, 60, 94, -18, 4, 0, 0},
	{3, -15, 58, 96, -18, 4, 0, 0}, {4, -15, 55, 98, -18, 4, 0, 0},
	{3, -14, 52, 100, -17, 4, 0, 0}, {3, -14, 50, 102, -17, 4, 0, 0},
	{3, -13, 47, 104, -17, 4, 0, 0}, {3, -13, 45, 106, -17, 4, 0, 0},
	{3, -12, 42, 108, -16, 3, 0, 0}, {3, -11, 40, 109, -16, 3, 0, 0},
	{3, -11, 37, 111, -15, 3, 0, 0}, {2, -10, 35, 113, -15, 3, 0, 0},
	{3, -10, 32, 114, -14, 3, 0, 0}, {2, -9, 29, 116, -13, 3, 0, 0},
	{2, -8, 27, 117, -13, 3, 0, 0}, {2, -8, 25, 119, -12, 2, 0, 0},
	{2, -7, 22, 120, -11, 2, 0, 0}, {1, -6, 20, 121, -10, 2, 0, 0},
	{1, -6, 18, 122, -9, 2, 0, 0}, {1, -5, 15, 123, -8, 2, 0, 0},
	{1, -4, 13, 124, -7, 1, 0, 0}, {1, -4, 11, 125, -6, 1, 0, 0},
	{1, -3, 8, 126, -5, 1, 0, 0}, {1, -2, 6, 126, -4, 1, 0, 0},
	{0, -1, 4, 127, -3, 1, 0, 0}, {0, 0, 2, 127, -1, 0, 0, 0},

	// [0, 1)
	{0, 0, 0, 127, 1, 0, 0, 0}, {0, 0, -1, 127, 2, 0, 0, 0},
	{0, 1, -3, 127, 4, -2, 1, 0}, {0, 1, -5, 127, 6, -2, 1, 0},
	{0, 2, -6, 126, 8, -3, 1, 0}, {-1, 2, -7, 126, 11, -4, 2, -1},
	{-1, 3, -8, 125, 13, -5, 2, -1}, {-1, 3, -10, 124, 16, -6, 3, -1},
	{-1, 4, -11, 123, 18, -7, 3, -1}, {-1, 4, -12, 122, 20, -7, 3, -1},
	{-1, 4, -13, 121, 23, -8, 3, -1}, {-2, 5, -14, 120, 25, -9, 4, -1},
	{-1, 5, -15, 119, 27, -10, 4, -1}, {-1, 5, -16, 118, 30, -11, 4, -1},
	{-2, 6, -17, 116, 33, -12, 5, -1}, {-2, 6, -17, 114, 35, -12, 5, -1},
	{-2, 6, -18, 113, 38, -13, 5, -1}, {-2, 7, -19, 111, 41, -14, 6, -2},
	{-2, 7, -19, 110, 43, -15, 6, -2}, {-2, 7, -20, 108, 46, -15, 6, -2},
	{-2, 7, -20, 106, 49, -16, 6, -2}, {-2, 7, -21, 104, 51, -16, 7, -2},
	{-2, 7, -21, 102, 54, -17, 7, -2}, {-2, 8, -21, 100, 56, -18, 7, -2},
	{-2, 8, -22, 98, 59, -18, 7, -2}, {-2, 8, -22, 96, 62, -19, 7, -2},
	{-2, 8, -22, 94, 64, -19, 7, -2}, {-2, 8, -22, 91, 67, -20, 8, -2},
	{-2, 8, -22, 89, 69, -20, 8, -2}, {-2, 8, -22, 87, 72, -21, 8, -2},
	{-2, 8, -21, 84, 74, -21, 8, -2}, {-2, 8, -22, 82, 77, -21, 8, -2},
	{-2, 8, -21, 79, 79, -21, 8, -2}, {-2, 8, -21, 77, 82, -22, 8, -2},
	{-2, 8, -21, 74, 84, -21, 8, -2}, {-2, 8, -21, 72, 87, -22, 8, -2},
	{-2, 8, -20, 69, 89, -22, 8, -2}, {-2, 8, -20, 67, 91, -22, 8, -2},
	{-2, 7, -19, 64, 94, -22, 8, -2}, {-2, 7, -19, 62, 96, -22, 8, -2},
	{-2, 7, -18, 59, 98, -22, 8, -2}, {-2, 7, -18, 56, 100, -21, 8, -2},
	{-2, 7, -17, 54, 102, -21, 7, -2}, {-2, 7, -16, 51, 104, -21, 7, -2},
	{-2, 6, -16, 49, 106, -20, 7, -2}, {-2, 6, -15, 46, 108, -20, 7, -2},
	{-2, 6, -15, 43, 110, -19, 7, -2}, {-2, 6, -14, 41, 111, -19, 7, -2},
	{-1, 5, -13, 38, 113, -18, 6, -2}, {-1, 5, -12, 35, 114, -17, 6, -2},
	{-1, 5, -12, 33, 116, -17, 6, -2}, {-1, 4, -11, 30, 118, -16, 5, -1},
	{-1, 4, -10, 27, 119, -15, 5, -1}, {-1, 4, -9, 25, 120, -14, 5, -2},
	{-1, 3, -8, 23, 121, -13, 4, -1}, {-1, 3, -7, 20, 122, -12, 4, -1},
	{-1, 3, -7, 18, 123, -11, 4, -1}, {-1, 3, -6, 16, 124, -10, 3, -1},
	{-1, 2, -5, 13, 125, -8, 3, -1}, {-1, 2, -4, 11, 126, -7, 2, -1},
	{0, 1, -3, 8, 126, -6, 2, 0}, {0, 1, -2, 6, 127, -5, 1, 0},
	{0, 1, -2, 4, 127, -3, 1, 0}, {0, 0, 0, 2, 127, -1, 0, 0},

	// [1, 2)
	{0, 0, 0, 1, 127, 0, 0, 0}, {0, 0, 0, -1, 127, 2, 0, 0},
	{0, 0, 1, -3, 127, 4, -1, 0}, {0, 0, 1, -4, 126, 6, -2, 1},
	{0, 0, 1, -5, 126, 8, -3, 1}, {0, 0, 1, -6, 125, 11, -4, 1},
	{0, 0, 1, -7, 124, 13, -4, 1}, {0, 0, 2, -8, 123, 15, -5, 1},
	{0, 0, 2, -9, 122, 18, -6, 1}, {0, 0, 2, -10, 121, 20, -6, 1},
	{0, 0, 2, -11, 120, 22, -7, 2}, {0, 0, 2, -12, 119, 25, -8, 2},
	{0, 0, 3, -13, 117, 27, -8, 2}, {0, 0, 3, -13, 116, 29, -9, 2},
	{0, 0, 3, -14, 114, 32, -10, 3}, {0, 0, 3, -15, 113, 35, -10, 2},
	{0, 0, 3, -15, 111, 37, -11, 3}, {0, 0, 3, -16, 109, 40, -11, 3},
	{0, 0, 3, -16, 108, 42, -12, 3}, {0, 0, 4, -17, 106, 45, -13, 3},
	{0, 0, 4, -17, 104, 47, -13, 3}, {0, 0, 4, -17, 102, 50, -14, 3},
	{0, 0, 4, -17, 100, 52, -14, 3}, {0, 0, 4, -18, 98, 55, -15, 4},
	{0, 0, 4, -18, 96, 58, -15, 3}, {0, 0, 4, -18, 94, 60, -16, 4},
	{0, 0, 4, -18, 91, 63, -16, 4}, {0, 0, 4, -18, 89, 65, -16, 4},
	{0, 0, 4, -18, 87, 68, -17, 4}, {0, 0, 4, -18, 85, 70, -17, 4},
	{0, 0, 4, -18, 82, 73, -17, 4}, {0, 0, 4, -18, 80, 75, -17, 4},
	{0, 0, 4, -18, 78, 78, -18, 4}, {0, 0, 4, -17, 75, 80, -18, 4},
	{0, 0, 4, -17, 73, 82, -18, 4}, {0, 0, 4, -17, 70, 85, -18, 4},
	{0, 0, 4, -17, 68, 87, -18, 4}, {0, 0, 4, -16, 65, 89, -18, 4},
	{0, 0, 4, -16, 63, 91, -18, 4}, {0, 0, 4, -16, 60, 94, -18, 4},
	{0, 0, 3, -15, 58, 96, -18, 4}, {0, 0, 4, -15, 55, 98, -18, 4},
	{0, 0, 3, -14, 52, 100, -17, 4}, {0, 0, 3, -14, 50, 102, -17, 4},
	{0, 0, 3, -13, 47, 104, -17, 4}, {0, 0, 3, -13, 45, 106, -17, 4},
	{0, 0, 3, -12, 42, 108, -16, 3}, {0, 0, 3, -11, 40, 109, -16, 3},
	{0, 0, 3, -11, 37, 111, -15, 3}, {0, 0, 2, -10, 35, 113, -15, 3},
	{0, 0, 3, -10, 32, 114, -14, 3}, {0, 0, 2, -9, 29, 116, -13, 3},
	{0, 0, 2, -8, 27, 117, -13, 3}, {0, 0, 2, -8, 25, 119, -12, 2},
	{0, 0, 2, -7, 22, 120, -11, 2}, {0, 0, 1, -6, 20, 121, -10, 2},
	{0, 0, 1, -6, 18, 122, -9, 2}, {0, 0, 1, -5, 15, 123, -8, 2},
	{0, 0, 1, -4, 13, 124, -7, 1}, {0, 0, 1, -4, 11, 125, -6, 1},
	{0, 0, 1, -3, 8, 126, -5, 1}, {0, 0, 1, -2, 6, 126, -4, 1},
	{0, 0, 0, -1, 4, 127, -3, 1}, {0, 0, 0, 0, 2, 127, -1, 0},
	// dummy (replicate row index 191)
	{0, 0, 0, 0, 2, 127, -1, 0},
}

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
	if bSize < BLOCK_8X8 {
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

	t.computePrediction(p)
}

// compoute_prediction()
func (t *TileGroup) computePrediction(p *Parser) {
	sbMask := 15
	if p.sequenceHeader.Use128x128SuperBlock {
		sbMask = 31
	}

	subBlockMiRow := p.MiRow & sbMask
	subBlockMiCol := p.MiCol & sbMask

	for plane := 0; plane < 1+Int(t.HasChroma)*2; plane++ {
		planeSz := t.getPlaneResidualSize(p.MiSize, plane, p)
		num4x4W := p.Num4x4BlocksWide[planeSz]
		num4x4H := p.Num4x4BlocksHigh[planeSz]
		log2W := MI_SIZE_LOG2 + Mi_Width_Log2[planeSz]
		log2H := MI_SIZE_LOG2 + Mi_Height_Log2[planeSz]
		subX := 0
		subY := 0
		if plane > 0 {
			subX = Int(p.sequenceHeader.ColorConfig.SubsamplingX)
			subY = Int(p.sequenceHeader.ColorConfig.SubsamplingY)
		}
		baseX := (p.MiCol >> subX) * MI_SIZE
		baseY := (p.MiRow >> subY) * MI_SIZE
		candRow := (p.MiRow >> subY) << subY
		candCol := (p.MiCol >> subX) << subX

		t.IsInterIntra = (Bool(t.IsInter) && p.RefFrame[1] == INTRA_FRAME)

		if t.IsInterIntra {
			var mode int
			if t.InterIntraMode == II_DC_PRED {
				mode = DC_PRED
			} else if t.InterIntraMode == II_V_PRED {
				mode = V_PRED
			} else if t.InterIntraMode == II_H_PRED {
				mode = H_PRED
			} else {
				mode = SMOOTH_PRED
			}
			haveLeft := p.AvailLChroma
			haveAbove := p.AvailUChroma
			if plane == 0 {
				haveLeft := p.AvailL
				haveAbove := p.AvailU
			}
			t.predictIntra(plane, baseX, baseY, haveLeft, haveAbove, p.BlockDecoded[plane][(subBlockMiRow>>subY)-1][(subBlockMiCol>>subX)+num4x4W], p.BlockDecoded[plane][(subBlockMiRow>>subY)+num4x4H][(subBlockMiCol>>subX)-1], mode, log2W, log2H, p)
		}

		if Bool(t.IsInter) {
			predW := t.Block_Width[p.MiSize] >> subX
			predH := t.Block_Height[p.MiSize] >> subY
			someUseIntra := false

			for r := 0; r < (num4x4H << subY); r++ {
				for c := 0; c < (num4x4W << subX); c++ {
					if p.RefFrames[candRow+r][candCol+c][0] == INTRA_FRAME {
						someUseIntra = true
					}
				}
			}

			if someUseIntra {
				predW = num4x4W * 4
				predH = num4x4H * 4
				candRow = p.MiRow
				candCol = p.MiCol
			}
			r := 0
			for y := 0; y < num4x4H; y += predH {
				c := 0
				for x := 0; x < num4x4W; x += predW {
					t.predictInter(plane, baseX+x, baseY+y, predW, predH, candRow+r, candCol+c, p)
				}
			}
		}
	}
}

// 7.11.3 Inter prediction process
func (t *TileGroup) predictInter(plane int, x int, y int, w int, h int, candRow int, candCol int, p *Parser) {
	isCompound := p.RefFrames[candRow][candCol][1] > INTRA_FRAME

	t.roundVariablesDerivationProcess(isCompound, p)

	if plane == 0 && t.MotionMode == LOCALWARP {
		t.warpEstimationProcess(p)
	}

	if plane == 0 && t.MotionMode == LOCALWARP && t.LocalValid {
		t.LocalValid, _, _, _, _ = t.setupShearProcess(t.LocalWarpParams)
	}

	refList := 0
	refFrame := p.RefFrames[candRow][candCol][refList]

	var globalValid bool
	if t.YMode == GLOBALMV || t.YMode == GLOBAL_GLOBALMV && p.GmType[refFrame] > TRANSLATION {
		globalValid, _, _, _, _ = t.setupShearProcess(p.uncompressedHeader.GmParams[refFrame])
	}

	useWarp := 0
	if w < 8 || h < 8 {
		useWarp = 0
	} else if p.uncompressedHeader.ForceIntegerMv {
		useWarp = 0
	} else if t.MotionMode == LOCALWARP && t.LocalValid {
		useWarp = 1
	} else if (t.YMode == GLOBALMV || t.YMode == GLOBAL_GLOBALMV) && p.GmType[refFrame] > TRANSLATION && !t.isScaled(refFrame, p) && globalValid {
		useWarp = 2
	}

	mv := t.Mvs[candRow][candCol][refList]

	var refIdx int
	if !Bool(t.useIntrabc) {
		refIdx = p.uncompressedHeader.ref_frame_idx[refFrame-LAST_FRAME]
	} else {
		refIdx = -1
		p.RefFrameWidth[len(p.RefFrameWidth)-1] = p.uncompressedHeader.FrameWidth
		p.RefFrameHeight[len(p.RefFrameHeight)-1] = p.uncompressedHeader.FrameHeight
		t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = p.uncompressedHeader.UpscaledWidth
	}

	startX, startY, stepX, stepY := t.motionVectorScalingProcess(plane, refIdx, x, y, mv, p)

	if Bool(t.useIntrabc) {
		p.RefFrameWidth[len(p.RefFrameWidth)-1] = p.MiCols * MI_SIZE
		p.RefFrameHeight[len(p.RefFrameHeight)-1] = p.MiRows * MI_SIZE
		t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = p.MiCols * MI_SIZE
	}

	var preds [][][]int
	if useWarp != 0 {
		for i8 := 0; i8 <= ((h - 1) >> 3); i8++ {
			for j8 := 0; j8 <= ((w - 1) >> 3); j8++ {
				// TODO: what exactly is supposed to happen here
				preds[refList] = t.blockWarpProcess(useWarp, plane, refList, x, y, i8, j8, w, h, p)
			}
		}
	}

	if useWarp == 0 {
		preds[refList] = t.blockInterPredictionProcess(plane, refIdx, startX, startY, stepX, stepY, w, h, candRow, candCol, p)
	}

	if isCompound {
		refList = 1

		refFrame := p.RefFrames[candRow][candCol][refList]

		var globalValid bool
		if t.YMode == GLOBALMV || t.YMode == GLOBAL_GLOBALMV && p.GmType[refFrame] > TRANSLATION {
			globalValid, _, _, _, _ = t.setupShearProcess(p.uncompressedHeader.GmParams[refFrame])
		}

		useWarp := 0
		if w < 8 || h < 8 {
			useWarp = 0
		} else if p.uncompressedHeader.ForceIntegerMv {
			useWarp = 0
		} else if t.MotionMode == LOCALWARP && t.LocalValid {
			useWarp = 1
		} else if (t.YMode == GLOBALMV || t.YMode == GLOBAL_GLOBALMV) && p.GmType[refFrame] > TRANSLATION && !t.isScaled(refFrame, p) && globalValid {
			useWarp = 2
		}

		mv := t.Mvs[candRow][candCol][refList]

		var refIdx int
		if !Bool(t.useIntrabc) {
			refIdx = p.uncompressedHeader.ref_frame_idx[refFrame-LAST_FRAME]
		} else {
			refIdx = -1
			p.RefFrameWidth[len(p.RefFrameWidth)-1] = p.uncompressedHeader.FrameWidth
			p.RefFrameHeight[len(p.RefFrameHeight)-1] = p.uncompressedHeader.FrameHeight
			t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = p.uncompressedHeader.UpscaledWidth
		}

		startX, startY, stepX, stepY := t.motionVectorScalingProcess(plane, refIdx, x, y, mv, p)

		if Bool(t.useIntrabc) {
			p.RefFrameWidth[len(p.RefFrameWidth)-1] = p.MiCols * MI_SIZE
			p.RefFrameHeight[len(p.RefFrameHeight)-1] = p.MiRows * MI_SIZE
			t.RefUpscaledWidth[len(t.RefUpscaledWidth)-1] = p.MiCols * MI_SIZE
		}

		var preds [][][]int
		if useWarp != 0 {
			for i8 := 0; i8 <= ((h - 1) >> 3); i8++ {
				for j8 := 0; j8 <= ((w - 1) >> 3); j8++ {
					// TODO: what exactly is supposed to happen here
					preds[refList] = t.blockWarpProcess(useWarp, plane, refList, x, y, i8, j8, w, h, p)
				}
			}
		}

		if useWarp == 0 {
			preds[refList] = t.blockInterPredictionProcess(plane, refIdx, startX, startY, stepX, stepY, w, h, candRow, candCol, p)
		}
	}

	if t.CompoundType == COMPOUND_WEDGE && plane == 0 {
		t.wedgeMaskProcess(w, h, p)
	} else if t.CompoundType == COMPOUND_INTRA {
		t.intraModeVariantMaskProcess(w, h)
	} else if t.CompoundType == COMPOUND_DIFFWTD {
		t.differenceWeightMaskProcess(preds, w, h, p)
	}

	if t.CompoundType == COMPOUND_DISTANCE {
		t.distanceWeightsProcess(candRow, candCol, p)
	}
}

var Quant_Dist_Weight = [][]int{
	{2, 3}, {2, 5}, {2, 7}, {1, MAX_FRAME_DISTANCE},
}

var Quant_Dist_Lookup = [][]int{
	{9, 7}, {11, 5}, {12, 4}, {13, 3},
}

// 7.11.3.15 Distance weights process
func (t *TileGroup) distanceWeightsProcess(candRow int, candCol int, p *Parser) {
	var dist []int
	for refList := 0; refList < 2; refList++ {
		h := p.uncompressedHeader.OrderHints[p.RefFrames[candRow][candCol][refList]]
		dist[refList] = Clip3(0, MAX_FRAME_DISTANCE, Abs(p.uncompressedHeader.getRelativeDist(h, p.uncompressedHeader.OrderHint, p)))
	}
	d0 := dist[1]
	d1 := dist[0]
	order := Int(d0 <= d1)

	if d0 == 0 || d1 == 0 {
		t.FwdWeight = Quant_Dist_Lookup[3][order]
		t.BckWeight = Quant_Dist_Lookup[3][1-order]
	} else {
		var i int
		for i = 0; i < 3; i++ {
			c0 := Quant_Dist_Weight[i][order]
			c1 := Quant_Dist_Weight[i][1-order]

			if Bool(order) {
				if d0*c0 > d1*c1 {
					break
				}
			} else {
				if d0*c0 < d1*c1 {
					break
				}
			}
		}
		t.FwdWeight = Quant_Dist_Lookup[i][order]
		t.BckWeight = Quant_Dist_Lookup[i][1-order]
	}
}

// 7.11.3.12 Difference weight mask process
func (t *TileGroup) differenceWeightMaskProcess(preds [][][]int, w int, h int, p *Parser) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			diff := Abs(preds[0][i][j] - preds[1][i][j])
			diff = Round2(diff, (p.sequenceHeader.ColorConfig.BitDepth-8)+t.InterPostRound)
			m := Clip3(0, 64, 38+diff/16)
			if Bool(t.MaskType) {
				t.Mask[i][j] = 64 - m
			} else {
				t.Mask[i][j] = m
			}
		}
	}
}

// 7.11.3.13 Intra mode variant mask proces
func (t *TileGroup) intraModeVariantMaskProcess(w int, h int) {
	sizeScale := MAX_SB_SIZE / Max(h, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if t.InterIntraMode == II_V_PRED {
				t.Mask[i][j] = Ii_Weights_1d[i*sizeScale]
			} else if t.InterIntraMode == II_H_PRED {
				t.Mask[i][j] = Ii_Weights_1d[j*sizeScale]
			} else if t.InterIntraMode == II_SMOOTH_PRED {
				t.Mask[i][j] = Ii_Weights_1d[Min(i, j)*sizeScale]
			} else {
				t.Mask[i][j] = 32
			}
		}
	}
}

// 7.11.3.11 Wedge mask process
func (t *TileGroup) wedgeMaskProcess(w int, h int, p *Parser) {
	t.InitialiseWedgeMaskTable(p)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			t.Mask[i][j] = WedgeMasks[p.MiSize][t.WedgeSign][t.WedgeIndex][i][j]
		}
	}
}

// 7.11.3.4 Block inter prediction process
func (t *TileGroup) blockInterPredictionProcess(plane int, refIdx int, x int, y int, xStep int, yStep int, w int, h int, candRow int, candCol int, p *Parser) [][]int {
	var ref [][][]int
	if refIdx == -1 {
		ref = p.CurrFrame
	} else {
		ref = t.FrameStore[refIdx]
	}

	subX := 0
	subY := 0
	if plane != 0 {
		subX = Int(p.sequenceHeader.ColorConfig.SubsamplingX)
		subY = Int(p.sequenceHeader.ColorConfig.SubsamplingY)
	}

	lastX := ((t.RefUpscaledWidth[refIdx] + subX) >> subX) - 1
	lastY := ((p.RefFrameHeight[refIdx] + subY) >> subY) - 1

	intermediateHeight := (((h-1)*yStep + (1 << SCALE_SUBPEL_BITS) - 1) >> SCALE_SUBPEL_BITS) + 8

	interpFilter := t.InterpFilters[candRow][candCol][1]
	if w <= 4 {
		if interpFilter == EIGHTTAP || interpFilter == EIGHTTAP_SHARP {
			interpFilter = 4
		} else if interpFilter == EIGHTTAP_SMOOTH {
			interpFilter = 5
		}
	}

	var intermediate [][]int
	for r := 0; r < intermediateHeight; r++ {
		for c := 0; c < w; c++ {
			s := 0
			p := x + xStep*c
			for t := 0; t < 8; t++ {
				s += Subpel_Filters[interpFilter][(p>>6)*SUBPEL_MASK][t] * ref[plane][Clip3(0, lastY, (y>>10)+r-3)][Clip3(0, lastX, (p>>10)+t-3)]
			}
			intermediate[r][c] = Round2(s, t.InterRound0)
		}
	}

	interpFilter = t.InterpFilters[candRow][candCol][0]
	if h <= 4 {
		if interpFilter == EIGHTTAP || interpFilter == EIGHTTAP_SHARP {
			interpFilter = 4
		} else if interpFilter == EIGHTTAP_SMOOTH {
			interpFilter = 5
		}
	}

	var pred [][]int
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			s := 0
			p := (y & 1023) + yStep*r
			for t := 0; t < 8; t++ {
				s += Subpel_Filters[interpFilter][(p>>6)*SUBPEL_MASK][t] * intermediate[(p>>10)+t][c]
			}
			pred[r][c] = Round2(s, t.InterRound1)
		}
	}

	return pred
}

// 7.11.3.5 Block warp process
func (t *TileGroup) blockWarpProcess(useWarp int, plane int, refList int, x int, y int, i8 int, j8 int, w int, h int, p *Parser) [][]int {
	var pred [][]int

	refIdx := p.uncompressedHeader.ref_frame_idx[p.RefFrame[refList]-LAST_FRAME]
	ref := t.FrameStore[refIdx]

	subX := 0
	subY := 0
	if plane != 0 {
		subX = Int(p.sequenceHeader.ColorConfig.SubsamplingX)
		subY = Int(p.sequenceHeader.ColorConfig.SubsamplingY)
	}

	lastX := ((t.RefUpscaledWidth[refIdx] + subX) >> subX) - 1
	lastY := ((p.RefFrameHeight[refIdx] + subY) >> subY) - 1

	srcX := (x + j8*8 + 4) << subX
	srcY := (y + i8*8 + 4) << subY

	var warpParams []int
	if useWarp == 1 {
		warpParams = t.LocalWarpParams
	} else {
		warpParams = p.uncompressedHeader.GmParams[p.RefFrame[refList]]
	}

	dstX := warpParams[2]*srcX + warpParams[3]*srcY + warpParams[0]
	dstY := warpParams[4]*srcX + warpParams[5]*srcY + warpParams[1]

	_, alpha, beta, gamma, delta := t.setupShearProcess(warpParams)

	x4 := dstX >> subX
	y4 := dstY >> subY
	ix4 := x4 >> WARPEDMODEL_PREC_BITS
	sx4 := x4 & ((1 << WARPEDMODEL_PREC_BITS) - 1)
	iy4 := y4 >> WARPEDMODEL_PREC_BITS
	sy4 := y4 & ((1 << WARPEDMODEL_PREC_BITS) - 1)

	var intermediate [][]int
	for i1 := -7; i1 < 8; i1++ {
		for i2 := -4; i2 < 4; i2++ {
			sx := sx4 + alpha*i2 + beta*i1
			offs := Round2(sx, WARPEDDIFF_PREC_BITS) + WARPEDPIXEL_PREC_SHIFTS
			s := 0

			for i3 := 0; i3 < 8; i3++ {
				s += Warped_Filters[offs][i3] * ref[plane][Clip3(0, lastY, iy4+i1)][Clip3(0, lastX, ix4+i2-3+i3)]
			}
			intermediate[(i1 + 7)][(i2 + 4)] = Round2(s, t.InterRound0)
		}

	}

	for i1 := -4; i1 < Min(4, h-i8*8-4); i1++ {
		for i2 := -4; i2 < Min(4, w-j8*8-4); i2++ {
			sy := sy4 + gamma*i2 + delta*i1
			offs := Round2(sy, WARPEDDIFF_PREC_BITS) + WARPEDPIXEL_PREC_SHIFTS
			s := 0

			for i3 := 0; i3 < 8; i3++ {
				s += Warped_Filters[offs][i3] * intermediate[i1+i3+4][i2+4]
			}
			pred[i8*8+i1+4][j8*8+i2+4] = Round2(s, t.InterRound1)
		}
	}

	return pred
}

// 7.11.3.3  Motion vector scaling process
func (t *TileGroup) motionVectorScalingProcess(plane int, refIdx int, x int, y int, mv []int, p *Parser) (int, int, int, int) {
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameWidth / 2)) / p.uncompressedHeader.FrameWidth
	yScale := ((t.RefUpscaledHeight[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameHeight / 2)) / p.uncompressedHeader.FrameHeight

	subX := 0
	subY := 0
	if plane != 0 {
		subX = Int(p.sequenceHeader.ColorConfig.SubsamplingX)
		subY = Int(p.sequenceHeader.ColorConfig.SubsamplingY)
	}

	halfSample := (1 << (SUBPEL_BITS - 1))
	origX := ((x << SUBPEL_BITS) + ((2 * mv[1]) >> subX) + halfSample)
	origY := ((y << SUBPEL_BITS) + ((2 * mv[0]) >> subY) + halfSample)

	baseX := (origX*xScale - (halfSample << REF_SCALE_SHIFT))
	baseY := (origY*yScale - (halfSample << REF_SCALE_SHIFT))

	off := ((1 << (SCALE_SUBPEL_BITS - SUBPEL_BITS)) / 2)

	startX := (Round2Signed(baseX, REF_SCALE_SHIFT+SUBPEL_BITS-SCALE_SUBPEL_BITS) + off)
	startY := (Round2Signed(baseY, REF_SCALE_SHIFT+SUBPEL_BITS-SCALE_SUBPEL_BITS) + off)

	stepX := Round2Signed(xScale, REF_SCALE_SHIFT-SCALE_SUBPEL_BITS)
	stepY := Round2Signed(yScale, REF_SCALE_SHIFT-SCALE_SUBPEL_BITS)

	return startX, startY, stepX, stepY
}

// 7.11.3.6 Setup shear process
func (t *TileGroup) setupShearProcess(warpParams []int) (bool, int, int, int, int) {
	alpha0 := Clip3(-32768, 32767, warpParams[2]-(1<<WARPEDMODEL_PREC_BITS))
	beta0 := Clip3(-32768, 32767, warpParams[3])

	divShift, divFactor := t.resolveDivisorProcess(warpParams[2])

	v := warpParams[4] << WARPEDMODEL_PREC_BITS
	gamma0 := Clip3(-32768, 32767, Round2Signed(v*divFactor, divShift))
	w := warpParams[3] * warpParams[4]
	delta0 := Clip3(-32768, 32767, warpParams[5]-Round2Signed(w*divFactor, divShift)-1<<WARPEDMODEL_PREC_BITS)

	alpha := Round2Signed(alpha0, WARP_PARAM_REDUCE_BITS) << WARP_PARAM_REDUCE_BITS
	beta := Round2Signed(beta0, WARP_PARAM_REDUCE_BITS) << WARP_PARAM_REDUCE_BITS
	gamma := Round2Signed(gamma0, WARP_PARAM_REDUCE_BITS) << WARP_PARAM_REDUCE_BITS
	delta := Round2Signed(delta0, WARP_PARAM_REDUCE_BITS) << WARP_PARAM_REDUCE_BITS

	warpValid := true
	if 4*Abs(alpha)+7*Abs(beta) >= (1 << WARPEDMODEL_PREC_BITS) {
		warpValid = false
	}

	if 4*Abs(gamma)+4*Abs(delta) >= (1 << WARPEDMODEL_PREC_BITS) {
		warpValid = false
	}

	return warpValid, alpha, beta, gamma, delta
}

// 7.11.3.8 Warp estimation process
func (t *TileGroup) warpEstimationProcess(p *Parser) {
	var A [][]int
	var Bx []int
	var By []int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			A[i][j] = 0
		}

		Bx[i] = 0
		By[i] = 0
	}

	// TODO: matrix is symmetrical, so an entry is omitted!
	w4 := p.Num4x4BlocksWide[p.MiSize]
	h4 := p.Num4x4BlocksHigh[p.MiSize]
	midY := p.MiRow*4 + h4*2 - 1
	midX := p.MiCol*4 + w4*2 - 1
	suy := midY * 8
	sux := midX * 8
	duy := suy + t.Mv[0][0]
	dux := sux + t.Mv[0][1]

	for i := 0; i < t.NumSamples; i++ {
		sy := t.CandList[i][0] - suy
		sx := t.CandList[i][1] - sux
		dy := t.CandList[i][2] - duy
		dx := t.CandList[i][3] - dux

		if Abs(sx-dx) < LS_MV_MAX && Abs(sy-dy) < LS_MV_MAX {
			A[0][0] += LsProduct(sx, sx) + 8
			A[0][1] += LsProduct(sx, sy) + 4
			A[1][1] += LsProduct(sy, sx) + 8
			Bx[0] += LsProduct(sx, dx) + 8
			Bx[1] += LsProduct(sy, dx) + 4
			Bx[0] += LsProduct(sx, dy) + 4
			Bx[1] += LsProduct(sy, dy) + 8
		}
	}

	det := A[0][0]*A[1][1] - A[0][1]*A[0][1]

	if det == 0 {
		t.LocalValid = false
		return
	} else {
		t.LocalValid = true
	}

	divShift, divFactor := t.resolveDivisorProcess(det)

	divShift -= WARPEDMODEL_PREC_BITS

	if divShift < 0 {
		divFactor = divFactor << (divShift)
		divShift = 0
	}

	t.LocalWarpParams[2] = t.diag(A[1][1]*Bx[0]-A[0][1]*Bx[1], divFactor, divShift)
	t.LocalWarpParams[3] = t.nondiag(-A[1][1]*Bx[0]+A[0][0]*Bx[1], divFactor, divShift)
	t.LocalWarpParams[4] = t.nondiag(A[1][1]*By[0]-A[0][1]*By[1], divFactor, divShift)
	t.LocalWarpParams[5] = t.diag(-A[1][1]*By[0]+A[0][0]*By[1], divFactor, divShift)

	mvx := t.Mv[0][1]
	mvy := t.Mv[0][0]
	vx := mvx*(1<<(WARPEDMODEL_PREC_BITS-3)) - (midX*(t.LocalWarpParams[2]-(1<<WARPEDMODEL_PREC_BITS)) + midY*t.LocalWarpParams[3])
	vy := mvy*(1<<(WARPEDMODEL_PREC_BITS-3)) - (midX * (t.LocalWarpParams[4] + midY + (t.LocalWarpParams[5]) - (1 << WARPEDMODEL_PREC_BITS)))

	t.LocalWarpParams[0] = Clip3(-WARPEDMODEL_TRANS_CLAMP, WARPEDMODEL_TRANS_CLAMP-1, vx)
	t.LocalWarpParams[1] = Clip3(-WARPEDMODEL_TRANS_CLAMP, WARPEDMODEL_TRANS_CLAMP-1, vy)
}

// nondiag(v)
func (t *TileGroup) nondiag(v int, divFactor int, divShift int) int {
	return Clip3(-WARPEDMODEL_NONDIAGAFFINE_CLAMP+1, WARPEDMODEL_NONDIAGAFFINE_CLAMP-1, Round2Signed(v*divFactor, divShift))
}

// diag(v)
func (t *TileGroup) diag(v int, divFactor int, divShift int) int {
	return Clip3((1<<WARPEDMODEL_PREC_BITS)-WARPEDMODEL_NONDIAGAFFINE_CLAMP+1, (1<<WARPEDMODEL_PREC_BITS)+WARPEDMODEL_NONDIAGAFFINE_CLAMP-1, Round2Signed(v*divFactor, divShift))
}

// 7.11.3.7 Resolve divisor process
func (t *TileGroup) resolveDivisorProcess(d int) (int, int) {
	n := FloorLog2(Abs(d))
	e := Abs(d) - (1 << n)

	var f int
	if n > DIV_LUT_BITS {
		f = Round2(e, n-DIV_LUT_BITS)
	} else {
		f = e << (DIV_LUT_BITS - n)
	}

	divShift := n + DIV_LUT_PREC_BITS
	var divFactor int
	if d < 0 {
		divFactor = -Div_Lut[f]
	} else {
		divFactor = Div_Lut[f]
	}

	return divShift, divFactor
}

// 7.11.3.2 Rounding variables derivation process
func (t *TileGroup) roundVariablesDerivationProcess(isCompound bool, p *Parser) {
	t.InterRound0 = 3
	t.InterRound1 = 3
	if isCompound {
		t.InterRound1 = 7
	} else {
		t.InterRound1 = 11
	}

	if p.sequenceHeader.ColorConfig.BitDepth == 12 {
		t.InterRound0 = t.InterRound0 + 2
	}

	if p.sequenceHeader.ColorConfig.BitDepth == 12 && !isCompound {
		t.InterRound1 = t.InterRound1 - 2
	}

}

// 7.11.2 Intra prediction process
// predict_intra( plane, x, y, haveLeft, haveAbove, haveAboveRight, haveBelowLeft, mode, log2W, log2H )
func (t *TileGroup) predictIntra(plane int, x int, y int, haveLeft bool, haveAbove bool, haveAboveRight int, haveBelowLeft int, mode int, log2W int, log2H int, p *Parser) {
	w := 1 << log2W
	h := 1 << log2H
	maxX := (p.MiCols * MI_SIZE) - 1
	maxY := (p.MiRows * MI_SIZE) - 1

	if plane > 0 {
		maxX = ((p.MiCols * MI_SIZE) >> Int(p.sequenceHeader.ColorConfig.SubsamplingX)) - 1
		maxY = ((p.MiRows * MI_SIZE) >> Int(p.sequenceHeader.ColorConfig.SubsamplingY)) - 1
	}

	for i := 0; i < w+h-1; i++ {
		if Int(haveAbove) == 0 && Int(haveLeft) == 1 {
			t.AboveRow[i] = p.CurrFrame[plane][y][x-1]
		} else if Int(haveAbove) == 0 && Int(haveLeft) == 0 {
			t.AboveRow[i] = (1 << (p.sequenceHeader.ColorConfig.BitDepth - 1)) - 1

		} else {
			aboveLimit := Min(maxX, x+w-1)
			if Bool(haveAboveRight) {
				aboveLimit = Min(maxX, x+2*w-1)
			}
			t.AboveRow[i] = p.CurrFrame[plane][y-1][Min(aboveLimit, x+i)]
		}
	}

	for i := 0; i < w+h-1; i++ {
		if Int(haveLeft) == 0 && Int(haveAbove) == 1 {
			t.LeftCol[i] = p.CurrFrame[plane][y-1][x]
		} else if Int(haveLeft) == 0 && Int(haveAbove) == 0 {
			t.AboveRow[i] = (1 << (p.sequenceHeader.ColorConfig.BitDepth - 1)) + 1

		} else {
			leftLimit := Min(maxY, y+h-1)
			if Bool(haveBelowLeft) {
				leftLimit = Min(maxY, y+2*h-1)
			}
			t.AboveRow[i] = p.CurrFrame[plane][Min(leftLimit, y+i)][x-1]
		}
	}

	if Int(haveAbove) == 1 && Int(haveLeft) == 1 {
		t.AboveRow[len(t.AboveRow)-1] = p.CurrFrame[plane][y-1][x-1]
	} else if Int(haveAbove) == 1 {
		t.AboveRow[len(t.AboveRow)-1] = p.CurrFrame[plane][y-1][x]
	} else if Int(haveLeft) == 1 {
		t.AboveRow[len(t.AboveRow)-1] = p.CurrFrame[plane][y][x-1]
	} else {
		t.AboveRow[len(t.AboveRow)-1] = 1 << (p.sequenceHeader.ColorConfig.BitDepth - 1)
	}

	t.LeftCol[len(t.LeftCol)-1] = t.AboveRow[len(t.AboveRow)-1]

	var pred [][]int
	if plane == 0 && Bool(t.UseFilterIntra) {
		pred = t.recursiveIntraPredictionProcess(w, h, p)
	} else if t.isDirectionalMode(mode) {
		pred = t.directionalIntraPredictionProcess(plane, x, y, Int(haveLeft), Int(haveAbove), mode, w, h, maxX, maxY, p)
	} else if mode == SMOOTH_PRED || mode == SMOOTH_V_PRED || mode == SMOOTH_H_PRED {
		pred = t.smoothIntraPredictionProcess(mode, log2W, log2H, w, h)
	} else if mode == DC_PRED {
		pred = t.dcIntraPredictionProcess(haveLeft, haveAbove, log2W, log2H, w, h, p)
	} else {
		pred = t.basicIntraPredictionProcess(w, h)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			p.CurrFrame[plane][y+i][x+j] = pred[i][j]
		}
	}
}

// 7.11.2.2 Basic intra prediction process
func (t *TileGroup) basicIntraPredictionProcess(w int, h int) [][]int {
	var pred [][]int

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			base := t.AboveRow[j] + t.LeftCol[i] - t.AboveRow[len(t.AboveRow)-1]
			pLeft := Abs(base - t.LeftCol[i])
			pTop := Abs(base - t.AboveRow[j])
			pTopLeft := Abs(base - t.AboveRow[len(t.AboveRow)-1])

			if pLeft <= pTop && pLeft <= pTopLeft {
				pred[i][j] = t.LeftCol[i]
			} else if pTop <= pTopLeft {
				pred[i][j] = t.AboveRow[j]
			} else {
				pred[i][j] = t.AboveRow[len(t.AboveRow)-1]
			}
		}
	}

	return pred
}

// 7.11.2.5 DC intra prediction process
func (t *TileGroup) dcIntraPredictionProcess(haveLeft bool, haveAbove bool, log2W int, log2H int, w int, h int, p *Parser) [][]int {
	var pred [][]int
	if haveLeft && haveAbove {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				sum := 0
				for k := 0; k < h; k++ {
					sum += t.LeftCol[k]
				}
				for k := 0; k < w; k++ {
					sum += t.AboveRow[k]
				}
				sum += (w + h) >> 1
				avg := sum / (w + h)
				pred[i][j] = avg
			}
		}
	} else if haveLeft && !haveAbove {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				sum := 0
				for k := 0; k < h; k++ {
					sum += t.LeftCol[k]
				}
				leftAvg := Clip1((sum+(h>>1))>>log2H, p)
				pred[i][j] = leftAvg
			}
		}

	} else if !haveLeft && !haveAbove {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				pred[i][j] = 1 << (p.sequenceHeader.ColorConfig.BitDepth - 1)
			}
		}

	}

	return pred
}

// 7.11.2.6 Smooth intra prediction process
func (t *TileGroup) smoothIntraPredictionProcess(mode int, log2W int, log2H int, w int, h int) [][]int {
	var pred [][]int
	if mode == SMOOTH_PRED {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var smWeightsX []int
				switch log2W {
				case 2:
					smWeightsX = Sm_Weights_Tx_4x4
				case 3:
					smWeightsX = Sm_Weights_Tx_8x8
				case 4:
					smWeightsX = Sm_Weights_Tx_16x16
				case 5:
					smWeightsX = Sm_Weights_Tx_32x32
				case 6:
					smWeightsX = Sm_Weights_Tx_64x64
				}

				var smWeightsY []int
				switch log2H {
				case 2:
					smWeightsY = Sm_Weights_Tx_4x4
				case 3:
					smWeightsY = Sm_Weights_Tx_8x8
				case 4:
					smWeightsY = Sm_Weights_Tx_16x16
				case 5:
					smWeightsY = Sm_Weights_Tx_32x32
				case 6:
					smWeightsY = Sm_Weights_Tx_64x64
				}

				smoothPred := smWeightsY[i]*t.AboveRow[j] + (256-smWeightsY[i])*t.LeftCol[h-1] + smWeightsX[j]*t.LeftCol[i] + (256-smWeightsX[j])*t.AboveRow[w-1]
				pred[i][j] = Round2(smoothPred, 9)
			}
		}
	} else if mode == SMOOTH_V_PRED {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var smWeights []int
				switch log2H {
				case 2:
					smWeights = Sm_Weights_Tx_4x4
				case 3:
					smWeights = Sm_Weights_Tx_8x8
				case 4:
					smWeights = Sm_Weights_Tx_16x16
				case 5:
					smWeights = Sm_Weights_Tx_32x32
				case 6:
					smWeights = Sm_Weights_Tx_64x64
				}

				smoothPred := smWeights[i]*t.AboveRow[j] + (256-smWeights[i])*t.LeftCol[h-1]
				pred[i][j] = Round2(smoothPred, 8)
			}
		}
	} else if mode == SMOOTH_H_PRED {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var smWeights []int
				switch log2W {
				case 2:
					smWeights = Sm_Weights_Tx_4x4
				case 3:
					smWeights = Sm_Weights_Tx_8x8
				case 4:
					smWeights = Sm_Weights_Tx_16x16
				case 5:
					smWeights = Sm_Weights_Tx_32x32
				case 6:
					smWeights = Sm_Weights_Tx_64x64
				}

				smoothPred := smWeights[j]*t.LeftCol[i] + (256-smWeights[j])*t.AboveRow[w-1]
				pred[i][j] = Round2(smoothPred, 8)
			}
		}
	}

	return pred
}

// 7.11.2.4. Directional intra prediction process
func (t *TileGroup) directionalIntraPredictionProcess(plane int, x int, y int, haveLeft int, haveAbove int, mode int, w int, h int, maxX int, maxY int, p *Parser) [][]int {
	var pred [][]int

	angleDelta := t.AngleDeltaUV
	if plane == 0 {
		angleDelta = t.AngleDeltaY
	}

	pAngle := Mode_To_Angle[mode] + angleDelta*ANGLE_STEP
	upsampleAbove := false
	upsampleLeft := false

	if Int(p.sequenceHeader.EnableIntraEdgeFilter) == 1 {
		var filterType int
		if pAngle != 90 && pAngle != 180 {
			if pAngle > 90 && pAngle < 180 && (w+h) >= 24 {
				t.LeftCol[len(t.LeftCol)] = t.filterCornerProcess()
				t.AboveRow[len(t.AboveRow)] = t.filterCornerProcess()
			}
			filterType = Int(t.getFilterType(plane, p))

			if haveAbove == 1 {
				strength := t.intraEdgeFilterStrengthSelectionProcess(w, h, filterType, pAngle-90)
				sumPart := 0
				if pAngle < 90 {
					sumPart = h
				}
				numPx := Min(w, (maxX-x+1)) + sumPart
				t.intraEdgeFilterProcess(numPx, strength, 0)
			}

			if haveLeft == 1 {
				strength := t.intraEdgeFilterStrengthSelectionProcess(w, h, filterType, pAngle-180)
				sumPart := 0
				if pAngle > 180 {
					sumPart = w
				}
				numPx := Min(h, (maxY-y+1)) + sumPart
				t.intraEdgeFilterProcess(numPx, strength, 1)
			}
		}

		upsampleAbove = t.intraEdgeUpsampleSelectionProcess(w, h, filterType, pAngle-90)

		sumPart := 0
		if pAngle < 90 {
			sumPart = h
		}
		numPx := w + sumPart

		if upsampleAbove {
			t.intraEdgeUpsampleProcess(numPx, false, p)
		}

		upsampleLeft = t.intraEdgeUpsampleSelectionProcess(w, h, filterType, pAngle-180)

		sumPart = 0
		if pAngle > 180 {
			sumPart = w
		}
		numPx = h + sumPart

		if upsampleLeft {
			t.intraEdgeUpsampleProcess(numPx, true, p)
		}

	}

	var dx int
	if pAngle < 90 {
		dx = Dr_Intra_Derivative[pAngle]
	} else if pAngle > 90 && pAngle < 180 {
		dx = Dr_Intra_Derivative[180-pAngle]
	}

	var dy int
	if pAngle > 90 && pAngle < 180 {
		dy = Dr_Intra_Derivative[pAngle-90]
	} else if pAngle > 180 {
		dy = Dr_Intra_Derivative[270-pAngle]
	}

	if pAngle < 90 {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				idx := (i + i) * dx
				base := (idx >> (6 - Int(upsampleAbove))) + (j << Int(upsampleAbove))
				shift := ((idx << Int(upsampleAbove)) >> 1) & 0x1F
				maxBaseX := (w + h - 1) << Int(upsampleAbove)

				if base < maxBaseX {
					pred[i][j] = Round2(t.AboveRow[base]*(32-shift)+t.AboveRow[base+1]*shift, 5)
				} else {
					pred[i][j] = t.AboveRow[maxBaseX]
				}
			}

		}
	} else if pAngle > 90 && pAngle < 180 {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				idx := (j << 6) - (i+1)*dx
				base := (idx >> (6 - Int(upsampleAbove)))

				if base >= -(1 << Int(upsampleAbove)) {
					shift := ((idx << Int(upsampleAbove)) >> 1) & 0x1F
					pred[i][j] = Round2(t.AboveRow[base]*(32-shift)+t.AboveRow[base+1]*shift, 5)
				} else {
					idx = (i << 6) - (j+1)*dy
					base = idx >> (6 - Int(upsampleLeft))
					shift := ((idx << Int(upsampleLeft)) >> 1) & 0x1F
					pred[i][j] = Round2(t.LeftCol[base]*(32-shift)+t.LeftCol[base+1]*shift, 5)
				}
			}
		}

	} else if pAngle > 180 {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				idx := (j + 1) * dy
				base := (idx >> (6 >> Int(upsampleLeft))) + (i << Int(upsampleLeft))
				shift := ((idx << Int(upsampleLeft)) >> 1) & 0x1F
				pred[i][j] = Round2(t.LeftCol[base]*(32-shift)+t.LeftCol[base+1]*shift, 5)
			}
		}

	} else if pAngle == 90 {
		for j := 0; j < w; j++ {
			for i := 0; i < h; i++ {
				pred[i][j] = t.AboveRow[j]
			}
		}
	} else if pAngle == 180 {
		for j := 0; j < w; j++ {
			for i := 0; i < h; i++ {
				pred[i][j] = t.LeftCol[j]
			}
		}
	}

	return pred
}

// 7.11.2.11 Intra edge upsample process
func (t *TileGroup) intraEdgeUpsampleProcess(numPx int, dir bool, p *Parser) {
	// does this actually modify those arrays?
	var buf []int
	if !dir {
		buf = t.AboveRow
	} else {
		buf = t.LeftCol
	}

	var dup []int
	dup[0] = buf[len(buf)-1]
	for i := -1; i < numPx; i++ {
		dup[i+2] = buf[i]
	}
	dup[numPx+2] = buf[numPx-1]

	buf[len(buf)-2] = dup[0]
	for i := 0; i < numPx; i++ {
		s := -dup[i] + (9 * dup[i+1]) + (9 * dup[i+2]) - dup[i+3]
		s = Clip1(Round2(s, 4), p)
		buf[2*i-1] = s
		buf[2*i] = dup[i+2]
	}
}

// 7.11.2.10 Intra edge upsample selection process
func (t *TileGroup) intraEdgeUpsampleSelectionProcess(w int, h int, filterType int, delta int) bool {
	d := Abs(delta)
	blkWh := w + h
	var useUpsample bool

	if d <= 0 || d >= 40 {
		useUpsample = false
	} else if filterType == 0 {
		useUpsample = blkWh <= 16
	} else {
		useUpsample = blkWh <= 8
	}

	return useUpsample
}

// 7.11.2.12 Intra edge filter process
func (t *TileGroup) intraEdgeFilterProcess(sz int, strength int, left int) {
	if strength == 0 {
		return
	}

	var edge []int
	for i := 0; i < sz; i++ {
		if Bool(left) {
			edge[i] = t.LeftCol[i-1]
		} else {
			edge[i] = t.AboveRow[i-1]
		}
	}

	for i := 0; i < sz; i++ {
		s := 0
		for j := 0; j < INTRA_EDGE_TAPS; j++ {
			k := Clip3(0, sz-12, i-2+j)
			s += Intra_Edge_Kernel[strength-1][j] * edge[k]
			if left == 1 {
				t.LeftCol[i-1] = (s + 8) >> 4
			}
			if left == 0 {
				t.AboveRow[i-1] = (s + 8) >> 4
			}
		}
	}
}

// 7.11.2.9. Intra edge filter strength selection process
func (t *TileGroup) intraEdgeFilterStrengthSelectionProcess(w int, h int, filterType int, delta int) int {
	d := Abs(delta)
	blkWh := w + h

	strength := 0
	if filterType == 0 {
		if blkWh <= 8 {
			if d >= 56 {
				strength = 1
			}
		} else if blkWh <= 12 {
			if d >= 40 {
				strength = 1
			}
		} else if blkWh <= 16 {
			if d >= 8 {
				strength = 1
			}
			if d >= 16 {
				strength = 2
			}
			if d >= 32 {
				strength = 3
			}
		} else if blkWh <= 32 {
			strength = 1
			if d >= 4 {
				strength = 2
			}
			if d >= 4 {
				strength = 3
			}
		}
	} else {
		if blkWh <= 8 {
			if d >= 40 {
				strength = 1
			}
			if d >= 64 {
				strength = 2
			}
		} else if blkWh <= 16 {
			if d >= 20 {
				strength = 1
			}
			if d >= 48 {
				strength = 2
			}
		} else if blkWh <= 24 {
			if d >= 4 {
				strength = 3
			}
		} else {
			strength = 3
		}
	}

	return strength
}

// 7.11.2.8. Intra filter type process
func (t *TileGroup) getFilterType(plane int, p *Parser) bool {
	aboveSmooth := false
	leftSmooth := false

	condition := p.AvailUChroma
	if plane == 0 {
		condition = p.AvailU
	}

	if condition {
		r := p.MiRow - 1
		c := p.MiCol

		if plane > 0 {
			if p.sequenceHeader.ColorConfig.SubsamplingX && Bool(p.MiCol&1) {
				c++
			}
			if p.sequenceHeader.ColorConfig.SubsamplingY && Bool(p.MiRow&1) {
				r--
			}
		}
		aboveSmooth = t.isSmooth(r, c, plane, p)
	}

	condition = p.AvailLChroma
	if plane == 0 {
		condition = p.AvailL
	}

	if condition {
		r := p.MiRow
		c := p.MiCol - 1

		if plane > 0 {
			if p.sequenceHeader.ColorConfig.SubsamplingX && Bool(p.MiCol&1) {
				c--
			}
			if p.sequenceHeader.ColorConfig.SubsamplingY && Bool(p.MiRow&1) {
				r++
			}
		}
		aboveSmooth = t.isSmooth(r, c, plane, p)
	}

	return aboveSmooth || leftSmooth
}

// is_smooth( row, col, plane )
func (t *TileGroup) isSmooth(row int, col int, plane int, p *Parser) bool {
	var mode int
	if plane == 0 {
		mode = t.YModes[row][col]
	} else {
		if p.RefFrames[row][col][0] > INTRA_FRAME {
			return false
		}
		mode = t.UVModes[row][col]
	}

	return mode == SMOOTH_PRED || mode == SMOOTH_V_PRED || mode == SMOOTH_H_PRED
}

// 7.11.2.7. Filter corner process
func (t *TileGroup) filterCornerProcess() int {
	s := t.LeftCol[0]*5 + t.AboveRow[len(t.AboveRow)-1]*6 + t.AboveRow[0]*5
	return Round2(s, 4)
}

// 7.11.2.3. Recursive intra prediction process
func (t *TileGroup) recursiveIntraPredictionProcess(w int, h int, parser *Parser) [][]int {
	w4 := w >> 2
	h2 := h >> 1

	var pred [][]int
	for i2 := 0; i2 <= h2-1; i2++ {
		for j4 := 0; j4 <= w4-1; j4++ {
			var p []int
			for i := 0; i <= 6; i++ {
				if i < 5 {
					if i2 == 0 {
						p[i] = t.AboveRow[(j4<<2)+i-1]
					} else if j4 == 0 && i == 0 {
						p[i] = t.LeftCol[(i2<<1)-1]
					} else {
						p[i] = pred[(i2<<1)-1][(j4<<2)+i-1]
					}
				} else {
					if j4 == 0 {
						p[i] = t.LeftCol[(i2<<1)+i-5]
					} else {
						p[i] = pred[(i2<<1)+i-5][(j4<<2)-1]

					}
				}
			}

			var pr int
			for i1 := 0; i1 <= 1; i1++ {
				for j1 := 0; j1 <= 3; j1++ {
					pr = 0
					for i := 0; i <= 6; i++ {
						pr += Intra_Filter_Taps[t.FilterIntraMode][(i1<<2)+j1][i] * p[i]
					}
					pred[(i2<<1)+i1][(j4<<2)+j1] = Clip1(Round2Signed(pr, INTRA_FILTER_SCALE_BITS), parser)
				}

			}
		}
	}

	return pred
}

// get_plane_residual_size( subsize, plane)
func (t *TileGroup) getPlaneResidualSize(subsize int, plane int, p *Parser) int {
	subx := 0
	suby := 0

	if plane > 0 {
		subx = Int(p.sequenceHeader.ColorConfig.SubsamplingX)
		suby = Int(p.sequenceHeader.ColorConfig.SubsamplingY)
	}

	return Subsampled_Size[subsize][subx][suby]
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
		p.MiSize > BLOCK_4X4 &&
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

	if p.MiSize > BLOCK_4X4 && allowSelect && p.uncompressedHeader.TxMode == TX_MODE_SELECT {
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
	if p.MiSize >= BLOCK_8X8 &&
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

		if p.MiSize >= BLOCK_8X8 && t.Block_Width[p.MiSize] <= 64 && t.Block_Height[p.MiSize] <= 64 && Bool(p.uncompressedHeader.AllowScreenContentTools) {
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
	if p.MiSize >= BLOCK_8X8 {
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
					sbSize = BLOCK_128X128
				} else {
					sbSize = BLOCK_64X64
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

	if p.MiSize >= BLOCK_8X8 {

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
		sbSize = BLOCK_128X128
	} else {
		sbSize = BLOCK_64X64
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
		sbSize = BLOCK_128X128
	} else {
		sbSize = BLOCK_64X64
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

	cdefSize4 := p.Num4x4BlocksWide[BLOCK_64X64]
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
