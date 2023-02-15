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
