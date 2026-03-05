package solutions

import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
	acc := make([]int, 0)
	for _, list := range lists {
		it := list
		for it != nil {
			acc = append(acc, it.Val)
			it = it.Next
		}

	}

	ans := &ListNode{Val: -1, Next: nil}
	fake := ans

	sort.Ints(acc)

	for _, val := range acc {
		node := &ListNode{Val: val, Next: nil}
		fake.Next = node
		fake = fake.Next
	}

	return ans.Next
}
