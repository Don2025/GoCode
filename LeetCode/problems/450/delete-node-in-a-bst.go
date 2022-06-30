package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/delete-node-in-a-bst/
//------------------------Leetcode Problem 450------------------------
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val, cur.Val = cur.Val, root.Val
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

//------------------------Leetcode Problem 450------------------------
/*
 * https://leetcode.cn/problems/delete-node-in-a-bst/
 * 执行用时：4ms 在所有Go提交中击败了99.53%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了99.81%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		key, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", deleteNode(root, key))
	}
}
