package p90

import (
	"github.com/stretchr/testify/assert"
	"localhost/leetcode/util"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution90(t *testing.T) {
	assert := assert.New(t)
	actual := subsetsWithDup([]int{1, 2, 2})
	expected := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	assert.Equal(len(actual), len(expected), "Solution90")
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if util.IntSliceEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution90")
	}
}
