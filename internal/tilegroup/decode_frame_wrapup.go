package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// 7.16 Upscaling process
func (t *TileGroup) upscalingProcess(inputFrame [3][9][9]int, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) [3][9][9]int {
	var outputFrame [3][9][9]int
	for plane := 0; plane < sh.ColorConfig.NumPlanes; plane++ {
		var subX, subY int
		if plane > 0 {
			subX = util.Int(sh.ColorConfig.SubsamplingX)
			subY = util.Int(sh.ColorConfig.SubsamplingY)
		} else {
			subX = 0
			subY = 0
		}

		downscaledPlaneW := util.Round2(uh.FrameWidth, subX)
		upscaledPlaneW := util.Round2(uh.UpscaledWidth, subX)
		planeH := util.Round2(uh.FrameHeight, subY)
		stepX := ((downscaledPlaneW << SUPERRES_SCALE_BITS) + (upscaledPlaneW / 2)) / upscaledPlaneW
		err := (upscaledPlaneW * stepX) - (downscaledPlaneW << SUPERRES_SCALE_BITS)
		initialSubpelX :=
			(-((upscaledPlaneW-downscaledPlaneW)<<(SUPERRES_SCALE_BITS-1))+upscaledPlaneW/2)/
				upscaledPlaneW +
				(1 << (SUPERRES_EXTRA_BITS - 1)) - err/2
		initialSubpelX &= SUPERRES_SCALE_MASK
		miW := state.MiCols >> subX
		minX := 0
		maxX := miW*MI_SIZE - 1
		for y := 0; y < planeH; y++ {
			for x := 0; x < upscaledPlaneW; x++ {
				srcX := -(1 << SUPERRES_SCALE_BITS) + initialSubpelX + x*stepX
				srcXPx := (srcX >> SUPERRES_SCALE_BITS)
				srcXSubpel := (srcX & SUPERRES_SCALE_MASK) >> SUPERRES_EXTRA_BITS
				sum := 0
				for k := 0; k < SUPERRES_FILTER_TAPS; k++ {
					sampleX := util.Clip3(minX, maxX, srcXPx+(k-SUPERRES_FILTER_OFFSET))
					px := inputFrame[plane][y][sampleX]
					sum += px * UPSCALE_FILTER[srcXSubpel][k]
				}
				outputFrame[plane][y][x] = util.Clip1(util.Round2(sum, FILTER_BITS), sh.ColorConfig.BitDepth)
			}
		}
	}

	return outputFrame
}

const FILTER_BITS = 7
const SUPERRES_FILTER_BITS = 6
const SUPERRES_FILTER_TAPS = 8
const SUPERRES_FILTER_OFFSET = 3
const SUPERRES_FILTER_SHIFTS = 1 << SUPERRES_FILTER_BITS
const SUPERRES_SCALE_BITS = 14
const SUPERRES_SCALE_MASK = (1 << 14) - 1
const SUPERRES_EXTRA_BITS = 8

var UPSCALE_FILTER = [SUPERRES_FILTER_SHIFTS][SUPERRES_FILTER_TAPS]int{
	{0, 0, 0, 128, 0, 0, 0, 0}, {0, 0, -1, 128, 2, -1, 0, 0},
	{0, 1, -3, 127, 4, -2, 1, 0}, {0, 1, -4, 127, 6, -3, 1, 0},
	{0, 2, -6, 126, 8, -3, 1, 0}, {0, 2, -7, 125, 11, -4, 1, 0},
	{-1, 2, -8, 125, 13, -5, 2, 0}, {-1, 3, -9, 124, 15, -6, 2, 0},
	{-1, 3, -10, 123, 18, -6, 2, -1}, {-1, 3, -11, 122, 20, -7, 3, -1},
	{-1, 4, -12, 121, 22, -8, 3, -1}, {-1, 4, -13, 120, 25, -9, 3, -1},
	{-1, 4, -14, 118, 28, -9, 3, -1}, {-1, 4, -15, 117, 30, -10, 4, -1},
	{-1, 5, -16, 116, 32, -11, 4, -1}, {-1, 5, -16, 114, 35, -12, 4, -1},
	{-1, 5, -17, 112, 38, -12, 4, -1}, {-1, 5, -18, 111, 40, -13, 5, -1},
	{-1, 5, -18, 109, 43, -14, 5, -1}, {-1, 6, -19, 107, 45, -14, 5, -1},
	{-1, 6, -19, 105, 48, -15, 5, -1}, {-1, 6, -19, 103, 51, -16, 5, -1},
	{-1, 6, -20, 101, 53, -16, 6, -1}, {-1, 6, -20, 99, 56, -17, 6, -1},
	{-1, 6, -20, 97, 58, -17, 6, -1}, {-1, 6, -20, 95, 61, -18, 6, -1},
	{-2, 7, -20, 93, 64, -18, 6, -2}, {-2, 7, -20, 91, 66, -19, 6, -1},
	{-2, 7, -20, 88, 69, -19, 6, -1}, {-2, 7, -20, 86, 71, -19, 6, -1},
	{-2, 7, -20, 84, 74, -20, 7, -2}, {-2, 7, -20, 81, 76, -20, 7, -1},
	{-2, 7, -20, 79, 79, -20, 7, -2}, {-1, 7, -20, 76, 81, -20, 7, -2},
	{-2, 7, -20, 74, 84, -20, 7, -2}, {-1, 6, -19, 71, 86, -20, 7, -2},
	{-1, 6, -19, 69, 88, -20, 7, -2}, {-1, 6, -19, 66, 91, -20, 7, -2},
	{-2, 6, -18, 64, 93, -20, 7, -2}, {-1, 6, -18, 61, 95, -20, 6, -1},
	{-1, 6, -17, 58, 97, -20, 6, -1}, {-1, 6, -17, 56, 99, -20, 6, -1},
	{-1, 6, -16, 53, 101, -20, 6, -1}, {-1, 5, -16, 51, 103, -19, 6, -1},
	{-1, 5, -15, 48, 105, -19, 6, -1}, {-1, 5, -14, 45, 107, -19, 6, -1},
	{-1, 5, -14, 43, 109, -18, 5, -1}, {-1, 5, -13, 40, 111, -18, 5, -1},
	{-1, 4, -12, 38, 112, -17, 5, -1}, {-1, 4, -12, 35, 114, -16, 5, -1},
	{-1, 4, -11, 32, 116, -16, 5, -1}, {-1, 4, -10, 30, 117, -15, 4, -1},
	{-1, 3, -9, 28, 118, -14, 4, -1}, {-1, 3, -9, 25, 120, -13, 4, -1},
	{-1, 3, -8, 22, 121, -12, 4, -1}, {-1, 3, -7, 20, 122, -11, 3, -1},
	{-1, 2, -6, 18, 123, -10, 3, -1}, {0, 2, -6, 15, 124, -9, 3, -1},
	{0, 2, -5, 13, 125, -8, 2, -1}, {0, 1, -4, 11, 125, -7, 2, 0},
	{0, 1, -3, 8, 126, -6, 2, 0}, {0, 1, -3, 6, 127, -4, 1, 0},
	{0, 1, -2, 4, 127, -3, 1, 0}, {0, 0, -1, 2, 128, -1, 0, 0},
}

