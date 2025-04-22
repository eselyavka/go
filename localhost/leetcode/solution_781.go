package solutions

func numRabbits(answers []int) int {
	count := make(map[int]int)
	for _, ans := range answers {
		count[ans]++
	}

	total := 0
	for k, v := range count {
		groupSize := k + 1
		numGroups := (v + groupSize - 1) / groupSize
		total += numGroups * groupSize
	}
	return total
}
