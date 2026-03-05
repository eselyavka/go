package solutions

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}

	right := n - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}

	if right == 0 {
		return 0
	}

	ans := right
	left := 0

	for left < right && (left == 0 || arr[left] >= arr[left-1]) {
		for right < n && arr[left] > arr[right] {
			right++
		}
		ans = MinInts([]int{ans, right - left - 1})
		left++
	}

	return ans
}
