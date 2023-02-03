package solutions

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	d := make(map[int]int)

	for _, num := range arr {
		d[num]++
	}

	pairs := make([][]int, 0)

	for k, v := range d {
		pairs = append(pairs, []int{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	diff := k
	for _, item := range pairs {
		num := item[0]
		cnt := item[1]
		diff -= cnt
		if diff >= 0 {
			delete(d, num)
		} else {
			return len(d)
		}
	}

	return len(d)
}
