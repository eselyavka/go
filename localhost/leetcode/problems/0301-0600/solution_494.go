package solutions

func targetWays(i, n, totalSum, currSum, target int, nums []int, memo [][]int) int {
	if currSum == target && i == n {
		return 1
	}

	if i >= n {
		return 0
	}

	if memo[i][currSum+totalSum] != MinInt {
		return memo[i][currSum+totalSum]
	}

	add := targetWays(i+1, n, totalSum, currSum+nums[i], target, nums, memo)
	subtract := targetWays(i+1, n, totalSum, currSum-nums[i], target, nums, memo)
	memo[i][currSum+totalSum] = add + subtract

	return memo[i][currSum+totalSum]
}

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	totalSum := sum(nums)

	memo := make([][]int, 0)
	for i := 0; i < n; i++ {
		arr := make([]int, 2*totalSum+1)
		for j := 0; j < 2*totalSum+1; j++ {
			arr[j] = MinInt
		}
		memo = append(memo, arr)
	}

	return targetWays(0, n, totalSum, 0, target, nums, memo)
}
