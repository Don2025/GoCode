package main

import (
	"bufio"
	"os"
	"strconv"
)

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		println(reverseLeftWords(s, n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了96.29%的用户
**/
