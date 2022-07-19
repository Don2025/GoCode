package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/di-string-match/
//------------------------Leetcode Problem 942------------------------
func diStringMatch(s string) []int {
	i, d := 0, len(s)
	var ans []int
	for _, x := range s {
		if x == 'I' {
			ans = append(ans, i)
			i++
		} else if x == 'D' {
			ans = append(ans, d)
			d--
		}
	}
	ans = append(ans, d)
	return ans
}

//------------------------Leetcode Problem 942------------------------
/*
 * https://leetcode.cn/problems/di-string-match/
 * 执行用时：4ms 在所有Go提交中击败了80.88%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了42.65%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ans := diStringMatch(scanner.Text())
		Printf("Output: %v\n", ans)
	}
}
