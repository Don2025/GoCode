package main

import "math/rand"

// https://leetcode.cn/problems/random-pick-with-blacklist/
//------------------------Leetcode Problem 710------------------------

type Solution struct {
	B2W   map[int]int
	Bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	bound := n - m
	black := make(map[int]bool)
	for _, v := range blacklist {
		if v >= bound {
			black[v] = true
		}
	}
	b2w := make(map[int]int, m-len(black))
	w := bound
	for _, b := range blacklist {
		if b < bound {
			for black[w] {
				w++
			}
			b2w[b] = w
			w++
		}
	}
	return Solution{b2w, bound}
}

func (s *Solution) Pick() int {
	x := rand.Intn(s.Bound)
	if s.B2W[x] > 0 {
		return s.B2W[x]
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
//------------------------Leetcode Problem 710------------------------
/*
 * https://leetcode.cn/problems/random-pick-with-blacklist/
 * 执行用时：96ms 在所有Go提交中击败了74.62%的用户
 * 占用内存：11.1MB 在所有Go提交中击败了90.00%的用户
**/
