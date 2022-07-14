package main

// https://leetcode.cn/problems/implement-magic-dictionary/
//------------------------Leetcode Problem 676------------------------

type Trie struct {
	Children [26]*Trie
	IsWord   bool
}

type MagicDictionary struct {
	*Trie
}

func Constructor() MagicDictionary {
	return MagicDictionary{&Trie{}}
}

func (d *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := d.Trie
		for _, ch := range word {
			idx := ch - 'a'
			if cur.Children[idx] == nil {
				cur.Children[idx] = &Trie{}
			}
			cur = cur.Children[idx]
		}
		cur.IsWord = true
	}
}

func (d *MagicDictionary) Search(searchWord string) bool {
	var dfs func(*Trie, string, bool) bool
	dfs = func(node *Trie, word string, modified bool) bool {
		if len(word) == 0 {
			return modified && node.IsWord
		}
		idx := word[0] - 'a'
		if node.Children[idx] != nil && dfs(node.Children[idx], word[1:], modified) {
			return true
		}
		if !modified {
			for i, child := range node.Children {
				if i != int(idx) && child != nil && dfs(child, word[1:], true) {
					return true
				}
			}
		}
		return false
	}
	return dfs(d.Trie, searchWord, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
//------------------------Leetcode Problem 676------------------------
/*
 * https://leetcode.cn/problems/implement-magic-dictionary/
 * 执行用时：36ms 在所有Go提交中击败了41.86%的用户
 * 占用内存：8MB 在所有Go提交中击败了23.26%的用户
**/
