package p3195

import (
	"localhost/leetcode/util"
	"math"
)

func minimumArea(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	xMin := util.MaxInt
	yMin := util.MaxInt
	xMax := util.MinInt
	yMax := util.MinInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				xMin = util.MinInts([]int{xMin, i})
				yMin = util.MinInts([]int{yMin, j})
				xMax = util.MaxInts([]int{xMax, i})
				yMax = util.MaxInts([]int{yMax, j})
			}
		}
	}

	ans := math.Abs(float64(xMax-xMin)+1) * math.Abs(float64(yMax-yMin)+1)
	return int(ans)
}
