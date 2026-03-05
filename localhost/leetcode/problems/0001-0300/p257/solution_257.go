package p257

import (
	"localhost/leetcode/util"
	"strconv"
	"strings"
)

func rec(node *util.TreeNode, path *[]string, res *[]string) {
	if node == nil {
		return
	}

	*path = append(*path, strconv.Itoa(node.Val))

	if node.Left == nil && node.Right == nil {
		*res = append(*res, strings.Join(*path, "->"))
	}

	rec(node.Left, path, res)
	rec(node.Right, path, res)

	n := len(*path) - 1
	*path = (*path)[:n]
}

func binaryTreePaths(root *util.TreeNode) []string {
	path := make([]string, 0)
	res := make([]string, 0)

	rec(root, &path, &res)

	return res
}
