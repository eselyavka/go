package solutions

func zeroFilledSubarray(nums []int) int64 {
	ans := int64(0)
	m := len(nums)

	i := 0

	for i < m {
		if nums[i] == 0 {
			n := 0
			k := i
			for k < m && nums[k] == 0 {
				n++
				k++
			}
			ans += int64(n*(n+1)) / 2
			i = k
			continue
		}
		i++
	}

	return ans
}

func zeroFilledSubarrayRef(nums []int) int64 {
	ans := int64(0)
	num_subarrays := int64(0)

	for _, num := range nums {
		if num == 0 {
			num_subarrays++
		} else {
			num_subarrays = 0
		}

		ans += num_subarrays
	}

	return ans
}
