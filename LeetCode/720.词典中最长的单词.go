package main

import (
	"bufio"
	"os"
	"strings"
)

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, x := range word {
		x -= 'a'
		if node.children[x] == nil {
			node.children[x] = &Trie{}
		}
		node = node.children[x]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, x := range word {
		x -= 'a'
		if node.children[x] == nil || !node.children[x].isEnd {
			return false
		}
		node = node.children[x]
	}
	return true
}

func longestWord(words []string) string {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	var ans string
	for _, word := range words {
		if t.Search(word) && (len(word) > len(ans) || len(word) == len(ans) && word < ans) {
			ans = word
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := strings.Fields(input.Text())
		println(longestWord(words))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了83.48%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了17.39%的用户
**/
