package main

type NumArray struct {
	Nums, Tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (na *NumArray) add(index, val int) {
	for ; index < len(na.Tree); index += index & -index {
		na.Tree[index] += val
	}
}

func (na *NumArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += na.Tree[index]
	}
	return
}

func (na *NumArray) Update(index, val int) {
	na.add(index+1, val-na.Nums[index])
	na.Nums[index] = val
}

func (na *NumArray) SumRange(left, right int) int {
	return na.prefixSum(right+1) - na.prefixSum(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

/*
 * 执行用时：456ms 在所有Go提交中击败了91.7S8%的用户
 * 占用内存：19MB 在所有Go提交中击败了97.26%的用户
**/
