package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.15. CDEF process
func (t *TileGroup) cdefProcess(state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	step4 := shared.NUM_4X4_BLOCKS_WIDE[shared.BLOCK_8X8]
	cdefSize4 := shared.NUM_4X4_BLOCKS_WIDE[shared.BLOCK_64X64]
	cdefMask4 := ^(cdefSize4 - 1)

	for r := 0; r < state.MiRows; r += step4 {
		for c := 0; c < state.MiCols; c += step4 {
			baseR := r & cdefMask4
			baseC := c & cdefMask4
			idx := state.Cdef.CdefIdx[baseR][baseC]
			t.cdefBlock(r, c, idx, state, sh, uh)
		}
	}
}

// 7.15.1. CDEF block process
func (t *TileGroup) cdefBlock(r int, c int, idx int, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	startY := r * MI_SIZE
	endY := startY + MI_SIZE*2
	startX := c * MI_SIZE
	endX := startX + MI_SIZE*2
	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			state.CdefFrame[0][y][x] = state.CurrFrame[0][y][x]
		}
	}
	if sh.ColorConfig.NumPlanes > 1 {
		startY >>= util.Int(sh.ColorConfig.SubsamplingY)
		endY >>= util.Int(sh.ColorConfig.SubsamplingY)
		startX >>= util.Int(sh.ColorConfig.SubsamplingX)
		endX >>= util.Int(sh.ColorConfig.SubsamplingX)
		for y := startY; y < endY; y++ {
			for x := startX; x < endX; x++ {
				state.CdefFrame[1][y][x] = state.CurrFrame[1][y][x]
				state.CdefFrame[2][y][x] = state.CurrFrame[2][y][x]
			}
		}
	}

	if idx == -1 {
		return
	}

	coeffShift := sh.ColorConfig.BitDepth - 8
	skip := util.Bool(t.Skips[r][c]) &&
		util.Bool(t.Skips[r+1][c]) &&
		util.Bool(t.Skips[r][c+1]) &&
		util.Bool(t.Skips[r+1][c+1])

	if !skip {
		yDir, varr := t.cdefDirection(r, c, state, sh)
		priStr := uh.CdefYPriStrength[idx] << coeffShift
		secStr := uh.CdefYSecStrength[idx] << coeffShift

		var dir int
		if priStr == 0 {
			dir = 0
		} else {
			dir = yDir
		}

		var varStr int
		if util.Bool(varr >> 6) {
			varStr = util.Min(util.FloorLog2(varr>>6), 12)
		} else {
			varStr = 0
		}

		if util.Bool(varr) {
			priStr = (priStr*(4+varStr) + 8) >> 4
		} else {
			priStr = 0
		}

		damping := uh.CdefDampening + coeffShift
		t.cdefFilter(0, r, c, priStr, secStr, damping, dir, state, sh, uh)

		if sh.ColorConfig.NumPlanes == 1 {
			return
		}

		priStr = uh.CdefUVPriStrength[idx] << coeffShift
		secStr = uh.CdefUVSecStrength[idx] << coeffShift

		if priStr == 0 {
			dir = 0
		} else {
			dir = CDEF_UV_DIR[util.Int(sh.ColorConfig.SubsamplingX)][util.Int(sh.ColorConfig.SubsamplingY)][yDir]
			damping = uh.CdefDampening + coeffShift - 1
		}

		t.cdefFilter(1, r, c, priStr, secStr, damping, dir, state, sh, uh)
		t.cdefFilter(2, r, c, priStr, secStr, damping, dir, state, sh, uh)
	}
}

var CDEF_UV_DIR = [2][2][8]int{
	{{0, 1, 2, 3, 4, 5, 6, 7},
		{1, 2, 2, 2, 3, 4, 6, 0}},
	{{7, 0, 2, 4, 5, 6, 6, 6},
		{0, 1, 2, 3, 4, 5, 6, 7}},
}

