package p926

import "localhost/leetcode/util"

func minFlipsMonoIncr(s string) int {
	zero_cnt := 0

	for _, c := range s {
		if c == '0' {
			zero_cnt++
		}
	}

	one_cnt := 0

	ans := len(s) - zero_cnt

	for _, c := range s {
		if c == '0' {
			zero_cnt--
		} else {
			ans = util.MinInts([]int{ans, zero_cnt + one_cnt})
			one_cnt++
		}
	}

	return ans
}
