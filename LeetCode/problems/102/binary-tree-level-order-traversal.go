package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/binary-tree-level-order-traversal/
//------------------------Leetcode Problem 102------------------------
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ans = append(ans, []int{})
		var q []*TreeNode
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			ans[i] = append(ans[i], node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		queue = q
	}
	return ans
}

//------------------------Leetcode Problem 102------------------------
/*
 * https://leetcode.cn/problems/binary-tree-level-order-traversal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了86.29%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", levelOrder(root))
	}
}
