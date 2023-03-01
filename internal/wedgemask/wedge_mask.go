package wedgemask

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
)

const MASK_MASTER_SIZE = 64
const WEDGE_TYPES = 16

const WEDGE_HORIZONTAL = 0
const WEDGE_VERTICAL = 1
const WEDGE_OBLIQUE27 = 2
const WEDGE_OBLIQUE63 = 3
const WEDGE_OBLIQUE117 = 4
const WEDGE_OBLIQUE153 = 5

var Wedge_Bits = []int{0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 0, 0}

var Wedge_Master_Oblique_Odd = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 6, 18,
	37, 53, 60, 63, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
}

var Wedge_Master_Oblique_Even = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 4, 11, 27,
	46, 58, 62, 63, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
}

var Wedge_Master_Vertical = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 7, 21,
	43, 57, 62, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
}

var Wedge_Codebook = [][][]int{
	{
		{WEDGE_OBLIQUE27, 4, 4}, {WEDGE_OBLIQUE63, 4, 4},
		{WEDGE_OBLIQUE117, 4, 4}, {WEDGE_OBLIQUE153, 4, 4},
		{WEDGE_HORIZONTAL, 4, 2}, {WEDGE_HORIZONTAL, 4, 4},
		{WEDGE_HORIZONTAL, 4, 6}, {WEDGE_VERTICAL, 4, 4},
		{WEDGE_OBLIQUE27, 4, 2}, {WEDGE_OBLIQUE27, 4, 6},
		{WEDGE_OBLIQUE153, 4, 2}, {WEDGE_OBLIQUE153, 4, 6},
		{WEDGE_OBLIQUE63, 2, 4}, {WEDGE_OBLIQUE63, 6, 4},
		{WEDGE_OBLIQUE117, 2, 4}, {WEDGE_OBLIQUE117, 6, 4},
	},
	{
		{WEDGE_OBLIQUE27, 4, 4}, {WEDGE_OBLIQUE63, 4, 4},
		{WEDGE_OBLIQUE117, 4, 4}, {WEDGE_OBLIQUE153, 4, 4},
		{WEDGE_VERTICAL, 2, 4}, {WEDGE_VERTICAL, 4, 4},
		{WEDGE_VERTICAL, 6, 4}, {WEDGE_HORIZONTAL, 4, 4},
		{WEDGE_OBLIQUE27, 4, 2}, {WEDGE_OBLIQUE27, 4, 6},
		{WEDGE_OBLIQUE153, 4, 2}, {WEDGE_OBLIQUE153, 4, 6},
		{WEDGE_OBLIQUE63, 2, 4}, {WEDGE_OBLIQUE63, 6, 4},
		{WEDGE_OBLIQUE117, 2, 4}, {WEDGE_OBLIQUE117, 6, 4},
	},
	{
		{WEDGE_OBLIQUE27, 4, 4}, {WEDGE_OBLIQUE63, 4, 4},
		{WEDGE_OBLIQUE117, 4, 4}, {WEDGE_OBLIQUE153, 4, 4},
		{WEDGE_HORIZONTAL, 4, 2}, {WEDGE_HORIZONTAL, 4, 6},
		{WEDGE_VERTICAL, 2, 4}, {WEDGE_VERTICAL, 6, 4},
		{WEDGE_OBLIQUE27, 4, 2}, {WEDGE_OBLIQUE27, 4, 6},
		{WEDGE_OBLIQUE153, 4, 2}, {WEDGE_OBLIQUE153, 4, 6},
		{WEDGE_OBLIQUE63, 2, 4}, {WEDGE_OBLIQUE63, 6, 4},
		{WEDGE_OBLIQUE117, 2, 4}, {WEDGE_OBLIQUE117, 6, 4},
	},
}

var WedgeMasks = [][][][][]int{}

func InitialiseWedgeMaskTable(state State) {
	w := MASK_MASTER_SIZE
	h := MASK_MASTER_SIZE

	var MasterMask [][][]int
	for j := 0; j < w; j++ {
		shift := MASK_MASTER_SIZE / 4
		for i := 0; i < h; i += 2 {
			MasterMask[WEDGE_OBLIQUE63][i][j] = Wedge_Master_Oblique_Even[util.Clip3(0, MASK_MASTER_SIZE-1, j-shift)]
			shift -= 1
			MasterMask[WEDGE_OBLIQUE63][i+1][j] = Wedge_Master_Oblique_Odd[util.Clip3(0, MASK_MASTER_SIZE-1, j-shift)]
			MasterMask[WEDGE_VERTICAL][i][j] = Wedge_Master_Vertical[j]
			MasterMask[WEDGE_VERTICAL][i+1][j] = Wedge_Master_Vertical[j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			msk := MasterMask[WEDGE_OBLIQUE63][i][j]
			MasterMask[WEDGE_OBLIQUE27][j][i] = msk
			MasterMask[WEDGE_OBLIQUE117][i][w-1-j] = 64 - msk
			MasterMask[WEDGE_OBLIQUE153][w-1-j][i] = 64 - msk
			MasterMask[WEDGE_HORIZONTAL][j][i] = MasterMask[WEDGE_VERTICAL][i][j]
		}
	}

	for bsize := shared.BLOCK_8X8; bsize < shared.BLOCK_SIZES; bsize++ {
		if Wedge_Bits[bsize] > 0 {
			w := state.BlockWidth[bsize]
			h := state.BlockHeight[bsize]

			for wedge := 0; wedge < WEDGE_TYPES; wedge++ {
				dir := getWedgeDirection(bsize, wedge, state)
				xoff := MASK_MASTER_SIZE/2 - ((getWedgeXoff(bsize, wedge, state) * w) >> 3)
				yoff := MASK_MASTER_SIZE/2 - ((getWedgeYoff(bsize, wedge, state) * h) >> 3)
				sum := 0
				for i := 0; i < w; i++ {
					sum += MasterMask[dir][yoff][xoff+i]
				}
				for i := 0; i < h; i++ {
					sum += MasterMask[dir][yoff+i][xoff]
				}
				avg := (sum + (w+h-1)/2) / (w + h - 1)
				flipSign := avg < 32
				for i := 0; i < h; i++ {
					for j := 0; j < w; j++ {
						WedgeMasks[bsize][util.Int(flipSign)][wedge][i][j] = MasterMask[dir][yoff+i][xoff+j]
						WedgeMasks[bsize][util.Int(!flipSign)][wedge][i][j] = 64 - MasterMask[dir][yoff+i][xoff+j]
					}
				}
			}
		}
	}
}

func getWedgeDirection(bsize int, index int, state State) int {
	return Wedge_Codebook[blockShape(bsize, state)][index][0]
}

func getWedgeXoff(bsize int, index int, state State) int {
	return Wedge_Codebook[blockShape(bsize, state)][index][1]
}

func getWedgeYoff(bsize int, index int, state State) int {
	return Wedge_Codebook[blockShape(bsize, state)][index][2]
}

func blockShape(bsize int, state State) int {
	w4 := state.Num4x4BlocksWide[bsize]
	h4 := state.Num4x4BlocksHigh[bsize]
	if h4 > w4 {
		return 0
	} else if h4 < w4 {
		return 1
	}
	return 2
}
