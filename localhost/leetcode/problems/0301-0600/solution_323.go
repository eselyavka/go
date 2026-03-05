package solutions

func dfs_323(graph map[int][]int, v []int, k int, visited map[int]bool) {
	visited[k] = true
	for _, neighbor := range v {
		if !visited[neighbor] {
			dfs_323(graph, graph[neighbor], neighbor, visited)
		}
	}
}
func countComponents(n int, edges [][]int) int {
	G := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)

	}
	ans := 0
	visited := make(map[int]bool, n)
	for k := 0; k < n; k++ {
		if !visited[k] {
			dfs_323(G, G[k], k, visited)
			ans++
		}
	}
	return ans
}
