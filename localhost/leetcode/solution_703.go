package solutions

import "sort"

type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	instance := KthLargest{k: k, nums: nums}

	return instance
}

func (this *KthLargest) BinarySearch(val int) int {
	l := 0
	r := len(this.nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if this.nums[mid] == val {
			return mid
		}

		if val > this.nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r + 1
}

func (this *KthLargest) Insert(idx int, val int) {
	if len(this.nums) == idx {
		this.nums = append(this.nums, val)
		return
	}

	this.nums = append(this.nums[:idx+1], this.nums[idx:]...)
	this.nums[idx] = val
}

func (this *KthLargest) Add(val int) int {
	idx := this.BinarySearch(val)
	this.Insert(idx, val)

	return this.nums[len(this.nums)-(this.k)]
}
