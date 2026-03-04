package solutions

func quickSelect(nums []int, l, r, k int) int {
	pivot := nums[r]
	p := l

	for i := l; i < r; i++ {
		if nums[i] <= pivot {
			buf := nums[i]
			nums[i] = nums[p]
			nums[p] = buf
			p++
		}
	}

	buf := nums[p]
	nums[p] = nums[r]
	nums[r] = buf

	if k < p {
		return quickSelect(nums, l, p-1, k)
	}

	if k > p {
		return quickSelect(nums, p+1, r, k)
	}

	return nums[p]
}
func findKthLargest(nums []int, k int) int {

	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}
