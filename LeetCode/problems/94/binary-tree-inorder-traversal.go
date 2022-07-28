package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
//------------------------Leetcode Problem 94------------------------
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	node := root
	var ans []int
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		node = node.Right
	}
	return ans
}

//------------------------Leetcode Problem 94------------------------
/*
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了45.59%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := structures.MakeTree(scanner.Text())
		Printf("Output: %v\n", inorderTraversal(root))
	}
}
