package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/add-one-row-to-tree/
//------------------------Leetcode Problem 623------------------------
func addOneRow(root *TreeNode, val, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	q := []*TreeNode{root}
	var curDepth int
	var depthNodes []*TreeNode
	for len(q) > 0 {
		curDepth++
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node != nil && curDepth == depth-1 {
				depthNodes = append(depthNodes, node)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	for _, node := range depthNodes {
		if node.Left != nil {
			node.Left = &TreeNode{Val: val, Left: node.Left}
		} else {
			node.Left = &TreeNode{Val: val}
		}
		if node.Right != nil {
			node.Right = &TreeNode{Val: val, Right: node.Right}
		} else {
			node.Right = &TreeNode{Val: val}
		}
	}
	return root
}

//------------------------Leetcode Problem 623------------------------
/*
 * https://leetcode.cn/problems/add-one-row-to-tree/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：5.5MB 在所有Go提交中击败了45.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		var val, depth int
		Sscanf(scanner.Text(), "%d %d", &val, &depth)
		Printf("Output: %v\n", addOneRow(root, val, depth))
	}
}
