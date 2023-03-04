package uncompressedheader

import (
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/tileinfo"
)

type State struct {
	Header         header.Header
	SequenceHeader sequenceheader.SequenceHeader
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

	LoopFilterDeltaEnabled bool
	LoopFilterRefDeltas    [8]int
	LoopFilterModeDeltas   [2]int

	CurrentQIndex int
}

func (s *State) newTileInfoState() tileinfo.State {
	return tileinfo.State{
		MiCols:        s.MiCols,
		MiRows:        s.MiRows,
		TileColsLog2:  s.TileColsLog2,
		TileRowsLog2:  s.TileRowsLog2,
		MiRowStarts:   s.MiRowStarts,
		MiColStarts:   s.MiColStarts,
		TileCols:      s.TileCols,
		TileRows:      s.TileRows,
		TileSizeBytes: s.TileSizeBytes,
	}
}

func (s *State) update(resultState tileinfo.State) {
	s.MiCols = resultState.MiCols
	s.MiRows = resultState.MiRows
	s.TileColsLog2 = resultState.TileColsLog2
	s.TileRowsLog2 = resultState.TileRowsLog2
	s.MiRowStarts = resultState.MiRowStarts
	s.MiColStarts = resultState.MiColStarts
	s.TileCols = resultState.TileCols
	s.TileRows = resultState.TileRows
	s.TileSizeBytes = resultState.TileSizeBytes
}
