package structures

import (
	"strconv"
	"strings"
)

// TreeNode is a binary tree node. Tree is a finite set of n (n >= 0) nodes.
// Val is the int value data we are keeping track of at this node.
// Left is the tree's left child node.
// Right is the tree's right child node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MakeTree returns a new *TreeNode with the given value.
func MakeTree(str string) *TreeNode {
	if len(str) == 0 {
		return nil
	}
	strs := strings.Fields(str)
	n := len(strs)
	if n == 0 {
		return nil
	}
	x, _ := strconv.Atoi(strs[0])
	root := &TreeNode{Val: x}
	queue := make([]*TreeNode, 1, n<<1)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && strs[i] != "null" {
			x, _ = strconv.Atoi(strs[i])
			node.Left = &TreeNode{Val: x}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && strs[i] != "null" {
			x, _ = strconv.Atoi(strs[i])
			node.Right = &TreeNode{Val: x}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
