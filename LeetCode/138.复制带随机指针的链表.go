package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	for p := head; p != nil; p = p.Next {
		mp[p] = &Node{p.Val, nil, nil}
	}
	for p := head; p != nil; p = p.Next {
		mp[p].Next = mp[p.Next]
		mp[p].Random = mp[p.Random]
	}
	return mp[head]
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了30.75%的用户
**/
