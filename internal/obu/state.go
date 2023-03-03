package obu

import (
	"github.com/m4tthewde/gov1/internal/cdef"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/tilegroup"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
)

type State struct {
	Header               header.Header
	SequenceHeader       sequenceheader.SequenceHeader
	OperatingPointIdc    int
	SeenFrameHeader      bool
	RenderWidth          int
	RenderHeight         int
	UpscaledWidth        int
	UpscaledHeight       int
	TileNum              int
	MiCols               int
	MiRows               int
	MiColStarts          []int
	MiRowStarts          []int
	MiCol                int
	MiRow                int
	MiSize               int
	MiSizes              [][]int
	MiRowStart           int
	MiColStart           int
	MiRowEnd             int
	MiColEnd             int
	TileColsLog2         int
	TileRowsLog2         int
	TileCols             int
	TileRows             int
	TileSizeBytes        int
	CurrentQIndex        int
	DeltaLF              []int
	RefLrWiener          [][][]int
	Num4x4BlocksWide     []int
	Num4x4BlocksHigh     []int
	ReadDeltas           bool
	Cdef                 cdef.Cdef
	BlockDecoded         [][][]int
	FrameRestorationType []int
	LoopRestorationSize  []int
	AvailU               bool
	AvailL               bool
	AvailUChroma         bool
	AvailLChroma         bool
	FeatureEnabled       [][]int
	FeatureData          [][]int
	RefFrame             []int
	RefFrames            [][][]int
	RefFrameWidth        []int
	RefFrameHeight       []int
	GmType               []int
	PrevGmParams         [][]int
	PrevSegmentIds       [][]int
	CurrFrame            [][][]int
	SymbolMaxBits        int
}

func NewState() State {
	// TODO: probably smart to set some sensible defaults here
	return State{}
}

func (s *State) newUncompressedHeaderState() uncompressedheader.State {
	return uncompressedheader.State{
		Header:         s.Header,
		SequenceHeader: s.SequenceHeader,
		MiCols:         s.MiCols,
		MiRows:         s.MiRows,
		TileColsLog2:   s.TileColsLog2,
		TileRowsLog2:   s.TileRowsLog2,
		MiRowStarts:    s.MiRowStarts,
		MiColStarts:    s.MiColStarts,
		TileCols:       s.TileCols,
		TileRows:       s.TileRows,
		TileSizeBytes:  s.TileSizeBytes,
		FeatureEnabled: s.FeatureEnabled,
		FeatureData:    s.FeatureData,
		GmType:         s.GmType,
		PrevGmParams:   s.PrevGmParams,
	}
}

func (s *State) newTileGroupState() tilegroup.State {
	return tilegroup.State{}
}

func (s *State) update(state uncompressedheader.State) {
	s.Header = state.Header
	s.SequenceHeader = state.SequenceHeader
	s.MiCols = state.MiCols
	s.MiRows = state.MiRows
	s.TileColsLog2 = state.TileColsLog2
	s.TileRowsLog2 = state.TileRowsLog2
	s.MiRowStarts = state.MiRowStarts
	s.MiColStarts = state.MiColStarts
	s.TileCols = state.TileCols
	s.TileRows = state.TileRows
	s.TileSizeBytes = state.TileSizeBytes
	s.FeatureEnabled = state.FeatureEnabled
	s.FeatureData = state.FeatureData
	s.GmType = state.GmType
	s.PrevGmParams = state.PrevGmParams
}
