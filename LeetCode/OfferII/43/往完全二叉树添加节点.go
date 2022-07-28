package main

import "github.com/Don2025/GoCode/structures"

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/NaqhDT/
// ------------------------剑指 Offer II Problem 43------------------------

type CBTInserter struct {
	CBTree *TreeNode
	Queue  []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
	for queue[0].Left != nil && queue[0].Right != nil {
		queue = append(queue, queue[0].Left)
		queue = append(queue, queue[0].Right)
		queue = queue[1:]
	}
	return CBTInserter{root, queue}
}

func (cbt *CBTInserter) Insert(v int) int {
	node := &TreeNode{v, nil, nil}
	var ans int
	if cbt.Queue[0].Left == nil {
		cbt.Queue[0].Left = node
		ans = cbt.Queue[0].Val
	} else if cbt.Queue[0].Right == nil {
		cbt.Queue[0].Right = node
		ans = cbt.Queue[0].Val
	}
	if cbt.Queue[0].Left != nil && cbt.Queue[0].Right != nil {
		cbt.Queue = append(cbt.Queue, cbt.Queue[0].Left)
		cbt.Queue = append(cbt.Queue, cbt.Queue[0].Right)
		cbt.Queue = cbt.Queue[1:]
	}
	return ans
}

func (cbt *CBTInserter) Get_root() *TreeNode {
	return cbt.CBTree
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
// ------------------------剑指 Offer II Problem 43------------------------
/*
 * https://leetcode.cn/problems/NaqhDT/
 * 执行用时：4ms 在所有Go提交中击败了98.51%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了64.68%的用户
**/
