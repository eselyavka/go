package solutions

func smallestNumber(n int) int {
	for true {
		if n&(n+1) == 0 {
			return n
		}
		n++
	}

	return -1
}
