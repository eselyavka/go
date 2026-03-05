package p2145

import "github.com/eseliavka/go/util"

func numberOfArrays(differences []int, lower int, upper int) int {
	ps := 0
	maxPs := 0
	minPs := 0

	for _, difference := range differences {
		ps += difference
		maxPs = util.MaxInts([]int{maxPs, ps})
		minPs = util.MinInts([]int{minPs, ps})
	}

	low := lower - minPs
	hi := upper - maxPs

	return util.MaxInts([]int{0, hi - low + 1})
}
