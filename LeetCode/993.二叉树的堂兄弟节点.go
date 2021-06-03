package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
		if xFound && yFound {
			return
		}
		dfs(node.Left, node, depth+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xParent != yParent
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strArr := strings.Fields(input.Text())
		var root *TreeNode = createTree(strArr, 0)
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		println(isCousins(root, x, y))
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
 * 占用内存：2.5MB 在所有Go提交中击败了80.02%的用户
**/
