package main

import (
	"math/rand"
	"sort"
)

// https://leetcode.cn/problems/2bCMpM/
// ------------------------剑指 Offer II Problem 71------------------------

type Solution struct {
	Pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (s *Solution) PickIndex() int {
	x := rand.Intn(s.Pre[len(s.Pre)-1]) + 1
	return sort.SearchInts(s.Pre, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

// ------------------------剑指 Offer II Problem 71------------------------
/*
 * https://leetcode.cn/problems/2bCMpM/
 * 执行用时：40ms 在所有Go提交中击败了98.97%的用户
 * 占用内存：7MB 在所有Go提交中击败了98.97%的用户
**/
