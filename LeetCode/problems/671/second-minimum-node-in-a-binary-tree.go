package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/second-minimum-node-in-a-binary-tree/
//------------------------Leetcode Problem 671------------------------
func findSecondMinimumValue(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, value int) int {
		if root == nil {
			return -1
		}
		if root.Val > value {
			return root.Val
		}
		l, r := dfs(root.Left, value), dfs(root.Right, value)
		if l >= 0 && r >= 0 {
			return min(l, r)
		}
		return max(l, r)
	}
	return dfs(root, root.Val)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 671------------------------
/*
 * https://leetcode.cn/problems/second-minimum-node-in-a-binary-tree/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了98.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", findSecondMinimumValue(root))
	}
}
