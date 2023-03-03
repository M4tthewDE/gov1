package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/cdef"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/wedgemask"
)

type State struct {
	UncompressedHeader   uncompressedheader.UncompressedHeader
	SequenceHeader       sequenceheader.SequenceHeader
	FrameRestorationType []int
	LoopRestorationSize  []int
	Num4x4BlocksWide     []int
	Num4x4BlocksHigh     []int
	UpscaledWidth        int
	MiSize               int
	MiSizes              [][]int
	MiCol                int
	MiRow                int
	MiCols               int
	MiRows               int
	MiColStart           int
	MiColEnd             int
	MiRowStart           int
	MiRowEnd             int
	MiColStarts          []int
	MiRowStarts          []int
	AvailU               bool
	AvailL               bool
	AvailUChroma         bool
	AvailLChroma         bool
	Skip                 int
	RefFrame             []int
	RefFrames            [][][]int
	IsInter              int
	BlockDecoded         [][][]int
	GmType               []int
	RefFrameWidth        []int
	RefFrameHeight       []int
	CurrFrame            [][][]int
	BlockWidth           []int
	BlockHeight          []int
	ReadDeltas           bool
	FeatureEnabled       [][]int
	Cdef                 cdef.Cdef
	CdefSize4            int
	CurrentQIndex        int
	DeltaLF              []int
	PrevSegmentIds       [][]int
	FeatureData          [][]int
	TileCols             int
	TileRows             int
	TileColsLog2         int
	TileRowsLog2         int
	TileNum              int
	TileSizeBytes        int
	SymbolMaxBits        int
}

func (s *State) newWedgeMaskState() wedgemask.State {
	return wedgemask.State{
		BlockWidth:       s.BlockWidth,
		BlockHeight:      s.BlockHeight,
		Num4x4BlocksWide: s.Num4x4BlocksWide,
		Num4x4BlocksHigh: s.Num4x4BlocksHigh,
	}
}
