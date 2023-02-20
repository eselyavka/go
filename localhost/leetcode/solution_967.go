package solutions

import (
	"strconv"
	"strings"
)

func rec_967(i, k, n int, acc []string, ans map[string]struct{}) {
	if len(acc) == n {
		ans[strings.Join(acc, "")] = struct{}{}
		return
	}

	if i < 0 || i >= 10 {
		return
	}

	acc = append(acc, strconv.Itoa(i))

	rec_967(i+k, k, n, acc, ans)
	rec_967(i-k, k, n, acc, ans)
}
func numsSameConsecDiff(n int, k int) []int {
	buf := make(map[string]struct{})
	acc := make([]string, 0)
	for i := 1; i < 10; i++ {
		rec_967(i, k, n, acc, buf)
	}

	ans := make([]int, 0)

	for k, _ := range buf {
		val, _ := strconv.Atoi(k)
		ans = append(ans, val)
	}

	return ans
}
