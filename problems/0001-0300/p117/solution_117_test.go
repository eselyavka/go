package p117

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

func TestSolution117(t *testing.T) {
	assert := assert.New(t)
	root := util.NodeNext{Val: 1, Left: nil, Right: nil, Next: nil}
	root.Left = &util.NodeNext{Val: 2, Left: nil, Right: nil, Next: nil}
	root.Left.Left = &util.NodeNext{Val: 4, Left: nil, Right: nil, Next: nil}
	root.Left.Right = &util.NodeNext{Val: 5, Left: nil, Right: nil, Next: nil}
	root.Right = &util.NodeNext{Val: 3, Left: nil, Right: nil, Next: nil}
	root.Right.Right = &util.NodeNext{Val: 7, Left: nil, Right: nil, Next: nil}

	_ = connect(&root)

	q := make([]*util.NodeNext, 0)
	q = append(q, &root)

	ans := make([]int, 0)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Next != nil {
			ans = append(ans, node.Next.Val)
		}
	}

	assert.Equal([]int{3, 5, 7}, ans, "Solution177")
}
