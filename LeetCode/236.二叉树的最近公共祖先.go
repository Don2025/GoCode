package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

/*
 * 执行用时：12ms 在所有Go提交中击败了43.24%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了47.33%的用户
**/
