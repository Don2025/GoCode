package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/iSwD2y/
// ------------------------剑指 Offer II Problem 65------------------------

type Trie struct {
	Child [26]*Trie
	Depth int
}

func Constructor() *Trie {
	return &Trie{[26]*Trie{}, 1}
}

func (t *Trie) Insert(word string) {
	curr := t
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		c := word[i] - 'a'
		if curr.Child[c] == nil {
			curr.Child[c] = Constructor()
		}
		curr.Child[c].Depth = curr.Depth + 1
		curr = curr.Child[c]
	}
}

func minimumLengthEncoding(words []string) int {
	t := Constructor()
	for _, w := range words {
		t.Insert(w)
	}
	var ans int
	var dfs func(node *Trie)
	dfs = func(node *Trie) {
		var hasChild bool
		for i := 0; i < 26; i++ {
			if node.Child[i] != nil {
				hasChild = true
				dfs(node.Child[i])
			}
		}
		if !hasChild {
			ans += node.Depth
		}
	}
	dfs(t)
	return ans
}

// ------------------------剑指 Offer II Problem 65------------------------
/*
 * https://leetcode.cn/problems/iSwD2y/
 * 执行用时：28ms 在所有Go提交中击败了76.16%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了61.59%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of strings separated by space:")
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", minimumLengthEncoding(words))
		Printf("Input a line of strings separated by space:")
	}
}
