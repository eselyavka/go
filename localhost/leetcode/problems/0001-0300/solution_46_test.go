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

func TestSolution46(t *testing.T) {
	assert := assert.New(t)
	actual := permute([]int{1, 2, 3})
	expected := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	var flag bool
	for _, a := range actual {
		flag = false
		for _, b := range expected {
			if intSliceEqual(a, b) {
				flag = true
			}
		}
		assert.True(flag, "Solution46")
	}
}
