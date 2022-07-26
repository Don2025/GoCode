package main

import "math/rand"

// https://leetcode.cn/problems/design-skiplist/
//------------------------Leetcode Problem 1206------------------------
const maxLevel = 32
const pFactor = 0.25

type SkiplistNode struct {
	Val     int
	Forward []*SkiplistNode
}

type Skiplist struct {
	Head  *SkiplistNode
	Level int
}

func Constructor() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

func (sl *Skiplist) randomLevel() int {
	level := 1
	for rand.Float64() < pFactor && level < maxLevel {
		level++
	}
	return level
}

func (sl *Skiplist) Search(target int) bool {
	cur := sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		for cur.Forward[i] != nil && cur.Forward[i].Val < target {
			cur = cur.Forward[i]
		}
	}
	cur = cur.Forward[0]
	return cur != nil && cur.Val == target
}

func (sl *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = sl.Head
	}
	cur := sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		for cur.Forward[i] != nil && cur.Forward[i].Val < num {
			cur = cur.Forward[i]
		}
		update[i] = cur
	}
	level := sl.randomLevel()
	sl.Level = max(sl.Level, level)
	newNode := &SkiplistNode{num, make([]*SkiplistNode, level)}
	for i, node := range update[:level] {
		newNode.Forward[i] = node.Forward[i]
		node.Forward[i] = newNode
	}
}

func (sl *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	cur := sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		for cur.Forward[i] != nil && cur.Forward[i].Val < num {
			cur = cur.Forward[i]
		}
		update[i] = cur
	}
	cur = cur.Forward[0]
	if cur == nil || cur.Val != num {
		return false
	}
	for i := 0; i < sl.Level && update[i].Forward[i] == cur; i++ {
		update[i].Forward[i] = cur.Forward[i]
	}
	for sl.Level > 1 && sl.Head.Forward[sl.Level-1] == nil {
		sl.Level--
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
//------------------------Leetcode Problem 1206------------------------
/*
 * https://leetcode.cn/problems/design-skiplist/
 * 执行用时：52ms 在所有Go提交中击败了89.27%的用户
 * 占用内存：9.9MB 在所有Go提交中击败了21.46%的用户
**/
