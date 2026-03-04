package solutions

func destCity(paths [][]string) string {
	G := make(map[string]string)

	for _, item := range paths {
		G[item[0]] = item[1]
	}

	ans := paths[0][0]

	for true {
		if _, ok := G[ans]; ok {
			ans = G[ans]
			continue
		}

		return ans
	}

	return ans
}
