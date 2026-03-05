package p938

import (
	"github.com/stretchr/testify/assert"
	"localhost/leetcode/util"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution938(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 10, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{3, nil, nil}
	root.Left.Right = &util.TreeNode{7, nil, nil}
	root.Right = &util.TreeNode{15, nil, nil}
	root.Right.Right = &util.TreeNode{18, nil, nil}
	assert.Equal(rangeSumBST(&root, 7, 15), 32, "Solution938")
}
