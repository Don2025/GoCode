package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]bool)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if set[k-node.Val] {
			return true
		}
		set[node.Val] = true
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return false
}

/*
 * 执行用时：16ms 在所有Go提交中击败了98.74%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了60.92%的用户
**/
