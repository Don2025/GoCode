package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/find-duplicate-subtrees/
// ------------------------Leetcode Problem 652------------------------
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := map[*TreeNode]bool{}
	seen := map[string]*TreeNode{}
	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		serial := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if p, ok := seen[serial]; ok {
			repeat[p] = true
		} else {
			seen[serial] = node
		}
		return serial
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}

//------------------------Leetcode Problem 652------------------------
/*
 * https://leetcode.cn/problems/find-duplicate-subtrees/
 * 执行用时：16ms 在所有Go提交中击败了45.06%的用户
 * 占用内存：16MB 在所有Go提交中击败了18.22%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		ans := findDuplicateSubtrees(root)
		for _, node := range ans {
			fmt.Printf("%v\n", structures.TreeToLevelOrderStrings(node))
		}
	}
}
