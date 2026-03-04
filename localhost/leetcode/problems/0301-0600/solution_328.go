package solutions

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd_s := ListNode{Val: MinInt, Next: nil}
	odd := &odd_s
	even_s := ListNode{Val: MinInt, Next: nil}
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
