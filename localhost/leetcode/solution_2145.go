package solutions

func numberOfArrays(differences []int, lower int, upper int) int {
	ps := 0
	maxPs := 0
	minPs := 0

	for _, difference := range differences {
		ps += difference
		maxPs = MaxInts([]int{maxPs, ps})
		minPs = MinInts([]int{minPs, ps})
	}

	low := lower - minPs
	hi := upper - maxPs

	return MaxInts([]int{0, hi - low + 1})
}
