package p1669

import "github.com/eseliavka/go/util"

func mergeInBetween(list1 *util.ListNode, a int, b int, list2 *util.ListNode) *util.ListNode {
	prevA := list1

	for i := 0; i < a-1; i++ {
		prevA = prevA.Next
	}

	afterB := prevA
	for i := 0; i < b-a+2; i++ {
		afterB = afterB.Next
	}

	tail2 := list2
	for tail2.Next != nil {
		tail2 = tail2.Next
	}

	prevA.Next = list2
	tail2.Next = afterB

	return list1
}
