package p769

import "github.com/eseliavka/go/util"

func maxChunksToSorted(arr []int) int {
	n := len(arr)

	ans := 0
	maxSoFar := -1

	for i := 0; i < n; i++ {
		maxSoFar = util.MaxInts([]int{maxSoFar, arr[i]})
		if maxSoFar == i {
			ans++
		}
	}

	return ans
}
