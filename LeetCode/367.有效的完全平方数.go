package main

import (
	"bufio"
	"os"
	"strconv"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(isPerfectSquare(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了58.72%的用户
**/
