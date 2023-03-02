package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
)

// read_block_tx_size()
func (t *TileGroup) readBlockTxSize(b *bitstream.BitStream) {
	bw4 := t.State.Num4x4BlocksWide[t.State.MiSize]
	bh4 := t.State.Num4x4BlocksHigh[t.State.MiSize]

	if t.State.UncompressedHeader.TxMode == shared.TX_MODE_SELECT &&
		t.State.MiSize > shared.BLOCK_4X4 &&
		util.Bool(t.IsInter) &&
		!util.Bool(t.Skip) &&
		!t.Lossless {
		maxTxSz := Max_Tx_Size_Rect[t.State.MiSize]
		txW4 := Tx_Width[maxTxSz] / MI_SIZE
		txH4 := Tx_Height[maxTxSz] / MI_SIZE

		for row := t.State.MiRow; row < t.State.MiRow+bh4; row += txH4 {
			for col := t.State.MiCol; col < t.State.MiCol+bw4; col += txW4 {
				t.readVarTxSize(row, col, maxTxSz, 0, b)
			}
		}
	} else {
		t.readTxSize(!util.Bool(t.Skip) || util.Bool(t.IsInter), b)
		for row := t.State.MiRow; row < t.State.MiRow+bh4; row++ {
			for col := t.State.MiCol; col < t.State.MiCol+bw4; col++ {
				t.InterTxSizes[row][col] = t.TxSize
			}
		}
	}

}

// read_tx_size( allowSelect )
func (t *TileGroup) readTxSize(allowSelect bool, b *bitstream.BitStream) {
	if t.Lossless {
		t.TxSize = TX_4X4
		return
	}

	maxRectTxSize := Max_Tx_Size_Rect[t.State.MiSize]
	// TODO: what is this for?
	//maxTxDepth := Max_Tx_Depth[p.MiSize]
	t.TxSize = maxRectTxSize

	if t.State.MiSize > shared.BLOCK_4X4 && allowSelect && t.State.UncompressedHeader.TxMode == shared.TX_MODE_SELECT {
		txDepth := b.S()
		for i := 0; i < txDepth; i++ {
			t.TxSize = Split_Tx_Size[t.TxSize]
		}
	}
}

// read_var_tx_size( row, col, txSz, depth )
func (t *TileGroup) readVarTxSize(row int, col int, txSz int, depth int, b *bitstream.BitStream) {
	if row >= t.State.MiRows || col >= t.State.MiCols {
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
				t.readVarTxSize(row+i, col+j, subTxSz, depth+1, b)
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
