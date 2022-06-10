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

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(isSymmetric(createTree(strings.Fields(input.Text()), 0)))
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
 * 执行用时：4ms 在所有Go提交中击败了73.32%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了68.14%的用户
**/
