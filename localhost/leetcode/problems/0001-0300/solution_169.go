package solutions

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n]++
		} else {
			m[n] = 1
		}
	}

	majority := len(nums) / 2

	for k, v := range m {
		if v > majority {
			return k
		}
	}
	return -1
}
