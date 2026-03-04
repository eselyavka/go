package solutions

func addDigits(num int) int {
	ans := num

	for ans >= 10 {
		ans = ans/10 + ans%10
	}

	return ans
}
