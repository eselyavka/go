package solutions

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	bfs := func(graph map[int][]int, n int) int {
		visited := make([]bool, n)
		queue := make([]int, 0)
		queue = append(queue, 0)
		visited[0] = true

		path := 0

		for len(queue) > 0 {
			cnt := len(queue)
			for cnt > 0 {
				node := queue[0]
				queue = queue[1:]
				if node == n-1 {
					return path
				}
				for _, neighbor := range graph[node] {
					if visited[neighbor] == false {
						queue = append(queue, neighbor)
						visited[neighbor] = true
					}
				}
				cnt--
			}
			path++
		}

		return path
	}

	G := make(map[int][]int)
	for i := 1; i < n; i++ {
		G[i-1] = append(G[i-1], i)
	}

	ans := make([]int, 0)

	for _, query := range queries {
		u := query[0]
		v := query[1]
		G[u] = append(G[u], v)
		ans = append(ans, bfs(G, n))
	}

	return ans
}
