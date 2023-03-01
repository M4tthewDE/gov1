package tileinfo

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/util"
)

const MAX_TILE_WIDTH = 4096
const MAX_TILE_AREA = 4096 * 2304
const MAX_TILE_COLS = 64
const MAX_TILE_ROWS = 64

type TileInfo struct {
	ContextUpdateTileId int
}

func NewTileInfo(b *bitstream.BitStream, s State) (TileInfo, State) {
	tileInfo := TileInfo{}

	sbCols := (s.MiCols + 15) >> 4
	if s.Use128x128SuperBlock {
		sbCols = (s.MiCols + 31) >> 5
	}

	sbRows := (s.MiRows + 15) >> 4
	if s.Use128x128SuperBlock {
		sbCols = (s.MiRows + 31) >> 5
	}

	sbShift := 4
	if s.Use128x128SuperBlock {
		sbShift = 5
	}

	sbSize := sbShift + 2
	maxTileWidthSb := MAX_TILE_WIDTH >> sbSize
	maxTileAreaSb := MAX_TILE_AREA >> (2 * sbSize)
	minLog2TileCols := util.TileLog2(maxTileWidthSb, sbCols)
	maxLog2TileCols := util.TileLog2(1, int(util.Min(sbCols, MAX_TILE_COLS)))
	maxLog2TileRows := util.TileLog2(1, util.Min(sbRows, MAX_TILE_ROWS))
	minLog2Tiles := util.Max(minLog2TileCols, util.TileLog2(maxTileAreaSb, sbRows*sbCols))

	uniformTileSpacing := b.F(1) != 0
	if uniformTileSpacing {
		s.TileColsLog2 = minLog2TileCols

		for s.TileColsLog2 < maxLog2TileCols {
			incrementTileColsLog2 := b.F(1) != 0
			if incrementTileColsLog2 {
				s.TileColsLog2++
			} else {
				break
			}

		}

		tileWidthSb := (sbCols + (1 << s.TileColsLog2) - 1) >> s.TileColsLog2
		i := 0
		for startSb := 0; startSb < sbCols; startSb += tileWidthSb {
			s.MiColStarts[i] = startSb << sbShift
			i += 1
		}

		s.MiColStarts[i] = s.MiCols
		s.TileCols = i

		minLog2TileRows := util.Max(minLog2Tiles-s.TileColsLog2, 0)
		s.TileRowsLog2 = minLog2TileRows
		for s.TileRowsLog2 > maxLog2TileRows {
			incrementTileRowsLog2 := b.F(1) != 0
			if incrementTileRowsLog2 {
				s.TileRowsLog2++
			} else {
				break
			}
		}
		tileHeightSb := (sbRows + (1 << s.TileRowsLog2) - 1) >> s.TileRowsLog2
		i = 0
		for startSb := 0; startSb < sbRows; startSb += tileHeightSb {
			s.MiRowStarts[i] = startSb << sbShift
			i += 1
		}
		s.MiRowStarts[i] = s.MiRows
		s.TileRows = i

	} else {
		widestTileSb := 0
		startSb := 0
		i := 0
		for startSb < sbCols {
			s.MiColStarts[i] = startSb << sbShift
			maxWidth := util.Min(sbCols-startSb, maxTileWidthSb)
			widthInSbsMinusOne := b.Ns(maxWidth)
			sizeSb := widthInSbsMinusOne + 1
			widestTileSb = util.Max(sizeSb, widestTileSb)
			startSb += sizeSb

			i++
		}
		s.MiColStarts[i] = s.MiCols
		s.TileCols = i
		s.TileColsLog2 = util.TileLog2(1, s.TileCols)

		if minLog2Tiles > 0 {
			maxTileAreaSb = (sbRows * sbCols) >> (minLog2Tiles + 1)
		} else {
			maxTileAreaSb = sbRows * sbCols
		}
		maxTileHeightSb := util.Max(maxTileAreaSb/widestTileSb, 1)

		startSb = 0
		i = 0
		for startSb < sbRows {
			s.MiRowStarts[i] = startSb << sbShift
			maxHeight := util.Min(sbRows-startSb, maxTileHeightSb)
			heightInSbsMinusOne := b.Ns(maxHeight)
			sizeSb := heightInSbsMinusOne + 1
			startSb += sizeSb

			i++
		}
		s.MiRowStarts[i] = s.MiRows
		s.TileRows = i
		s.TileRowsLog2 = util.TileLog2(1, s.TileRows)
	}
	if s.TileColsLog2 > 0 || s.TileRowsLog2 > 0 {

		tileInfo.ContextUpdateTileId = b.F(s.TileRowsLog2 + s.TileColsLog2)
		tileSizeBytesMinusOne := b.F(2)
		s.TileSizeBytes = tileSizeBytesMinusOne + 1
	} else {
		tileInfo.ContextUpdateTileId = 0
	}

	return tileInfo, s
}
