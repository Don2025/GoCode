package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil {
			return false
		}
		if root.Val != head.Val {
			return false
		}
		return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

/*
 * 执行用时：16ms 在所有Go提交中击败了85.51%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了69.57%的用户
**/
