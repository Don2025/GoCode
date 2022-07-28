package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/pOCWxh/
// ------------------------剑指 Offer II Problem 47------------------------
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

// ------------------------剑指 Offer II Problem 47------------------------
/*
 * https://leetcode.cn/problems/pOCWxh/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了66.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		ans := structures.TreeToLevelOrderStrings(pruneTree(root))
		Printf("Output: %s\n", ans)
	}
}
