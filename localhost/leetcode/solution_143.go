package solutions

func reorderList(head *ListNode) {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var second, next, prev *ListNode

	second = slow.Next
	slow.Next = nil

	for second != nil {
		next = second.Next
		second.Next = prev
		prev = second
		second = next
	}

	first := head
	second = prev
	var tmp1, tmp2 *ListNode
	for second != nil {
		tmp1 = first.Next
		tmp2 = second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}

/*
   arr = []

   it = head

   while it:
       arr.append(it.val)
       it = it.next

   l = 0
   r = len(arr) - 1

   ans = head
   i=0
   while ans:
       if i % 2 == 0:
           ans.val = arr[l]
           l +=1
       else:
           ans.val = arr[r]
           r -= 1
       i+=1
       ans = ans.next
   return head
*/
func reorderListMem(head *ListNode) {
	arr := make([]int, 0)
	it := head
	for it != nil {
		arr = append(arr, it.Val)
		it = it.Next
	}

	l := 0
	r := len(arr) - 1

	ans := head
	i := 0
	for ans != nil {
		if i%2 == 0 {
			ans.Val = arr[l]
			l++
		} else {
			ans.Val = arr[r]
			r--
		}
		i++
		ans = ans.Next
	}
}
