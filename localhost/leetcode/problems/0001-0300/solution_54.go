package solutions

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	ans := make([]int, 0)

	left := 0
	right := n
	top := 0
	bottom := m

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		for i := top; i < bottom; i++ {
			ans = append(ans, matrix[i][right-1])
		}
		right--

		if !(left < right && top < bottom) {
			break
		}

		for i := right - 1; i > left-1; i-- {
			ans = append(ans, matrix[bottom-1][i])
		}
		bottom--

		for i := bottom - 1; i > top-1; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
	}
	return ans
}
