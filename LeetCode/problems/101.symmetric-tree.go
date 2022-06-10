package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root, root)
}

/*
 * https://leetcode.cn/problems/symmetric-tree/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了69.62%的用户
**/

func main() {
	// The standard input is a line, for example:
	// 1 2 2 null 3 null 3
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		fmt.Println(isSymmetric(root))
	}
}
