package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
// ------------------------剑指 Offer I Problem 55-I------------------------
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	var depth int
	for len(queue) > 0 {
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

// ------------------------剑指 Offer I Problem 55-I------------------------
/*
 * https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
 * 执行用时：4ms 在所有Go提交中击败了83.94%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了33.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("%v\n", maxDepth(root))
	}
}
