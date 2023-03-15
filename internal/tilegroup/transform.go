package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/util"
)

var Transform_Row_Shift = []int{
	0, 1, 2, 2, 2, 0, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 2, 2,
}

// 7.13.3 2D inverse transform process
func (t *TileGroup) InverseTranform(txSz int, sh sequenceheader.SequenceHeader) {
	log2W := TX_WIDTH_LOG2[txSz]
	log2H := TX_HEIGHT_LOG2[txSz]
	w := 1 << log2W
	h := 1 << log2H
	var rowShift int
	if t.Lossless {
		rowShift = 0
	} else {
		rowShift = Transform_Row_Shift[txSz]
	}

	colShift := 4
	if t.Lossless {
		colShift = 0
	}

	rowClampRange := sh.ColorConfig.BitDepth + 8
	colClampRange := util.Max(sh.ColorConfig.BitDepth+6, 16)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i < 32 && j < 32 {
				t.T[j] = t.Dequant[i][j]
			} else {
				t.T[j] = 0
			}
		}
		if util.Abs(log2W-log2H) == 1 {
			for j := 0; j < w; j++ {
				t.T[j] = util.Round2(t.T[j]*2896, 12)
			}
		}

		if t.Lossless {
			t.inverseWalshHadamardTransform(2)
		} else if t.PlaneTxType == DCT_DCT ||
			t.PlaneTxType == ADST_DCT ||
			t.PlaneTxType == FLIPADST_DCT ||
			t.PlaneTxType == H_DCT {

			t.inverseDCT(log2W, rowClampRange)
		} else if t.PlaneTxType == DCT_ADST ||
			t.PlaneTxType == ADST_ADST ||
			t.PlaneTxType == DCT_FLIPADST ||
			t.PlaneTxType == FLIPADST_FLIPADST ||
			t.PlaneTxType == ADST_FLIPADST ||
			t.PlaneTxType == FLIPADST_ADST ||
			t.PlaneTxType == H_ADST ||
			t.PlaneTxType == H_FLIPADST {

			t.inverseADST(log2W, rowClampRange)
		} else {
			t.inverseIdentityTransform(log2W)
		}
		for j := 0; j < w; j++ {
			t.Residual[i][j] = util.Round2(t.T[j], rowShift)
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			t.Residual[i][j] = util.Clip3(-(1 << (colClampRange - 1)), (1<<(colClampRange-1))-1, t.Residual[i][j])
		}
	}

	for j := 0; j < w; j++ {
		if t.Lossless {
			t.inverseWalshHadamardTransform(0)
		} else if t.PlaneTxType == DCT_DCT ||
			t.PlaneTxType == DCT_ADST ||
			t.PlaneTxType == DCT_FLIPADST ||
			t.PlaneTxType == V_DCT {
			t.inverseDCT(log2H, colClampRange)
		} else if t.PlaneTxType == ADST_DCT ||
			t.PlaneTxType == ADST_ADST ||
			t.PlaneTxType == FLIPADST_DCT ||
			t.PlaneTxType == FLIPADST_FLIPADST ||
			t.PlaneTxType == ADST_FLIPADST ||
			t.PlaneTxType == FLIPADST_ADST ||
			t.PlaneTxType == V_ADST ||
			t.PlaneTxType == V_FLIPADST {
			t.inverseADST(log2H, colClampRange)
		} else {
			t.inverseIdentityTransform(log2H)
		}

		for i := 0; i < h; i++ {
			t.Residual[i][j] = util.Round2(t.T[j], colShift)
		}
	}
}

func (t *TileGroup) b(a int, b int, angle int, flag int, r int) {
	if flag == 1 {
		t.b(a, b, angle, 0, r)
		tmp := t.T[a]
		t.T[a] = t.T[b]
		t.T[b] = tmp
		return
	}
	x := t.T[a]*util.Cos128(angle) - t.T[b]*util.Sin128(angle)
	y := t.T[a]*util.Sin128(angle) - t.T[b]*util.Cos128(angle)
	t.T[a] = util.Round2(x, 12)
	t.T[b] = util.Round2(y, 12)
}

