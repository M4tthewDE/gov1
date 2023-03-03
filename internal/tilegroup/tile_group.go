package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/parser"
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
	66, 59, 52, 45, 39, 34, 29, 25, 21, 17, 14, 12, 10, 9, 8, 8}
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
	LocalWarpParams []int

	FrameStore [][][][]int
	Mask       [][]int

	FwdWeight int
	BckWeight int
}

func NewTileGroup(p *parser.Parser, sz int) TileGroup {
	t := TileGroup{}
	t.build(p, sz)
	return t
}

func (t *TileGroup) build(p *parser.Parser, sz int) {
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
		t.decodeTile(p)
	}
}

// 8.2.4 Exit process for symbol decoder
func (t *TileGroup) exitSymbol(p *parser.Parser) {
	if p.SymbolMaxBits < -14 {
		panic("Violating bitstream conformance!")
	}

	p.position += Max(0, p.SymbolMaxBits)

	if !p.uncompressedHeader.DisableFrameEndUpdateCdf && p.TileNum == p.uncompressedHeader.TileInfo.ContextUpdateTileId {
		// TODO: whatever is supposed to happen ehre
	}
}

// decode_tile()
func (t *TileGroup) decodeTile(p *parser.Parser) {
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

// inter_frame_mode_info()
func (t *TileGroup) interFrameModeInfo(p *parser.Parser) {
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
func (t *TileGroup) intraBlockModeInfo(p *parser.Parser) {
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
func (t *TileGroup) interBlockModeInfo(p *parser.Parser) {
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
func (t *TileGroup) needsInterpFilter(p *parser.Parser) bool {
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
func (t *TileGroup) readCompoundType(isCompound bool, p *parser.Parser) {
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
func (t *TileGroup) readMotionMode(isCompound bool, p *parser.Parser) {
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
func (t *TileGroup) isScaled(refFrame int) bool {
	refIdx := p.uncompressedHeader.ref_frame_idx[refFrame-LAST_FRAME]
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameWidth / 2)) / p.uncompressedHeader.FrameWidth
	yScale := ((t.RefUpscaledHeight[refIdx] << REF_SCALE_SHIFT) + (p.uncompressedHeader.FrameHeight / 2)) / p.uncompressedHeader.FrameHeight
	noScale := 1 << REF_SCALE_SHIFT

	return xScale != noScale || yScale != noScale
}

// find_warp_samples() 7.10.4.
func (t *TileGroup) findWarpSamples(p *parser.Parser) {
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
func (t *TileGroup) addSample(deltaRow int, deltaCol int, p *parser.Parser) {
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
func (t *TileGroup) readInterIntraMode(isCompound bool, p *parser.Parser) {
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
func (t *TileGroup) readRefFrames(p *parser.Parser) {
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
func (t *TileGroup) readIsInter(p *parser.Parser) {
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
func (t *TileGroup) interSegmentId(preSkip int, p *parser.Parser) {
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
func (t *TileGroup) readSkipMode(p *parser.Parser) {
	if t.segFeatureActive(SEG_LVL_SKIP, p) || t.segFeatureActive(SEG_LVL_REF_FRAME, p) || t.segFeatureActive(SEG_LVL_GLOBALMV, p) || !Bool(p.uncompressedHeader.SkipModePresent) || t.Block_Width[p.MiSize] < 8 || t.Block_Height[p.MiSize] < 8 {
		t.SkipMode = 0
	} else {
		t.SkipMode = p.S()
	}
}

// get_segment_id( )
func (t *TileGroup) getSegmentId(p *parser.Parser) int {
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

// is_directional_mode( mode )
func (t *TileGroup) isDirectionalMode(mode int) bool {
	return (mode >= V_PRED) && (mode <= D67_PRED)
}

// clear_block_decoded_flags( r, c, sbSize4 )
func (t *TileGroup) clearBlockDecodedFlags(r int, c int, sbSize4 int, p *parser.Parser) {
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
