package solutions

func lastVisitedIntegers(nums []int) []int {
	ans := make([]int, 0)
	seen := make([]int, 0)

	k := 0
	for _, num := range nums {
		if num >= 0 {
			seen = append([]int{num}, seen...)
			k = 0
		} else {
			k++
			if k <= len(seen) {
				ans = append(ans, seen[k-1])
			} else {
				ans = append(ans, -1)
			}
		}
	}

	return ans
}
