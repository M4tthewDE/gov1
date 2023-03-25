package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.18 Output process
func (t *TileGroup) outputProcess(state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	if state.OperatingPointIdc != 0 {
		panic("not implemented")
	}

	w, h, subX, subY := t.intermediateOutputPreparation(uh, state, &sh)

	if sh.FilmGrainParamsPresent && uh.ApplyGrain {
		t.filmGrainSynthesis(w, h, subX, subY)
	}
}

// 7.18.2. Intermediate output preparation process
func (t *TileGroup) intermediateOutputPreparation(uh uncompressedheader.UncompressedHeader, state *state.State, sh *sequenceheader.SequenceHeader) (int, int, int, int) {
	var w, h, subX, subY int
	if uh.ShowExistingFrame {
		w := t.RefUpscaledWidth[uh.FrameToShowMapIdx]
		h := t.RefFrameHeight[uh.FrameToShowMapIdx]
		subX := state.RefSubsamplingX[uh.FrameToShowMapIdx]
		subY := state.RefSubsamplingY[uh.FrameToShowMapIdx]

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				t.OutY[y][x] = t.FrameStore[uh.FrameToShowMapIdx][0][y][x]
			}
		}

		for x := 0; x < (w+subX)>>subX; x++ {
			for y := 0; y < (h+subY)>>subY; y++ {
				t.OutU[y][x] = t.FrameStore[uh.FrameToShowMapIdx][1][x][y]
			}
		}

		for x := 0; x < (w+subX)>>subX; x++ {
			for y := 0; y < (h+subY)>>subY; y++ {
				t.OutV[y][x] = t.FrameStore[uh.FrameToShowMapIdx][2][x][y]
			}
		}

		// TODO: does this assignment work? normall sequenceheader doesn't get modified
		sh.ColorConfig.BitDepth = state.RefBitDepth[uh.FrameToShowMapIdx]
	} else {
		w := uh.UpscaledWidth
		h := uh.FrameHeight
		subX := util.Int(sh.ColorConfig.SubsamplingX)
		subY := util.Int(sh.ColorConfig.SubsamplingY)

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				t.OutY[y][x] = state.LrFrame[0][y][x]
			}
		}

		for x := 0; x < (w+subX)>>subX; x++ {
			for y := 0; y < (h+subY)>>subY; y++ {
				t.OutU[y][x] = state.LrFrame[1][y][x]
			}
		}

		for x := 0; x < (w+subX)>>subX; x++ {
			for y := 0; y < (h+subY)>>subY; y++ {
				t.OutV[y][x] = state.LrFrame[2][y][x]
			}
		}
	}

	return w, h, subX, subY
}

// 7.18.3. Film grain syntehsis
func (t *TileGroup) filmGrainSynthesis(w int, h int, subX int, subY int) {
	panic("not implemented")
}
