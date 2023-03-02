package solutions

import "fmt"

func compress(chars []byte) int {
	i := 0
	j := 0

	n := len(chars)

	for i < n {

		cnt := 0
		chars[j] = chars[i]
		for i < n && chars[i] == chars[j] {
			cnt++
			i++
		}
		if cnt > 1 {
			tmp := fmt.Sprintf("%d", cnt)
			for _, c := range tmp {
				j++
				chars[j] = byte(c)
			}
		}
		j++
	}

	return j
}
