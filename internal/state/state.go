package state

import (
	"github.com/m4tthewde/gov1/internal/shared"
)

type State struct {
	MiCols         int
	MiRows         int
	TileColsLog2   int
	TileRowsLog2   int
	MiRowStarts    []int
	MiColStarts    []int
	TileCols       int
	TileRows       int
	TileSizeBytes  int
	FeatureEnabled [shared.MAX_SEGMENTS][shared.SEG_LVL_MAX]int
	FeatureData    [shared.MAX_SEGMENTS][shared.SEG_LVL_MAX]int
	GmType         [shared.ALTREF_FRAME + 1]int
	PrevGmParams   [shared.ALTREF_FRAME + 1][6]int
	PrevSegmentIds [][]int
	RefFrameType   [7]int
	CurrentQIndex  int
}
