package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	getSuccessor := func(node *TreeNode) *TreeNode {
		ans := node.Right
		for ans.Left != nil && ans.Left != node {
			ans = ans.Left
		}
		return ans
	}
	var sum int
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			p := getSuccessor(node)
			if p.Left == nil {
				p.Left = node
				node = node.Right
			} else {
				p.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

/*
 * 执行用时：8ms 在所有Go提交中击败了94.76%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了64.63%的用户
**/
