package solutions

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	freq := make([][]int, len(nums)+1, len(nums)+1)

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	res := make([]int, 0)

	for i := len(freq) - 1; i != 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
