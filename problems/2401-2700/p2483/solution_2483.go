package p2483

func bestClosingTime(customers string) int {
	n := len(customers)
	var b2i = map[bool]int{false: 0, true: 1}

	prefixSum := make([]int, n+1, n+1)
	var is_absent bool

	for i := 1; i < n+1; i++ {
		is_absent = string(customers[i-1]) == "N"
		prefixSum[i] = b2i[is_absent] + prefixSum[i-1]
	}

	suffixSum := make([]int, n+1, n+1)
	var is_present bool
	for j := n - 1; j >= 0; j-- {
		is_present = string(customers[j]) == "Y"
		suffixSum[j] = b2i[is_present] + suffixSum[j+1]
	}

	ans := [2]int{0, prefixSum[0] + suffixSum[0]}
	for i := 0; i < n+1; i++ {
		if prefixSum[i]+suffixSum[i] < ans[1] {
			ans[0] = i
			ans[1] = prefixSum[i] + suffixSum[i]
		}
	}

	return ans[0]
}
