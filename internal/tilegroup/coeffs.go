package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// coeffs( plane, startX, startY, txSz )
func (t *TileGroup) coeffs(plane int, startX int, startY int, txSz int, b *bitstream.BitStream) {
	x4 := startX >> 2
	y4 := startY >> 2
	w4 := Tx_Width[txSz] >> 2
	h4 := Tx_Height[txSz] >> 2
	txSzCtx := (TX_SIZE_SQR[txSz] + TX_SIZE_SQR_UP[txSz] + 1) >> 1
	ptype := plane > 0
	var segEob int
	if txSz == TX_16X64 || txSz == TX_64X16 {
		segEob = 512
	} else {
		segEob = util.Min(1024, Tx_Width[txSz]*Tx_Height[txSz])
	}

	for c := 0; c < segEob; c++ {
		t.Quant[c] = 0
	}
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			t.Dequant[i][j] = 0
		}
	}
	eob := 0
	culLevel := 0
	dcCategory := 0

	allZero := b.S()
}

// transform_type( x4, y4, txSz)
func (t *TileGroup) transformType(x4 int, y4 int, txSz int, uh uncompressedheader.UncompressedHeader, state *state.State, b *bitstream.BitStream) {
	set := t.getTxSet(txSz, uh)

	qIndex := uh.BaseQIdx
	if uh.SegmentationEnabled {
		qIndex = uh.GetQIndex(1, t.SegmentId, state)
	}
	if set > 0 && qIndex > 0 {
		if util.Bool(t.IsInter) {
			interTxType := b.S()
			if set == TX_SET_INTER_1 {
				TxType := Tx_TYpe
			}
		}
	}
}

// get_tx_set( txSz )
func (t *TileGroup) getTxSet(txSz int, uh uncompressedheader.UncompressedHeader) int {
	txSzSqr := TX_SIZE_SQR[txSz]
	txSzSqrUp := TX_SIZE_SQR_UP[txSz]
	if txSzSqrUp > TX_32X32 {
		return TX_SET_DCTONLY
	}

	if util.Bool(t.IsInter) {
		if uh.ReducedTxSet || txSzSqrUp == TX_32X32 {
			return TX_SET_INTER_3
		} else if txSzSqr == TX_16X16 {
			return TX_SET_INTER_2
		}
		return TX_SET_INTER_1
	} else {
		if txSzSqrUp == TX_32X32 {
			return TX_SET_DCTONLY
		} else if uh.ReducedTxSet {
			return TX_SET_INTRA_2
		} else if txSzSqr == TX_16X16 {
			return TX_SET_INTRA_2
		}
		return TX_SET_INTRA_1
	}
}

var Tx_Type_Intra_Inv_Set1 = []int{IDTX, DCT_DCT, V_DCT, H_DCT, ADST_ADST, ADST_DCT, DCT_ADST}
var Tx_Type_Intra_Inv_Set2 = []int{IDTX, DCT_DCT, ADST_ADST, ADST_DCT, DCT_ADST}
var Tx_Type_Inter_Inv_Set1 = []int{IDTX, V_DCT, H_DCT, V_ADST, H_ADST, V_FLIPADST, H_FLIPADST,
	DCT_DCT, ADST_DCT, DCT_ADST, FLIPADST_DCT, DCT_FLIPADST, ADST_ADST,
	FLIPADST_FLIPADST, ADST_FLIPADST, FLIPADST_ADST}
var Tx_Type_Inter_Inv_Set2 = []int{IDTX, V_DCT, H_DCT, DCT_DCT, ADST_DCT, DCT_ADST, FLIPADST_DCT,
	DCT_FLIPADST, ADST_ADST, FLIPADST_FLIPADST, ADST_FLIPADST,
	FLIPADST_ADST}
var Tx_Type_Inter_Inv_Set3 = []int{IDTX, DCT_DCT}

const TX_SET_DCTONLY = 0
const TX_SET_INTER_1 = 1
const TX_SET_INTER_2 = 2
const TX_SET_INTER_3 = 3
const TX_SET_INTRA_1 = 4
const TX_SET_INTRA_2 = 5
const TX_SET_INTRA_3 = 6

var TX_SIZE_SQR = []int{
	TX_4X4,
	TX_8X8,
	TX_16X16,
	TX_32X32,
	TX_64X64,
	TX_4X4,
	TX_4X4,
	TX_8X8,
	TX_8X8,
	TX_16X16,
	TX_16X16,
	TX_32X32,
	TX_32X32,
	TX_4X4,
	TX_4X4,
	TX_8X8,
	TX_8X8,
	TX_16X16,
	TX_16X16,
}

var TX_SIZE_SQR_UP = []int{
	TX_4X4,
	TX_8X8,
	TX_16X16,
	TX_32X32,
	TX_64X64,
	TX_8X8,
	TX_8X8,
	TX_16X16,
	TX_16X16,
	TX_32X32,
	TX_32X32,
	TX_64X64,
	TX_64X64,
	TX_16X16,
	TX_16X16,
	TX_32X32,
	TX_32X32,
	TX_64X64,
	TX_64X64,
}
