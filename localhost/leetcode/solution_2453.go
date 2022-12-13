package solutions

func destroyTargets(nums []int, space int) int {
	d := make(map[int]int)

	for _, num := range nums {
		d[num%space]++
	}

	max := MinInt

	for _, val := range d {
		if max < val {
			max = val
		}
	}

	ans := MaxInt

	for _, num := range nums {
		if d[num%space] == max {
			ans = MinInts([]int{ans, num})
		}
	}
	return ans
}
