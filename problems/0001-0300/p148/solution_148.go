package p148

import (
	"github.com/eseliavka/go/util"
	"sort"
)

func sortList(head *util.ListNode) *util.ListNode {
	if head == nil {
		return head
	}

	arr := make([]int, 0)
	it := head
	for it != nil {
		arr = append(arr, it.Val)
		it = it.Next
	}
	sort.Ints(arr)
	new_head := util.ListNode{Val: util.MaxInt, Next: nil}
	tail := &new_head
	for _, num := range arr {
		tail.Next = &util.ListNode{Val: num, Next: nil}
		tail = tail.Next
	}

	return new_head.Next
}
