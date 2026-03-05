package solutions

func minimumSwaps(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	max_pos := make([]int, 2, 2)
	min_pos := make([]int, 2, 2)

	max := MinInt
	min := MaxInt

	for idx, num := range nums {
		if num >= max {
			max_pos = []int{idx, num}
			max = num
		}

		if min > num {
			min_pos = []int{idx, num}
			min = num
		}
	}

	i := max_pos[0] + 1
	ans := 0

	for i < n {
		if intSliceEqual([]int{i, nums[i]}, min_pos) {
			min_pos = []int{i - 1, nums[i]}
		}

		nums[i-1], nums[i] = nums[i], nums[i-1]
		ans++
		i++
	}

	ans += min_pos[0]

	return ans
}
