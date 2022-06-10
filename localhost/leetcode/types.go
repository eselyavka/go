package solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SparseVector struct {
	storage map[int]int
}
