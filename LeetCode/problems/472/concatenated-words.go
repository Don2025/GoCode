package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strings"
)

// https://leetcode.cn/problems/concatenated-words/
//------------------------Leetcode Problem 472------------------------

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func (root *Trie) insert(word string) {
	node := root
	for _, x := range word {
		x -= 'a'
		if node.Children[x] == nil {
			node.Children[x] = &Trie{}
		}
		node = node.Children[x]
	}
	node.IsEnd = true
}

func (root *Trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, x := range word {
		node = node.Children[x-'a']
		if node == nil {
			return false
		}
		if node.IsEnd && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var ans []string
	root := &Trie{}
	for _, word := range words {
		if word == "" {
			continue
		}
		if root.dfs(word) {
			ans = append(ans, word)
		} else {
			root.insert(word)
		}
	}
	return ans
}

//------------------------Leetcode Problem 472------------------------
/*
 * https://leetcode.cn/problems/concatenated-words/
 * 执行用时：112ms 在所有Go提交中击败了67.86%的用户
 * 占用内存：23.4MB 在所有Go提交中击败了85.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findAllConcatenatedWordsInADict(words))
	}
}
