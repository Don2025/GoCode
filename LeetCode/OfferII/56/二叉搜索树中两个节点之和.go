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

// https://leetcode.cn/problems/opLdQZ/
//-------------------------剑指 Offer II Problem 56------------------------
func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]bool)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = true
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}

//-------------------------剑指 Offer II Problem 56------------------------
/*
 * https://leetcode.cn/problems/opLdQZ/
 * 执行用时：16ms 在所有Go提交中击败了96.80%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了63.60%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %t\n", findTarget(root, k))
	}
}
