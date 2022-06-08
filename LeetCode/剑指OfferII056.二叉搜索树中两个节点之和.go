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
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = true
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}

/*
 * 执行用时：16ms 在所有Go提交中击败了96.80%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了63.60%的用户
**/
