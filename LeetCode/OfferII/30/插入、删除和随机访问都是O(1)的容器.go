package main

import "math/rand"

type RandomizedSet struct {
	Nums    []int
	Indices map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.Indices[val]; ok {
		return false
	}
	rs.Indices[val] = len(rs.Nums)
	rs.Nums = append(rs.Nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.Indices[val]
	if !ok {
		return false
	}
	last := len(rs.Nums) - 1
	rs.Nums[id] = rs.Nums[last]
	rs.Indices[rs.Nums[id]] = id
	rs.Nums = rs.Nums[:last]
	delete(rs.Indices, val)
	return true
}

/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	return rs.Nums[rand.Intn(len(rs.Nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// ------------------------剑指 Offer II Problem 30------------------------
/*
 * https://leetcode.cn/problems/FortPu/
 * 执行用时：132ms 在所有Go提交中击败了93.17%的用户
 * 占用内存：55.4MB 在所有Go提交中击败了36.10%的用户
**/
