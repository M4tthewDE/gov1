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

func Round2(x int, n int) int {
	if n == 0 {
		return x
	}

	return (x + (1 << (n - 1))) >> n
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
