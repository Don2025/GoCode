package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}
	stack := []*Node{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.8MB 在所有Go提交中击败了85.84%的用户
**/
