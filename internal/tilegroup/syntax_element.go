package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/symbol"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
)

func (t *TileGroup) singleRefP1Symbol(state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	fwdCount := t.countRefs(shared.LAST_FRAME, state)
	fwdCount += t.countRefs(shared.LAST2_FRAME, state)
	fwdCount += t.countRefs(shared.LAST3_FRAME, state)
	fwdCount += t.countRefs(shared.GOLDEN_FRAME, state)
	bwdCount := t.countRefs(shared.BWDREF_FRAME, state)
	bwdCount += t.countRefs(shared.ALTREF2_FRAME, state)
	bwdCount += t.countRefs(shared.ALTREF_FRAME, state)
	ctx := refCountCtx(fwdCount, bwdCount)
	return symbol.ReadSymbol(state.TileSingleRefCdf[ctx][0], state, b, uh)
}

func (t *TileGroup) countRefs(frameType int, state *state.State) int {
	c := 0
	if state.AvailU {
		if t.AboveRefFrame[0] == frameType {
			c++
		}
		if t.AboveRefFrame[1] == frameType {
			c++
		}
	}

	if state.AvailL {
		if t.LeftRefFrame[0] == frameType {
			c++
		}
		if t.LeftRefFrame[1] == frameType {
			c++
		}
	}

	return c
}

func refCountCtx(counts0 int, counts1 int) int {
	if counts0 < counts1 {
		return 0
	} else if counts0 == counts1 {
		return 1
	}

	return 2
}

func (t *TileGroup) singleRefP3Symbol(state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	last12Count := t.countRefs(shared.LAST_FRAME, state) + t.countRefs(shared.LAST2_FRAME, state)
	last3GoldCount := t.countRefs(shared.LAST3_FRAME, state) + t.countRefs(shared.GOLDEN_FRAME, state)
	ctx := refCountCtx(last12Count, last3GoldCount)
	return symbol.ReadSymbol(state.TileSingleRefCdf[ctx][2], state, b, uh)
}

func (t *TileGroup) singleRefP4Symbol(state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	fwdCount := t.countRefs(shared.LAST_FRAME, state)
	fwdCount += t.countRefs(shared.LAST2_FRAME, state)
	fwdCount += t.countRefs(shared.LAST3_FRAME, state)
	fwdCount += t.countRefs(shared.GOLDEN_FRAME, state)
	bwdCount := t.countRefs(shared.BWDREF_FRAME, state)
	bwdCount += t.countRefs(shared.ALTREF2_FRAME, state)
	bwdCount += t.countRefs(shared.ALTREF_FRAME, state)
	ctx := refCountCtx(fwdCount, bwdCount)
	return symbol.ReadSymbol(state.TileSingleRefCdf[ctx][3], state, b, uh)
}
