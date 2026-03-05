package solutions

func islandPerimeter(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	ans := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				if i == 0 {
					ans++
				}
				if i == row-1 {
					ans++
				}
				if j == 0 {
					ans++
				}
				if j == col-1 {
					ans++
				}
				if 0 <= i-1 && i-1 < row && grid[i-1][j] == 0 {
					ans++
				}
				if 0 <= i+1 && i+1 < row && grid[i+1][j] == 0 {
					ans++
				}
				if 0 <= j-1 && j-1 < col && grid[i][j-1] == 0 {
					ans++
				}
				if 0 <= j+1 && j+1 < col && grid[i][j+1] == 0 {
					ans++
				}
			}
		}
	}

	return ans
}
