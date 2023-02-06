package solutions

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	curr := root
	stack := make([]*TreeNode, 0)
	flag := false

	for true {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else if len(stack) > 0 {
			if flag {
				return stack[len(stack)-1]
			}
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node == p {
				flag = true
			}
			curr = node.Right
		} else {
			break
		}
	}

	return nil
}
