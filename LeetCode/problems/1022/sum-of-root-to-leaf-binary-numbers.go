package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"os"
	"strings"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
//------------------------Leetcode Problem 1022------------------------
func sumRootToLeaf(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, val int) int {
		if node == nil {
			return 0
		}
		val = val<<1 | node.Val
		if node.Left == nil && node.Right == nil {
			return val
		}
		return dfs(node.Left, val) + dfs(node.Right, val)
	}
	return dfs(root, 0)
}

//------------------------Leetcode Problem 1022------------------------
/*
 * https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了80.30%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := structures.CreateTree(strings.Fields(scanner.Text()), 0)
		Printf("Output: %v\n", sumRootToLeaf(root))
	}
}
