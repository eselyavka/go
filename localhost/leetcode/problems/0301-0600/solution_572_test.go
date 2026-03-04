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

func TestSolution572(t *testing.T) {
	assert := assert.New(t)

	root1 := TreeNode{Val: 3, Left: nil, Right: nil}
	root1.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root1.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root1.Left.Right = &TreeNode{Val: 2, Left: nil, Right: nil}
	root1.Right = &TreeNode{Val: 5, Left: nil, Right: nil}

	root2 := TreeNode{Val: 4, Left: nil, Right: nil}
	root2.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root2.Right = &TreeNode{Val: 2, Left: nil, Right: nil}

	actual := isSubtree(&root1, &root2)

	assert.True(actual, "Solution572")

}
