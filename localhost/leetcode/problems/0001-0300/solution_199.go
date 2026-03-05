package solutions

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	ans := make([]int, 0)

	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		k := len(queue)
		for k > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			k--
		}
	}

	return ans
}
