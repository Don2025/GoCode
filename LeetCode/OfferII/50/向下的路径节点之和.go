package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/6eUYwP/
//-------------------------剑指 Offer II Problem 50------------------------
func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	preSum := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cur int) {
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

//-------------------------剑指 Offer II Problem 50------------------------
/*
 * https://leetcode.cn/problems/6eUYwP/
 * 执行用时：4ms 在所有Go提交中击败了94.57%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了31.01%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		targetSum, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", pathSum(root, targetSum))
	}
}
