package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func IsBalanced_Solution(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return treeDepth(pRoot) > -1
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	if leftDepth == -1 || rightDepth == -1 || abs(rightDepth-leftDepth) > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 运行时间：7ms 超过46.09%用Go提交的代码
 * 占用内存：2152KB 超过80.00%用Go提交的代码
**/
