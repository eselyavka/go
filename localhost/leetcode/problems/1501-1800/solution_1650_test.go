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

func TestSolution1650(t *testing.T) {
	assert := assert.New(t)
	root := Node{Val: 3, Left: nil, Right: nil, Parent: nil}
	root.Left = &Node{Val: 5, Left: nil, Right: nil, Parent: &root}
	root.Left.Left = &Node{Val: 6, Left: nil, Right: nil, Parent: root.Left}
	root.Left.Right = &Node{Val: 2, Left: nil, Right: nil, Parent: root.Left}
	root.Left.Right.Left = &Node{Val: 7, Left: nil, Right: nil, Parent: root.Left.Right}
	root.Left.Right.Right = &Node{Val: 4, Left: nil, Right: nil, Parent: root.Left.Right}
	root.Right = &Node{Val: 1, Left: nil, Right: nil, Parent: &root}
	root.Right.Left = &Node{Val: 0, Left: nil, Right: nil, Parent: root.Right}
	root.Right.Right = &Node{Val: 8, Left: nil, Right: nil, Parent: root.Right}

	p := root.Left
	q := root.Left.Right.Right

	assert.Equal(lowestCommonAncestor(p, q), root.Left, "Solution1650")
}
