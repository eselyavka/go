package p199

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

func TestSolution199(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 4, Left: nil, Right: nil}

	actual := rightSideView(&root)
	assert.Equal([]int{1, 3, 4}, actual, "Solution199")
}
