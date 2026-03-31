package p3043

import "strconv"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := make(map[int]struct{})

	for _, num := range arr1 {
		for num > 0 {
			prefixes[num] = struct{}{}
			num /= 10
		}
	}

	ans := 0

	for _, num := range arr2 {
		for num > 0 {
			if _, ok := prefixes[num]; ok {
				numStr := strconv.Itoa(num)
				if len(numStr) > ans {
					ans = len(numStr)
					break
				}
			}
			num /= 10
		}
	}

	return ans
}
