package solutions

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	i, j := m-1, n-1
	total := m + n - 1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[total] = nums2[j]
			j--
		} else {
			nums1[total] = nums1[i]
			i--
		}
		total--
	}

	for j >= 0 {
		nums1[total] = nums2[j]
		total--
		j--
	}
}
