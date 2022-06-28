package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/valid-perfect-square/
//------------------------Leetcode Problem 367------------------------
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)>>1
		square := mid * mid
		if square < num {
			l = mid + 1
		} else if square > num {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}

//------------------------Leetcode Problem 367------------------------
/*
 * https://leetcode.cn/problems/valid-perfect-square/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了58.72%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", isPerfectSquare(n))
	}
}