func (t *TileGroup) h(a int, b int, flag int, r int) {
	if flag == 1 {
		t.h(b, a, 0, r)
		return
	}
	x := t.T[a]
	y := t.T[b]
	t.T[a] = util.Clip3(-(1 << (r - 1)), (1<<(r-1))-1, x+y)
	t.T[b] = util.Clip3(-(1 << (r - 1)), (1<<(r-1))-1, x-y)
}

// 7.13.2.2 Inverse DCT array permutation process
func (t *TileGroup) inverseDCTArrayPermutation(n int) {
	copyT := t.T
	for i := 0; i < (1 << n); i++ {
		t.T[i] = copyT[util.Brev(n, i)]
	}
}

// 7.13.2.3 Inverse DCT process
func (t *TileGroup) inverseDCT(n int, r int) {
	t.inverseDCTArrayPermutation(n)

	if n == 6 {
		for i := 0; i <= 15; i++ {
			t.b(32+i, 63-i, 63-4*util.Brev(4, i), 0, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 7; i++ {
			t.b(16+i, 31-i, 6+(util.Brev(3, 7-i)<<3), 0, r)
		}
	}

	if n == 6 {
		for i := 0; i <= 15; i++ {
			t.h(32+i*2, 33+i*2, i&1, r)
		}

	}

	if n >= 4 {
		for i := 0; i <= 3; i++ {
			t.b(8+i, 15-i, 12+(util.Brev(2, 3-i)<<4), 0, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 7; i++ {
			t.h(16+2*i, 17+2*i, i&1, r)
		}
	}

	if n == 6 {
		for i := 0; i <= 3; i++ {
			for j := 0; j <= 1; j++ {
				t.b(62-i*4-j, 33+i*4+j, 60-16*util.Brev(2, i)+64*j, 1, r)
			}
		}
	}

	if n >= 3 {
		for i := 0; i <= 1; i++ {
			t.b(4+i, 7-i, 56-32*i, 0, r)
		}

	}

	if n >= 4 {
		for i := 0; i <= 3; i++ {
			t.h(8+2*i, 9+2*i, i&1, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 1; j++ {
				t.b(30-4*i-j, 17+4*i+j, 24+(j<<6)+((1-i)<<5), 1, r)
			}
		}
	}

	if n == 6 {
		for i := 0; i <= 7; i++ {
			for j := 0; j <= 1; j++ {
				t.h(32+i*4+j, 35+i*4-j, i&1, r)
			}
		}
	}

	for i := 0; i <= 1; i++ {
		t.b(2*i, 2*i+1, 32+16*i, 1-i, r)
	}

	if n >= 3 {
		for i := 0; i <= 1; i++ {
			t.h(4+2*i, 5+2*i, i, r)
		}
	}

	if n >= 4 {
		for i := 0; i <= 1; i++ {
			t.b(14-i, 9+i, 48+64*i, 1, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 3; i++ {
			for j := 0; j <= 1; j++ {
				t.h(16+4*i+j, 19+4*i-j, i&1, r)
			}
		}
	}

	if n == 6 {
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 3; j++ {
				t.b(61-i*8-j, 34+i*8+j, 56-i*32+(j>>1)*64, 1, r)
			}
		}
	}

	for i := 0; i <= 1; i++ {
		t.h(i, 3-i, 0, r)
	}

	if n >= 3 {
		t.b(6, 5, 32, 1, r)
	}

	if n >= 4 {
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 1; j++ {
				t.h(8+4*i+j, 11+4*i-j, i, r)
			}
		}
	}

	if n >= 5 {
		for i := 0; i <= 3; i++ {
			t.b(29-i, 18+i, 48+(i>>1)*64, 1, r)
		}
	}

	if n == 6 {
		for i := 0; i <= 3; i++ {
			for j := 0; j <= 3; j++ {
				t.h(32+8*i+j, 39+8*i-j, i&1, r)
			}
		}
	}

	if n >= 3 {
		for i := 0; i <= 3; i++ {
			t.h(i, 7-i, 0, r)
		}
	}

	if n >= 4 {
		for i := 0; i <= 1; i++ {
			t.b(13-i, 10+i, 32, 1, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 1; i++ {
			for j := 0; j <= 3; j++ {
				t.h(16+i*8+j, 23+i*8-j, i, r)
			}
		}
	}

	if n == 6 {
		for i := 0; i <= 7; i++ {
			if i < 4 {
				t.b(59-i, 36+i, 48, 1, r)
			} else {
				t.b(59-i, 36+i, 112, 1, r)
			}
		}
	}
	if n >= 4 {
		for i := 0; i <= 7; i++ {
			t.h(i, 15-i, 0, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 3; i++ {
			t.b(27-i, 20+i, 32, 1, r)
		}
	}

	if n == 6 {
		for i := 0; i <= 7; i++ {
			t.h(32+i, 47-i, 0, r)
			t.h(48+i, 63-i, 1, r)
		}
	}

	if n >= 5 {
		for i := 0; i <= 15; i++ {
			t.h(i, 31-i, 0, r)
		}
	}

	if n == 6 {
		for i := 0; i <= 7; i++ {
			t.b(55-i, 40+i, 32, 1, r)
		}
	}

	if n == 6 {
		for i := 0; i <= 31; i++ {
			t.h(i, 63-i, 0, r)
		}
	}
}

// 7.13.2.4 Inverse ADST input array permutation process
func (t *TileGroup) inverseADSTInputArrayPermutation(n int) {
	n0 := 1 << n
	copyT := t.T
	for i := 0; i < n0; i++ {
		var idx int
		if util.Bool(i & 1) {
			idx = i - 1
		} else {
			idx = n0 - i - 1
		}
		t.T[i] = copyT[idx]
	}
}

// 7.13.2.5 inverse ADST output array permutation process
func (t *TileGroup) inverseADSTOutputArrayPermutation(n int) {
	n0 := 1 << n
	copyT := t.T

	for i := 0; i < n0; i++ {
		a := ((i >> 3) & 1)
		b := ((i >> 2) & 1) ^ ((i >> 3) & 1)
		c := ((i >> 1) & 1) ^ ((i >> 2) & 1)
		d := (i & 1) ^ ((i >> 1) & 1)
		idx := ((d << 3) | (c << 2) | (b << 1) | a) >> (4 - n)

		if util.Bool(i & 1) {
			t.T[i] = -copyT[idx]
		} else {
			t.T[i] = copyT[idx]
		}
	}

}

const SINPI_1_9 = 1321
const SINPI_2_9 = 2482
const SINPI_3_9 = 3344
const SINPI_4_9 = 3803

// 7.13.2.6 Inverse ADST4 process
func (t *TileGroup) inverseADST4(r int) {
	// TODO: bistream conformance precision ???
	s := []int{}
	x := []int{}

	s[0] = SINPI_1_9 * t.T[0]
	s[1] = SINPI_2_9 * t.T[0]
	s[2] = SINPI_3_9 * t.T[1]
	s[3] = SINPI_4_9 * t.T[2]
	s[4] = SINPI_1_9 * t.T[2]
	s[5] = SINPI_2_9 * t.T[3]
	s[6] = SINPI_4_9 * t.T[3]
	a7 := t.T[0] - t.T[2]
	b7 := a7 + t.T[3]
	s[0] = s[0] + s[3]
	s[1] = s[1] - s[4]
	s[3] = s[2]
	s[2] = SINPI_3_9 * b7
	s[0] = s[0] + s[5]
	s[1] = s[1] - s[6]
	x[0] = s[0] + s[3]
	x[1] = s[1] + s[3]
	x[2] = s[2]
	x[3] = s[0] + s[1]
	x[3] = x[3] - s[3]
	t.T[0] = util.Round2(x[0], 12)
	t.T[1] = util.Round2(x[1], 12)
	t.T[2] = util.Round2(x[2], 12)
	t.T[3] = util.Round2(x[3], 12)
}

// 7.13.2.7 Inverse ADST8 process
func (t *TileGroup) inverseADST8(r int) {
	t.inverseADSTInputArrayPermutation(3)
	for i := 0; i <= 3; i++ {
		t.b(2*i, 2*i+1, 60-16*i, 1, r)
	}

	for i := 0; i <= 3; i++ {
		t.h(i, 4+i, 0, r)
	}

	for i := 0; i <= 1; i++ {
		t.b(4+3*i, 5+i, 48-32*i, 1, r)
	}

	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			t.h(4*j+i, 2+4*j+i, 0, r)
		}
	}

	for i := 0; i <= 1; i++ {
		t.b(2+4*i, 3+4*i, 32, 1, r)
	}

	t.inverseADSTOutputArrayPermutation(4)
}

// 7.13.2.8 Inverse ADST16 process
func (t *TileGroup) inverseADST16(r int) {
	t.inverseADSTInputArrayPermutation(4)
	for i := 0; i <= 7; i++ {
		t.b(2*i, 2*i+1, 62-8*i, 1, r)
	}

	for i := 0; i <= 7; i++ {
		t.h(i, 8+i, 0, r)
	}

	for i := 0; i <= 1; i++ {
		t.b(8+2*i, 9+2*i, 56-32*i, 1, r)
		t.b(13+2*i, 12+2*i, 8+32*i, 1, r)
	}

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 1; j++ {
			t.h(8*j+i, 4+8*j+i, 0, r)
		}
	}

	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			t.b(4+8*j+3*i, 5+8*j+i, 48-32*i, 1, r)
		}
	}

	for i := 0; i <= 1; i++ {
		for j := 0; j <= 3; j++ {
			t.h(4*j+i, 2+4*j+i, 0, r)
		}
	}

	for i := 0; i <= 3; i++ {
		t.b(2+4*i, 3+4*i, 32, 1, r)
	}

	t.inverseADSTOutputArrayPermutation(4)
}

// 7.13.2.9 Inverse ADST process
func (t *TileGroup) inverseADST(n int, r int) {
	if n == 2 {
		t.inverseADST4(r)
	} else if n == 3 {
		t.inverseADST8(r)
	} else {
		t.inverseADST16(r)
	}

}

// 7.13.2.10 Inverse Walsh-Hadamard transform process
func (t *TileGroup) inverseWalshHadamardTransform(shift int) {
	a := t.T[0] >> shift
	c := t.T[1] >> shift
	d := t.T[2] >> shift
	b := t.T[3] >> shift
	a += c
	d -= b
	e := (a - d) >> 1
	b = e - b
	c = e - c
	a -= b
	d += c
	t.T[0] = a
	t.T[1] = b
	t.T[2] = c
	t.T[3] = d
}

// 7.13.2.11 Inverse identity transform 4 process
func (t *TileGroup) inverseIdentitytTransform4() {
	for i := 0; i <= 3; i++ {
		t.T[i] = util.Round2(t.T[i]*5793, 12)
	}
}

// 7.13.2.12 Inverse identity transform 8 process
func (t *TileGroup) inverseIdentitytTransform8() {
	for i := 0; i <= 7; i++ {
		t.T[i] = t.T[i] * 2
	}
}

// 7.13.2.13 Inverse identity transform 16 process
func (t *TileGroup) inverseIdentitytTransform16() {
	for i := 0; i <= 15; i++ {
		t.T[i] = util.Round2(t.T[i]*11586, 12)
	}
}

// 7.13.2.14 Inverse identity transform 32 process
func (t *TileGroup) inverseIdentitytTransform32() {
	for i := 0; i <= 31; i++ {
		t.T[i] = t.T[i] * 4
	}
}

// 7.13.2.15 Inverse identity transform process
func (t *TileGroup) inverseIdentityTransform(n int) {
	if n == 2 {
		t.inverseIdentitytTransform4()
	} else if n == 3 {
		t.inverseIdentitytTransform8()
	} else if n == 4 {
		t.inverseIdentitytTransform16()
	} else {
		t.inverseIdentitytTransform32()
	}
}
