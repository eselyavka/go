package p1578

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

func TestSolution1578(t *testing.T) {
	assert := assert.New(t)
	actual := minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1})
	assert.Equal(actual, 26, "SolutionOther1578")
}
