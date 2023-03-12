package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// read_symbol()
func ReadSymbol(cdf []int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	cur := state.SymbolRange
	symbol := -1

	N := len(cdf) - 1
	var prev int
	for {
		symbol++
		prev = cur
		f := (1 << 15) - cdf[symbol]
		cur := ((state.SymbolRange >> 8) * (f >> EC_PROB_SHIFT)) >> (7 - EC_PROB_SHIFT)
		cur += EC_MIN_PROB * (N - symbol - 1)

		if state.SymbolValue < cur {
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
