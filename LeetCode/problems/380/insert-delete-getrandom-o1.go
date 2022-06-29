package _80

import "math/rand"

// https://leetcode.cn/problems/insert-delete-getrandom-o1/
//------------------------Leetcode Problem 380------------------------

type RandomizedSet struct {
	Nums     []int
	GetIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.GetIndex[val]; ok {
		return false
	}
	this.Nums = append(this.Nums, val)
	this.GetIndex[val] = len(this.Nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.GetIndex[val]; !ok {
		return false
	}
	index := this.GetIndex[val]
	lastIndex := len(this.Nums) - 1
	this.GetIndex[this.Nums[lastIndex]] = index
	this.Nums[index], this.Nums[lastIndex] = this.Nums[lastIndex], this.Nums[index]
	this.Nums = this.Nums[0:lastIndex]
	delete(this.GetIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.Nums[rand.Intn(len(this.Nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//------------------------Leetcode Problem 380------------------------
/*
 * https://leetcode.cn/problems/insert-delete-getrandom-o1/
 * 执行用时：160ms 在所有Go提交中击败了29.12%的用户
 * 占用内存：43.5MB 在所有Go提交中击败了99.73%的用户
**/
