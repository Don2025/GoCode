package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var cnt int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		cnt += root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) + cnt
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了61.71%的用户
**/
