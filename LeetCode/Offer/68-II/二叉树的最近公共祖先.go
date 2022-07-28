package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// ------------------------剑指 Offer I Problem 68-I------------------------
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// ------------------------剑指 Offer I Problem 68-I------------------------
/*
 * https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
 * 执行用时：8ms 在所有Go提交中击败了90.89%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了95.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		p := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		q := utils.StringToTreeNode(scanner.Text())
		Printf("%v\n", lowestCommonAncestor(root, p, q))
	}
}
