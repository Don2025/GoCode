package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, value int) int {
	if root == nil {
		return -1
	}
	if root.Val > value {
		return root.Val
	}
	l, r := dfs(root.Left, value), dfs(root.Right, value)
	if l >= 0 && r >= 0 {
		return min(l, r)
	}
	return max(l, r)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(findSecondMinimumValue(createTree(strings.Fields(input.Text()), 0)))
	}
}

func createTree(arr []string, index int) *TreeNode {
	p := new(TreeNode)
	if arr[index] == "null" {
		return nil
	}
	p.Val, _ = strconv.Atoi(arr[index])
	if index < len(arr) && (2*index+1) < len(arr) {
		p.Left = createTree(arr, 2*index+1)
	}
	if index < len(arr) && (2*index+2) < len(arr) {
		p.Right = createTree(arr, 2*index+2)
	}
	return p
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了98.90%的用户
**/
