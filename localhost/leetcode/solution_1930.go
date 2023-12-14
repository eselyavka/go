package solutions

func countPalindromicSubsequence(s string) int {
	if len(s) == 3 {
		if s[0] == s[2] {
			return 1
		} else {
			return 0
		}
	}

	d := make(map[rune][2]int)

	for i, c := range s {
		if _, ok := d[c]; ok {
			arr := d[c]
			arr[1] = i
			d[c] = arr
		} else {
			d[c] = [2]int{i, -1}
		}
	}

	ans := 0

	for _, value := range d {
		used := make(map[rune]struct{})
		idx_s := make(map[int]struct{})
		idx_s[value[0]] = struct{}{}
		idx_s[value[1]] = struct{}{}

		for i := value[0]; i < value[len(value)-1]; i++ {
			if _, ok := idx_s[i]; !ok {
				if _, ok := used[rune(s[i])]; !ok {
					ans += 1
					used[rune(s[i])] = struct{}{}
				}
			}
		}
	}

	return ans
}
