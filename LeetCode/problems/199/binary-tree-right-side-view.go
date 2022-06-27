package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/binary-tree-right-side-view/
//------------------------Leetcode Problem 199------------------------
func rightSideView(root *TreeNode) []int {
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > len(ans) {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 1)
	return ans
}

//------------------------Leetcode Problem 199------------------------
/*
 * https://leetcode.cn/problems/binary-tree-right-side-view/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了77.17%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v", rightSideView(root))
	}
}
