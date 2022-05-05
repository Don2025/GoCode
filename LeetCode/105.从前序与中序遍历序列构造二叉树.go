package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

/*
 * 执行用时：8ms 在所有Go提交中击败了22.39%的用户
 * 占用内存：4MB 在所有Go提交中击败了77.78%的用户
**/
