package main

import (
	"bufio"
	"os"
	"strings"
)

func isAnagram(s string, t string) bool {
	cnt := [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for _, x := range cnt {
		if x != 0 {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		s, t := arr[0], arr[1]
		println(isAnagram(s, t))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了99.81%的用户
**/
