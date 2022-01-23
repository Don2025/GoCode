package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

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

/*
 * 执行用时：24ms 在所有Go提交中击败了94.89%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了63.61%的用户
**/
