package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/w6cpku/
//-------------------------剑指 Offer II Problem 54------------------------
func convertBST(root *TreeNode) *TreeNode {
	getSuccessor := func(node *TreeNode) *TreeNode {
		ans := node.Right
		for ans.Left != nil && ans.Left != node {
			ans = ans.Left
		}
		return ans
	}
	var sum int
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			p := getSuccessor(node)
			if p.Left == nil {
				p.Left = node
				node = node.Right
			} else {
				p.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

//-------------------------剑指 Offer II Problem 54------------------------
/*
 * https://leetcode.cn/problems/w6cpku/
 * 执行用时：8ms 在所有Go提交中击败了94.76%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了64.63%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		ans := structures.TreeToLevelOrderStrings(convertBST(root))
		Printf("Output: %v\n", ans)
	}
}
