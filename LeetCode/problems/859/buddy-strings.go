package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/buddy-strings/
//------------------------Leetcode Problem 859------------------------
func buddyStrings(s string, goal string) bool {
	if len(s) == len(goal) {
		if s == goal {
			cnt := [26]int{}
			for _, x := range s {
				cnt[x-'a']++
				if cnt[x-'a'] > 1 {
					return true
				}
			}
			return false
		}
		x, y := -1, -1
		for i := range s {
			if s[i] != goal[i] {
				if x == -1 {
					x = i
				} else if y == -1 {
					y = i
				} else {
					return false
				}
			}
		}
		return y != -1 && s[x] == goal[y] && s[y] == goal[x]
	}
	return false
}

//------------------------Leetcode Problem 859------------------------
/*
 * https://leetcode.cn/problems/buddy-strings/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了58.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		s, goal := arr[0], arr[1]
		Printf("Output: %t\n", buddyStrings(s, goal))
	}
}
