package solutions

func findChampionEasy(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	arr := make([]int, m, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				arr[i]++
			}
		}
	}

	champ := -1
	champIdx := -1

	for i := 0; i < m; i++ {
		if arr[i] > champ {
			champ = arr[i]
			champIdx = i
		}
	}

	return champIdx
}
