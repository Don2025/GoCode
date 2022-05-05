package main

import (
	"bufio"
	"os"
	"strconv"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(countDigitOne(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了58.31%的用户
**/
