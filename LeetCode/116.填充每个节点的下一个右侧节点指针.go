package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if i < n-1 {
				node.Next = q[0]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

/*
 * 执行用时：4ms 在所有Go提交中击败了92.25%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了35.17%的用户
**/
