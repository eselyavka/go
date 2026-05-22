package p129

import (
	"testing"

	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
)

func TestSolution129(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}

	actual := sumNumbers(&root)

	assert.Equal(25, actual, "Solution129")
}

func TestSolution129NestedPaths(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 9, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 0, Left: nil, Right: nil}

	actual := sumNumbers(&root)

	assert.Equal(1026, actual, "Solution129")
}

func TestSolution129NilRoot(t *testing.T) {
	assert := assert.New(t)

	actual := sumNumbers(nil)

	assert.Equal(0, actual, "Solution129")
}
