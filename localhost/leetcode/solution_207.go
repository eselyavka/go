package solutions

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		adj[a] = append(adj[a], b)
	}

	visited := make([]bool, numCourses)
	inStack := make([]bool, numCourses)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		visited[u] = true
		inStack[u] = true

		for _, v := range adj[u] {
			if !visited[v] {
				if dfs(v) {
					return true
				}
			} else if inStack[v] {
				return true
			}
		}

		inStack[u] = false
		return false
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] && dfs(i) {
			return false
		}
	}
	return true
}
