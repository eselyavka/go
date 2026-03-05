package p328

import "github.com/eseliavka/go/util"

func oddEvenList(head *util.ListNode) *util.ListNode {
	if head == nil {
		return nil
	}
	odd_s := util.ListNode{Val: util.MinInt, Next: nil}
	odd := &odd_s
	even_s := util.ListNode{Val: util.MinInt, Next: nil}
	even := &even_s

	curr := head
	idx := 0
	for curr != nil {
		if idx%2 == 1 {
			odd.Next = curr
			odd = odd.Next
		} else {
			even.Next = curr
			even = even.Next
		}
		idx++
		buf := curr.Next
		curr.Next = nil
		curr = buf
	}
	even.Next = odd_s.Next

	return even_s.Next
}
