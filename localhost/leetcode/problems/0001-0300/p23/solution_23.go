package p23

import (
	"localhost/leetcode/util"
	"sort"
)

func mergeKLists(lists []*util.ListNode) *util.ListNode {
	acc := make([]int, 0)
	for _, list := range lists {
		it := list
		for it != nil {
			acc = append(acc, it.Val)
			it = it.Next
		}

	}

	ans := &util.ListNode{Val: -1, Next: nil}
	fake := ans

	sort.Ints(acc)

	for _, val := range acc {
		node := &util.ListNode{Val: val, Next: nil}
		fake.Next = node
		fake = fake.Next
	}

	return ans.Next
}
