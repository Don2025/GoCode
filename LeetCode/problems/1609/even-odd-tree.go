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

// https://leetcode.cn/problems/even-odd-tree/
//------------------------Leetcode Problem 1609------------------------
func isEvenOddTree(root *TreeNode) bool {
	var odd bool // when odd=false, means even
	q := []*TreeNode{root}
	for len(q) > 0 {
		var prev int
		if odd {
			prev = math.MaxInt32
		}
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if odd && (prev <= node.Val || node.Val%2 == 1) {
				return false
			}
			if !odd && (prev >= node.Val || node.Val%2 == 0) {
				return false
			}
			prev = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		odd = !odd
	}
	return true
}

/*
 * https://leetcode.cn/problems/even-odd-tree/
 * 执行用时：176ms 在所有Go提交中击败了66.67%的用户
 * 占用内存：17.1MB 在所有Go提交中击败了71.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", isEvenOddTree(root))
	}
}
