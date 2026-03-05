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

func TestSolution1574(t *testing.T) {
	assert := assert.New(t)
	actual := findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})
	assert.Equal(actual, 3, "Solution1574")
}
