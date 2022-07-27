package _845

import "container/heap"

// https://leetcode.cn/problems/seat-reservation-manager/
//------------------------Leetcode Problem 1845------------------------
type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Top() int            { return h[0] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type SeatManager struct {
	Seats *hp
}

func Constructor(n int) SeatManager {
	m := SeatManager{&hp{}}
	heap.Init(m.Seats)
	for i := 1; i <= n; i++ {
		heap.Push(m.Seats, i)
	}
	return m
}

func (this *SeatManager) Reserve() int {
	x := this.Seats.Top()
	heap.Pop(this.Seats)
	return x
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.Seats, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
//------------------------Leetcode Problem 1845------------------------
/*
 * https://leetcode.cn/problems/seat-reservation-manager/
 * 执行用时：404ms 在所有Go提交中击败了20.59%的用户
 * 占用内存：31.2MB 在所有Go提交中击败了73.51%的用户
**/
