package tilegroup

import "github.com/m4tthewde/gov1/internal/shared"

func (t *TileGroup) isInside(candidateR int, candidateC int) bool {
	return candidateC >= t.State.MiColStart &&
		candidateC < t.State.MiColEnd &&
		candidateR >= t.State.MiRowStart &&
		candidateR < t.State.MiRowEnd
}

// is_scaled( refFrame )
func (t *TileGroup) isScaled(refFrame int) bool {
	refIdx := t.State.UncompressedHeader.RefFrameIdx[refFrame-shared.LAST_FRAME]
	xScale := ((t.RefUpscaledWidth[refIdx] << REF_SCALE_SHIFT) + (t.State.UncompressedHeader.FrameWidth / 2)) / t.State.UncompressedHeader.FrameWidth
	yScale := ((t.RefUpscaledHeight[refIdx] << REF_SCALE_SHIFT) + (t.State.UncompressedHeader.FrameHeight / 2)) / t.State.UncompressedHeader.FrameHeight
	noScale := 1 << REF_SCALE_SHIFT

	return xScale != noScale || yScale != noScale
}

// is_directional_mode( mode )
func (t *TileGroup) isDirectionalMode(mode int) bool {
	return (mode >= V_PRED) && (mode <= D67_PRED)
}
