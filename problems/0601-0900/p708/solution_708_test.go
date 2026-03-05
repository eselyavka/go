package p708

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

func TestSolution708(t *testing.T) {
	assert := assert.New(t)
	head := util.ListNode{Val: 3, Next: nil}
	head.Next = &util.ListNode{Val: 4, Next: nil}
	head.Next.Next = &util.ListNode{Val: 1, Next: nil}
	head.Next.Next.Next = &head

	res := insert(&head, 2)
	acc := make([]int, 0)

	curr := res.Next
	prev := res

	for curr != res {
		acc = append(acc, prev.Val)
		prev = curr
		curr = curr.Next
	}

	acc = append(acc, prev.Val)

	assert.Equal(acc, []int{3, 4, 1, 2}, "Solution708")
}
