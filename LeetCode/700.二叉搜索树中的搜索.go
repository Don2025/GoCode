package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

/*
 * 执行用时：20ms 在所有Go提交中击败了91.76%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了13.67%的用户
**/
