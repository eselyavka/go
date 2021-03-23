package solutions

import "math"

func maxProfit(prices []int) int {
	res := float64(0)
	min_so_far := math.Pow(10, 4) + 1
	for _, price := range prices {
		min_so_far = math.Min(float64(min_so_far), float64(price))
		max_so_far := float64(price) - min_so_far
		res = math.Max(res, max_so_far)
	}
	return int(res)
}
