package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/sum-of-left-leaves/
//------------------------Leetcode Problem 404------------------------
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var cnt int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		cnt += root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) + cnt
}

//------------------------Leetcode Problem 404------------------------
/*
 * https://leetcode.cn/problems/sum-of-left-leaves/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了61.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", sumOfLeftLeaves(root))
	}
}
