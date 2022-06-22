package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/find-bottom-left-tree-value/
//------------------------Leetcode Problem 513------------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}

//------------------------Leetcode Problem 513------------------------
/*
 * https://leetcode.cn/problems/find-bottom-left-tree-value/
 * 执行用时：4ms 在所有Go提交中击败了89.17%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了60.20%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", findBottomLeftValue(root))
		Printf("Input a line of numbers separated by space:")
	}
}
