package main

import (
	"bufio"
	"os"
	"regexp"
)

// https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
// ------------------------剑指 Offer I Problem 20------------------------
func isNumber(s string) bool {
	ans, _ := regexp.MatchString(`^\s*[+-]?((\d*\.\d+)|(\d+\.?\d*))([eE][+-]?\d+)?\s*$`, s)
	return ans
}

// ------------------------剑指 Offer I Problem 20------------------------
/*
 * https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
 * 执行用时：20ms 在所有Go提交中击败了6.35%的用户
 * 占用内存：7MB 在所有Go提交中击败了5.69%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		println(isNumber(scanner.Text()))
	}
}
