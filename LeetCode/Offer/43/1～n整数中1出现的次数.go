package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
// ------------------------剑指 Offer I Problem 43------------------------
func countDigitOne(n int) int {
	digitNum, ans := 1, 0
	low, high, cur := 0, n/10, n%10
	for high != 0 || cur != 0 {
		if cur < 1 {
			ans += high * digitNum
		} else if cur == 1 {
			ans += high*digitNum + low + 1
		} else {
			ans += (high + 1) * digitNum
		}
		low = low + cur*digitNum
		high, cur = high/10, high%10
		digitNum *= 10
	}
	return ans
}

// ------------------------剑指 Offer I Problem 43------------------------
/*
 * https://leetcode.cn/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了58.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", countDigitOne(n))
	}
}
