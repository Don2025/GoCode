package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Same(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && Same(node1.Left, node2.Right) && Same(node1.Right, node2.Left)
}

func isSymmetrical(pRoot *TreeNode) bool {
	return Same(pRoot, pRoot)
}

/*
 * 运行时间：10ms 超过0.00%用Go提交的代码
 * 占用内存：8922KB 超过13.68%用Go提交的代码
**/
