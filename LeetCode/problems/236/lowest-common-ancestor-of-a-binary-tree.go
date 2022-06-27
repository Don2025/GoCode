package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
//------------------------Leetcode Problem 236------------------------
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

//------------------------Leetcode Problem 236------------------------
/*
 * https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
 * 执行用时：12ms 在所有Go提交中击败了43.24%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了47.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		p := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		q := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", lowestCommonAncestor(root, p, q))
	}
}
