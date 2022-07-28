package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/hPov7L/
// ------------------------剑指 Offer II Problem 44------------------------
func largestValues(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		val := math.MinInt32
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			val = max(val, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, val)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 44------------------------
/*
 * https://leetcode.cn/problems/hPov7L/
 * 执行用时：4ms 在所有Go提交中击败了97.32%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了89.26%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", largestValues(root))
	}
}
