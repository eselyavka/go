package util

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SparseVector is a sparse vector representation.
type SparseVector struct {
	Storage map[int]int
}

// Node represents a binary tree node with parent pointer.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// NodeNext represents a binary tree node with next pointer.
type NodeNext struct {
	Val   int
	Left  *NodeNext
	Right *NodeNext
	Next  *NodeNext
}

// Tuple keeps a tree node and a column index.
type Tuple struct {
	Node *TreeNode
	Col  int
}

// TupleInt is an integer pair.
type TupleInt struct {
	Row int
	Col int
}

// TupleChar is a rune/count pair.
type TupleChar struct {
	C   uint8
	Num int
}

// Tuple3 keeps three integers.
type Tuple3 struct {
	Num1 int
	Num2 int
	Num3 int
}

// Tuple4 keeps count and sum values.
type Tuple4 struct {
	Count int
	Sum   int
}

// Tuple5 keeps board state metadata.
type Tuple5 struct {
	Steps int
	Board [][]int
}

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PickIdx stores prefix sums for weighted picks.
type PickIdx struct {
	PrefixSum []int
	TotalSum  int
}
