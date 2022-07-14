package solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SparseVector struct {
	storage map[int]int
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Parent *Node
}