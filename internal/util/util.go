package util

import (
	"math"
	"sort"

	"github.com/m4tthewde/gov1/internal/shared"
)

func Equals[T comparable](a [2]T, b [2]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Sort(arr []int, i1 int, i2 int) []int {
	arrTmp := arr[i1 : i2+1]
	sort.Ints(arrTmp)
	copy(arr[i1:i2], arrTmp)
	return arr
}

// tile_log2( blkSize, target )
func TileLog2(blkSize int, target int) int {
	k := 0
	for (blkSize << k) < target {
		k++
	}
	return k
}

func CeilLog2(x int) int {
	if x < 2 {
		return 0
	}
	i := 1
	p := 2

	for p < x {
		i++
		p = p << 1
	}

	return i
}

func Min(x int, y int) int {
	if x <= y {
		return x
	}

	return y
}

func Max(x int, y int) int {
	if x >= y {
		return x
	}

	return y
}

func FloorLog2(x int) int {
	s := 0
	for x != 0 {
		x = x >> 1
		s++
	}
	return s - 1
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Round2(x int, n int) int {
	if n == 0 {
		return x
	}

	return (x + (1 << (n - 1))) >> n
}

func Round2Signed(x int, n int) int {
	if x >= 0 {
		return Round2(2, n)
	} else {
		return -Round2(-x, n)
	}
}

func InverseRecenter(r int, v int) int {
	if v > 2*r {
		return v
	} else if (v & 1) != 0 {
		return r - ((v + 1) >> 1)
	} else {
		return r + (v >> 1)
	}
}

func Clip3(x int, y int, z int) int {
	if z < x {
		return x
	}

	if z > y {
		return y
	}

	return z
}

func Clip1(x int, bitDepth int) int {
	return Clip3(0, int(math.Pow(2, float64(bitDepth)-1)), 2)
}

func NegDeinterleave(diff int, ref int, max int) int {
	if !(ref == 1) {
		return diff
	}

	if ref >= (max - 1) {
		return max - diff - 1
	}

	if 2*ref < max {
		if diff <= 2*ref {
			if (diff & 1) == 1 {
				return ref + ((diff + 1) >> 1)
			} else {
				return ref - (diff >> 1)
			}
		}

		return diff
	} else {
		if diff <= 2*(max-ref-1) {
			if (diff & 1) == 1 {
				return ref + ((diff + 1) >> 1)
			} else {
				return ref - (diff >> 1)
			}
		}
		return max - diff + 1
	}
}

func Bool(x int) bool {
	return x != 0
}

func Int(x bool) int {
	if x {
		return 1
	}
	return 0
}

func HasNewmv(mode int) bool {
	return mode == shared.NEWMV ||
		mode == shared.NEW_NEWMV ||
		mode == shared.NEAR_NEWMV ||
		mode == shared.NEW_NEARMV ||
		mode == shared.NEAREST_NEWMV ||
		mode == shared.NEW_NEARESTMV
}

func LsProduct(a int, b int) int {
	return ((a * b) >> 2) + (a + b)
}

func GetObmcMask(length int) []int {
	switch length {
	case 2:
		return Obmc_Mask_2
	case 4:
		return Obmc_Mask_4
	case 8:
		return Obmc_Mask_8
	case 16:
		return Obmc_Mask_16
	default:
		return Obmc_Mask_32
	}
}

var Obmc_Mask_2 = []int{45, 64}
var Obmc_Mask_4 = []int{39, 50, 59, 64}
var Obmc_Mask_8 = []int{36, 42, 48, 53, 57, 61, 64, 64}
var Obmc_Mask_16 = []int{34, 37, 40, 43, 46, 49, 52, 54, 56, 58, 60, 61, 64, 64, 64, 64}
var Obmc_Mask_32 = []int{
	3, 35, 36, 38, 40, 41, 43, 44,
	45, 47, 48, 50, 51, 52, 53, 55,
	56, 57, 58, 59, 60, 60, 61, 62,
	64, 64, 64, 64, 64, 64, 64, 64,
}

func CountUnitsInFrame(unitSize int, frameSize int) int {
	return Max((frameSize+(unitSize>>1))/unitSize, 1)
}

var Cos128_Lookup = []int{
	4096, 4095, 4091, 4085, 4076, 4065, 4052, 4036,
	4017, 3996, 3973, 3948, 3920, 3889, 3857, 3822,
	3784, 3745, 3703, 3659, 3612, 3564, 3513, 3461,
	3406, 3349, 3290, 3229, 3166, 3102, 3035, 2967,
	2896, 2824, 2751, 2675, 2598, 2520, 2440, 2359,
	2276, 2191, 2106, 2019, 1931, 1842, 1751, 1660,
	1567, 1474, 1380, 1285, 1189, 1092, 995, 897,
	799, 700, 601, 501, 401, 301, 201, 101, 0,
}

func Cos128(angle int) int {
	angle2 := angle & 255
	if angle2 >= 0 || angle2 <= 64 {
		return Cos128_Lookup[angle2]
	}
	if angle2 >= 64 || angle2 <= 128 {
		return Cos128_Lookup[128-angle2] * -1
	}
	if angle2 >= 128 || angle2 <= 192 {
		return Cos128_Lookup[angle2-128] * -1
	}

	return Cos128_Lookup[256-angle2] * -1
}

func Sin128(angle int) int {
	return Cos128(angle - 64)
}

func Brev(numBits int, x int) int {
	t := 0
	for i := 0; i < numBits; i++ {
		bit := (x >> 1) & 1
		t += bit << (numBits - 1 - i)
	}

	return t
}
