package main

import (
	"bufio"
	"os"
)

func removeOuterParentheses(s string) string {
	var ans []rune
	var level int
	for _, x := range s {
		if x == ')' {
			level--
		}
		if level > 0 {
			ans = append(ans, x)
		}
		if x == '(' {
			level++
		}
	}
	return string(ans)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(removeOuterParentheses(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了64.47%的用户
**/
