package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/reformat-phone-number/
//------------------------Leetcode Problem 1694------------------------
func reformatNumber(number string) string {
	s := strings.ReplaceAll(number, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	ans := []string{}
	i := 0
	for ; i+4 < len(s); i += 3 {
		ans = append(ans, s[i:i+3])
	}
	s = s[i:]
	if len(s) < 4 {
		ans = append(ans, s)
	} else {
		ans = append(ans, s[:2], s[2:])
	}
	return strings.Join(ans, "-")
}

//------------------------Leetcode Problem 1694------------------------
/*
 * https://leetcode.cn/problems/reformat-phone-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了38.56%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %s\n", reformatNumber(scanner.Text()))
	}
}
