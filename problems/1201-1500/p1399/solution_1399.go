package p1399

import "github.com/eseliavka/go/util"

func countLargestGroup(n int) int {
	digitSums := make(map[int]int)
	currMax := 0

	for i := 1; i <= n; i++ {
		s := util.SumDigits(i)
		digitSums[s]++
		currMax = util.MaxInts([]int{currMax, digitSums[s]})
	}

	ans := 0
	for _, count := range digitSums {
		if count == currMax {
			ans++
		}
	}
	return ans
}
