package solutions

import "math/rand"

func ConstructorPickIndex(w []int) PickIdx {
	prefix_sum := make([]int, 0)

	total_sum := 0
	for _, weight := range w {
		total_sum += weight
		prefix_sum = append(prefix_sum, total_sum)
	}

	return PickIdx{prefix_sum, total_sum}
}

func (this *PickIdx) PickIndex() int {
	target := rand.Intn(this.total_sum)
	for i, ps_sum := range this.prefix_sum {
		if target < ps_sum {
			return i
		}
	}
	return -1
}
