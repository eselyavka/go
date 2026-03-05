package solutions

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution40(t *testing.T) {
	assert := assert.New(t)
	actual := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	expected := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
	assert.Equal(len(actual), len(expected), "Solution40")
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if intSliceEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution40")
	}
}
