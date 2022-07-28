package main

type DLinkedNode struct {
	Key, Value int
	Prev, Next *DLinkedNode
}

type LRUCache struct {
	Size, Capacity int
	Cache          map[int]*DLinkedNode
	Head, Tail     *DLinkedNode
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{Key: key, Value: val}
}

func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.Prev = lru.Head
	node.Next = lru.Head.Next
	lru.Head.Next.Prev = node
	lru.Head.Next = node
}

func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() (node *DLinkedNode) {
	node = lru.Tail.Prev
	lru.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Cache:    map[int]*DLinkedNode{},
		Head:     initDLinkedNode(0, 0),
		Tail:     initDLinkedNode(0, 0),
		Capacity: capacity,
	}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
	return l
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.Cache[key]; !ok {
		return -1
	}
	node := lru.Cache[key]
	lru.moveToHead(node)
	return node.Value
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.Cache[key]; !ok {
		node := initDLinkedNode(key, value)
		lru.Cache[key] = node
		lru.addToHead(node)
		lru.Size++
		if lru.Size > lru.Capacity {
			removed := lru.removeTail()
			delete(lru.Cache, removed.Key)
			lru.Size--
		}
	} else {
		node := lru.Cache[key]
		node.Value = value
		lru.moveToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// ------------------------剑指 Offer II Problem 31------------------------
/*
 * https://leetcode.cn/problems/OrIXps/
 * 执行用时：460ms 在所有Go提交中击败了50.73%的用户
 * 占用内存：73.1MB 在所有Go提交中击败了86.86%的用户
**/
