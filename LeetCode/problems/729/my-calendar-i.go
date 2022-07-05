package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

// https://leetcode.cn/problems/my-calendar-i/
//------------------------Leetcode Problem 729------------------------

type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil)
	return MyCalendar{t}
}

func (c *MyCalendar) Book(start int, end int) bool {
	node, _ := c.Ceiling(end)
	it := c.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		c.Put(start, end)
		return true
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
//------------------------Leetcode Problem 729------------------------
/*
 * https://leetcode.cn/problems/my-calendar-i/
 * 执行用时：72ms 在所有Go提交中击败了64.07%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了52.10%的用户
**/
