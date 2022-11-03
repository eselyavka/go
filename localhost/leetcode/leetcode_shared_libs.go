package solutions

func MaxInts(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func Reverse(input []int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}

	return result
}

func isPalindrome(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func binarySearch(l int, r int, nums []int, target int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[l] <= nums[mid] {
		if nums[l] <= target && nums[mid] >= target {
			return binarySearch(l, mid-1, nums, target)
		}

		return binarySearch(mid+1, r, nums, target)
	}

	if nums[mid] < target && nums[r] >= target {
		return binarySearch(mid+1, r, nums, target)
	}

	return binarySearch(l, mid-1, nums, target)
}

func binarySearchMin(l, r int, nums []int) int {
	mid := (l + r) / 2

	elem := nums[mid]
	var idx_left, idx_right int

	if mid-1 < 0 {
		idx_left = len(nums) - 1
	} else {
		idx_left = mid - 1
	}

	if mid+1 == len(nums) {
		idx_right = 0
	} else {
		idx_right = mid + 1
	}

	if nums[idx_left] > elem && nums[idx_right] > elem {
		return elem
	}

	if nums[r] <= elem {
		return binarySearchMin(mid+1, r, nums)
	} else {
		return binarySearchMin(l, mid-1, nums)
	}
}
