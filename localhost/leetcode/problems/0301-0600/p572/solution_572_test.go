package p572

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

func TestSolution572(t *testing.T) {
	assert := assert.New(t)

	root1 := util.TreeNode{Val: 3, Left: nil, Right: nil}
	root1.Left = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root1.Left.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root1.Left.Right = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root1.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}

	root2 := util.TreeNode{Val: 4, Left: nil, Right: nil}
	root2.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root2.Right = &util.TreeNode{Val: 2, Left: nil, Right: nil}

	actual := isSubtree(&root1, &root2)

	assert.True(actual, "Solution572")

}
