package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/
//------------------------Leetcode Problem 1161------------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	ans, maxSum := 1, root.Val
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		size := len(q)
		var sum int
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}

//------------------------Leetcode Problem 1161------------------------
/*
 * https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/
 * 执行用时：92ms 在所有Go提交中击败了94.00%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了38.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %d\n", maxLevelSum(root))
	}
}
