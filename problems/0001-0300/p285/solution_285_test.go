package p285

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

func TestSolution285(t *testing.T) {
	assert := assert.New(t)
	root := util.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &util.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}
	actual := inorderSuccessor(&root, root.Left)
	assert.Equal(2, actual.Val, "Solution285")
}
