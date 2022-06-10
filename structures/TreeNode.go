package structures

// TreeNode is a binary tree node. Tree is a finite set of n (n >= 0) nodes.
// Val is the int value data we are keeping track of at this node.
// Left is the tree's left child node.
// Right is the tree's right child node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
