package main

import "github.com/Don2025/GoCode/LeetCode/LCP/34"

func pathSum(root *_4.TreeNode, targetSum int) int {
	var ans int
	preSum := map[int]int{0: 1}
	var dfs func(*_4.TreeNode, int)
	dfs = func(node *_4.TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		ans += preSum[cur-targetSum]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]--
		return
	}
	dfs(root, 0)
	return ans
}

/*
 * 执行用时：4ms 在所有Go提交中击败了94.57%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了31.01%的用户
**/
