package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"math"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
//------------------------Leetcode Problem 515------------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		maxVal := math.MinInt32
		for _, node := range queue {
			if node.Val > maxVal {
				maxVal = node.Val
			}
		}
		ans = append(ans, maxVal)
		var newQueue []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return
}

//------------------------Leetcode Problem 515------------------------
/*
 * https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
 * 执行用时：8ms 在所有Go提交中击败了92.77%的用户
 * 占用内存：6MB 在所有Go提交中击败了40.96%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := structures.MakeTree(scanner.Text())
		// root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", largestValues(root))
	}
}
