package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type MedianFinder struct {
	small, large *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	s := &IntHeap{}
	heap.Init(s)
	l := &IntHeap{}
	heap.Init(l)
	return MedianFinder{s, l}
}

func (this *MedianFinder) AddNum(num int) {
	if this.large.Len() == 0 || num > (*this.large)[0] {
		heap.Push(this.large, num)
	} else {
		heap.Push(this.small, -num)
	}
	// 调节大小栈，确保 FindMedian() 得到正确的结果
	if this.large.Len() > this.small.Len()+1 {
		heap.Push(this.small, -heap.Pop(this.large).(int))
	} else if this.small.Len() > this.large.Len()+1 {
		heap.Push(this.large, -heap.Pop(this.small).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.large.Len() < this.small.Len() {
		return float64(-(*this.small)[0])
	} else if this.large.Len() > this.small.Len() {
		return float64((*this.large)[0])
	}
	return float64(-(*this.small)[0]+(*this.large)[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/*
 * 执行用时：72ms 在所有Go提交中击败了98.52%的用户
 * 占用内存：12.4MB 在所有Go提交中击败了88.50%的用户
**/
