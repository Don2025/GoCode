package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var depth int
	for _, child := range root.Children {
		depth = max(depth, maxDepth(child))
	}
	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了69.76%的用户
**/
