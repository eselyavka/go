package solutions

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	i := 0
	j := 0

	res := make([]int, 0)

	for i < n && j < m {
		if nums1[i] > nums2[j] {
			res = append(res, nums2[j])
			j++
		} else {
			res = append(res, nums1[i])
			i++
		}
	}

	for i < n {
		res = append(res, nums1[i])
		i++
	}

	for j < m {
		res = append(res, nums2[j])
		j++
	}

	mid := len(res) / 2

	if len(res)%2 == 1 {
		return float64(res[mid])
	}

	return float64(res[mid-1]+res[mid]) / 2.0

}
