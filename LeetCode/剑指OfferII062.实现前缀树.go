package main

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, x := range word {
		x -= 'a'
		if node.Children[x] == nil {
			node.Children[x] = &Trie{}
		}
		node = node.Children[x]
	}
	node.IsEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, x := range prefix {
		x -= 'a'
		if node.Children[x] == nil {
			return nil
		}
		node = node.Children[x]
	}
	return node
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
 * 执行用时：40ms 在所有Go提交中击败了96.12%的用户
 * 占用内存：18MB 在所有Go提交中击败了42.24%的用户
**/
