package solutions

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	q := make([]int, 0)
	l := 0
	r := 0
	n := len(nums)
	for r < n {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if l > q[0] {
			q = q[1:]
		}

		if r+1 >= k {
			ans = append(ans, nums[q[0]])
			l++
		}
		r++
	}

	return ans
}
