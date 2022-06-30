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

// https://leetcode.cn/problems/path-sum-iii/
//------------------------Leetcode Problem 437------------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return rootSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var ans int
	if root.Val == targetSum {
		ans++
	}
	ans += rootSum(root.Left, targetSum-root.Val) + rootSum(root.Right, targetSum-root.Val)
	return ans
}

//------------------------Leetcode Problem 437------------------------
/*
 * https://leetcode.cn/problems/path-sum-iii/
 * 执行用时：28ms 在所有Go提交中击败了24.12%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了64.30%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		targetSum, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", pathSum(root, targetSum))
	}
}
