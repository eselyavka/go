package solutions

func trap(height []int) int {
	n := len(height)
	left := make([]int, n, n)
	right := make([]int, n, n)

	left[0] = height[0]
	right[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		left[i] = MaxInts([]int{left[i-1], height[i]})
	}

	for j := n - 2; j > 0; j-- {
		right[j] = MaxInts([]int{right[j+1], height[j]})
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += MinInts([]int{left[i], right[i]}) - height[i]
	}

	return ans
}
