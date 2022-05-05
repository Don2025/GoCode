package main

import "math"

type MinStack struct {
	Nums, MinValue []int
}

func Constructor() MinStack {
	return MinStack{MinValue: []int{math.MaxInt32}}
}

func (this *MinStack) Push(val int) {
	this.Nums = append(this.Nums, val)
	this.MinValue = append(this.MinValue, min(this.MinValue[len(this.MinValue)-1], val))
}

func (this *MinStack) Pop() {
	this.Nums = this.Nums[:len(this.Nums)-1]
	this.MinValue = this.MinValue[:len(this.MinValue)-1]
}

func (this *MinStack) Top() int {
	return this.Nums[len(this.Nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinValue[len(this.MinValue)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*
 * 执行用时：12ms 在所有Go提交中击败了95.71%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了73.12%的用户
**/
