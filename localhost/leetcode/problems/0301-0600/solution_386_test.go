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

func TestSolution386(t *testing.T) {
	assert := assert.New(t)
	actual := largestDivisibleSubset([]int{4, 8, 10, 240})
	assert.Equal(actual, []int{4, 8, 240}, "Solution386")
}
