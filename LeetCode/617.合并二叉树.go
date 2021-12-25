package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}
	if root1 != nil {
		return root1
	}
	return root2
}

/*
 * 执行用时：16ms 在所有Go提交中击败了79.00%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了46.82%的用户
**/
