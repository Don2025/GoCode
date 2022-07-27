package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/latest-time-by-replacing-hidden-digits/
//------------------------Leetcode Problem 1736------------------------
func maximumTime(time string) string {
	ans := []byte(time)
	if ans[0] == '?' {
		if ans[1] >= '4' && ans[1] <= '9' {
			ans[0] = '1'
		} else {
			ans[0] = '2'
		}
	}
	if ans[1] == '?' {
		if ans[0] == '2' {
			ans[1] = '3'
		} else {
			ans[1] = '9'
		}
	}
	if ans[3] == '?' {
		ans[3] = '5'
	}
	if ans[4] == '?' {
		ans[4] = '9'
	}
	return string(ans)
}

//------------------------Leetcode Problem 1736------------------------
/*
 * https://leetcode.cn/problems/latest-time-by-replacing-hidden-digits/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了70.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", maximumTime(scanner.Text()))
	}
}
