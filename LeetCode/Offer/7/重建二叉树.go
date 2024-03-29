package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/
// ------------------------剑指 Offer I Problem 7------------------------
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	t := len(inorder[:i]) + 1
	root.Left = buildTree(preorder[1:t], inorder[:i])
	root.Right = buildTree(preorder[t:], inorder[i+1:])
	return root
}

// ------------------------剑指 Offer I Problem 7------------------------
/*
 * https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/
 * 执行用时：4ms 在所有Go提交中击败了94.29%的用户
 * 占用内存：4MB 在所有Go提交中击败了69.53%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		preorder := utils.StringToInts(scanner.Text())
		scanner.Scan()
		inorder := utils.StringToInts(scanner.Text())
		root := buildTree(preorder, inorder)
		Printf("Output: %v\n", structures.TreeToLevelOrderStrings(root))
	}
}
