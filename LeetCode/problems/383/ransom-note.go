package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/ransom-note/
//------------------------Leetcode Problem 383------------------------
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

//------------------------Leetcode Problem 383------------------------
/*
 * https://leetcode.cn/problems/ransom-note/
 * 执行用时：4ms 在所有Go提交中击败了82.42%的用户
 * 占用内存：3.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		ransomNote, magazine := arr[0], arr[1]
		Printf("Output: %v\n", canConstruct(ransomNote, magazine))
	}
}
