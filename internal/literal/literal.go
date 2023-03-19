package literal

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/symbol"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

func L(n int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	x := 0
	for i := 0; i < n; i++ {
		x = 2*x + readBool(state, b, uh)
	}

	return x
}

// NS( n )
func NS(n int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	w := util.FloorLog2(n) + 1
	m := (1 << w) - n
	v := L(w-1, state, b, uh)
	if v < m {
		return v
	}
	extraBit := L(1, state, b, uh)
	return (v << 1) - m + extraBit
}

func readBool(state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	cdf := []int{
		1 << 14,
		1 << 15,
		0,
	}

	return symbol.ReadSymbol(cdf, state, b, uh)
}
