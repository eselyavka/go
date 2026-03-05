package p876

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

func TestSolution876(t *testing.T) {
	assert := assert.New(t)
	head := util.ListNode{Val: 1, Next: nil}
	head.Next = &util.ListNode{Val: 2, Next: nil}
	head.Next.Next = &util.ListNode{Val: 3, Next: nil}
	head.Next.Next.Next = &util.ListNode{Val: 4, Next: nil}
	head.Next.Next.Next.Next = &util.ListNode{Val: 5, Next: nil}

	assert.Equal(middleNode(&head), head.Next.Next, "Solution876")
}
