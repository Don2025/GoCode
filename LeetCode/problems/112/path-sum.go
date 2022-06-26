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

// https://leetcode.cn/problems/path-sum/
//------------------------Leetcode Problem 112------------------------
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

//------------------------Leetcode Problem 112------------------------
/*
 * https://leetcode.cn/problems/path-sum/
 * 执行用时：4ms 在所有Go提交中击败了91.84%的用户
 * 占用内存：4.6MB 在所有Go提交中击败了65.61%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		targetSum, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", hasPathSum(root, targetSum))
	}
}
