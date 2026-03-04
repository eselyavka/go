package solutions

import "math"

func minimumArea(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	xMin := MaxInt
	yMin := MaxInt
	xMax := MinInt
	yMax := MinInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				xMin = MinInts([]int{xMin, i})
				yMin = MinInts([]int{yMin, j})
				xMax = MaxInts([]int{xMax, i})
				yMax = MaxInts([]int{yMax, j})
			}
		}
	}

	ans := math.Abs(float64(xMax-xMin)+1) * math.Abs(float64(yMax-yMin)+1)
	return int(ans)
}
