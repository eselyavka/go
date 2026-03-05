package p2943

import (
	"localhost/leetcode/util"
	"sort"
)

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)

	best := 1
	curr := 1

	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			curr++
		} else {
			curr = 1
		}
		best = util.MaxInts([]int{curr, best})
	}

	hGap := best + 1

	best = 1
	curr = 1

	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			curr++
		} else {
			curr = 1
		}
		best = util.MaxInts([]int{curr, best})
	}

	vGap := best + 1

	ans := util.MinInts([]int{hGap, vGap})

	return ans * ans
}
