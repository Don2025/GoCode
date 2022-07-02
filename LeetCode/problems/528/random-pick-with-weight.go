package _28

import (
	"math/rand"
	"sort"
)

// https://leetcode.cn/problems/random-pick-with-weight/
//------------------------Leetcode Problem 528------------------------

type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
//------------------------Leetcode Problem 528------------------------
/*
 * https://leetcode.cn/problems/random-pick-with-weight/
 * 执行用时：48ms 在所有Go提交中击败了63.16%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了53.95%的用户
**/
