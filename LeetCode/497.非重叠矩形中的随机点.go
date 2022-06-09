package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	Rects [][]int
	Sum   []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}
	return Solution{rects, sum}
}

func (s *Solution) Pick() []int {
	k := rand.Intn(s.Sum[len(s.Sum)-1])
	rectIndex := sort.SearchInts(s.Sum, k+1) - 1
	r := s.Rects[rectIndex]
	a, b, y := r[0], r[1], r[3]
	da := (k - s.Sum[rectIndex]) / (y - b + 1)
	db := (k - s.Sum[rectIndex]) % (y - b + 1)
	return []int{a + da, b + db}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

/*
 * 执行用时：40ms 在所有Go提交中击败了11.12%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了62.60%的用户
**/
