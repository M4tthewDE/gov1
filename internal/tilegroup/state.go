package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
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
	MiCol                int
	MiRow                int
	MiCols               int
	MiRows               int
	MiColStart           int
	MiColEnd             int
	MiRowStart           int
	MiRowEnd             int
	AvailU               bool
	AvailL               bool
	AvailUChroma         bool
	AvailLChroma         bool
	Skip                 int
	RefFrame             []int
	RefFrames            [][][]int
	IsInter              int
}
