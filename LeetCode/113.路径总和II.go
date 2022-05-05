package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if sum == 0 && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, targetSum)
	return ans
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了73.33%的用户
**/
