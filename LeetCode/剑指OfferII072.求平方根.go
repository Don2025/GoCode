package main

import (
	"bufio"
	"os"
	"strconv"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(mySqrt(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了100.00%的用户
**/
