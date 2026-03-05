package p2033

import (
	"localhost/leetcode/util"
	"math"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	flattenArr := util.Flatten2DArray(grid)
	rem := flattenArr[0] % x

	for i := 1; i < len(flattenArr); i++ {
		if flattenArr[i]%x != rem {
			return -1
		}
	}

	sort.Ints(flattenArr)

	midIdx := len(flattenArr) / 2

	goal := flattenArr[midIdx]
	ops := 0

	for i := 0; i < len(flattenArr); i++ {
		ops += int(math.Abs(float64(goal-flattenArr[i]))) / x
	}

	return ops
}
