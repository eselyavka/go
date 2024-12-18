package solutions

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	stack := make([]int, 0)
	n := len(prices)

	copy(ans, prices)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= prices[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] -= prices[i]
		}
		stack = append(stack, i)
	}

	return ans
}
