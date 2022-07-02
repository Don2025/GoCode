package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/subtree-of-another-tree/
//------------------------Leetcode Problem 572------------------------

type TreeNode = structures.TreeNode

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	var check func(a, b *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val == b.Val {
			return check(a.Left, b.Left) && check(a.Right, b.Right)
		}
		return false
	}
	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

//------------------------Leetcode Problem 572------------------------
/*
 * https://leetcode.cn/problems/subtree-of-another-tree/
 * 执行用时：12ms 在所有Go提交中击败了86.61%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了37.12%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		subRoot := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", isSubtree(root, subRoot))
	}
}
