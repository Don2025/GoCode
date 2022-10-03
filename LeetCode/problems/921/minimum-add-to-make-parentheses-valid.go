package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
//------------------------Leetcode Problem 921------------------------
func minAddToMakeValid(s string) int {
	var cnt, ans int
	for _, c := range s {
		if c == '(' {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				ans++
			}
		}
	}
	return ans + cnt
}

//------------------------Leetcode Problem 921------------------------
/*
 * https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了21.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %d\n", minAddToMakeValid(scanner.Text()))
	}
}
