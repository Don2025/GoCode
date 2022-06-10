package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func (h *MinHeap) Top() interface{} {
	return (*h)[0]
}

type KthLargest struct {
	Heap *MinHeap
	Size int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{&MinHeap{}, k}
	for _, x := range nums {
		kl.Add(x)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	if len(*kl.Heap) < kl.Size || val > kl.Heap.Top().(int) {
		heap.Push(kl.Heap, val)
	}
	if len(*kl.Heap) > kl.Size {
		heap.Pop(kl.Heap)
	}
	return kl.Heap.Top().(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

/*
 * 执行用时：24ms 在所有Go提交中击败了92.70%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了23.60%的用户
**/
