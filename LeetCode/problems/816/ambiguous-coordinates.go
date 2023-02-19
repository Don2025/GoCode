package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/ambiguous-coordinates/
//------------------------Leetcode Problem 816------------------------
func ambiguousCoordinates(s string) (ans []string) {
	s = s[1 : len(s)-1]
	getPos := func(s string) (pos []string) {
		if s[0] != '0' || s == "0" {
			pos = append(pos, s)
		}
		for i := 1; i < len(s); i++ {
			if i != 1 && s[0] == '0' || s[len(s)-1] == '0' {
				continue
			}
			pos = append(pos, s[:i]+"."+s[i:])
		}
		return
	}
	for i := 1; i < len(s); i++ {
		left, right := getPos(s[:i]), getPos(s[i:])
		if len(left) == 0 || len(right) == 0 {
			continue
		}
		for _, l := range left {
			for _, r := range right {
				ans = append(ans, "("+l+", "+r+")")
			}
		}
	}
	return
}

//------------------------Leetcode Problem 816------------------------
/*
 * https://leetcode.cn/problems/ambiguous-coordinates/
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", ambiguousCoordinates(scanner.Text()))
	}
}
