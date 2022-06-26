package main

type Node struct {
	Val               int
	Left, Right, Next *Node
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
//------------------------Leetcode Problem 117------------------------
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i < n-1 {
				node.Next = node
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

//------------------------Leetcode Problem 117------------------------
/*
 * https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
 * 执行用时：4ms 在所有Go提交中击败了73.17%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了53.80%的用户
**/
