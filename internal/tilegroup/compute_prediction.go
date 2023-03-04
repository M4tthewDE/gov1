package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
)

// compute_prediction()
func (t *TileGroup) computePrediction() {
	sbMask := 15
	if t.State.SequenceHeader.Use128x128SuperBlock {
		sbMask = 31
	}

	subBlockMiRow := t.State.MiRow & sbMask
	subBlockMiCol := t.State.MiCol & sbMask

	for plane := 0; plane < 1+util.Int(t.HasChroma)*2; plane++ {
		planeSz := t.getPlaneResidualSize(t.State.MiSize, plane)
		num4x4W := t.State.Num4x4BlocksWide[planeSz]
		num4x4H := t.State.Num4x4BlocksHigh[planeSz]
		log2W := shared.MI_SIZE_LOG2 + shared.MI_WIDTH_LOG2[planeSz]
		log2H := shared.MI_SIZE_LOG2 + shared.MI_HEIGHT_LOG2[planeSz]
		subX := 0
		subY := 0
		if plane > 0 {
			subX = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingX)
			subY = util.Int(t.State.SequenceHeader.ColorConfig.SubsamplingY)
		}
		baseX := (t.State.MiCol >> subX) * MI_SIZE
		baseY := (t.State.MiRow >> subY) * MI_SIZE
		candRow := (t.State.MiRow >> subY) << subY
		candCol := (t.State.MiCol >> subX) << subX

		t.IsInterIntra = (util.Bool(t.IsInter) && t.State.RefFrame[1] == shared.INTRA_FRAME)

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
			haveLeft := t.State.AvailLChroma
			haveAbove := t.State.AvailUChroma
			if plane == 0 {
				haveLeft = t.State.AvailL
				haveAbove = t.State.AvailU
			}
			t.predictIntra(plane, baseX, baseY, haveLeft, haveAbove, t.State.BlockDecoded[plane][(subBlockMiRow>>subY)-1][(subBlockMiCol>>subX)+num4x4W], t.State.BlockDecoded[plane][(subBlockMiRow>>subY)+num4x4H][(subBlockMiCol>>subX)-1], mode, log2W, log2H)
		}

		if util.Bool(t.IsInter) {
			predW := t.Block_Width[t.State.MiSize] >> subX
			predH := t.Block_Height[t.State.MiSize] >> subY
			someUseIntra := false

			for r := 0; r < (num4x4H << subY); r++ {
				for c := 0; c < (num4x4W << subX); c++ {
					if t.State.RefFrames[candRow+r][candCol+c][0] == shared.INTRA_FRAME {
						someUseIntra = true
					}
				}
			}

			if someUseIntra {
				predW = num4x4W * 4
				predH = num4x4H * 4
				candRow = t.State.MiRow
				candCol = t.State.MiCol
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