// 7.15.2. CDEF direction process
func (t *TileGroup) cdefDirection(r int, c int, state *state.State, sh sequenceheader.SequenceHeader) (int, int) {
	var cost [8]int
	var partial [8][15]int
	for i := 0; i < 8; i++ {
		cost[i] = 0
		for j := 0; j < 15; j++ {
			partial[i][j] = 0
		}
	}

	bestCost := 0
	yDir := 0
	x0 := c << shared.MI_SIZE_LOG2
	y0 := r << shared.MI_SIZE_LOG2
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x := (state.CurrFrame[0][y0+i][x0+j] >> (sh.ColorConfig.BitDepth - 8)) - 128
			partial[0][i+j] += x
			partial[1][i+j/2] += x
			partial[2][i] += x
			partial[3][3+i-j/2] += x
			partial[4][7+i-j] += x
			partial[5][3-i/2+j] += x
			partial[6][j] += x
			partial[7][i/2+j] += x
		}
	}
	for i := 0; i < 8; i++ {
		cost[2] += partial[2][i] * partial[2][i]
		cost[6] += partial[6][i] * partial[6][i]
	}
	cost[2] *= DIV_TABLE[8]
	cost[6] *= DIV_TABLE[8]
	for i := 0; i < 7; i++ {
		cost[0] += (partial[0][i]*partial[0][i] +
			partial[0][14-i]*partial[0][14-i]) *
			DIV_TABLE[i+1]
		cost[4] += (partial[4][i]*partial[4][i] +
			partial[4][14-i]*partial[4][14-i]) *
			DIV_TABLE[i+1]
	}
	cost[0] += partial[0][7] * partial[0][7] * DIV_TABLE[8]
	cost[4] += partial[4][7] * partial[4][7] * DIV_TABLE[8]
	for i := 1; i < 8; i += 2 {
		for j := 0; j < 4+1; j++ {
			cost[i] += partial[i][3+j] * partial[i][3+j]
		}
		cost[i] *= DIV_TABLE[8]
		for j := 0; j < 4-1; j++ {
			cost[i] += (partial[i][j]*partial[i][j] +
				partial[i][10-j]*partial[i][10-j]) *
				DIV_TABLE[2*j+2]
		}
	}
	for i := 0; i < 8; i++ {
		if cost[i] > bestCost {
			bestCost = cost[i]
			yDir = i
		}
	}

	return yDir, (bestCost - cost[(yDir+4)&7]) >> 10
}

var DIV_TABLE = [9]int{
	0, 840, 420, 280, 210, 168, 140, 120, 105,
}

// 7.15.3. CDEF filter process
func (t *TileGroup) cdefFilter(plane int, r int, c int, priStr int, secStr int, damping int, dir int, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	// what does this mean?
	// "MiColStart, MiRowStart, MiColEnd, MiRowEnd are set equal to the values
	// they had when the syntax element MiSizes[r][c] was written."

	coeffShift := sh.ColorConfig.BitDepth
	var subX, subY int
	if plane > 0 {
		subX = util.Int(sh.ColorConfig.SubsamplingX)
		subY = util.Int(sh.ColorConfig.SubsamplingY)
	} else {
		subX = 0
		subY = 0
	}
	x0 := (c * MI_SIZE) >> subX
	y0 := (r * MI_SIZE) >> subY
	w := 8 >> subX
	h := 8 >> subY
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			sum := 0
			x := state.CurrFrame[plane][y0+i][x0+j]
			max := x
			min := x
			for k := 0; k < 2; k++ {
				for sign := -1; sign <= 1; sign += 2 {
					p := t.cdefGetAt(plane, x0, y0, i, j, dir, k, sign, subX, subY, state)
					if t.CdefAvailable {
						sum += CDEF_PRI_TAPS[(priStr>>coeffShift)&1][k] * constrain(p-x, priStr,
							damping)
						max = util.Max(p, max)
						min = util.Min(p, min)
					}
					for dirOff := -2; dirOff <= 2; dirOff += 4 {
						s := t.cdefGetAt(plane, x0, y0, i, j, (dir+dirOff)&7, k, sign, subX, subY, state)
						if t.CdefAvailable {
							sum += CDEF_SEC_TAPS[(priStr>>coeffShift)&1][k] * constrain(s-x,
								secStr, damping)
							max = util.Max(s, max)
							min = util.Min(s, min)
						}
					}
				}
			}
			state.CdefFrame[plane][y0+i][x0+j] = util.Clip3(min, max, x+((8+sum-util.Int(sum < 0))>>4))
		}
	}
}

func (t *TileGroup) cdefGetAt(plane int, x0 int, y0 int, i int, j int, dir int, k int, sign int, subX int, subY int, state *state.State) int {
	y := y0 + i + sign*CDEF_DIRECTIONS[dir][k][0]
	x := x0 + j + sign*CDEF_DIRECTIONS[dir][k][1]
	candidateR := (y << subY) >> shared.MI_SIZE_LOG2
	candidateC := (x << subX) >> shared.MI_SIZE_LOG2
	if isInsideFilterRegion(candidateR, candidateC, state) {
		t.CdefAvailable = true
		return state.CurrFrame[plane][y][x]
	} else {
		t.CdefAvailable = false
		return 0
	}
}

var CDEF_DIRECTIONS = [8][2][2]int{
	{{-1, 1}, {-2, 2}},
	{{0, 1}, {-1, 2}},
	{{0, 1}, {0, 2}},
	{{0, 1}, {1, 2}},
	{{1, 1}, {2, 2}},
	{{1, 0}, {2, 1}},
	{{1, 0}, {2, 0}},
	{{1, 0}, {2, -1}},
}

func constrain(diff int, threshold int, damping int) int {
	if !util.Bool(threshold) {
		return 0
	}
	dampingAdj := util.Max(0, damping-util.FloorLog2(threshold))
	sign := 1
	if diff < 0 {
		sign = -1
	}
	return sign * util.Clip3(0, util.Abs(diff), threshold-(util.Abs(diff)>>dampingAdj))
}

var CDEF_PRI_TAPS = [2][2]int{
	{4, 2}, {3, 3},
}
var CDEF_SEC_TAPS = [2][2]int{
	{2, 1}, {2, 1},
}
