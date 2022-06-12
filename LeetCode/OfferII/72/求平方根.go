package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/jJ0w9p/
// ------------------------剑指 Offer II Problem 72------------------------
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 72------------------------
/*
 * https://leetcode.cn/problems/jJ0w9p/
 * 执行用时：4ms 在所有Go提交中击败了40.79%的用户
 * 占用内存：2MB 在所有Go提交中击败了61.37%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a number:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", mySqrt(n))
		Printf("Input a number:")
	}
}
