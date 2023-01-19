package solutions

func getModifiedArray(length int, updates [][]int) []int {
	ans := make([]int, length, length)

	for _, t := range updates {
		sidx := t[0]
		eidx := t[1]
		inc := t[2]

		ans[sidx] += inc

		if eidx < length-1 {
			ans[eidx+1] -= inc
		}
	}

	for i := 1; i < length; i++ {
		ans[i] += ans[i-1]
	}

	return ans
}
