package p426

import "localhost/leetcode/util"

func inOrderTraversal(root *util.Node, acc *[]*util.Node) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, acc)
	(*acc) = append((*acc), root)
	inOrderTraversal(root.Right, acc)
}

func treeToDoublyList(root *util.Node) *util.Node {
	if root == nil {
		return nil
	}

	acc := make([]*util.Node, 0)
	inOrderTraversal(root, &acc)
	n := len(acc)

	for i, val := range acc {
		if i == 0 {
			val.Left = acc[n-1]
			if n == 1 {
				val.Right = val
			} else {
				val.Right = acc[1]
			}
		} else if i == n-1 {
			val.Left = acc[i-1]
			val.Right = acc[0]
		} else {
			val.Left = acc[i-1]
			val.Right = acc[i+1]
		}
	}

	return acc[0]
}
