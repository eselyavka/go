package p2265

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

func TestSolution2265(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 8, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 6, Left: nil, Right: nil}

	actual := averageOfSubtree(&root)
	assert.Equal(actual, 5, "Solution2265")
}
