package main

type Node struct {
	Val        int
	Prev, Next *Node
}

type MyLinkedList struct {
	root, tail *Node
}

func Constructor() MyLinkedList {
	l := MyLinkedList{&Node{-1, nil, nil}, &Node{-1, nil, nil}}
	l.root.Next = l.tail
	l.tail.Prev = l.root
	return l
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.root
	for i := 0; i <= index; i++ {
		if cur.Next == this.tail {
			return -1
		}
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, this.root, this.root.Next}
	this.root.Next.Prev = node
	this.root.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, this.tail.Prev, this.tail}
	this.tail.Prev.Next = node
	this.tail.Prev = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.root
	for i := 0; i <= index; i++ {
		if cur.Next == this.tail {
			if i == index {
				this.AddAtTail(val)
				return
			} else {
				return
			}
		}
		cur = cur.Next
	}
	node := &Node{val, cur.Prev, cur}
	cur.Prev.Next = node
	cur.Prev = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	cur := this.root
	for i := 0; i <= index; i++ {
		if cur.Next == this.tail {
			return
		}
		cur = cur.Next
	}
	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/*
 * 执行用时：20ms 在所有Go提交中击败了97.86%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了70.95%的用户
**/
