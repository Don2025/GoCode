package main

import (
	"bufio"
	"os"
)

func reverseWords(s string) string {
	var i int
	ans := []byte(s)
	for i < len(s) {
		start := i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		l, r := start, i-1
		for l < r {
			ans[l], ans[r] = s[r], s[l]
			l++
			r--
		}
		for i < len(s) && s[i] == ' ' {
			i++
		}
	}
	return string(ans)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(reverseWords(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了80.51%的用户
**/
