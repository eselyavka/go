package solutions

import "sort"

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	diags := make(map[int][]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			diags[i-j] = append(diags[i-j], grid[i][j])
		}
	}

	for _, value := range diags {
		sort.Ints(value)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i-j >= 0 {
				grid[i][j] = diags[i-j][len(diags[i-j])-1]
				diags[i-j] = diags[i-j][:len(diags[i-j])-1]
			} else {
				grid[i][j] = diags[i-j][0]
				diags[i-j] = diags[i-j][1:]
			}
		}
	}

	return grid
}
