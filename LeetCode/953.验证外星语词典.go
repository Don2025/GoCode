package main

import (
	"bufio"
	"os"
	"strings"
)

func isAlienSorted(words []string, order string) bool {
	book := [26]int{}
	for i := 0; i < len(order); i++ {
		book[order[i]-'a'] = i
	}
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		prefix := true
		for j := 0; j < min(len(w1), len(w2)); j++ {
			if w1[j] != w2[j] {
				if book[w1[j]-'a'] > book[w2[j]-'a'] {
					return false
				}
				prefix = false
				break
			}
		}
		if prefix && len(w1) > len(w2) {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := strings.Fields(input.Text())
		input.Scan()
		order := input.Text()
		println(isAlienSorted(words, order))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了100.00%的用户
**/
