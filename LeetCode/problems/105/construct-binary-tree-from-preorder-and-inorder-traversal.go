package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
//------------------------Leetcode Problem 105------------------------
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

//------------------------Leetcode Problem 105------------------------
/*
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 * 执行用时：8ms 在所有Go提交中击败了22.39%的用户
 * 占用内存：4MB 在所有Go提交中击败了77.78%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		preorder := utils.StringToInts(scanner.Text())
		scanner.Scan()
		inorder := utils.StringToInts(scanner.Text())
		ans := structures.TreeToLevelOrderStrings(buildTree(preorder, inorder))
		Printf("Output: %v\n", ans)
	}
}
