package main

import (
	"bufio"
	"os"
	"strings"
)

func toGoatLatin(sentence string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	var ans string
	for i, word := range strings.Split(sentence, " ") {
		if _, ok := vowels[word[0]]; !ok {
			word = word[1:] + word[:1]
		}
		ans += word + "ma" + strings.Repeat("a", i+1) + " "
	}
	return strings.TrimRight(ans, " ")
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(toGoatLatin(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了28.21%的用户
**/
