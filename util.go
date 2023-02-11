package main

func SliceAssign[T any](s []T, i int, v T) []T {
	for i >= len(s) {
		s = append(s, make([]T, 1)...)
	}

	s[i] = v
	return s
}
