package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

// https://leetcode.cn/problems/reformat-the-string/
//------------------------Leetcode Problem 1417------------------------
func reformat(s string) string {
	var cntDigit int
	for _, x := range s {
		if unicode.IsDigit(x) {
			cntDigit++
		}
	}
	cntAlpha := len(s) - cntDigit
	if abs(cntDigit-cntAlpha) > 1 {
		return ""
	}
	flag := cntDigit > cntAlpha
	t := []byte(s)
	for i, j := 0, 1; i < len(t); i += 2 {
		if unicode.IsDigit(rune(t[i])) != flag {
			for unicode.IsDigit(rune(t[j])) != flag {
				j += 2
			}
			t[i], t[j] = t[j], t[i]
		}
	}
	return string(t)
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 1417------------------------
/*
 * https://leetcode.cn/problems/reformat-the-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", reformat(scanner.Text()))
	}
}
