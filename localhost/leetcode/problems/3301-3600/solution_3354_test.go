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

func TestSolution3354(t *testing.T) {
	assert := assert.New(t)
	actual := countValidSelections([]int{1, 0, 2, 0, 3})
	assert.Equal(actual, 2, "SolutionOther3354")
}
