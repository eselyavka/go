package p528

import (
	"math/rand"
)

type PickIdx struct {
	prefix_sum []int
	total_sum  int
}

func ConstructorPickIndex(w []int) PickIdx {
	prefixSum := make([]int, 0, len(w))

	totalSum := 0
	for _, weight := range w {
		totalSum += weight
		prefixSum = append(prefixSum, totalSum)
	}

	return PickIdx{prefix_sum: prefixSum, total_sum: totalSum}
}

func (this *PickIdx) PickIndex() int {
	target := rand.Intn(this.total_sum)
	for i, psSum := range this.prefix_sum {
		if target < psSum {
			return i
		}
	}
	return -1
}
