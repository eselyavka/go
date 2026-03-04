package solutions

import "sort"

func threeSumSmaller(nums []int, target int) int {
	n := len(nums)

	if n < 3 {
		return 0
	}

	sort.Ints(nums)

	ans := 0

	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] >= target {
				k--
			} else {
				ans += k - j
				j++
			}
		}
	}

	return ans
}
