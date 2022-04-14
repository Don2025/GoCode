package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil {
			return false
		}
		return A.Val == B.Val && dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
	}
	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

/*
 * 执行用时：20ms 在所有Go提交中击败了66.51%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了85.80%的用户
**/
