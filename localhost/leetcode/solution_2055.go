package solutions

import "sort"

func platesBetweenCandles(s string, queries [][]int) []int {
	candles := make([]int, 0)

	for idx, c := range s {
		if c == '|' {
			candles = append(candles, idx)
		}
	}

	ans := make([]int, 0)

	for i := 0; i < len(queries); i++ {
		s := queries[i][0]
		e := queries[i][1]

		start_idx := sort.Search(len(candles), func(i int) bool { return candles[i] >= s })
		end_idx := sort.Search(len(candles), func(i int) bool { return candles[i] > e }) - 1

		if end_idx <= start_idx {
			ans = append(ans, 0)
			continue
		}
		ans = append(ans, candles[end_idx]-candles[start_idx]-(end_idx-start_idx))
	}

	return ans
}
