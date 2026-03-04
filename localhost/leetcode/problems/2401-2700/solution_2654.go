package solutions

func minOperations_2654(nums []int) int {
	n := len(nums)

	ones := 0
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}

	if ones > 0 {
		return n - ones
	}

	best := MaxInt

	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i + 1; j < n; j++ {
			g = GCD(g, nums[j])
			if g == 1 {
				best = MinInts([]int{best, j - i + 1})
				break
			}
		}
	}

	if best == MaxInt {
		return -1
	}

	return (best - 1) + (n - 1)
}
