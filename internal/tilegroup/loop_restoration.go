package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.17. Loop restoration process
func (t *TileGroup) loopRestoration(state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	state.LrFrame = state.UpscaledCdefFrame

	if !uh.UsesLr {
		return
	}

	for y := 0; y < uh.FrameHeight; y += MI_SIZE {
		for x := 0; x < uh.UpscaledWidth; x += MI_SIZE {
			for plane := 0; plane < sh.ColorConfig.NumPlanes; plane++ {
				if uh.FrameRestorationType[plane] != shared.RESTORE_NONE {
					row := y >> shared.MI_SIZE_LOG2
					col := x >> shared.MI_SIZE_LOG2
					t.loopRestoreBlock(plane, row, col, state, sh, uh)
				}
			}
		}
	}
}

// 7.17.1. Loop restore block process
func (t *TileGroup) loopRestoreBlock(plane int, row int, col int, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	lumaY := row * MI_SIZE
	stripeNum := (lumaY + 8) / 64

	var subX, subY int
	if plane == 0 {
		subX = 0
		subY = 0
	} else {
		subX = util.Int(sh.ColorConfig.SubsamplingX)
		subY = util.Int(sh.ColorConfig.SubsamplingY)
	}

	t.StripeStartY = ((-8 + stripeNum*64) >> subY)
	t.StripeEndY = t.StripeStartY + (64 >> subY) - 1

	unitSize := state.LoopRestorationSize[plane]
	unitRows := util.CountUnitsInFrame(unitSize, util.Round2(uh.FrameHeight, subY))
	unitCols := util.CountUnitsInFrame(unitSize, util.Round2(uh.UpscaledWidth, subX))

	unitRow := util.Min(unitRows-1, ((row*MI_SIZE+8)>>subY)/unitSize)
	unitCol := util.Min(unitCols-1, ((col*MI_SIZE+8)>>subX)/unitSize)

	t.PlaneEndX = util.Round2(uh.UpscaledWidth, subX) - 1
	t.PlaneEndY = util.Round2(uh.FrameHeight, subY) - 1

	x := col * MI_SIZE >> subX
	y := row * MI_SIZE >> subY
	w := util.Min(MI_SIZE>>subX, t.PlaneEndX-x+1)
	h := util.Min(MI_SIZE>>subY, t.PlaneEndY-y+1)
	rType := t.LrType[plane][unitRow][unitCol]

	if rType == shared.RESTORE_WIENER {
		t.wienerFilter(plane, unitRow, unitCol, x, y, w, h, state, sh)
	} else if rType == shared.RESTORE_SGRPROJ {
		t.selfGuidedFilter(plane, unitRow, unitCol, x, y, w, h, state, sh)
	}

}

// 7.17.2. Self guided filter process
func (t *TileGroup) selfGuidedFilter(plane int, unitRow int, unitCol int, x int, y int, w int, h int, state *state.State, sh sequenceheader.SequenceHeader) {
	set := t.LrSgrSet[plane][unitRow][unitCol]
	pass := false
	flt0 := t.boxFilter(plane, x, y, w, h, set, pass, state, sh)
	pass = true
	flt1 := t.boxFilter(plane, x, y, w, h, set, pass, state, sh)

	w0 := t.LrSgrXqd[plane][unitRow][unitCol][0]
	w1 := t.LrSgrXqd[plane][unitRow][unitCol][1]
	w2 := (1 << SGRPROJ_PRJ_BITS) - w0 - w1
	r0 := util.Bool(SGR_PARAMS[set][0])
	r1 := util.Bool(SGR_PARAMS[set][2])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			u := state.UpscaledCdefFrame[plane][y+i][x+j] << SGRPROJ_RST_BITS
			v := w1 * u
			if r0 {
				v += w0 * flt0[i][j]
			} else {
				v += w0 * u
			}
			if r1 {
				v += w2 * flt1[i][j]
			} else {
				v += w2 * u
			}
			s := util.Round2(v, SGRPROJ_RST_BITS+SGRPROJ_PRJ_BITS)
			state.LrFrame[plane][y+i][x+j] = util.Clip1(s, sh.ColorConfig.BitDepth)
		}
	}
}

