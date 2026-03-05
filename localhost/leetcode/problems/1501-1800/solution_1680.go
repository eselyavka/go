package solutions

func concatenatedBinary(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var ans uint
	var length uint
	for i := 1; i < n+1; i++ {
		if i&(i-1) == 0 {
			length++
		}
		ans = ((ans << length) | uint(i)) % 1000000007
	}
	return int(ans)
}
