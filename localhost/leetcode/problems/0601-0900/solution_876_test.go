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

func TestSolution876(t *testing.T) {
	assert := assert.New(t)
	head := ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: 3, Next: nil}
	head.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	head.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

	assert.Equal(middleNode(&head), head.Next.Next, "Solution876")
}
