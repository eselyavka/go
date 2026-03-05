package p2275

import "github.com/eseliavka/go/util"

func largestCombination(candidates []int) int {
	ans := 0
	for b := 0; b < 32; b++ {
		cnt := 0
		mask := 1 << b
		for _, x := range candidates {
			if x&mask > 0 {
				cnt++
			}
		}
		ans = util.MaxInts([]int{ans, cnt})
	}

	return ans
}
