package solutions

func kthSmallest(root *TreeNode, k int) int {
	curr := root
	stack := make([]*TreeNode, 0)

	for true {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return node.Val
			}
			curr = node.Right
		}
	}
	return -1
}
