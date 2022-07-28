package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/3Etpl5/
//-------------------------剑指 Offer II Problem 49------------------------
func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, preSum int) int {
		if node == nil {
			return 0
		}
		sum := preSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
	return dfs(root, 0)
}

//-------------------------剑指 Offer II Problem 49------------------------
/*
 * https://leetcode.cn/problems/3Etpl5/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了98.77%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %d\n", sumNumbers(root))
	}
}
