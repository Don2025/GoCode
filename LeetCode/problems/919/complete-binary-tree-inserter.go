package main

import "github.com/Don2025/GoCode/structures"

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/complete-binary-tree-inserter/
//------------------------Leetcode Problem 919------------------------

type CBTInserter struct {
	Tree *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{root}
}

func (cbt *CBTInserter) Insert(val int) int {
	q := []*TreeNode{cbt.Tree}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
			return node.Val
		}
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
			return node.Val
		}
		q = append(q, node.Left, node.Right)
	}
	return 0
}

func (cbt *CBTInserter) Get_root() *TreeNode {
	return cbt.Tree
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//------------------------Leetcode Problem 919------------------------
/*
 * https://leetcode.cn/problems/complete-binary-tree-inserter/
 * 执行用时：24ms 在所有Go提交中击败了8.11%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了5.41%的用户
**/
