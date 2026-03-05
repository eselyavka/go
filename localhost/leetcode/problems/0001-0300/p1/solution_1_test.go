package p1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum_1([]int{2, 7, 11, 15}, 9), "Solution1")
}
