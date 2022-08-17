package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/deepest-leaves-sum/
//------------------------Leetcode Problem 1302------------------------

func deepestLeavesSum(root *TreeNode) int {
	q := []*TreeNode{root}
	var sum int
	for len(q) > 0 {
		sum = 0
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return sum
}

//------------------------Leetcode Problem 1302------------------------
/*
 * https://leetcode.cn/problems/deepest-leaves-sum/
 * 执行用时：40ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：8.4MB 在所有Go提交中击败了6.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", deepestLeavesSum(root))
	}
}
