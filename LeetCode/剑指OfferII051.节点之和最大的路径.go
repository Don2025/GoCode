package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(0, maxGain(node.Left))
		rightGain := max(0, maxGain(node.Right))
		priceNewPath := node.Val + leftGain + rightGain
		maxSum = max(maxSum, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 执行用时：12ms 在所有Go提交中击败了96.48%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了51.76%的用户
**/
