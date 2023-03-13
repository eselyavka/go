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
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type tuple struct {
	node *TreeNode
	col  int
}

type tupleInt struct {
	row int
	col int
}

type tupleChar struct {
	c   uint8
	num int
}

type tuple3 struct {
	num1 int
	num2 int
	num3 int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type PickIdx struct {
	prefix_sum []int
	total_sum  int
}
