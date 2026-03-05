package p53

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

func TestSolution53(t *testing.T) {
	assert := assert.New(t)
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	actual := maxSubArray(nums)
	assert.Equal(actual, 6, "Solution53")
}
