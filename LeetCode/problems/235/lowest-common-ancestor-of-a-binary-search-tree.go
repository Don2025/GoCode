package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
//------------------------Leetcode Problem 235------------------------
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

//------------------------Leetcode Problem 235------------------------
/*
 * https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
 * 执行用时：20ms 在所有Go提交中击败了66.16%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了28.34%的用户
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
