package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	var mid int
	for i, x := range vin {
		if x == pre[0] {
			mid = i
			break
		}
	}
	root.Left = reConstructBinaryTree(pre[1:mid+1], vin[0:mid])
	root.Right = reConstructBinaryTree(pre[mid+1:len(pre)], vin[mid+1:len(vin)])
	return root
}

/*
 * 运行时间：9ms 超过1.18%用Go提交的代码
 * 占用内存：2188KB 超过66.12%用Go提交的代码
**/
