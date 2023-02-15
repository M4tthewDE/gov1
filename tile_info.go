package main

const MAX_TILE_WIDTH = 4096
const MAX_TILE_AREA = 4096 * 2304
const MAX_TILE_COLS = 64
const MAX_TILE_ROWS = 64

type TileInfo struct {
}

func NewTileInfo(p *Parser, s SequenceHeader) TileInfo {
	t := TileInfo{}
	t.Build(p, s)

	return t
}

func (t *TileInfo) Build(p *Parser, s SequenceHeader) {
	sbCols := (p.MiCols + 15) >> 4
	if s.Use128x128SuperBlock {
		sbCols = (p.MiCols + 31) >> 5
	}

	sbRows := (p.MiRows + 15) >> 4
	if s.Use128x128SuperBlock {
		sbCols = (p.MiRows + 31) >> 5
	}

	sbShift := 4
	if s.Use128x128SuperBlock {
		sbShift = 5
	}

	sbSize := sbShift + 2
	maxTileWidthSb := MAX_TILE_WIDTH >> sbSize
	maxTileAreaSb := MAX_TILE_AREA >> (2 * sbSize)
	minLog2TileCols := tileLog2(maxTileWidthSb, sbCols)
	maxLog2TileCols := tileLog2(1, int(Min(sbCols, MAX_TILE_COLS)))
	maxLog2TileRows := tileLog2(1, Min(sbRows, MAX_TILE_ROWS))
	minLog2Tiles := Max(minLog2TileCols, tileLog2(maxTileAreaSb, sbRows*sbCols))
}
