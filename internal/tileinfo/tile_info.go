package tileinfo

const MAX_TILE_WIDTH = 4096
const MAX_TILE_AREA = 4096 * 2304
const MAX_TILE_COLS = 64
const MAX_TILE_ROWS = 64

type TileInfo struct {
	ContextUpdateTileId int
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

	uniformTileSpacing := p.f(1) != 0
	if uniformTileSpacing {
		p.TileColsLog2 = minLog2TileCols

		for p.TileColsLog2 < maxLog2TileCols {
			incrementTileColsLog2 := p.f(1) != 0
			if incrementTileColsLog2 {
				p.TileColsLog2++
			} else {
				break
			}

		}

		tileWidthSb := (sbCols + (1 << p.TileColsLog2) - 1) >> p.TileColsLog2
		i := 0
		for startSb := 0; startSb < sbCols; startSb += tileWidthSb {
			p.MiColStarts = SliceAssign(p.MiColStarts, i, startSb<<sbShift)
			i += 1
		}

		p.MiColStarts = SliceAssign(p.MiColStarts, i, p.MiCols)
		p.TileCols = i

		minLog2TileRows := Max(minLog2Tiles-p.TileColsLog2, 0)
		p.TileRowsLog2 = minLog2TileRows
		for p.TileRowsLog2 > maxLog2TileRows {
			incrementTileRowsLog2 := p.f(1) != 0
			if incrementTileRowsLog2 {
				p.TileRowsLog2++
			} else {
				break
			}
		}
		tileHeightSb := (sbRows + (1 << p.TileRowsLog2) - 1) >> p.TileRowsLog2
		i = 0
		for startSb := 0; startSb < sbRows; startSb += tileHeightSb {
			p.MiRowStarts = SliceAssign(p.MiRowStarts, i, startSb<<sbShift)
			i += 1
		}
		p.MiRowStarts = SliceAssign(p.MiRowStarts, i, p.MiRows)
		p.TileRows = i

	} else {
		widestTileSb := 0
		startSb := 0
		i := 0
		for startSb < sbCols {
			p.MiColStarts = SliceAssign(p.MiColStarts, i, startSb<<sbShift)
			maxWidth := Min(sbCols-startSb, maxTileWidthSb)
			widthInSbsMinusOne := p.ns(maxWidth)
			sizeSb := widthInSbsMinusOne + 1
			widestTileSb = Max(sizeSb, widestTileSb)
			startSb += sizeSb

			i++
		}
		p.MiColStarts = SliceAssign(p.MiColStarts, i, p.MiCols)
		p.TileCols = i
		p.TileColsLog2 = tileLog2(1, p.TileCols)

		if minLog2Tiles > 0 {
			maxTileAreaSb = (sbRows * sbCols) >> (minLog2Tiles + 1)
		} else {
			maxTileAreaSb = sbRows * sbCols
		}
		maxTileHeightSb := Max(maxTileAreaSb/widestTileSb, 1)

		startSb = 0
		i = 0
		for startSb < sbRows {
			p.MiRowStarts = SliceAssign(p.MiRowStarts, i, startSb<<sbShift)
			maxHeight := Min(sbRows-startSb, maxTileHeightSb)
			heightInSbsMinusOne := p.ns(maxHeight)
			sizeSb := heightInSbsMinusOne + 1
			startSb += sizeSb

			i++
		}
		p.MiRowStarts = SliceAssign(p.MiRowStarts, i, p.MiRows)
		p.TileRows = i
		p.TileRowsLog2 = tileLog2(1, p.TileRows)
	}
	if p.TileColsLog2 > 0 || p.TileRowsLog2 > 0 {

		t.ContextUpdateTileId = p.f(p.TileRowsLog2 + p.TileColsLog2)
		tileSizeBytesMinusOne := p.f(2)
		p.TileSizeBytes = tileSizeBytesMinusOne + 1
	} else {
		t.ContextUpdateTileId = 0
	}
}
