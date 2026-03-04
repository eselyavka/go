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

func TestSolution4(t *testing.T) {
	assert := assert.New(t)
	res := findMedianSortedArrays([]int{1, 2}, []int{3})
	assert.Equal(res, 2.0, "Solution4")
}
