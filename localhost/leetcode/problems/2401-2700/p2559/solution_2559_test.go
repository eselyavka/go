package p2559

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

func TestSolution2559(t *testing.T) {
	assert := assert.New(t)
	actual := vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}})
	assert.Equal(actual, []int{2, 3, 0}, "Solution2559")
}
