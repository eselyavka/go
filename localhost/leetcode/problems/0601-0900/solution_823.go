package solutions

import (
	"math"
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	mod := int(math.Pow(10, 9) + 7)

	sort.Ints(arr)

	d := make(map[int]int)

	for idx, num := range arr {
		d[num] = 1
		for j := 0; j < idx; j++ {
			_, ok := d[num/arr[j]]
			if num%arr[j] == 0 && ok {
				d[num] += d[arr[j]] * d[num/arr[j]]
			}
		}
	}

	ans := 0
	for _, v := range d {
		ans += v
	}

	return ans % mod
}
