package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var pick int

// https://leetcode.cn/problems/guess-number-higher-or-lower/
//------------------------Leetcode Problem 374------------------------
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(n int) int {
	if n > pick {
		return -1
	} else if n < pick {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		ans := guess(mid)
		if ans > 0 {
			l = mid + 1
		} else if ans < 0 {
			r = mid - 1
		} else {
			return mid
		}
	}
	return 0
}

//------------------------Leetcode Problem 374------------------------
/*
 * https://leetcode.cn/problems/guess-number-higher-or-lower/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了91.23%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		pick, _ = strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", guessNumber(n))
	}
}
