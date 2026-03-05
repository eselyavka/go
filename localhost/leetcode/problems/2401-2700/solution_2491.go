package solutions

import "sort"

func dividePlayers(skill []int) int64 {
	n := len(skill)

	if n == 2 {
		return int64(skill[0] * skill[1])
	}

	l := 0
	r := n - 1

	sort.Ints(skill)

	d := make(map[int][]int)

	for l <= r {
		s := skill[l] + skill[r]
		d[s] = append(d[s], skill[l]*skill[r])
		l++
		r--
	}

	if len(d) > 1 {
		return -1
	}

	for _, val := range d {
		return int64(sum(val))
	}

	return -1
}
