package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/first-unique-character-in-a-string/
//------------------------Leetcode Problem 387------------------------
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", firstUniqChar(scanner.Text()))
	}
}

//------------------------Leetcode Problem 387------------------------
/*
 * https://leetcode.cn/problems/first-unique-character-in-a-string/
 * 执行用时：4ms 在所有Go提交中击败了95.45%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了100.00%的用户
**/
