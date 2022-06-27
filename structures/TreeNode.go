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

// CreateTree creates a binary tree from a one-dimensional []string array with null values.
func CreateTree(arr []string, index int) *TreeNode {
	p := new(TreeNode)
	if arr[index] == "null" {
		return nil
	}
	p.Val, _ = strconv.Atoi(arr[index])
	if index < len(arr) && (2*index+1) < len(arr) {
		p.Left = CreateTree(arr, 2*index+1)
	}
	if index < len(arr) && (2*index+2) < len(arr) {
		p.Right = CreateTree(arr, 2*index+2)
	}
	return p
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

// TreeToLevelOrderStrings converts *TreeNode into []string by level order traversal.
func TreeToLevelOrderStrings(root *TreeNode) []string {
	var ans []string
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	var level int
	for level = 0; len(queue) > 0; level++ {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				ans = append(ans, "null")
			} else {
				ans = append(ans, strconv.Itoa(node.Val))
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}
	level--
	return ans[:1<<level-1]
}
