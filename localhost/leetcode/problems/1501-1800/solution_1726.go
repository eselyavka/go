package solutions

func tupleSameProduct(nums []int) int {
	n := len(nums)

	if n < 3 {
		return 0
	}

	products := make(map[int]map[int]struct{})
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			if nums[i] != nums[j] {
				if _, exists := products[product]; exists {
					products[product][nums[i]] = struct{}{}
					products[product][nums[j]] = struct{}{}
				} else {
					products[product] = map[int]struct{}{nums[i]: {}, nums[j]: {}}
				}
			}
		}
	}
	ans := 0
	for _, product := range products {
		tupleLen := len(product)
		if tupleLen >= 4 {
			ans += tupleLen * (tupleLen - 2)
		}
	}

	return ans
}
