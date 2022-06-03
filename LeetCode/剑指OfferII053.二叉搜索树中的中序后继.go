package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

/*
 * 执行用时：16ms 在所有Go提交中击败了81.85%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了87.90%的用户
**/
