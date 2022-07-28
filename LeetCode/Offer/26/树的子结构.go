package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
// ------------------------剑指 Offer I Problem 26------------------------
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil {
			return false
		}
		return A.Val == B.Val && dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
	}
	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// ------------------------剑指 Offer I Problem 26------------------------
/*
 * https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
 * 执行用时：20ms 在所有Go提交中击败了66.51%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了85.80%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		A := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		B := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", isSubStructure(A, B))
	}
}
