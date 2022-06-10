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

var ans, cnt int

func kthLargest(root *TreeNode, k int) int {
	dfs(root, k)
	return ans
}

func dfs(root *TreeNode, k int) {
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		root := createTree(strings.Fields(input.Text()), 0)
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(kthLargest(root, k))
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
