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

// ------------------------剑指 Offer I Problem 54------------------------
// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func kthLargest(root *TreeNode, k int) int {
	var ans, cnt int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		dfs(root.Right, k)
		cnt++
		if cnt == k {
			ans = root.Val
			return
		}
		dfs(root.Left, k)
	}
	dfs(root, k)
	return ans
}

// ------------------------剑指 Offer I Problem 54------------------------
/*
 * https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
 * 执行用时：8ms 在所有Go提交中击败了84.96%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了98.35%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", kthLargest(root, k))
	}
}
