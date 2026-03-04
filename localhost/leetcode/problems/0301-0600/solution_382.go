package solutions

import "math/rand"

type Solution struct {
	h      *ListNode
	length int
}

func Constructor_382(head *ListNode) Solution {
	it := head
	length := 0
	for it != nil {
		length++
		it = it.Next
	}

	return Solution{h: head, length: length}

}

func (this *Solution) GetRandom() int {
	rnd := rand.Intn(this.length)
	curr := 0
	it := this.h
	for it != nil {
		if curr == rnd {
			return it.Val
		}
		it = it.Next
		curr++
	}

	return this.h.Val
}
