package main

import (
	"bufio"
	"os"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	var cnt [26]int
	for _, x := range magazine {
		cnt[x-'a']++
	}
	for _, x := range ransomNote {
		cnt[x-'a']--
		if cnt[x-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		ransomNote, magazine := arr[0], arr[1]
		println(canConstruct(ransomNote, magazine))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了82.42%的用户
 * 占用内存：3.8MB 在所有Go提交中击败了100.00%的用户
**/
