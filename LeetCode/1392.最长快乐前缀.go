package main

import (
	"bufio"
	"os"
)

func longestPrefix(s string) string {
	n := len(s)
	next := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return s[0:next[n-1]]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(longestPrefix(input.Text()))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了85.94%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了37.50%的用户
**/
