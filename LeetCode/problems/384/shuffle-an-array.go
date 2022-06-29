package _84

import "math/rand"

// https://leetcode.cn/problems/shuffle-an-array/
//------------------------Leetcode Problem 384------------------------

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
//------------------------Leetcode Problem 384------------------------
/*
 * https://leetcode.cn/problems/shuffle-an-array/
 * 执行用时：44ms 在所有Go提交中击败了38.64%的用户
 * 占用内存：9MB 在所有Go提交中击败了85.23%的用户
**/
