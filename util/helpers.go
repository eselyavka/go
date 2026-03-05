package util

import "fmt"

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

func StringSlicesEqual(a, b []string) bool {
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

func IntSliceEqual(a, b []int) bool {
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

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func NTrue(array []bool) int {
	n := 0
	for _, v := range array {
		if v {
			n++
		}
	}
	return n
}

func IsPalindrome(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func BinarySearch(l int, r int, nums []int, target int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[l] <= nums[mid] {
		if nums[l] <= target && nums[mid] >= target {
			return BinarySearch(l, mid-1, nums, target)
		}
		return BinarySearch(mid+1, r, nums, target)
	}
	if nums[mid] < target && nums[r] >= target {
		return BinarySearch(mid+1, r, nums, target)
	}
	return BinarySearch(l, mid-1, nums, target)
}

func BinarySearchMin(l, r int, nums []int) int {
	mid := (l + r) / 2
	elem := nums[mid]
	var idxLeft, idxRight int
	if mid-1 < 0 {
		idxLeft = len(nums) - 1
	} else {
		idxLeft = mid - 1
	}
	if mid+1 == len(nums) {
		idxRight = 0
	} else {
		idxRight = mid + 1
	}
	if nums[idxLeft] > elem && nums[idxRight] > elem {
		return elem
	}
	if nums[r] <= elem {
		return BinarySearchMin(mid+1, r, nums)
	}
	return BinarySearchMin(l, mid-1, nums)
}

func InitLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: -1}
	tail := head
	for _, num := range nums {
		tail.Next = &ListNode{Val: num}
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

func ReverseInt(n int) int {
	rev := 0
	for n != 0 {
		mod := n % 10
		rev = rev*10 + mod
		n /= 10
	}
	return rev
}

func ReverseString(s string) string {
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

func InStringArray(array []string, target string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false
}

func In2DIntArray(array [][]int, target []int) bool {
	for _, arr := range array {
		if IntSliceEqual(arr, target) {
			return true
		}
	}
	return false
}

func ArrayUniq(array []byte) bool {
	set := make(map[byte]struct{})
	for _, num := range array {
		set[num] = struct{}{}
	}
	return len(array) == len(set)
}

func BinaryTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := BinaryTreeHeight(root.Left)
	right := BinaryTreeHeight(root.Right)
	return 1 + MaxInts([]int{left, right})
}

func Int2DSliceIsEqual(this, that [][]int) bool {
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		flag := false
		for j := 0; j < len(that); j++ {
			if IntSliceEqual(this[i], that[j]) {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func Bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func CountSetBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func BinaryTreeBFS(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	arr := make([]int, 0)
	queue := []*TreeNode{node}
	current := 0
	for current != len(queue) {
		node := queue[current]
		current++
		if node == nil {
			continue
		}
		arr = append(arr, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	return arr
}

func FlipArrayXOR(original []int) []int {
	flipped := make([]int, len(original))
	for i, v := range original {
		flipped[i] = v ^ 1
	}
	return flipped
}

func SliceToString(slice []int) string {
	return fmt.Sprintf("%v", slice)
}

func TwoDSliceToString(slice [][]int) string {
	return fmt.Sprintf("%v", slice)
}

func Copy2DArray(original [][]int) [][]int {
	dup := make([][]int, len(original))
	for i := range original {
		dup[i] = make([]int, len(original[i]))
		CopyRow(dup[i], original[i])
	}
	return dup
}

func CopyRow(dest, src []int) {
	for i := range src {
		dest[i] = src[i]
	}
}

func Flatten2DArray(twoD [][]int) []int {
	total := 0
	for _, row := range twoD {
		total += len(row)
	}
	flattened := make([]int, 0, total)
	for _, row := range twoD {
		flattened = append(flattened, row...)
	}
	return flattened
}

func SumDigits(num int) int {
	total := 0
	for num > 0 {
		total += num % 10
		num /= 10
	}
	return total
}
