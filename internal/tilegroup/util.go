package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
)

func (t *TileGroup) isInside(candidateR int, candidateC int, state *state.State) bool {
	return candidateC >= state.MiColStart &&
		candidateC < state.MiColEnd &&
		candidateR >= state.MiRowStart &&
		candidateR < state.MiRowEnd
}

// is_scaled( refFrame )
func (t *TileGroup) isScaled(refFrame int, uh uncompressedheader.UncompressedHeader) bool {
	refIdx := uh.RefFrameIdx[refFrame-shared.LAST_FRAME]
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (uh.FrameWidth / 2)) / uh.FrameWidth
	yScale := ((t.RefFrameHeight[refIdx] << REF_SCALE_SHIFT) + (uh.FrameHeight / 2)) / uh.FrameHeight
	noScale := 1 << REF_SCALE_SHIFT

	return xScale != noScale || yScale != noScale
}

// is_directional_mode( mode )
func (t *TileGroup) isDirectionalMode(mode int) bool {
	return (mode >= V_PRED) && (mode <= D67_PRED)
}

func isInsideFilterRegion(candidateR int, candidateC int, state *state.State) bool {
	colStart := 0
	colEnd := state.MiCols
	rowStart := 0
	rowEnd := state.MiRows
	return candidateC >= colStart &&
		candidateC < colEnd &&
		candidateR >= rowStart &&
		candidateR < rowEnd
}
