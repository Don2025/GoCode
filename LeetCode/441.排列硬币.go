package main

import (
	"bufio"
	"os"
	"strconv"
)

func arrangeCoins(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)>>1
		num := mid * (mid + 1)
		if num == n*2 {
			return mid
		} else if num < n*2 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(arrangeCoins(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/
