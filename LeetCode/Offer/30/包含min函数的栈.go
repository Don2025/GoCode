package main

import "math"

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
// ------------------------剑指 Offer I Problem 30------------------------

type MinStack struct {
	stack, minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
// ------------------------剑指 Offer I Problem 30------------------------
/*
 * https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
 * 执行用时：16ms 在所有Go提交中击败了59.8%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了61.8%的用户
**/
