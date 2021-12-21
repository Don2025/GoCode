package main

import (
	"bufio"
	"os"
)

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := []byte(input.Text())
		reverseString(s)
		println(string(s))
	}
}

/*
 * 执行用时：28ms 在所有Go提交中击败了95.56%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了78.62%的用户
**/
