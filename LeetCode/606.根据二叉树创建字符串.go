package main

import "strconv"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var ans string
	if root == nil {
		return ans
	}
	ans = strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	if root.Left != nil {
		ans += "(" + tree2str(root.Left) + ")"
	} else {
		ans += "()"
	}
	if root.Right != nil {
		ans += "(" + tree2str(root.Right) + ")"
	}
	return ans
}

/*
 * 执行用时：8ms 在所有Go提交中击败了87.40%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了81.10%的用户
**/
