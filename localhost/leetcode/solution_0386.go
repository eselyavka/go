package solutions

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 0)
		dp[i] = append(dp[i], nums[i])
	}

	res := make([]int, 0)
	for j := n - 1; j > -1; j-- {
		for k := j + 1; k < n; k++ {
			if nums[k]%nums[j] == 0 {
				tmp := []int{nums[j]}
				for _, item := range dp[k] {
					tmp = append(tmp, item)
				}
				if len(tmp) > len(dp[j]) {
					dp[j] = tmp
				}
			}
		}
		if len(dp[j]) > len(res) {
			res = dp[j]
		}
	}

	return res
}
