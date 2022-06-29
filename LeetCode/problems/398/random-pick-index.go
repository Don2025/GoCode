package _98

import "math/rand"

// https://leetcode.cn/problems/random-pick-index/
//------------------------Leetcode Problem 398------------------------

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

func (nums Solution) Pick(target int) int {
	var cnt, ans int
	for i, num := range nums {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
//------------------------Leetcode Problem 398------------------------
/*
 * https://leetcode.cn/problems/random-pick-index/
 * 执行用时：68ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：8MB 在所有Go提交中击败了97.74%的用户
**/
