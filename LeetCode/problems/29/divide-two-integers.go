package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/divide-two-integers/
//------------------------Leetcode Problem 29------------------------
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 { // Otherwise, the 991/992 test case: -2147483648 -1 would get Wrong Answer.
		return math.MaxInt32
	}
	var negative bool = (dividend ^ divisor) < 0
	dividend = abs(dividend)
	divisor = abs(divisor)
	var ans int
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			ans += 1 << i
			dividend -= divisor << i
		}
	}
	if negative {
		ans = -ans
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 29------------------------
/*
 * https://leetcode.cn/problems/divide-two-integers/
 * 执行用时：4ms 在所有Go提交中击败了55.36%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		dividend, _ := strconv.Atoi(arr[0])
		divisor, _ := strconv.Atoi(arr[1])
		println(divide(dividend, divisor))
	}
}
