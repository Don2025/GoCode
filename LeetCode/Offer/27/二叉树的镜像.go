package main

import (
	"bufio"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
// ------------------------剑指 Offer I Problem 27------------------------
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

// ------------------------剑指 Offer I Problem 27------------------------
/*
 * https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了70.91%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		mirrorTree(root)
		structures.TreeToLevelOrderStrings(root)
	}
}
