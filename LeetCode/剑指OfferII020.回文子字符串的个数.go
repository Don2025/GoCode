package main

import (
	"bufio"
	"os"
)

func countSubstrings(s string) int {
	n := len(s)
	var ans int
	for i := 0; i < 2*n-1; i++ {
		l, r := i>>1, i>>1+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(countSubstrings(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了58.64%的用户
**/
