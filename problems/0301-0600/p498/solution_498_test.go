package p498

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

func TestSolution498(t *testing.T) {
	assert := assert.New(t)
	res := findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	assert.Equal(res, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}, "Solution498")
}
