package util

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

func TestSliceAssignTwoDimensionsInBounds(t *testing.T) {
	test := [][]int{{0, 1, 2, 3}, {4, 5}}

	test = SliceAssignNested(test, 0, 2, 10)
	test = SliceAssignNested(test, 1, 1, 10)

	assert.Equal(t, 2, len(test))
	assert.Equal(t, 4, len(test[0]))
	assert.Equal(t, 2, len(test[1]))
	assert.Equal(t, []int{0, 1, 10, 3}, test[0])
	assert.Equal(t, []int{4, 10}, test[1])
}

func TestSliceAssignTwoDimensionsOutOfBounds(t *testing.T) {
	test := [][]int{}

	test = SliceAssignNested(test, 0, 2, 10)

	assert.Equal(t, 1, len(test))
	assert.Equal(t, 3, len(test[0]))
	assert.Equal(t, []int{0, 0, 10}, test[0])
	assert.Equal(t, 10, test[0][2])
}

func TestSort(t *testing.T) {
	test := []int{7, 5, 8, 3, 10}

	test = Sort(test, 1, 3)
	assert.Equal(t, []int{7, 3, 5, 8, 10}, test)
}
