package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(depth(node.Left), depth(node.Right)) + 1
	}
	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 执行用时：4ms 在所有Go提交中击败了93.87%的用户
 * 占用内存：5.6MB 在所有Go提交中击败了37.53%的用户
**/
