package solutions

import "sort"

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}
	i := 0
	j := 0
	ans := 0

	row_max := make([]int, 0)

	for j < m*n {
		if i == m {
			i = 0
			ans += MaxInts(row_max)
			row_max = make([]int, 0)
		}
		row_max = append(row_max, grid[i][len(grid[i])-1])
		grid[i] = grid[i][:len(grid[i])-1]
		i++
		j++
	}

	ans += MaxInts(row_max)

	return ans
}
