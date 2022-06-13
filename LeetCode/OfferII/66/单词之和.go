package main

// https://leetcode.cn/problems/z1R5dt/
// ------------------------剑指 Offer II Problem 66------------------------

type TrieNode struct {
	Children [26]*TrieNode
	Val      int
}

type MapSum struct {
	Node *TrieNode
	Cnt  map[string]int
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}, map[string]int{}}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if m.Cnt[key] > 0 {
		delta -= m.Cnt[key]
	}
	m.Cnt[key] = val
	node := m.Node
	for _, ch := range key {
		ch -= 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &TrieNode{}
		}
		node = node.Children[ch]
		node.Val += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	node := m.Node
	for _, ch := range prefix {
		ch -= 'a'
		if node.Children[ch] == nil {
			return 0
		}
		node = node.Children[ch]
	}
	return node.Val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

// ------------------------剑指 Offer II Problem 66------------------------
/*
 * https://leetcode.cn/problems/z1R5dt/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了63.56%的用户
**/
