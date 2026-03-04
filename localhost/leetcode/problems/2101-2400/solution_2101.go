package solutions

import (
	"math"
)

func dfs2101(visited map[int]struct{}, graph map[int][]int, node int) {
	if _, ok := visited[node]; !ok {
		visited[node] = struct{}{}
		for _, neighbour := range graph[node] {
			dfs2101(visited, graph, neighbour)
		}
	}
}
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	if n == 1 {
		return 1
	}

	G := make(map[int][]int)

	for i := 0; i < n; i++ {
		x1 := bombs[i][0]
		y1 := bombs[i][1]
		r1 := bombs[i][2]
		G[i] = append(G[i], i)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2 := bombs[j][0]
			y2 := bombs[j][1]
			d := int(math.Ceil(math.Sqrt(math.Pow(float64((x2-x1)), 2) + math.Pow(float64((y2-y1)), 2))))
			if r1 >= d {
				G[i] = append(G[i], j)
			}
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		visited := make(map[int]struct{})
		dfs2101(visited, G, i)
		ans = MaxInts([]int{ans, len(visited)})
	}
	return ans
}