// 7.17.3. Box filter process
func (t *TileGroup) boxFilter(plane int, x int, y int, w int, h int, set int, pass bool, state *state.State, sh sequenceheader.SequenceHeader) [1][1]int {

	r := SGR_PARAMS[set][util.Int(pass)*2+0]
	if r == 0 {
		return [1][1]int{}
	}

	eps := SGR_PARAMS[set][util.Int(pass)*2+1]

	n := (2*r + 1) * (2*r + 1)
	n2e := n * n * eps
	s := (((1 << SGRPROJ_MTABLE_BITS) + n2e/2) / n2e)

	var A, B [3][3]int
	for i := -1; i < h+1; i++ {
		for j := -1; j < w+1; j++ {
			a := 0
			b := 0
			for dy := -r; dy <= r; dy++ {
				for dx := -r; dx <= r; dx++ {
					c := t.getSourceSample(plane, x+j+dx, y+i+dy, state)
					a += c * c
					b += c
				}
			}
			a = util.Round2(a, 2*(sh.ColorConfig.BitDepth-8))
			d := util.Round2(b, sh.ColorConfig.BitDepth-8)
			p := util.Max(0, a*n-d*d)
			z := util.Round2(p*s, SGRPROJ_MTABLE_BITS)
			var a2 int
			if z >= 255 {
				a2 = 256
			} else if z == 0 {
				a2 = 1
			} else {
				a2 = ((z << SGRPROJ_SGR_BITS) + (z / 2)) / (z + 1)
			}
			oneOverN := ((1 << SGRPROJ_RECIP_BITS) + (n / 2)) / n
			b2 := ((1 << SGRPROJ_SGR_BITS) - a2) * b * oneOverN

			// ugly negative indexing workaround
			if i == -1 && j == -1 {
				A[len(A)-1][len(A[len(A)-1])-1] = a2
				B[len(B)-1][len(B[len(B)-1])-1] = util.Round2(b2, SGRPROJ_RECIP_BITS)
			} else if i == -1 {
				A[len(A)-1][j] = a2
				B[len(B)-1][j] = util.Round2(b2, SGRPROJ_RECIP_BITS)
			} else if j == -1 {
				A[i][len(A[i])-1] = a2
				B[i][len(B[i])-1] = util.Round2(b2, SGRPROJ_RECIP_BITS)

			} else {
				A[i][j] = a2
				B[i][j] = util.Round2(b2, SGRPROJ_RECIP_BITS)
			}
		}
	}

	var F [1][1]int
	for i := 0; i < h; i++ {
		shift := 5
		if !pass && util.Bool(i&1) {
			shift = 4
		}
		for j := 0; j < w; j++ {
			a := 0
			b := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					var weight int
					if !pass {
						if util.Bool((i + dy) & 1) {
							if dx == 0 {
								weight = 6
							} else {
								weight = 5
							}
						} else {
							weight = 0
						}
					} else {
						if dx == 0 || dy == 0 {
							weight = 4
						} else {
							weight = 3
						}
					}

					// ugly negative indexing workaround
					if i+dy == -1 && j+dx == -1 {
						a += weight * A[len(A)-1][len(A[len(A)-1])-1]
						b += weight * B[len(B)-1][len(B[len(B)-1])-1]
					} else if i+dy == -1 {
						a += weight * A[len(A)-1][j+dx]
						b += weight * B[len(B)-1][j+dx]
					} else if j+dx == -1 {
						a += weight * A[i+dy][len(A[len(A)-1])-1]
						b += weight * B[i+dy][len(B[len(B)-1])-1]
					} else {
						a += weight * A[i+dy][j+dx]
						b += weight * B[i+dy][j+dx]

					}
				}
			}
			v := a*state.UpscaledCdefFrame[plane][y+i][x+j] + b
			F[i][j] = util.Round2(v, SGRPROJ_SGR_BITS+shift-SGRPROJ_RST_BITS)
		}
	}

	return F
}

const SGRPROJ_PRJ_BITS = 7
const SGRPROJ_MTABLE_BITS = 20
const SGRPROJ_SGR_BITS = 8
const SGRPROJ_RECIP_BITS = 12
const SGRPROJ_RST_BITS = 4

// 7.7.4 Wiener filter process
func (tg *TileGroup) wienerFilter(plane int, unitRow int, unitCol int, x int, y int, w int, h int, state *state.State, sh sequenceheader.SequenceHeader) {
	tg.roundVariablesDerivationProcess(false, sh)

	vfilter := tg.wienerCoefficient(tg.LrWiener[plane][unitRow][unitCol][0])
	hfilter := tg.wienerCoefficient(tg.LrWiener[plane][unitRow][unitCol][1])

	offset := (1 << (sh.ColorConfig.BitDepth + FILTER_BITS - tg.InterRound0 - 1))
	limit := (1 << (sh.ColorConfig.BitDepth + 1 + FILTER_BITS - tg.InterRound0)) - 1

	var intermediate [][]int
	for r := 0; r < h+6; r++ {
		for c := 0; c < w; c++ {
			s := 0
			for t := 0; t < 7; t++ {
				s += hfilter[t] * tg.getSourceSample(plane, x+c+t-3, y+r-3, state)
			}
			v := util.Round2(s, tg.InterRound0)
			intermediate[r][c] = util.Clip3(-offset, limit-offset, v)
		}
	}

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			s := 0
			for t := 0; t < 7; t++ {
				s += vfilter[t] * intermediate[r+t][c]
			}
			v := util.Round2(s, tg.InterRound1)
			state.LrFrame[plane][y+r][x+c] = util.Clip1(v, sh.ColorConfig.BitDepth)
		}
	}
}

func (t *TileGroup) wienerCoefficient(coeff [3]int) [7]int {
	var filter [7]int
	filter[3] = 128
	for i := 0; i < 3; i++ {
		c := coeff[i]
		filter[i] = c
		filter[6-1] = c
		filter[3] -= 2 * c
	}

	return filter
}

// 7.17.6. Get source sample process
func (t *TileGroup) getSourceSample(plane int, x int, y int, state *state.State) int {
	x = util.Min(t.PlaneEndX, x)
	x = util.Max(0, x)
	y = util.Min(t.PlaneEndY, y)
	y = util.Max(0, y)
	if y < t.StripeStartY {
		y = util.Max(t.StripeStartY-2, y)
		return state.UpscaledCurrFrame[plane][y][x]
	} else if y > t.StripeEndY {
		y = util.Min(t.StripeEndY+2, y)
		return state.UpscaledCurrFrame[plane][y][x]
	} else {
		return state.UpscaledCdefFrame[plane][y][x]
	}
}
