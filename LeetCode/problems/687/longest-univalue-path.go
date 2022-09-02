package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/longest-univalue-path/
//------------------------Leetcode Problem 687------------------------
func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if root.Left != nil && root.Left.Val == root.Val {
			left += 1
		} else {
			left = 0
		}
		if root.Right != nil && root.Right.Val == root.Val {
			right += 1
		} else {
			right = 0
		}
		ans = max(ans, left+right)
		return max(left, right)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 687------------------------
/*
 * https://leetcode.cn/problems/longest-univalue-path/
 * 执行用时：56ms 在所有Go提交中击败了64.93%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了97.76%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", longestUnivaluePath(root))
	}
}
