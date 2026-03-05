package p2453

import "github.com/eseliavka/go/util"

func destroyTargets(nums []int, space int) int {
	d := make(map[int]int)

	for _, num := range nums {
		d[num%space]++
	}

	max := util.MinInt

	for _, val := range d {
		if max < val {
			max = val
		}
	}

	ans := util.MaxInt

	for _, num := range nums {
		if d[num%space] == max {
			ans = util.MinInts([]int{ans, num})
		}
	}
	return ans
}
