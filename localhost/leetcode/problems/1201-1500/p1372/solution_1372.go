package p1372

import "localhost/leetcode/util"

func rec(root *util.TreeNode, goLeft bool, steps int, ans *int) {
	if root == nil {
		*ans = util.MaxInts([]int{steps - 1, *ans})
		return
	}

	if goLeft {
		rec(root.Left, false, steps+1, ans)
		rec(root.Right, true, 1, ans)
	} else {
		rec(root.Left, false, 1, ans)
		rec(root.Right, true, steps+1, ans)
	}
}

func longestZigZag(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0

	rec(root, false, 0, &ans)
	rec(root, true, 0, &ans)

	return ans
}
