package solutions

func isArraySpecial(nums []int, queries [][]int) []bool {
	noParityIndices := make([]int, 0)
	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i]%2 == nums[i-1]%2 {
			noParityIndices = append(noParityIndices, i)
		}
	}

	binarySearchInt := func(s, e int, arr []int) bool {
		l := 0
		r := len(arr) - 1

		for l <= r {
			mid := l + (r-l)/2
			idx := arr[mid]

			if idx < s {
				l = mid + 1
			} else if idx > e {
				r = mid - 1
			} else {
				return true
			}
		}

		return false
	}

	ans := make([]bool, 0)

	for _, query := range queries {
		left := query[0]
		right := query[1]

		badIdxExists := binarySearchInt(left+1, right, noParityIndices)
		if badIdxExists {
			ans = append(ans, false)
		} else {
			ans = append(ans, true)
		}
	}

	return ans
}
