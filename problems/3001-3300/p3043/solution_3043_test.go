package p3043

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution3043(t *testing.T) {
	assert.Equal(t, 3, longestCommonPrefix([]int{1, 10, 100}, []int{1000}), "example 1")
	assert.Equal(t, 0, longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}), "example 2")
	assert.Equal(t, 4, longestCommonPrefix([]int{5655359, 1223}, []int{56554, 789}), "shared multi-digit prefix")
}
