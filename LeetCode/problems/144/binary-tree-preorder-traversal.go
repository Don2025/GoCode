package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
//------------------------Leetcode Problem 144------------------------
func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	node := root
	var ans []int
	for node != nil || len(stack) > 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}

//------------------------Leetcode Problem 144------------------------
/*
 * https://leetcode.cn/problems/binary-tree-preorder-traversal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了69.27%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("%v\n", preorderTraversal(root))
	}
}
