package p94

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

func TestSolution94(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}

	assert.Equal(inorderTraversal(&root), []int{1, 3, 2}, "Solution94")
}
