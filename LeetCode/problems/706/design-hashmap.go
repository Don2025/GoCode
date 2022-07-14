package _06

import "container/list"

// https://leetcode.cn/problems/design-hashmap/
//------------------------Leetcode Problem 706------------------------
const base = 769

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
//------------------------Leetcode Problem 706------------------------
/*
 * https://leetcode.cn/problems/design-hashmap/
 * 执行用时：84ms 在所有Go提交中击败了96.98%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了54.03%的用户
**/
