package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/symbol"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// read_block_tx_size()
func (t *TileGroup) readBlockTxSize(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]

	if uh.TxMode == shared.TX_MODE_SELECT &&
		state.MiSize > shared.BLOCK_4X4 &&
		util.Bool(t.IsInter) &&
		!util.Bool(t.Skip) &&
		!t.Lossless {
		maxTxSz := Max_Tx_Size_Rect[state.MiSize]
		txW4 := Tx_Width[maxTxSz] / MI_SIZE
		txH4 := Tx_Height[maxTxSz] / MI_SIZE

		for row := state.MiRow; row < state.MiRow+bh4; row += txH4 {
			for col := state.MiCol; col < state.MiCol+bw4; col += txW4 {
				t.readVarTxSize(row, col, maxTxSz, 0, b, state)
			}
		}
	} else {
		t.readTxSize(!util.Bool(t.Skip) || util.Bool(t.IsInter), b, state, uh)
		for row := state.MiRow; row < state.MiRow+bh4; row++ {
			for col := state.MiCol; col < state.MiCol+bw4; col++ {
				t.InterTxSizes[row][col] = t.TxSize
			}
		}
	}

}

// read_tx_size( allowSelect )
func (t *TileGroup) readTxSize(allowSelect bool, b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	if t.Lossless {
		t.TxSize = TX_4X4
		return
	}

	maxRectTxSize := Max_Tx_Size_Rect[state.MiSize]
	maxTxDepth := Max_Tx_Depth[state.MiSize]
	t.TxSize = maxRectTxSize

	if state.MiSize > shared.BLOCK_4X4 && allowSelect && uh.TxMode == shared.TX_MODE_SELECT {
		txDepth := t.txDepthSymbol(maxRectTxSize, maxTxDepth, state, b, uh)
		for i := 0; i < txDepth; i++ {
			t.TxSize = Split_Tx_Size[t.TxSize]
		}
	}
}

func (t *TileGroup) txDepthSymbol(maxRectTxSize int, maxTxDepth int, state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	maxTxWidth := Tx_Width[maxRectTxSize]
	maxTxHeight := Tx_Height[maxRectTxSize]

	var aboveW int
	if state.AvailU && util.Bool(t.IsInters[state.MiRow-1][state.MiCol]) {
		aboveW = shared.BLOCK_WIDTH[state.MiSizes[state.MiRow-1][state.MiCol]]
	} else if state.AvailU {
		aboveW = t.getAboveTxWidth(state.MiRow, state.MiCol, state)
	} else {
		aboveW = 0
	}

	var leftH int
	if state.AvailL && util.Bool(t.IsInters[state.MiRow][state.MiCol-1]) {
		leftH = shared.BLOCK_HEIGHT[state.MiSizes[state.MiRow][state.MiCol-1]]
	} else if state.AvailU {
		leftH = t.getLeftTxHeight(state.MiRow, state.MiCol, state)
	} else {
		leftH = 0
	}

	ctx := util.Int(aboveW >= maxTxWidth) + util.Int(leftH >= maxTxHeight)

	if maxTxDepth == 4 {
		return symbol.ReadSymbol(state.TileTx64x64Cdf[ctx], state, b, uh)

	} else if maxTxDepth == 3 {
		return symbol.ReadSymbol(state.TileTx32x32Cdf[ctx], state, b, uh)
	} else if maxTxDepth == 2 {
		return symbol.ReadSymbol(state.TileTx16x16Cdf[ctx], state, b, uh)
	} else {
		return symbol.ReadSymbol(state.TileTx8x8Cdf[ctx], state, b, uh)
	}
}

func (t *TileGroup) getAboveTxWidth(row int, col int, state *state.State) int {
	if row == state.MiRow {
		if !state.AvailU {
			return 64
		} else if util.Bool(t.Skips[row-1][col]) && util.Bool(t.IsInters[row-1][col]) {
			return shared.BLOCK_WIDTH[state.MiSizes[row-1][col]]
		}
	}
	return Tx_Width[t.InterTxSizes[row-1][col]]
}

