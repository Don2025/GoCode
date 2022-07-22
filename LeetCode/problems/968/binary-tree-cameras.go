package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/binary-tree-cameras/
//------------------------Leetcode Problem 968------------------------
func minCameraCover(root *TreeNode) int {
	const INF = 0x3f3f3f3f
	var dfs func(*TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return INF, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = min(a, la+rb, ra+lb)
		c = min(a, lb+rb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}

func min(a ...int) int {
	val := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < val {
			val = a[i]
		}
	}
	return val
}

//------------------------Leetcode Problem 968------------------------
/*
 * https://leetcode.cn/problems/binary-tree-cameras/
 * 执行用时：4ms 在所有Go提交中击败了86.41%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了29.35%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", minCameraCover(root))
	}
}
