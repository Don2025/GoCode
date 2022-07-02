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

// https://leetcode.cn/problems/search-in-a-binary-search-tree/
//------------------------Leetcode Problem 700------------------------
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

//------------------------Leetcode Problem 700------------------------
/*
 * https://leetcode.cn/problems/search-in-a-binary-search-tree/
 * 执行用时：20ms 在所有Go提交中击败了91.76%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了13.67%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		ans := searchBST(root, val)
		Printf("Output: %v\n", structures.TreeToLevelOrderStrings(ans))
	}
}
