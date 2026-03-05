package solutions

func numPairsDivisibleBy60(time []int) int {
	n := len(time)
	if n == 1 {
		return 0
	}

	ans := 0
	arr := make([]int, 60, 60)

	for _, t := range time {
		div_60 := t % 60
		if div_60 == 0 {
			ans += arr[0]
		} else {
			diff := 60 - div_60
			ans += arr[diff]
		}
		arr[div_60]++
	}

	return ans
}
