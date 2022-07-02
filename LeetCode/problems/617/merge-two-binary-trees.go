package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/merge-two-binary-trees
//------------------------Leetcode Problem 617------------------------
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}
	if root1 != nil {
		return root1
	}
	return root2
}

//------------------------Leetcode Problem 617------------------------
/*
 * https://leetcode.cn/problems/merge-two-binary-trees/
 * 执行用时：16ms 在所有Go提交中击败了79.00%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了46.82%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root1 := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		root2 := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", mergeTrees(root1, root2))
	}
}
