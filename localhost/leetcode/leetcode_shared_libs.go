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

func MinInts(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
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

func nTrue(array []bool) int {
	n := 0
	for _, v := range array {
		if v {
			n++
		}
	}
	return n
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

func initLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tail := head

	for _, num := range nums {
		tail.Next = &ListNode{Val: num, Next: nil}
		tail = tail.Next
	}
	return head.Next
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverseInt(n int) int {
	rev := 0
	for n != 0 {
		mod := n % 10
		rev = rev*10 + mod
		n = n / 10
	}

	return rev
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func IntSliceReverse(s []int) {
	n := len(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func inStringArray(array []string, target string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}

	return false
}

func arrayUniq(array []byte) bool {
	set := make(map[byte]struct{})
	for _, num := range array {
		set[num] = struct{}{}
	}

	return len(array) == len(set)
}

func binaryTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := binaryTreeHeight(root.Left)
	right := binaryTreeHeight(root.Right)

	return 1 + MaxInts([]int{left, right})
}

func int2dSliceIsEqual(this, that [][]int) bool {
	if len(this) != len(that) {
		return false
	}

	for i := 0; i < len(this); i++ {
		flag := false
		for j := 0; j < len(that); j++ {
			if intSliceEqual(this[i], that[j]) {
				flag = true
			}
		}

		if !flag {
			return false
		}
	}

	return true
}
