package main

import (
	"bufio"
	"os"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	var cnt int
	ans := []rune(s)
	for i := range s {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			cnt--
			if cnt < 0 {
				ans[i] = '0'
				cnt = 0
			}
		}
	}
	cnt = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			cnt++
		} else if s[i] == '(' {
			cnt--
			if cnt < 0 {
				ans[i] = '0'
				cnt = 0
			}
		}
	}
	return strings.ReplaceAll(string(ans), "0", "")
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(minRemoveToMakeValid(input.Text()))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了57.76%的用户
 * 占用内存：6MB 在所有Go提交中击败了86.34%的用户
**/
