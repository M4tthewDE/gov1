package tileinfo

type State struct {
	MiCols               int
	MiRows               int
	Use128x128SuperBlock bool
	TileColsLog2         int
	TileRowsLog2         int
	MiRowStarts          []int
	MiColStarts          []int
	TileCols             int
	TileRows             int
	TileSizeBytes        int
}
