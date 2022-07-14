package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	root := &Trie{}
	for _, d := range dictionary {
		root.Insert(d)
	}
	words := strings.Fields(sentence)
	for i, word := range words {
		if w, ok := root.Search(word); ok {
			words[i] = w
		}
	}
	return strings.Join(words, " ")
}

type Trie struct {
	Children [26]*Trie
	IsWord   bool
}

func (root *Trie) Insert(word string) {
	node := root
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &Trie{}
		}
		node = node.Children[idx]
	}
	node.IsWord = true
}

func (root *Trie) Search(word string) (string, bool) {
	node := root
	for i, c := range word {
		if node.IsWord {
			return word[:i], true
		}
		idx := c - 'a'
		if node.Children[idx] == nil {
			break
		}
		node = node.Children[idx]
	}
	return word, node.IsWord
}

//------------------------Leetcode Problem 648------------------------
/*
 * https://leetcode.cn/problems/replace-words/
 * 执行用时：20ms 在所有Go提交中击败了83.06%的用户
 * 占用内存：20.5MB 在所有Go提交中击败了29.84%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dictionary := strings.Fields(scanner.Text())
		scanner.Scan()
		sentence := scanner.Text()
		Printf("Output: %v\n", replaceWords(dictionary, sentence))
	}
}
