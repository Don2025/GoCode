package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/8Zf90G/
//-------------------------剑指 Offer II Problem 52------------------------
func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	curNode := dummyNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		curNode.Right = node
		node.Left = nil
		curNode = node
		inorder(node.Right)
	}
	inorder(root)
	return dummyNode.Right
}

//-------------------------剑指 Offer II Problem 52------------------------
/*
 * https://leetcode.cn/problems/8Zf90G/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了42.97%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		ans := structures.TreeToLevelOrderStrings(increasingBST(root))
		Printf("Output: %s\n", ans)
	}
}
