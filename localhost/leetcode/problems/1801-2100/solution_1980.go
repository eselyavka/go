package solutions

func findDifferentBinaryString(nums []string) string {
	n := len(nums[0])
	seen := make(map[string]bool)

	for _, num := range nums {
		seen[num] = true
	}

	var ans string
	var found bool

	var generatePermutations func(chars []rune, n int, current string)
	generatePermutations = func(chars []rune, n int, current string) {
		if found {
			return
		}
		if n == 0 {
			if !seen[current] {
				ans = current
				found = true
			}
			return
		}
		for _, char := range chars {
			generatePermutations(chars, n-1, current+string(char))
		}
	}

	generatePermutations([]rune{'0', '1'}, n, "")

	return ans
}
