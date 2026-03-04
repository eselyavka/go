package solutions

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	ans := 0
	if root.Val >= low && root.Val <= high {
		ans += root.Val
	}

	if root.Val > low {
		ans += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		ans += rangeSumBST(root.Right, low, high)
	}
	return ans
}
