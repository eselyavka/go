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

func TestSolution56(t *testing.T) {
	assert := assert.New(t)
	actual := mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	assert.Equal([][]int{{1, 6}, {8, 10}, {15, 18}}, actual, "Solution56")
}
