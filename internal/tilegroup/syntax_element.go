package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// read_symbol()
func ReadSymbol(cdf []int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	N := len(cdf) - 1

	// precondition
	if N < 1 || cdf[N-1] != 1<<15 {
		panic("Violated precondition")
	}

	cur := state.SymbolRange
	symbol := -1

	var prev int
	var f int
	for {
		symbol++
		prev = cur
		f = (1 << 15) - cdf[symbol]
		cur = ((state.SymbolRange >> 8) * (f >> EC_PROB_SHIFT)) >> (7 - EC_PROB_SHIFT)
		cur += EC_MIN_PROB * (N - symbol - 1)

		if state.SymbolValue >= cur {
			break
		}
	}

	state.SymbolRange = prev - cur
	state.SymbolValue = state.SymbolValue - cur

	bits := 15 - util.FloorLog2(state.SymbolRange)
	state.SymbolRange = state.SymbolRange << bits
	numBits := util.Min(bits, util.Max(0, state.SymbolMaxBits))
	newData := b.F(numBits)
	paddedData := newData << (bits - numBits)
	state.SymbolValue = paddedData ^ (((state.SymbolValue + 1) << bits) - 1)
	state.SymbolMaxBits = state.SymbolMaxBits - bits

	if !uh.DisableCdfUpdate {
		rate := 3 + util.Int(cdf[N] > 15) + util.Int(cdf[N] > 31) + util.Min(util.FloorLog2(N), 2)
		tmp := 0
		for i := 0; i < N-1; i++ {
			if i == symbol {
				tmp = 1 << 15
			}
			if tmp < cdf[i] {
				cdf[i] -= ((cdf[i] - tmp) >> rate)
			} else {
				cdf[i] += ((tmp - cdf[i]) >> rate)
			}

		}
		cdf[N] += util.Int(cdf[N] < 32)
	}

	return symbol
}

const EC_PROB_SHIFT = 6
const EC_MIN_PROB = 4

func (t *TileGroup) singleRefP1Symbol(state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	fwdCount := t.countRefs(shared.LAST_FRAME, state)
	fwdCount += t.countRefs(shared.LAST2_FRAME, state)
	fwdCount += t.countRefs(shared.LAST3_FRAME, state)
	fwdCount += t.countRefs(shared.GOLDEN_FRAME, state)
	bwdCount := t.countRefs(shared.BWDREF_FRAME, state)
	bwdCount += t.countRefs(shared.ALTREF2_FRAME, state)
	bwdCount += t.countRefs(shared.ALTREF_FRAME, state)
	ctx := refCountCtx(fwdCount, bwdCount)
	return ReadSymbol(state.TileSingleRefCdf[ctx][0], state, b, uh)
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
	return ReadSymbol(state.TileSingleRefCdf[ctx][2], state, b, uh)
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
	return ReadSymbol(state.TileSingleRefCdf[ctx][3], state, b, uh)
}
