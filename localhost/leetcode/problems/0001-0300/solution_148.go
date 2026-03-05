package solutions

import "sort"

func sortList(head *ListNode) *ListNode {
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
	new_head := ListNode{Val: MaxInt, Next: nil}
	tail := &new_head
	for _, num := range arr {
		tail.Next = &ListNode{Val: num, Next: nil}
		tail = tail.Next
	}

	return new_head.Next
}
