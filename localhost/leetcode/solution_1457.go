package solutions

func pseudoPalindromicPaths(root *TreeNode) int {
	s := make(map[int]struct{})
	res := rec_1457(root, s)
	return res
}
func rec_1457(root *TreeNode, s map[int]struct{}) int {
	if root == nil {
		return 0
	}

	_, prs := s[root.Val]

	if prs {
		delete(s, root.Val)
	} else {
		s[root.Val] = struct{}{}
	}
	res := 0

	if root.Left == nil && root.Right == nil {
		if len(s) <= 1 {
			res = 1
		} else {
			res = 0
		}
	} else {
		res = res + rec_1457(root.Left, s) + rec_1457(root.Right, s)
	}

	_, prs = s[root.Val]

	if prs {
		delete(s, root.Val)
	} else {
		s[root.Val] = struct{}{}
	}

	return res
}
