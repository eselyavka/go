package solutions

func sortedArrayToBSTRec(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := TreeNode{Val: nums[mid], Left: nil, Right: nil}
	node.Left = sortedArrayToBSTRec(nums, left, mid-1)
	node.Right = sortedArrayToBSTRec(nums, mid+1, right)

	return &node
}
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTRec(nums, 0, len(nums)-1)
}
