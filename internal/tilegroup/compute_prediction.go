package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/util"
)

// compute_prediction()
func (t *TileGroup) computePrediction(state *state.State, sh sequenceheader.SequenceHeader) {
	sbMask := 15
	if sh.Use128x128SuperBlock {
		sbMask = 31
	}

	subBlockMiRow := state.MiRow & sbMask
	subBlockMiCol := state.MiCol & sbMask

	for plane := 0; plane < 1+util.Int(t.HasChroma)*2; plane++ {
		planeSz := t.getPlaneResidualSize(state.MiSize, plane)
		num4x4W := state.Num4x4BlocksWide[planeSz]
		num4x4H := state.Num4x4BlocksHigh[planeSz]
		log2W := shared.MI_SIZE_LOG2 + shared.MI_WIDTH_LOG2[planeSz]
		log2H := shared.MI_SIZE_LOG2 + shared.MI_HEIGHT_LOG2[planeSz]
		subX := 0
		subY := 0
		if plane > 0 {
			subX = util.Int(sh.ColorConfig.SubsamplingX)
			subY = util.Int(sh.ColorConfig.SubsamplingY)
		}
		baseX := (state.MiCol >> subX) * MI_SIZE
		baseY := (state.MiRow >> subY) * MI_SIZE
		candRow := (state.MiRow >> subY) << subY
		candCol := (state.MiCol >> subX) << subX

		t.IsInterIntra = (util.Bool(t.IsInter) && state.RefFrame[1] == shared.INTRA_FRAME)

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
			haveLeft := state.AvailLChroma
			haveAbove := state.AvailUChroma
			if plane == 0 {
				haveLeft = state.AvailL
				haveAbove = state.AvailU
			}
			t.predictIntra(plane, baseX, baseY, haveLeft, haveAbove, state.BlockDecoded[plane][(subBlockMiRow>>subY)-1][(subBlockMiCol>>subX)+num4x4W], state.BlockDecoded[plane][(subBlockMiRow>>subY)+num4x4H][(subBlockMiCol>>subX)-1], mode, log2W, log2H)
		}

		if util.Bool(t.IsInter) {
			predW := t.Block_Width[state.MiSize] >> subX
			predH := t.Block_Height[state.MiSize] >> subY
			someUseIntra := false

			for r := 0; r < (num4x4H << subY); r++ {
				for c := 0; c < (num4x4W << subX); c++ {
					if state.RefFrames[candRow+r][candCol+c][0] == shared.INTRA_FRAME {
						someUseIntra = true
					}
				}
			}

			if someUseIntra {
				predW = num4x4W * 4
				predH = num4x4H * 4
				candRow = state.MiRow
				candCol = state.MiCol
			}
			r := 0
			for y := 0; y < num4x4H; y += predH {
				c := 0
				for x := 0; x < num4x4W; x += predW {
					t.predictInter(plane, baseX+x, baseY+y, predW, predH, candRow+r, candCol+c)
				}
			}
		}
	}
}
