package main

type node struct {
	prev, next *node
	value      int
}

type MyCircularDeque struct {
	head, tail     *node
	capacity, size int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{capacity: k}
}

func (cq *MyCircularDeque) InsertFront(value int) bool {
	if cq.IsFull() {
		return false
	}
	n := &node{value: value}
	if cq.IsEmpty() {
		cq.head = n
		cq.tail = n
	} else {
		n.next = cq.head
		cq.head.prev = n
		cq.head = n
	}
	cq.size++
	return true
}

func (cq *MyCircularDeque) InsertLast(value int) bool {
	if cq.IsFull() {
		return false
	}
	n := &node{value: value}
	if cq.IsEmpty() {
		cq.head = n
		cq.tail = n
	} else {
		n.prev = cq.tail
		cq.tail.next = n
		cq.tail = n
	}
	cq.size++
	return true
}

func (cq *MyCircularDeque) DeleteFront() bool {
	if cq.IsEmpty() {
		return false
	}
	cq.head = cq.head.next
	if cq.head != nil {
		cq.head.prev = nil
	}
	cq.size--
	return true
}

func (cq *MyCircularDeque) DeleteLast() bool {
	if cq.IsEmpty() {
		return false
	}
	cq.tail = cq.tail.prev
	if cq.tail != nil {
		cq.tail.next = nil
	}
	cq.size--
	return true
}

func (cq *MyCircularDeque) GetFront() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.head.value
}

func (cq *MyCircularDeque) GetRear() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.tail.value
}

func (cq *MyCircularDeque) IsEmpty() bool {
	return cq.size == 0
}

func (cq *MyCircularDeque) IsFull() bool {
	return cq.size == cq.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//------------------------Leetcode Problem 641------------------------
/*
 * https://leetcode.cn/problems/design-circular-deque/
 * 执行用时：12ms 在所有Go提交中击败了86.83%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了37.13%的用户
**/
