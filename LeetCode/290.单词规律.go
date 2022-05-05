package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		println(wordPattern(arr[0], arr[1]))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了96.53%的用户
**/
