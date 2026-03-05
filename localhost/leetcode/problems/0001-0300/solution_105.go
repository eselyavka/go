package solutions

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := TreeNode{Val: preorder[0], Left: nil, Right: nil}

	mid := -1
	for idx, num := range inorder {
		if num == preorder[0] {
			mid = idx
		}
	}

	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return &root
}
