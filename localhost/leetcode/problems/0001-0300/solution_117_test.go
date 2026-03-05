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

func TestSolution117(t *testing.T) {
	assert := assert.New(t)
	root := NodeNext{Val: 1, Left: nil, Right: nil, Next: nil}
	root.Left = &NodeNext{Val: 2, Left: nil, Right: nil, Next: nil}
	root.Left.Left = &NodeNext{Val: 4, Left: nil, Right: nil, Next: nil}
	root.Left.Right = &NodeNext{Val: 5, Left: nil, Right: nil, Next: nil}
	root.Right = &NodeNext{Val: 3, Left: nil, Right: nil, Next: nil}
	root.Right.Right = &NodeNext{Val: 7, Left: nil, Right: nil, Next: nil}

	_ = connect(&root)

	q := make([]*NodeNext, 0)
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
