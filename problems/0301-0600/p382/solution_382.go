package p382

import (
	"github.com/eseliavka/go/util"
	"math/rand"
)

type Solution struct {
	h      *util.ListNode
	length int
}

func Constructor(head *util.ListNode) Solution {
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
