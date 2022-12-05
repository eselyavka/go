package solutions

func bestClosingTime(customers string) int {
	n := len(customers)
	var b2i = map[bool]int{false: 0, true: 1}

	prefix_sum := make([]int, n+1, n+1)
	var is_absent bool

	for i := 1; i < n+1; i++ {
		is_absent = string(customers[i-1]) == "N"
		prefix_sum[i] = b2i[is_absent] + prefix_sum[i-1]
	}

	suffix_sum := make([]int, n+1, n+1)
	var is_present bool
	for j := n - 1; j >= 0; j-- {
		is_present = string(customers[j]) == "Y"
		suffix_sum[j] = b2i[is_present] + suffix_sum[j+1]
	}

	ans := [2]int{0, prefix_sum[0] + suffix_sum[0]}
	for i := 0; i < n+1; i++ {
		if prefix_sum[i]+suffix_sum[i] < ans[1] {
			ans[0] = i
			ans[1] = prefix_sum[i] + suffix_sum[i]
		}
	}

	return ans[0]
}
