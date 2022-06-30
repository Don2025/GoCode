package main

import (
	"math"
	"math/rand"
)

// https://leetcode.cn/problems/generate-random-point-in-a-circle/
//------------------------Leetcode Problem 478------------------------

type Solution struct {
	Radius, X, Y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (s *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64())
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{s.X + r*cos*s.Radius, s.Y + r*sin*s.Radius}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
//------------------------Leetcode Problem 478------------------------
/*
 * https://leetcode.cn/problems/generate-random-point-in-a-circle/
 * 执行用时：44ms 在所有Go提交中击败了87.10%的用户
 * 占用内存：12.7MB 在所有Go提交中击败了22.58%的用户
**/
