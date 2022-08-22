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

// https://leetcode.cn/problems/print-binary-tree/
//------------------------Leetcode Problem 655------------------------
func printTree(root *TreeNode) [][]string {
	depth := getDepth(root)
	m := depth + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, i, j int) {
		ans[i][j] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, i+1, j-1<<(depth-i-1))
		}
		if node.Right != nil {
			dfs(node.Right, i+1, j+1<<(depth-i-1))
		}
	}
	dfs(root, 0, (n-1)>>1)
	return ans
}

func getDepth(root *TreeNode) (depth int) {
	if root.Left != nil {
		depth = getDepth(root.Left) + 1
	}
	if root.Right != nil {
		depth = max(depth, getDepth(root.Right)+1)
	}
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 655------------------------
/*
 * https://leetcode.cn/problems/print-binary-tree/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了80.95%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", printTree(root))
	}
}
