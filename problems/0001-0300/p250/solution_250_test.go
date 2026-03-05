package p250

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

func TestSolution250(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 5, Left: nil, Right: nil}

	actual := countUnivalSubtrees(&root)

	assert.Equal(4, actual, "Solution250")
}
