package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val, cur.Val = cur.Val, root.Val
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

/*
 * 执行用时：4ms 在所有Go提交中击败了99.53%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了99.81%的用户
**/
