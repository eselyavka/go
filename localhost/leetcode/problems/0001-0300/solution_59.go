package solutions

func generateMatrix(n int) [][]int {
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		buf := make([]int, n, n)
		ans = append(ans, buf)
	}

	i := 2
	prev_x := 0
	prev_y := 0

	ans[prev_x][prev_y] = 1

	for i <= (n * n) {
		// right
		for true {
			new_x := MinInts([]int{prev_x, n - 1})
			new_y := MinInts([]int{prev_y + 1, n - 1})
			if ans[new_x][new_y] == 0 {
				ans[new_x][new_y] = i
				i++
				prev_x, prev_y = new_x, new_y
			} else {
				break
			}
		}
		// down
		for true {
			new_x := MinInts([]int{prev_x + 1, n - 1})
			new_y := MinInts([]int{prev_y, n - 1})
			if ans[new_x][new_y] == 0 {
				ans[new_x][new_y] = i
				i++
				prev_x, prev_y = new_x, new_y
			} else {
				break
			}
		}
		// left
		for true {
			new_x := MinInts([]int{prev_x, n - 1})
			new_y := MaxInts([]int{prev_y - 1, 0})
			if ans[new_x][new_y] == 0 {
				ans[new_x][new_y] = i
				i++
				prev_x, prev_y = new_x, new_y
			} else {
				break
			}
		}
		// up
		for true {
			new_x := MinInts([]int{prev_x - 1, n - 1})
			new_y := MaxInts([]int{prev_y, 0})
			if ans[new_x][new_y] == 0 {
				ans[new_x][new_y] = i
				i++
				prev_x, prev_y = new_x, new_y
			} else {
				break
			}
		}
	}

	return ans
}
