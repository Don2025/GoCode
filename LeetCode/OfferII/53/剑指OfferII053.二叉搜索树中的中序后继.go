package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/P5rCT8/
//-------------------------剑指 Offer II Problem 53------------------------
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	if p.Right != nil {
		ans = p.Right
		for ans.Left != nil {
			ans = ans.Left
		}
		return ans
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			ans = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return ans
}

//-------------------------剑指 Offer II Problem 53------------------------
/*
 * https://leetcode.cn/problems/P5rCT8/
 * 执行用时：16ms 在所有Go提交中击败了81.85%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了87.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		p := utils.StringToTreeNode(scanner.Text())
		ans := structures.TreeToLevelOrderStrings(inorderSuccessor(root, p))
		Printf("Output: %s\n", ans)
	}
}
