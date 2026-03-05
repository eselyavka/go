package solutions

func countGood(nums []int, k int) int64 {
	n := len(nums)
	d := make(map[int]int)
	running_pairs := 0
	ans := 0
	r := 0

	for l := 0; l < n; l++ {
		for r < n && running_pairs < k {
			x := nums[r]
			running_pairs += d[x]
			d[x]++
			r++
		}

		if running_pairs < k {
			break
		}

		ans += n - r + 1
		y := nums[l]
		d[y]--
		running_pairs -= d[y]
	}

	return int64(ans)
}
