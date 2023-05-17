package solutions

/*
class Solution(object):

	def pairSum(self, head):
	    """
	    :type head: Optional[ListNode]
	    :rtype: int
	    """
	    fast = slow = head

	    while slow and fast and fast.next:
	        slow = slow.next
	        fast = fast.next.next

	    curr, prev = slow, None
	    while curr:
	        curr.next, prev, curr = prev, curr, curr.next

	    ans = 0
	    while prev:
	        ans = max(ans, head.val + prev.val)
	        head = head.next
	        prev = prev.next

	    return ans
*/
func pairSum(head *ListNode) int {
	fast := head
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow
	var prev *ListNode = nil

	for curr != nil {
		buf := curr.Next
		curr.Next = prev
		prev = curr
		curr = buf
	}

	ans := 0

	for prev != nil {
		ans = MaxInts([]int{ans, head.Val + prev.Val})
		head = head.Next
		prev = prev.Next
	}

	return ans
}
