package p3653

func xorAfterQueries(nums []int, queries [][]int) int {
	const mod = 1_000_000_007

	for _, query := range queries {
		l, r, k, v := query[0], query[1], query[2], query[3]
		for idx := l; idx <= r; idx += k {
			nums[idx] = nums[idx] * v % mod
		}
	}

	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	return ans
}
