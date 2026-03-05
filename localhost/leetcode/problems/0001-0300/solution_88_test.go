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

func TestSolution88(t *testing.T) {
	assert := assert.New(t)
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var m = 3
	var nums2 = []int{2, 5, 6}
	var n = 3
	merge(nums1, m, nums2, n)
	assert.Equal(nums1, []int{1, 2, 2, 3, 5, 6}, "Solution88")
}
