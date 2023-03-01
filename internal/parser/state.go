package parser

import (
	"github.com/m4tthewde/gov1/internal"
	"github.com/m4tthewde/gov1/internal/header"
)

type State struct {
	header               header.Header
	operatingPointIdc    int
	seenFrameHeader      bool
	renderWidth          int
	renderHeight         int
	upscaledWidth        int
	upscaledHeight       int
	TileNum              int
	miCols               int
	miRows               int
	miColStarts          []int
	miRowStarts          []int
	miCol                int
	miRow                int
	miSize               int
	miSizes              [][]int
	miRowStart           int
	miColStart           int
	miRowEnd             int
	miColEnd             int
	tileColsLog2         int
	tileRowsLog2         int
	tileCols             int
	tileRows             int
	tileSizeBytes        int
	currentQIndex        int
	deltaLF              []int
	refLrWiener          [][][]int
	num4x4BlocksWide     []int
	num4x4BlocksHigh     []int
	readDeltas           bool
	cdef                 internal.Cdef
	blockDecoded         [][][]int
	frameRestorationType []int
	loopRestorationSize  []int
	availU               bool
	availL               bool
	availUChroma         bool
	availLChroma         bool
	featureEnabled       [][]int
	featureData          [][]int
	refFrame             []int
	refFrames            [][][]int
	refFrameWidth        []int
	refFrameHeight       []int
	gmType               []int
	prevGmParams         [][]int
	prevSegmentIds       [][]int
	currFrame            [][][]int
	symbolMaxBits        int
}
