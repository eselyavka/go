package solutions

import (
	"strconv"
	"strings"
)

func rec_257(node *TreeNode, path *[]string, res *[]string) {
	if node == nil {
		return
	}

	*path = append(*path, strconv.Itoa(node.Val))

	if node.Left == nil && node.Right == nil {
		*res = append(*res, strings.Join(*path, "->"))
	}

	rec_257(node.Left, path, res)
	rec_257(node.Right, path, res)

	n := len(*path) - 1
	*path = (*path)[:n]
}

func binaryTreePaths(root *TreeNode) []string {
	path := make([]string, 0)
	res := make([]string, 0)

	rec_257(root, &path, &res)

	return res
}
