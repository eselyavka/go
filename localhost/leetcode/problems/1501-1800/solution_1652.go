package solutions

import "math"

func decrypt(code []int, k int) []int {
	n := len(code)

	if k == 0 {
		return make([]int, n, n)
	}

	res := make([]int, 0)
	if k > 0 {
		for i := 0; i < n; i++ {
			j := 1
			acc := 0
			for l := 0; l < k; l++ {
				idx := (i + j) % n
				j++
				acc += code[idx]
			}
			res = append(res, acc)
		}
	}
	if k < 0 {
		for i := 0; i < n; i++ {
			j := n - 1
			acc := 0
			for l := 0; l < int(math.Abs(float64(k))); l++ {
				idx := (i + j) % n
				j--
				acc += code[idx]
			}
			res = append(res, acc)
		}
	}
	return res
}
