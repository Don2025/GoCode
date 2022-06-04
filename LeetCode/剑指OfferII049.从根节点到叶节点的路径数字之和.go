package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, preSum int) int {
		if node == nil {
			return 0
		}
		sum := preSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
	return dfs(root, 0)
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了98.77%的用户
**/
