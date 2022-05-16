package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}
	if node := inorderSuccessor(root.Left, p); node != nil {
		return node
	}
	return root
}

/*
 * 执行用时：16ms 在所有Go提交中击败了75.80%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了51.60%的用户
**/
