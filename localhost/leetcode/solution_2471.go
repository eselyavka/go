package solutions

import "sort"

type tupleValIdx struct {
	value int
	index int
}

func minimumOperations(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0

	for len(q) > 0 {
		size := len(q)
		arr := make([]int, 0)
		for size > 0 {
			node := q[0]
			q = q[1:]
			arr = append(arr, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}

		n := len(arr)
		if n > 1 {
			visited := make([]bool, n)

			indexedArr := make([]tupleValIdx, 0)
			for i := 0; i < n; i++ {
				indexedArr = append(indexedArr, tupleValIdx{arr[i], i})
			}

			sort.Slice(indexedArr, func(i, j int) bool {
				return indexedArr[i].value < indexedArr[j].value
			})

			for i := 0; i < n; i++ {
				if visited[i] || indexedArr[i].index == i {
					continue
				}
				cycles := 0
				j := i

				for !visited[j] {
					visited[j] = true
					j = indexedArr[j].index
					cycles += 1
				}

				if cycles > 1 {
					ans += cycles - 1
				}
			}
		}
	}
	return ans
}
