package solutions

func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return MaxInts(nums)
	}

	if n == 3 {
		return MaxInts(nums)
	}

	dp1 := make([]int, n-1)
	dp1[0] = nums[0]
	dp2 := make([]int, n-1)
	dp2[0] = nums[1]

	for i := 1; i < n-1; i++ {
		num := nums[i]
		if i >= 2 {
			num += dp1[i-2]
		}
		if num > dp1[i-1] {
			dp1[i] = num
		} else {
			dp1[i] = dp1[i-1]
		}
	}

	arr := nums[1:]
	for i := 1; i < n-1; i++ {
		num := arr[i]
		if i >= 2 {
			num += dp2[i-2]
		}
		if num > dp2[i-1] {
			dp2[i] = num
		} else {
			dp2[i] = dp2[i-1]
		}
	}

	return MaxInts([]int{dp1[n-2], dp2[n-2]})
}
