package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	test := []int{7, 5, 8, 3, 10}

	test = Sort(test, 1, 3)
	assert.Equal(t, []int{7, 3, 5, 8, 10}, test)
}
