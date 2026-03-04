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

func TestSolution94(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 3, Left: nil, Right: nil}

	assert.Equal(inorderTraversal(&root), []int{1, 3, 2}, "Solution94")
}
