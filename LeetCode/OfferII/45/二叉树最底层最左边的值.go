package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/LwUNpT/
// ------------------------剑指 Offer II Problem 45------------------------
func findBottomLeftValue(root *TreeNode) int {
	var bottomLeft int
	if root == nil {
		return bottomLeft
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				bottomLeft = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return bottomLeft
}

// ------------------------剑指 Offer II Problem 45------------------------
/*
 * https://leetcode.cn/problems/LwUNpT/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了57.88%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %d\n", findBottomLeftValue(root))
	}
}
