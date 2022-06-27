package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/valid-anagram/
//------------------------Leetcode Problem 242------------------------
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

//------------------------Leetcode Problem 242------------------------
/*
 * https://leetcode.cn/problems/valid-anagram/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了99.81%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		s, t := arr[0], arr[1]
		Printf("Output: %v\n", isAnagram(s, t))
	}
}
