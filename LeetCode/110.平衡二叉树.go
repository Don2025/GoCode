package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var process func(*TreeNode) int
	process = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := process(root.Left), process(root.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	return process(root) != -1
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
 * 执行用时：8ms 在所有Go提交中击败了45.95%的用户
 * 占用内存：5.6MB 在所有Go提交中击败了46.90%的用户
**/
