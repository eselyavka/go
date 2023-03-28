package solutions

func distance(a, b string) int {
	n := len(a)

	distance := 0
	ra := []rune(a)
	rb := []rune(b)
	for i := 0; i < n; i++ {
		if ra[i] != rb[i] {
			distance++
		}
	}

	return distance
}
func minMutation(startGene string, endGene string, bank []string) int {
	G := make(map[string]map[string]struct{})

	arr := append(bank, startGene)

	for _, gene := range arr {
		for _, item := range bank {
			if distance(gene, item) == 1 {
				if data, ok := G[gene]; !ok {
					m := make(map[string]struct{})
					m[item] = struct{}{}
					G[gene] = m
				} else {
					data[item] = struct{}{}
					G[gene] = data
				}
			}
		}
	}

	ans := MaxInt
	q := make([][]string, 0)
	path := make([]string, 0)
	path = append(path, startGene)
	q = append(q, path[:])

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		last := p[len(p)-1]
		if last == endGene {
			ans = MinInts([]int{len(p) - 1, ans})
		}

		for vertex, _ := range G[last] {
			if !inStringArray(p, vertex) {
				newpath := p[:]
				newpath = append(newpath, vertex)
				q = append(q, newpath)
			}
		}
	}

	if ans == MaxInt {
		return -1
	}

	return ans
}
