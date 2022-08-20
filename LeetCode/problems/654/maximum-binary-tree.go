package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/maximum-binary-tree/
//------------------------Leetcode Problem 654------------------------
func constructMaximumBinaryTree(nums []int) *TreeNode {
	tree := make([]*TreeNode, len(nums))
	var stack []int
	for i, num := range nums {
		tree[i] = &TreeNode{Val: num}
		for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
			tree[i].Left = tree[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			tree[stack[len(stack)-1]].Right = tree[i]
		}
		stack = append(stack, i)
	}
	return tree[stack[0]]
}

//------------------------Leetcode Problem 654------------------------
/*
 * https://leetcode.cn/problems/maximum-binary-tree/
 * 执行用时：16ms 在所有Go提交中击败了98.74%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了60.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", constructMaximumBinaryTree(nums))
	}
}
