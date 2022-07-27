package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode
type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/linked-list-in-binary-tree/
//------------------------Leetcode Problem 1367------------------------
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil {
			return false
		}
		if root.Val != head.Val {
			return false
		}
		return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

//------------------------Leetcode Problem 1367------------------------
/*
 * https://leetcode.cn/problems/linked-list-in-binary-tree/
 * 执行用时：16ms 在所有Go提交中击败了85.51%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了69.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", isSubPath(head, root))
	}
}
