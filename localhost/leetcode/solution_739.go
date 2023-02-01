package solutions

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	stack := make([]int, 0)
	stack = append(stack, 0)
	ans := make([]int, n, n)

	for i := 1; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] = i - idx
		}
		stack = append(stack, i)
	}

	return ans
}
