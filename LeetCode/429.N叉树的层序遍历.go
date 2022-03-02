package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		var arr []int
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				arr = append(arr, node.Val)
			}
			if node.Children != nil {
				for _, child := range node.Children {
					queue = append(queue, child)
				}
			}
		}
		if len(arr) > 0 {
			ans = append(ans, arr)
		}
	}
	return ans
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.3MB 在所有Go提交中击败了71.69%的用户
**/
