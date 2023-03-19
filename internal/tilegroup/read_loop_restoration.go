package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/literal"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// read_lr( r, c, bSize )
func (t *TileGroup) readLr(r int, c int, bSize int, b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	if uh.AllowIntraBc {
		return
	}

	w := shared.NUM_4X4_BLOCKS_WIDE[bSize]
	h := shared.NUM_4X4_BLOCKS_HIGH[bSize]

	for plane := 0; plane < sh.ColorConfig.NumPlanes; plane++ {
		if uh.FrameRestorationType[plane] != shared.RESTORE_NONE {
			subX := 0
			subY := 0

			if sh.ColorConfig.SubsamplingX {
				subX = 1
			}

			if sh.ColorConfig.SubsamplingY {
				subY = 1
			}

			unitSize := state.LoopRestorationSize[plane]
			unitRows := util.CountUnitsInFrame(unitSize, util.Round2(uh.FrameHeight, subY))
			unitCols := util.CountUnitsInFrame(unitSize, util.Round2(state.UpscaledWidth, subX))
			unitRowStart := (r*(MI_SIZE>>subY) + unitSize - 1) / unitSize
			unitRowEnd := util.Min(unitRows, ((r+h)*(MI_SIZE>>subY)+unitSize-1)/unitSize)

			var numerator int
			var denominator int
			if uh.UseSuperRes {
				numerator = (MI_SIZE >> subX) * uh.SuperResDenom
				denominator = unitSize * shared.SUPERRES_NUM
			} else {
				numerator = MI_SIZE >> subX
				denominator = unitSize
			}
			unitColStart := (c*numerator + denominator - 1) / denominator
			unitColEnd := util.Min(unitCols, ((c+w)*numerator+denominator-1)/denominator)

			for unitRow := unitRowStart; unitRow < unitRowEnd; unitRow++ {
				for unitCol := unitColStart; unitCol < unitColEnd; unitCol++ {
					t.readLrUnit(plane, unitRow, unitCol, state, b, uh)
				}
			}
		}
	}
}

// read_lr_unit(plane, unitRow, unitCol)
func (t *TileGroup) readLrUnit(plane int, unitRow int, unitCol int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) {
	var restorationType int
	if uh.FrameRestorationType[plane] == shared.RESTORE_WIENER {
		useWiener := b.S()
		restorationType = shared.RESTORE_NONE
		if useWiener == 1 {
			restorationType = shared.RESTORE_WIENER
		}
	} else if uh.FrameRestorationType[plane] == shared.RESTORE_SGRPROJ {
		useSgrproj := b.S()
		restorationType = shared.RESTORE_NONE
		if useSgrproj == 1 {
			restorationType = shared.RESTORE_SGRPROJ
		}
	} else {
		restorationType = b.S()
	}

	t.LrType[plane][unitRow][unitCol] = restorationType

	if restorationType == shared.RESTORE_WIENER {
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
				v := t.decodeSignedSubexpWithRefBool(min, max+1, k, t.RefLrWiener[plane][pass][j], state, b, uh)
				t.LrWiener[plane][unitRow][unitCol][pass][j] = v
				t.RefLrWiener[plane][pass][j] = v
			}
		}
	} else if restorationType == shared.RESTORE_SGRPROJ {
		lrSgrSet := literal.L(SGRPROJ_PARAMS_BITS, state, b, uh)
		t.LrSgrSet[plane][unitRow][unitCol] = lrSgrSet

		for i := 0; i < 2; i++ {
			radius := SgrParams[lrSgrSet][i*2]
			min := Sgrproj_Xqd_Min[i]
			max := Sgrproj_Xqd_Max[i]

			var v int
			if radius != 0 {
				v = t.decodeSignedSubexpWithRefBool(min, max+1, SGRPROJ_PRJ_SUBEXP_K, t.RefSgrXqd[plane][i], state, b, uh)
			} else {
				v = 0
				if i == 1 {
					v = util.Clip3(min, max, (1<<SGRPROJ_BITS)-t.RefSgrXqd[plane][0])
				}
			}

			t.LrSgrXqd[plane][unitRow][unitCol][i] = v
			t.RefSgrXqd[plane][i] = v
		}
	}

}

func (t *TileGroup) decodeSignedSubexpWithRefBool(low int, high int, k int, r int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	x := t.decodeUnsignedSubexpWithRefBool(high-low, k, r-low, state, b, uh)
	return x + low

}

func (t *TileGroup) decodeUnsignedSubexpWithRefBool(mx int, k int, r int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	v := t.decodeSubexpBool(mx, k, state, b, uh)
	if (r << 1) <= mx {
		return util.InverseRecenter(r, v)
	} else {
		return mx - 1 - util.InverseRecenter(mx-1-r, v)
	}
}

func (t *TileGroup) decodeSubexpBool(numSyms int, k int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	i := 0
	mk := 0
	for {
		b2 := k
		if i == 1 {
			b2 = k + i - 1
		}

		a := 1 << b2

		if numSyms <= -mk+3*a {
			subexpUnifBools := literal.L(1, state, b, uh)
			return subexpUnifBools + mk
		} else {
			subexpMoreBools := literal.L(1, state, b, uh) != 0
			if subexpMoreBools {
				i++
				mk += a
			} else {
				subexpBools := literal.L(b2, state, b, uh)
				return subexpBools + mk
			}
		}
	}
}

var Sgrproj_Xqd_Mid = []int{-32, 31}
var Sgrproj_Xqd_Min = []int{-96, -32}
var Sgrproj_Xqd_Max = []int{31, 95}
var Wiener_Taps_Mid = []int{3, -7, 15}
var Wiener_Taps_Min = []int{-5, -23, -17}
var Wiener_Taps_Max = []int{10, 8, 46}
var Wiener_Taps_K = []int{1, 2, 3}

// TODO: why is this empty?
var SgrParams = [][]int{}
