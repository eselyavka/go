package p155

import "localhost/leetcode/util"

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	obj := MinStack{stack: make([]int, 0), min: make([]int, 0)}

	return obj
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min) > 0 {
		this.min = append(this.min, util.MinInts([]int{val, this.min[len(this.min)-1]}))
	} else {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() int {
	this.min = this.min[:len(this.min)-1]
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	return val
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
