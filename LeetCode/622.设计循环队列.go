package main

type MyCircularQueue struct {
	Nums       []int
	Head, Tail int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k+1), 0, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Nums[this.Tail] = value
	this.Tail = (this.Tail + 1) % len(this.Nums)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1) % len(this.Nums)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Nums[this.Head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Nums[(len(this.Nums)+this.Tail-1)%len(this.Nums)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.Tail+1)%len(this.Nums) == this.Head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

/*
 * 执行用时：12ms 在所有Go提交中击败了79.25%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了34.91%的用户
**/
