package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAssignInBounds(t *testing.T) {
	test := []int{0, 1, 2, 3}
	test = SliceAssign(test, 0, 1)

	assert.Equal(t, 4, len(test))
	assert.Equal(t, 1, test[0])
}

func TestSliceAssignOutOfBounds(t *testing.T) {
	test := []int{}

	test = SliceAssign(test, 2, 10)

	assert.Equal(t, 3, len(test))
	assert.Equal(t, 0, test[0])
	assert.Equal(t, 0, test[1])
	assert.Equal(t, 10, test[2])
}
