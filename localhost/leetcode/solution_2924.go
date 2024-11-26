package solutions

func findChampion(n int, edges [][]int) int {
	arr := make([]int, n, n)

	for _, edge := range edges {
		arr[edge[1]]++
	}
	champ := -1
	champCount := 0

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			champCount++
			champ = i
		}
	}

	if champCount == 1 {
		return champ
	}

	return -1
}
