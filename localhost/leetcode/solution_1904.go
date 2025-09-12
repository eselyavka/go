package solutions

import (
	"strconv"
	"strings"
)

func numberOfRounds(loginTime string, logoutTime string) int {
	toMin := func(s string) int {
		tokens := strings.Split(s, ":")
		hour, _ := strconv.Atoi(tokens[0])
		minutes, _ := strconv.Atoi(tokens[1])

		return hour*60 + minutes
	}

	start := toMin(loginTime)
	end := toMin(logoutTime)

	if end < start {
		end += 24 * 60
	}

	start = ((start + 14) / 15) * 15
	end = (end / 15) * 15

	return MaxInts([]int{0, (end - start) / 15})
}
