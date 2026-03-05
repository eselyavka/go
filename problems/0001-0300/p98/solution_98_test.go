package p98

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

func TestSolution98(t *testing.T) {
	assert := assert.New(t)

	root := util.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Left = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right.Right = &util.TreeNode{Val: 7, Left: nil, Right: nil}

	actual := isValidBST(&root)

	assert.False(actual, "Solution98")

}
