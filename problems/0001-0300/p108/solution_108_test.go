package p108

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution108(t *testing.T) {
	assert := assert.New(t)
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	actual := util.BinaryTreeBFS(root)
	assert.Equal(actual, []int{0, -10, 5, -3, 9}, "Solution108")
}
