package main

import "math/rand"

type Solution struct {
	m, n, total int
	mp          map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m * n, map[int]int{}}
}

func (this *Solution) Flip() []int {
	var ans []int
	x := rand.Intn(this.total)
	this.total--
	if y, ok := this.mp[x]; ok {
		ans = []int{y / this.n, y % this.n}
	} else {
		ans = []int{x / this.n, x % this.n}
	}
	if y, ok := this.mp[this.total]; ok {
		this.mp[x] = y
	} else {
		this.mp[x] = this.total
	}
	return ans
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	this.mp = map[int]int{}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */

/*
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了72.73%的用户
**/
