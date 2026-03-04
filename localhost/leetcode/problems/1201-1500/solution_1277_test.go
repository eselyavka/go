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

func TestSolution1277(t *testing.T) {
	assert := assert.New(t)
	actual := countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}})
	assert.Equal(actual, 15, "Solution1277")
}
