package solutions

func rec_1372(root *TreeNode, goLeft bool, steps int, ans *int) {
	if root == nil {
		*ans = MaxInts([]int{steps - 1, *ans})
		return
	}

	if goLeft {
		rec_1372(root.Left, false, steps+1, ans)
		rec_1372(root.Right, true, 1, ans)
	} else {
		rec_1372(root.Left, false, 1, ans)
		rec_1372(root.Right, true, steps+1, ans)
	}
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0

	rec_1372(root, false, 0, &ans)
	rec_1372(root, true, 0, &ans)

	return ans
}
