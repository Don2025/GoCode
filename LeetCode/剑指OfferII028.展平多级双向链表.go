package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	var arr []*Node
	var dfs func(*Node)
	dfs = func(head *Node) {
		if head == nil {
			return
		}
		arr = append(arr, head)
		dfs(head.Child)
		dfs(head.Next)
	}
	dfs(root)
	for i := 0; i < len(arr)-1; i++ {
		pre, cur := arr[i], arr[i+1]
		cur.Prev = pre
		pre.Next = cur
		pre.Child = nil
	}
	return root
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了5.49%的用户
**/
