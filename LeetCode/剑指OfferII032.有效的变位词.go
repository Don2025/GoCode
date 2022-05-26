package main

import (
	"bufio"
	"os"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	var cnt1, cnt2 [26]int
	for _, x := range s {
		cnt1[x-'a']++
	}
	for _, x := range t {
		cnt2[x-'a']++
	}
	return cnt1 == cnt2
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		println(isAnagram(arr[0], arr[1]))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了100.00%的用户
**/
