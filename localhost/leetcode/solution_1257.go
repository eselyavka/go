package solutions

func dfs_1257(root string, graph map[string][]string, path []string, r string, paths map[string][]string) {
	if _, ok := graph[root]; !ok {
		return
	}

	path = append(path, root)

	if inStringArray(graph[root], r) {
		copy_path := path[:]
		copy_path = append(copy_path, r)
		paths[r] = copy_path
		return
	}

	for _, item := range graph[root] {
		dfs_1257(item, graph, path, r, paths)
	}

	path = path[:len(path)-1]
}

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	root := regions[0][0]
	G := make(map[string][]string)

	for _, region := range regions {
		G[region[0]] = region[1:]
	}

	paths := make(map[string][]string)

	path := make([]string, 0)
	dfs_1257(root, G, path, region1, paths)
	path = make([]string, 0)
	dfs_1257(root, G, path, region2, paths)

	arr1 := paths[region1]
	arr2 := paths[region2]

	prev := arr1[0]
	n := MinInts([]int{len(arr1), len(arr2)})
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return prev
		}
		prev = arr1[i]
	}

	return prev
}
