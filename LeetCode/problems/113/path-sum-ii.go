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

// https://leetcode.cn/problems/path-sum-ii/
//------------------------Leetcode Problem 113------------------------
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if sum == 0 && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, targetSum)
	return ans
}

//------------------------Leetcode Problem 113------------------------
/*
 * https://leetcode.cn/problems/path-sum-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了73.33%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		targetSum, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", pathSum(root, targetSum))
	}
}
