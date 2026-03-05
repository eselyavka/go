package p2471

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

func TestSolution2471(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 7, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left.Left = &util.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right.Left = &util.TreeNode{Val: 10, Left: nil, Right: nil}

	actual := minimumOperations(&root)
	assert.Equal(actual, 3, "Solution2471")
}
