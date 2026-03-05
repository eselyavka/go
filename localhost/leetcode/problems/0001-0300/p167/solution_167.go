package p167

import "localhost/leetcode/util"

func twoSum_167(numbers []int, target int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		diff := target - numbers[i]
		idx := util.BinarySearch(0, n-1, numbers, diff)
		if idx >= 0 && idx != i {
			return []int{util.MinInts([]int{i, idx}) + 1, util.MaxInts([]int{i, idx}) + 1}
		}
	}

	return []int{}
}
