package main

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64 // 不能设置MinInt32因为有个测试用例[-2147483648]
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

/*
 * 执行用时：8ms 在所有Go提交中击败了21.01%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了59.62%的用户
**/
