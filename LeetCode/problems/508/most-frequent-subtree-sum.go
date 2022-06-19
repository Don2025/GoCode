package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/most-frequent-subtree-sum/
//------------------------Leetcode Problem 508------------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	cnt := make(map[int]int)
	var maxCnt int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[sum]++
		maxCnt = max(maxCnt, cnt[sum])
		return sum
	}
	dfs(root)
	var ans []int
	for k, v := range cnt {
		if v == maxCnt {
			ans = append(ans, k)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 508------------------------
/*
 * https://leetcode.cn/problems/most-frequent-subtree-sum/
 * 执行用时：4ms 在所有Go提交中击败了94.23%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了50.00%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", findFrequentTreeSum(root))
		Printf("Input a line of numbers separated by space:")
	}
}
