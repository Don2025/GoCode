package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/first-bad-version/
//------------------------Leetcode Problem 278------------------------
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)>>1
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

//------------------------Leetcode Problem 278------------------------
/*
 * https://leetcode.cn/problems/first-bad-version/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了91.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", firstBadVersion(n))
	}
}
