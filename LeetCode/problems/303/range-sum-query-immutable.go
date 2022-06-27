package _03

// https://leetcode.cn/problems/range-sum-query-immutable/
//------------------------Leetcode Problem 303------------------------

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, x := range nums {
		sums[i+1] = sums[i] + x
	}
	return NumArray{sums}
}

func (a *NumArray) SumRange(left int, right int) int {
	return a.sums[right+1] - a.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

//------------------------Leetcode Problem 303------------------------
/*
 * https://leetcode.cn/problems/range-sum-query-immutable/
 * 执行用时：24ms 在所有Go提交中击败了85.36%的用户
 * 占用内存：9.3MB 在所有Go提交中击败了59.59%的用户
**/
