package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var ans, cnt int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		dfs(root.Right, k)
		cnt++
		if cnt == k {
			ans = root.Val
			return
		}
		dfs(root.Left, k)
	}
	dfs(root, k)
	return ans
}

/*
 * 执行用时：8ms 在所有Go提交中击败了84.96%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了98.35%的用户
**/
