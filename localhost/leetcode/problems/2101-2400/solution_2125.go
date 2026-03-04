package solutions

import (
	"strings"
)

func numberOfBeams(bank []string) int {
	rows := len(bank)

	d := make(map[int]int)

	for idx, item := range bank {
		d[idx] = strings.Count(item, "1")
	}

	ans := 0

	for rows > -1 {
		k := rows - 1
		for k > -1 {
			if d[k] > 0 {
				break
			}
			k--
		}
		ans += d[rows] * d[k]
		rows--
	}

	return ans
}

func numberOfBeamsOther(bank []string) int {

	prevRowDevCnt := 0
	ans := 0
	var noBeams bool
	for _, item := range bank {
		runningDevCnt := 0
		noBeams = true
		for _, c := range item {
			if c == '1' {
				noBeams = false
				runningDevCnt++
				ans += prevRowDevCnt
			}
		}
		if !noBeams {
			prevRowDevCnt = runningDevCnt
		}
	}

	return ans
}
