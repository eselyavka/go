package p78

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution78(t *testing.T) {
	assert := assert.New(t)
	actual := subsets([]int{1, 2, 3})
	expected := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 2, 3}, {1, 3}, {2, 3}}
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if util.IntSliceEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution78")
	}
}
