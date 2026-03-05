package p1650

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

func TestSolution1650(t *testing.T) {
	assert := assert.New(t)
	root := util.Node{Val: 3, Left: nil, Right: nil, Parent: nil}
	root.Left = &util.Node{Val: 5, Left: nil, Right: nil, Parent: &root}
	root.Left.Left = &util.Node{Val: 6, Left: nil, Right: nil, Parent: root.Left}
	root.Left.Right = &util.Node{Val: 2, Left: nil, Right: nil, Parent: root.Left}
	root.Left.Right.Left = &util.Node{Val: 7, Left: nil, Right: nil, Parent: root.Left.Right}
	root.Left.Right.Right = &util.Node{Val: 4, Left: nil, Right: nil, Parent: root.Left.Right}
	root.Right = &util.Node{Val: 1, Left: nil, Right: nil, Parent: &root}
	root.Right.Left = &util.Node{Val: 0, Left: nil, Right: nil, Parent: root.Right}
	root.Right.Right = &util.Node{Val: 8, Left: nil, Right: nil, Parent: root.Right}

	p := root.Left
	q := root.Left.Right.Right

	assert.Equal(lowestCommonAncestor(p, q), root.Left, "Solution1650")
}
