package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

// https://leetcode.cn/problems/range-module/
//------------------------Leetcode Problem 715------------------------

type RangeModule struct {
	*redblacktree.Tree
}

func Constructor() RangeModule {
	return RangeModule{redblacktree.NewWithIntComparator()}
}

func (rm *RangeModule) AddRange(left int, right int) {
	if node, ok := rm.Floor(left); ok {
		r := node.Value.(int)
		if r >= right {
			return
		}
		if r >= left {
			left = node.Key.(int)
			rm.Remove(left)
		}
	}
	for node, ok := rm.Ceiling(left); ok && node.Key.(int) <= right; node, ok = rm.Ceiling(left) {
		right = max(right, node.Value.(int))
		rm.Remove(node.Key)
	}
	rm.Put(left, right)
}

func (rm *RangeModule) QueryRange(left int, right int) bool {
	node, ok := rm.Floor(left)
	return ok && node.Value.(int) >= right
}

func (rm *RangeModule) RemoveRange(left int, right int) {
	if node, ok := rm.Floor(left); ok {
		l, r := node.Key.(int), node.Value.(int)
		if r >= right {
			if l == left {
				rm.Remove(l)
			} else {
				node.Value = left
			}
			if right != r {
				rm.Put(right, r)
			}
			return
		}
		if r > left {
			node.Value = left
		}
	}
	for node, ok := rm.Ceiling(left); ok && node.Key.(int) <= right; node, ok = rm.Ceiling(left) {
		r := node.Value.(int)
		rm.Remove(node.Key)
		if r > right {
			rm.Put(right, r)
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
//------------------------Leetcode Problem 715------------------------
/*
 * https://leetcode.cn/problems/range-module/
 * 执行用时：184ms 在所有Go提交中击败了77.50%的用户
 * 占用内存：8.6MB 在所有Go提交中击败了62.50%的用户
**/
