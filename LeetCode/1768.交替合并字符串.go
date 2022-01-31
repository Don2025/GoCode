package main

import (
	"bufio"
	"os"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	sb := strings.Builder{}
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		w1 := input.Text()
		input.Scan()
		w2 := input.Text()
		println(mergeAlternately(w1, w2))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了86.75%的用户
**/
