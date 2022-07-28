package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/dKk3P7/
// ------------------------剑指 Offer II Problem 32------------------------
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

// ------------------------剑指 Offer II Problem 32------------------------
/*
 * https://leetcode.cn/problems/dKk3P7/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var s, t string
		Sscanf(scanner.Text(), "%s %s", &s, &t)
		Printf("Output: %v\n", isAnagram(s, t))
	}
}
