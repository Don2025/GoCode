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

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
//------------------------Leetcode Problem 230------------------------
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return node.Val
			}
			root = node.Right
		}
	}
	return 0
}

//------------------------Leetcode Problem 230------------------------
/*
 * https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
 * 执行用时：12ms 在所有Go提交中击败了73.80%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了58.73%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", kthSmallest(root, k))
	}
}
