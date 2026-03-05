package solutions

import "sort"

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)

	res := 0

	for _, n := range nums {
		if n > nums[res] {
			res++
		}
	}

	return res
}
