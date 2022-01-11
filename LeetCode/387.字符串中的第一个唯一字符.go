package main

import (
	"bufio"
	"os"
)

func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _, x := range s {
		cnt[x-'a']++
	}
	for i := range s {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(firstUniqChar(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了95.45%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了100.00%的用户
**/
