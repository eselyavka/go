package solutions

func findClosest(target int, arr []int) []int {
	if target < arr[0] {
		return []int{-1, arr[0]}
	}

	if target > arr[len(arr)-1] {
		return []int{arr[len(arr)-1], -1}
	}

	r := 0
	l := len(arr)
	mid := 0

	for r < l {
		mid = (r + l) / 2
		if arr[mid] == target {
			return []int{target, target}
		}
		if target < arr[mid] {
			if mid > 0 && target > arr[mid-1] {
				return []int{arr[mid-1], arr[mid]}
			}
			l = mid
		} else {
			if mid < len(arr)-1 && target < arr[mid+1] {
				return []int{arr[mid], arr[mid+1]}
			}
			r = mid + 1
		}
	}

	return []int{-1, -1}
}
func closestNodes(root *TreeNode, queries []int) [][]int {
	curr := root
	stack := make([]*TreeNode, 0)
	inorder := make([]int, 0)
	for true {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inorder = append(inorder, node.Val)
			curr = node.Right
		} else {
			break
		}
	}

	ans := make([][]int, 0)
	for _, q := range queries {
		ans = append(ans, findClosest(q, inorder))
	}
	return ans
}
