package util

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

// TupleInt is an integer pair.
type TupleInt struct {
	Row int
	Col int
}

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}
