package solutions

func getHappyString(n int, k int) string {
	var res []string
	chars := []rune{'a', 'b', 'c'}
	var generatePermutations func(chars []rune, n int, current string)

	generatePermutations = func(chars []rune, n int, current string) {
		if n == 0 {
			res = append(res, current)
			return
		}

		for _, char := range chars {
			if len(current) == 0 || current[len(current)-1] != byte(char) {
				generatePermutations(chars, n-1, current+string(char))
			}
		}
	}

	generatePermutations(chars, n, "")

	if k > len(res) {
		return ""
	}

	return res[k-1]
}
