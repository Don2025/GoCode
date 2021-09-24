package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	for p := root; p != nil; p = p.Next {
		if p.Child != nil { //若存在子链表则进行递归
			next := p.Next            //保留该结点的下一个结点
			child := flatten(p.Child) //得到扁平化后的子链表与之相连
			p.Next = child
			child.Prev = p
			p.Child = nil
			if next != nil { //连接该结点原来的下一个结点next
				for p.Next != nil {
					p = p.Next
				}
				p.Next = next
				next.Prev = p
			}
		}
	}
	return root
}
