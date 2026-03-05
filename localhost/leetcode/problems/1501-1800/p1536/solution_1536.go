package p1536

func minSwaps(grid [][]int) int {
	n := len(grid)
	tz := make([]int, n)

	for i := 0; i < n; i++ {
		traiZeros := 0
		for j := n - 1; j > -1; j-- {
			if grid[i][j] == 1 {
				break
			}
			traiZeros++
		}
		tz[i] = traiZeros
	}

	ans := 0
	var idx int
	for i := 0; i < n-1; i++ {
		req := n - i - 1
		idx = -1

		for j := i; j < n; j++ {
			if tz[j] >= req {
				idx = j
				break
			}
		}
		if idx == -1 {
			return -1
		}

		for idx > i {
			ans++
			tmp := tz[idx]
			tz[idx] = tz[idx-1]
			tz[idx-1] = tmp
			idx--
		}
	}
	return ans
}