// 7.19
func (t *TileGroup) motionVectorStorageProcess() {
	panic("not implemented")
}

// 7.20 Reference frame update process
func (t *TileGroup) referenceFrameUpdate(state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	for i := 0; i < shared.NUM_REF_FRAMES; i++ {
		if ((uh.RefreshFrameFlags >> i) & 1) == 1 {
			state.RefValid[i] = 1
			// might be concerning that this is not stored in the state
			t.RefUpscaledWidth[i] = uh.UpscaledWidth
			state.RefFrameWidth[i] = uh.FrameWidth
			state.RefFrameHeight[i] = uh.FrameHeight
			state.RefRenderWidth[i] = uh.RenderWidth
			state.RefRenderHeight[i] = uh.RenderHeight
			state.RefMiCols[i] = state.MiCols
			state.RefMiRows[i] = state.MiRows
			state.RefFrameType[i] = uh.FrameType
			state.RefSubsamplingX[i] = util.Int(sh.ColorConfig.SubsamplingX)
			state.RefSubsamplingY[i] = util.Int(sh.ColorConfig.SubsamplingY)
			state.RefBitDepth[i] = sh.ColorConfig.BitDepth

			for j := 0; j < shared.REFS_PER_FRAME; j++ {
				state.SavedOrderHints[i][j+shared.LAST_FRAME] = uh.OrderHints[j+shared.LAST_FRAME]
			}

			for x := 0; x < uh.UpscaledWidth; x++ {
				for y := 0; y < uh.FrameHeight; y++ {
					t.FrameStore[i][0][y][x] = state.LrFrame[0][y][x]
				}
			}

			for plane := 1; plane <= 2; plane++ {
				for x := 0; x < (uh.UpscaledWidth + util.Int(sh.ColorConfig.SubsamplingX)>>util.Int(sh.ColorConfig.SubsamplingX)); x++ {
					for y := 0; y < (uh.FrameHeight + util.Int(sh.ColorConfig.SubsamplingY)>>util.Int(sh.ColorConfig.SubsamplingY)); y++ {
						t.FrameStore[i][plane][y][x] = state.LrFrame[plane][y][x]
					}
				}
			}

			for row := 0; row < state.MiRows; row++ {
				for col := 0; col < state.MiCols; col++ {
					state.SavedRefFrames[i][row][col] = state.MfRefFrames[row][col]
				}
			}

			for comp := 0; comp <= 1; comp++ {
				for row := 0; row < state.MiRows; row++ {
					for col := 0; col < state.MiCols; col++ {
						state.SavedMvs[i][row][col][comp] = state.MfMvs[row][col][comp]
					}
				}
			}

			for ref := shared.LAST_FRAME; ref <= shared.ALTREF_FRAME; ref++ {
				for j := 0; j <= 5; j++ {
					state.SavedGmParams[i][ref][j] = uh.GmParams[ref][j]
				}
			}

			for row := 0; row < state.MiRows; row++ {
				for col := 0; col < state.MiCols; col++ {
					state.SavedSegmentIds[i][row][col] = t.SegmentIds[row][col]
				}
			}

			saveCdfs(i)

			if sh.FilmGrainParamsPresent {
				saveGrainParams(i)
			}

			saveLoopFilterParams(i)
			saveSegmentationParams(i)

			for i := 0; i < shared.NUM_REF_FRAMES; i++ {
				if ((uh.RefreshFrameFlags >> i) & 1) == 1 {
					state.RefOrderHint[i] = uh.OrderHint
				}
			}
		}
	}

}

// save_cdfs( i )
func saveCdfs(ctx int) {
	// TODO: implement
	//panic("not implemented")
}

// save_grain_params( i )
func saveGrainParams(ctx int) {
	// TODO: implement
	//panic("not implemented")
}

// save_loop_filter_params( i )
func saveLoopFilterParams(ctx int) {
	// TODO: implement
	//panic("not implemented")
}

// save_segmentation_params( i )
func saveSegmentationParams(ctx int) {
	// TODO: implement
	//panic("not implemented")
}

// 7.21
func (t *TileGroup) referenceFrameLoadingProcess() {
	panic("not implemented")
}
