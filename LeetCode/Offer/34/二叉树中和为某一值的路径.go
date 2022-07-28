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

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
// ------------------------剑指 Offer I Problem 34------------------------
func pathSum(root *TreeNode, target int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		target -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, target)
		dfs(node.Right, target)
	}
	dfs(root, target)
	return ans
}

// ------------------------剑指 Offer I Problem 34------------------------
/*
 * https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
 * 执行用时：4ms 在所有Go提交中击败了85.77%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了62.99%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", pathSum(root, target))
	}
}
