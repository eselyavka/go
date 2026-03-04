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

func TestSolution59(t *testing.T) {
	assert := assert.New(t)
	actual := generateMatrix(3)
	assert.Equal(actual, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, "Solution59")
}
