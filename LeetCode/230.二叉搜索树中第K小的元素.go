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

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return node.Val
			}
			root = node.Right
		}
	}
	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		root := createTree(strings.Fields(input.Text()), 0)
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(kthSmallest(root, k))
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
 * 执行用时：12ms 在所有Go提交中击败了73.80%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了58.73%的用户
**/
