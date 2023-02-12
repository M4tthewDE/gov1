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
