package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/UhWRSj/
// ------------------------剑指 Offer II Problem 63------------------------

type Trie struct {
	Child [26]*Trie
	IsEnd bool
}

func Constructor() *Trie {
	return &Trie{[26]*Trie{}, false}
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, c := range word {
		if curr.Child[c-'a'] == nil {
			curr.Child[c-'a'] = Constructor()
		}
		curr = curr.Child[c-'a']
	}
	curr.IsEnd = true
}

func (t *Trie) FindPrefix(word string) string {
	curr := t
	var sb strings.Builder
	for _, c := range word {
		if curr.Child[c-'a'] == nil {
			return word
		}
		sb.WriteRune(c)
		curr = curr.Child[c-'a']
		if curr.IsEnd {
			return sb.String()
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	tree := Constructor()
	for _, d := range dictionary {
		tree.Insert(d)
	}
	ans := strings.Fields(sentence)
	for i := range ans {
		ans[i] = tree.FindPrefix(ans[i])
	}
	return strings.Join(ans, " ")
}

// ------------------------剑指 Offer II Problem 63------------------------
/*
 * https://leetcode.cn/problems/UhWRSj/
 * 执行用时：24ms 在所有Go提交中击败了44.68%的用户
 * 占用内存：20.5MB 在所有Go提交中击败了42.55%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of strings separated by space:")
	for scanner.Scan() {
		dictionary := strings.Fields(scanner.Text())
		Printf("Input a line of string(sentence):")
		scanner.Scan()
		sentence := scanner.Text()
		Printf("Output: %v\n", replaceWords(dictionary, sentence))
		Printf("Input a line of strings separated by space:")
	}
}
