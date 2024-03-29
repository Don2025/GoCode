package _30

// https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list/
//------------------------Leetcode Problem 430------------------------

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

//------------------------Leetcode Problem 430------------------------
/*
 * https://leetcode.cn/problems/flatten-a-multilevel-doubly-linked-list/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了72.37%的用户
**/
