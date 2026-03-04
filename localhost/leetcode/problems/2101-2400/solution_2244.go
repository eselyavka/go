package solutions

func minimumRounds(tasks []int) int {
	if len(tasks) == 1 {
		return -1
	}

	freq := make(map[int]int)

	for _, v := range tasks {
		freq[v]++
	}

	ans := 0

	for _, f := range freq {
		if f == 1 {
			return -1
		}
		if f%3 == 0 {
			ans += f / 3
		} else {
			ans += f/3 + 1
		}
	}

	return ans
}
