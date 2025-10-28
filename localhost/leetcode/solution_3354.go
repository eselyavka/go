package solutions

func countValidSelections(nums []int) int {
	n := len(nums)
	total := sum(nums)

	simulate := func(start, step int) bool {
		arr := make([]int, n)
		copy(arr, nums)
		rem := total
		i := start
		s := step
		for i >= 0 && i < n {
			if arr[i] != 0 {
				arr[i]--
				rem -= 1
				s = -1 * s
			}
			i += s
		}
		return rem == 0
	}

	ans := 0

	for idx, x := range nums {
		if x == 0 {
			if simulate(idx, -1) {
				ans++
			}
			if simulate(idx, 1) {
				ans++
			}
		}
	}

	return ans
}
