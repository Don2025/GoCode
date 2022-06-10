package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//输入一棵二叉树的根结点，求该树的深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
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
		var root *TreeNode = createTree(strings.Fields(input.Text()), 0)
		println(maxDepth(root))
	}
}

//根据带null的一维数组创建二叉树
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
 * 占用内存：4.2MB 在所有Go提交中击败了53.74%的用户
**/
