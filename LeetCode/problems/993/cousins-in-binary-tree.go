package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"os"
	"strings"
)

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/cousins-in-binary-tree/
//------------------------Leetcode Problem 993------------------------
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

//------------------------Leetcode Problem 993------------------------
/*
 * https://leetcode.cn/problems/cousins-in-binary-tree/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了80.02%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strArr := strings.Fields(scanner.Text())
		root := structures.CreateTree(strArr, 0)
		scanner.Scan()
		var x, y int
		Sscanf(scanner.Text(), "%d %d", &x, &y)
		Printf("Output: %v\n", isCousins(root, x, y))
	}
}
