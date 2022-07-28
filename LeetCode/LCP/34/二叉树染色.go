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

// https://leetcode.cn/problems/er-cha-shu-ran-se-UGC/
// ------------------------LeetCode Cup Problem 34------------------------
func maxValue(root *TreeNode, k int) int {
	var dp func(*TreeNode) []int
	dp = func(node *TreeNode) []int {
		ans := make([]int, k+1)
		if node == nil {
			return ans
		}
		l, r := dp(node.Left), dp(node.Right)
		for i := 0; i < k; i++ {
			for j := 0; j < k-i; j++ {
				ans[i+j+1] = max(ans[i+j+1], l[i]+r[j]+node.Val)
			}
		}
		for i := 0; i <= k; i++ {
			for j := 0; j <= k; j++ {
				ans[0] = max(ans[0], l[i]+r[j])
			}
		}
		return ans
	}
	ans := dp(root)
	largest := ans[0]
	for i := 1; i < len(ans); i++ {
		if largest < ans[i] {
			largest = ans[i]
		}
	}
	return largest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------LeetCode Cup Problem 34------------------------
/*
 * https://leetcode.cn/problems/er-cha-shu-ran-se-UGC/
 * 执行用时：68ms 在所有Go提交中击败了50.00%的用户
 * 占用内存：9.7MB 在所有Go提交中击败了60.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", maxValue(root, k))
	}
}
