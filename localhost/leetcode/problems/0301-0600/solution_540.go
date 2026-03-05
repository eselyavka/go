package solutions

func binary_search_540(nums []int, l, r int) int {
	if l > r {
		return -1
	}

	if l == r {
		return nums[l]
	}

	mid := (l + r) / 2

	if mid%2 == 0 {
		if nums[mid] == nums[mid+1] {
			return binary_search_540(nums, mid+2, r)
		}
		return binary_search_540(nums, l, mid)
	}

	if nums[mid] == nums[mid-1] {
		return binary_search_540(nums, mid+1, r)
	}

	return binary_search_540(nums, l, mid-1)
}
func singleNonDuplicate(nums []int) int {
	return binary_search_540(nums, 0, len(nums)-1)
}
