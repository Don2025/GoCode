package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/MPnaiL/
// ------------------------剑指 Offer II Problem 14------------------------
func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}
	var cnt1, cnt2 [26]int
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n1; i < n2; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n1]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

// ------------------------剑指 Offer II Problem 14------------------------
/*
 * https://leetcode.cn/problems/MPnaiL/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了77.48%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var s1, s2 string
		Sscanf(scanner.Text(), "%s %s", &s1, &s2)
		Printf("Output: %v\n", checkInclusion(s1, s2))
	}
}
