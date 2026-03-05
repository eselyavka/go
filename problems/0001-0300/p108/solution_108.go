package p108

import "github.com/eseliavka/go/util"

func sortedArrayToBSTRec(nums []int, left, right int) *util.TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := util.TreeNode{Val: nums[mid], Left: nil, Right: nil}
	node.Left = sortedArrayToBSTRec(nums, left, mid-1)
	node.Right = sortedArrayToBSTRec(nums, mid+1, right)

	return &node
}
func sortedArrayToBST(nums []int) *util.TreeNode {
	return sortedArrayToBSTRec(nums, 0, len(nums)-1)
}
