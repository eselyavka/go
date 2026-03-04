package solutions

func groupThePeople(groupSizes []int) [][]int {

	buckets := make(map[int][]int)
	var res [][]int

	for i, size := range groupSizes {
		buckets[size] = append(buckets[size], i)

		if len(buckets[size]) == size {
			res = append(res, buckets[size])
			buckets[size] = nil
		}
	}

	return res
}
