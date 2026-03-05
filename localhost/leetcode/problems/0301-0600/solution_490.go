package solutions

func dfs_409(maze [][]int, visited [][]bool, pos []int, dest []int) bool {
	if visited[pos[0]][pos[1]] {
		return false
	}

	if intSliceEqual(pos, dest) {
		return true
	}

	visited[pos[0]][pos[1]] = true

	r := pos[1] + 1
	l := pos[1] - 1
	u := pos[0] + 1
	d := pos[0] - 1

	for r < len(maze[0]) && maze[pos[0]][r] == 0 {
		r++
	}
	if dfs_409(maze, visited, []int{pos[0], r - 1}, dest) {
		return true
	}

	for l >= 0 && maze[pos[0]][l] == 0 {
		l--
	}
	if dfs_409(maze, visited, []int{pos[0], l + 1}, dest) {
		return true
	}

	for u < len(maze) && maze[u][pos[1]] == 0 {
		u++
	}
	if dfs_409(maze, visited, []int{u - 1, pos[1]}, dest) {
		return true
	}

	for d >= 0 && maze[d][pos[1]] == 0 {
		d--
	}
	if dfs_409(maze, visited, []int{d + 1, pos[1]}, dest) {
		return true
	}

	return false
}

func hasPath(maze [][]int, start []int, destination []int) bool {
	m := len(maze)
	n := len(maze[0])

	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		vector := make([]bool, n)
		visited[i] = vector
	}

	return dfs_409(maze, visited, start, destination)
}
