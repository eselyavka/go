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

func TestSolution938(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 10, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{7, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{18, nil, nil}
	assert.Equal(rangeSumBST(&root, 7, 15), 32, "Solution938")
}
