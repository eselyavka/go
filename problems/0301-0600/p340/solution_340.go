package p340

import "github.com/eseliavka/go/util"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	n := len(s)

	left := 0
	right := 0

	ans := util.MinInt
	d := make(map[string]int)

	for right < n {
		d[string(s[right])] = right
		right++

		if len(d) == k+1 {
			mins := make([]int, 0)
			for _, val := range d {
				mins = append(mins, val)
			}
			min_idx := util.MinInts(mins)
			delete(d, string(s[min_idx]))
			left = min_idx + 1
		}
		ans = util.MaxInts([]int{ans, right - left})
	}
	return ans
}
