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

func TestSolution108(t *testing.T) {
	assert := assert.New(t)
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	actual := binaryTreeBFS(root)
	assert.Equal(actual, []int{0, -10, 5, -3, 9}, "Solution108")
}
