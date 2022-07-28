package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
// ------------------------剑指 Offer I Problem 55-II------------------------
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(depth(node.Left), depth(node.Right)) + 1
	}
	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer I Problem 55-II------------------------
/*
 * https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
 * 执行用时：4ms 在所有Go提交中击败了93.87%的用户
 * 占用内存：5.6MB 在所有Go提交中击败了37.53%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("%v\n", isBalanced(root))
	}
}
