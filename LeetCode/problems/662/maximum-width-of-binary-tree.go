package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/maximum-width-of-binary-tree/
//------------------------Leetcode Problem 662------------------------
func widthOfBinaryTree(root *TreeNode) int {
	type Pair struct {
		node  *TreeNode
		index int
	}
	ans := 1
	queue := []Pair{{root, 0}}
	for len(queue) > 0 {
		ans = max(ans, queue[len(queue)-1].index-queue[0].index+1)
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i].node.Left != nil {
				queue = append(queue, Pair{queue[i].node.Left, queue[i].index * 2})
			}
			if queue[i].node.Right != nil {
				queue = append(queue, Pair{queue[i].node.Right, queue[i].index*2 + 1})
			}
		}
		queue = queue[size:]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 662------------------------
/*
 * https://leetcode.cn/problems/maximum-width-of-binary-tree/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了57.86%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", widthOfBinaryTree(root))
	}
}
