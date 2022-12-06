package solutions

/*
class Solution(object):
    def oddEvenList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return

        odd_s = odd = ListNode("-inf")
        even_s = even = ListNode("-inf")

        curr = head
        i = 0
        while curr:
            if i % 2 == 1:
                odd.next = curr
                odd = odd.next
            else:
                even.next = curr
                even = even.next

            i += 1
            buf = curr.next
            curr.next = None
            curr = buf

        even.next=odd_s.next
        return even_s.next

*/

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
