package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	var preOrder func(*TreeNode, int) bool
	preOrder = func(root *TreeNode, k int) bool {
		if root == nil {
			return false
		}
		if m[k-root.Val] {
			return true
		}
		m[root.Val] = true
		return preOrder(root.Left, k) || preOrder(root.Right, k)
	}
	return preOrder(root, k)
}

/*
 * 执行用时：16ms 在所有Go提交中击败了98.28%的用户
 * 占用内存：7.6MB 在所有Go提交中击败了25.52%的用户
**/
