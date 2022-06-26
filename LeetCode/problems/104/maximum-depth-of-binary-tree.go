package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
//------------------------Leetcode Problem 104------------------------
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 104------------------------
/*
 * https://leetcode.cn/problems/maximum-depth-of-binary-tree/
 * 执行用时：4ms 在所有Go提交中击败了89.79%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了63.06%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", maxDepth(root))
	}
}
