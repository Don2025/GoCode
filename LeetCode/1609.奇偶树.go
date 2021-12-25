package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	var odd bool // when odd=false, means even
	q := []*TreeNode{root}
	for len(q) > 0 {
		var prev int
		if odd {
			prev = math.MaxInt32
		}
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if odd && (prev <= node.Val || node.Val%2 == 1) {
				return false
			}
			if !odd && (prev >= node.Val || node.Val%2 == 0) {
				return false
			}
			prev = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		odd = !odd
	}
	return true
}

/*
 * 执行用时：176ms 在所有Go提交中击败了66.67%的用户
 * 占用内存：17.1MB 在所有Go提交中击败了71.43%的用户
**/
