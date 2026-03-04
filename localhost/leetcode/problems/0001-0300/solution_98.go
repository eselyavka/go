package solutions

func isValidBST(root *TreeNode) bool {
	acc := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root

	for true {
		if len(acc) >= 2 && acc[len(acc)-2] >= acc[len(acc)-1] {
			return false
		}

		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			acc = append(acc, node.Val)
			curr = node.Right
		} else {
			break
		}
	}

	return true
}
