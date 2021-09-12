package main

import (
	"bufio"
	"os"
)

func checkValidString(s string) bool {
	l, r := 0, 0
	for _, x := range s {
		if x == '(' {
			l++
			r++
		} else if x == ')' {
			l = max(0, l-1)
			if r--; r < 0 {
				return false
			}
		} else {
			l = max(l-1, 0)
			r++
		}
	}
	return l == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(checkValidString(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
