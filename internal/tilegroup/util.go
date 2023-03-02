package tilegroup

func (t *TileGroup) isInside(candidateR int, candidateC int) bool {
	return candidateC >= t.State.MiColStart &&
		candidateC < t.State.MiColEnd &&
		candidateR >= t.State.MiRowStart &&
		candidateR < t.State.MiRowEnd
}
