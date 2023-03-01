package parser

import (
	"github.com/m4tthewde/gov1/internal"
)

type Parser struct {
	OperatingPointIdc    int
	seenFrameHeader      bool
	renderWidth          int
	renderHeight         int
	upscaledWidth        int
	upscaledHeight       int
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
	Cdef                 internal.Cdef
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

func NewParser() Parser {
	return Parser{}
}

// temporal_unit( sz )
func (p *Parser) temporalUnit(sz int) {
	for sz > 0 {
		frameUnitSize := p.leb128()
		sz -= p.leb128Bytes
		p.frameUnit(frameUnitSize)
		sz -= frameUnitSize
	}
}

// frame_unit( sz )
func (p *Parser) frameUnit(sz int) {
	for sz > 0 {
		obuLength := p.leb128()
		sz -= p.leb128Bytes
		p.parseObu(obuLength)
		sz -= obuLength

	}
}

// init_symbol( x )
func (p *Parser) initSymbol(a int) {
	panic("not implemented: init_symbol()")
}

// clear_above_context()
func (p *Parser) clearAboveContext() {
	panic("not implemented: clear_above_context()")
}

// clear_left_context( x )
func (p *Parser) clearLeftContext() {
	panic("not implemented: clear_left_context()")
}

func (p *Parser) isInside(candidateR int, candidateC int) bool {
	return candidateC >= p.MiColStart && candidateC < p.MiColEnd && candidateR >= p.MiRowStart && candidateR < p.MiRowEnd
}

// choose_operating_point()
func (p *Parser) ChooseOperatingPoint() int {
	// TODO: implement properly
	return 0
}
