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

// https://leetcode.cn/problems/insert-into-a-binary-search-tree/
//------------------------Leetcode Problem 701------------------------
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right
		}
	}
	return root
}

//------------------------Leetcode Problem 701------------------------
/*
 * https://leetcode.cn/problems/insert-into-a-binary-search-tree/
 * 执行用时：24ms 在所有Go提交中击败了94.89%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了63.61%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", insertIntoBST(root, val))
	}
}
