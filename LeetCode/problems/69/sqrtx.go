package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/sqrtx/
//------------------------Leetcode Problem 69------------------------
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

//------------------------Leetcode Problem 69------------------------
/*
 * https://leetcode.cn/problems/sqrtx/
 * 执行用时：4ms 在所有Go提交中击败了50.27%的用户
 * 占用内存：2MB 在所有Go提交中击败了99.96%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", mySqrt(x))
	}
}
