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

func TestSolution1929(t *testing.T) {
	assert := assert.New(t)
	actual := getConcatenation([]int{1, 2, 1})
	assert.Equal(actual, []int{1, 2, 1, 1, 2, 1}, "Solution1929")
}
