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

func TestSolution3314(t *testing.T) {
	assert := assert.New(t)
	actual := minBitwiseArray([]int{2, 3, 5, 7})
	assert.Equal(actual, []int{-1, 1, 4, 3}, "SolutionOther3314")
}
