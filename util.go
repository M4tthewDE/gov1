package main

func SliceAssign[T any](s []T, i int, v T) []T {
	for i >= len(s) {
		s = append(s, make([]T, 1)...)
	}

	s[i] = v
	return s
}

func SliceAssignNested[T any](s [][]T, i int, j int, v T) [][]T {
	for i >= len(s) {
		s = append(s, make([]T, 1))
	}

	for j >= len(s[i]) {
		s[i] = append(s[i], make([]T, 1)...)
	}

	s[i][j] = v
	return s
}

func Equals[T comparable](a []T, b []T) bool {
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

// tile_log2( blkSize, target )
func tileLog2(blkSize int, target int) int {
	k := 0
	for (blkSize << k) < target {
		k++
	}
	return k
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
	return mode == NEWMV || mode == NEW_NEWMV || mode == NEAR_NEWMV || mode == NEW_NEARMV || mode == NEAREST_NEWMV || mode == NEW_NEARESTMV
}
