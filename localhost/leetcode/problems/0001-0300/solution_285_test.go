package solutions

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution285(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	actual := inorderSuccessor(&root, root.Left)
	assert.Equal(2, actual.Val, "Solution285")
}
