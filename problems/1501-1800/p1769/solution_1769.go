package p1769

import "github.com/eseliavka/go/util"

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n, n)

	s := make(map[int]struct{})

	for idx, c := range boxes {
		if string(c) == "1" {
			s[idx] = struct{}{}
		}
	}

	for i := 0; i < n; i++ {
		buf := 0
		for idx, _ := range s {
			if i == idx {
				continue
			}
			buf += util.Abs(i - idx)
		}
		ans[i] = buf
	}

	return ans
}
