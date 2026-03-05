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

func TestSolution39(t *testing.T) {
	assert := assert.New(t)
	actual := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.True(int2dSliceIsEqual(actual, [][]int{{2, 2, 3}, {7}}), "Solution39")
}
