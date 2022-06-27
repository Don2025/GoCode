package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/word-pattern/
//------------------------Leetcode Problem 290------------------------
func wordPattern(pattern string, s string) bool {
	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		x := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != x || ch2word[x] != "" && ch2word[x] != word {
			return false
		}
		word2ch[word] = x
		ch2word[x] = word
	}
	return true
}

//------------------------Leetcode Problem 290------------------------
/*
 * https://leetcode.cn/problems/word-pattern/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了96.53%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %v\n", wordPattern(arr[0], arr[1]))
	}
}
