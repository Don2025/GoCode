package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	curNode := dummyNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		curNode.Right = node
		node.Left = nil
		curNode = node
		inorder(node.Right)
	}
	inorder(root)
	return dummyNode.Right
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了42.97%的用户
**/
