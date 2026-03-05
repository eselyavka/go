package solutions

import "sort"

func recoverTree(root *TreeNode) {
	arr := make([]int, 0)
	curr := root
	stack := make([]*TreeNode, 0)

	for true {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr = append(arr, node.Val)
			curr = node.Right
		} else {
			break
		}
	}

	sorted_arr := make([]int, len(arr))
	copy(sorted_arr, arr)
	sort.Ints(sorted_arr)
	n := len(sorted_arr)

	var swap_from int
	var swap_to int

	for i := 0; i < n; i++ {
		if arr[i] != sorted_arr[i] {
			swap_from = sorted_arr[i]
			swap_to = arr[i]
		}
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		if node.Val == swap_to {
			node.Val = swap_from
		} else if node.Val == swap_from {
			node.Val = swap_to
		}

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
}
