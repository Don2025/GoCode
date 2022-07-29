package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/jJ0w9p/
// ------------------------剑指 Offer II Problem 72------------------------
func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := l + (r-l)>>1 + 1
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// ------------------------剑指 Offer II Problem 72------------------------
/*
 * https://leetcode.cn/problems/jJ0w9p/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		println(mySqrt(n))
	}
}
