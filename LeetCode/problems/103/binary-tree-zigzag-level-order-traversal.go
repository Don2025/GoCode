package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
//------------------------Leetcode Problem 103------------------------
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		var vals []int
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if i%2 == 1 {
			for j, n := 0, len(vals); j < n/2; j++ {
				vals[j], vals[n-1-j] = vals[n-1-j], vals[j]
			}
		}
		ans = append(ans, vals)
	}
	return ans
}

//------------------------Leetcode Problem 103------------------------
/*
 * https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了86.97%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", zigzagLevelOrder(root))
	}
}