func (t *TileGroup) getLeftTxHeight(row int, col int, state *state.State) int {
	if row == state.MiCol {
		if !state.AvailL {
			return 64
		} else if util.Bool(t.Skips[row][col-1]) && util.Bool(t.IsInters[row][col-1]) {
			return shared.BLOCK_HEIGHT[state.MiSizes[row][col-1]]
		}
	}
	return Tx_Height[t.InterTxSizes[row][col-1]]
}

// read_var_tx_size( row, col, txSz, depth )
func (t *TileGroup) readVarTxSize(row int, col int, txSz int, depth int, b *bitstream.BitStream, state *state.State) {
	if row >= state.MiRows || col >= state.MiCols {
		return
	}

	var txfmSplit int
	if txSz == TX_4X4 || depth == MAX_VARTX_DEPTH {
		txfmSplit = 0
	} else {
		txfmSplit = b.S()
	}

	w4 := Tx_Width[txSz] / MI_SIZE
	h4 := Tx_Height[txSz] / MI_SIZE

	if util.Bool(txfmSplit) {
		subTxSz := Split_Tx_Size[txSz]
		stepW := Tx_Width[subTxSz] / MI_SIZE
		stepH := Tx_Height[subTxSz] / MI_SIZE

		for i := 0; i < h4; i += stepH {
			for j := 0; j < w4; j += stepW {
				t.readVarTxSize(row+i, col+j, subTxSz, depth+1, b, state)
			}
		}
	} else {
		for i := 0; i < h4; i++ {
			for j := 0; j < w4; j++ {
				t.InterTxSizes[row+i][col+i] = txSz
			}
		}
		t.TxSize = txSz
	}
}

const MAX_VARTX_DEPTH = 2

const TX_4X4 = 0
const TX_8X8 = 1
const TX_16X16 = 2
const TX_32X32 = 3
const TX_64X64 = 4
const TX_4X8 = 5
const TX_8X4 = 6
const TX_8X16 = 7
const TX_16X8 = 8
const TX_16X32 = 9
const TX_32X16 = 10
const TX_32X64 = 11
const TX_64X32 = 12
const TX_4X16 = 13
const TX_16X4 = 14
const TX_8X32 = 15
const TX_32X8 = 16
const TX_16X64 = 17
const TX_64X16 = 18

var Max_Tx_Size_Rect = []int{
	TX_4X4, TX_4X8, TX_8X4, TX_8X8,
	TX_8X16, TX_16X8, TX_16X16, TX_16X32,
	TX_32X16, TX_32X32, TX_32X64, TX_64X32,
	TX_64X64, TX_64X64, TX_64X64, TX_64X64,
	TX_4X16, TX_16X4, TX_8X32, TX_32X8,
	TX_16X64, TX_64X16,
}

var Max_Tx_Depth = []int{
	0, 1, 1, 1,
	2, 2, 2, 3,
	3, 3, 4, 4,
	4, 4, 4, 4,
	2, 2, 3, 3,
	4, 4,
}

var Split_Tx_Size = []int{
	TX_4X4,
	TX_4X4,
	TX_8X8,
	TX_16X16,
	TX_32X32,
	TX_4X4,
	TX_4X4,
	TX_8X8,
	TX_8X8,
	TX_16X16,
	TX_16X16,
	TX_32X32,
	TX_32X32,
	TX_4X8,
	TX_8X4,
	TX_8X16,
	TX_16X8,
	TX_16X32,
	TX_32X16,
}

var Tx_Width = []int{4, 8, 16, 32, 64, 4, 8, 8, 16, 16, 32, 32, 64, 4, 16, 8, 32, 16, 64}
var Tx_Height = []int{4, 8, 16, 32, 64, 8, 4, 16, 8, 32, 16, 64, 32, 16, 4, 32, 8, 64, 16}
