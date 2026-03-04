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

func TestSolution1448(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}

	actual := goodNodes(&root)
	assert.Equal(4, actual, "Solution1448")
}
