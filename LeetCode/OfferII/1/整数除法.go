package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

// https://leetcode.cn/problems/xoh6Oh/
// ------------------------剑指 Offer II Problem 1------------------------
func divide(a int, b int) int {
	if b == 1 {
		return a
	}
	if b == -1 {
		if a == math.MinInt32 {
			return math.MaxInt32
		}
		return -a
	}
	if a == 0 {
		return 0
	}
	positive := (a ^ b) >= 0
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}
	var quotient int
	for a <= b {
		base, divisor := 1, b
		for a-divisor <= divisor {
			divisor <<= 1
			base <<= 1
		}
		quotient += base
		a -= divisor
	}
	if !positive {
		quotient = -quotient
	}
	return quotient
}

// ------------------------剑指 Offer II Problem 1------------------------
/*
 * https://leetcode.cn/problems/xoh6Oh/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了100.00%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var a, b int
		Sscanf(scanner.Text(), "%d/%d", &a, &b)
		Printf("Output: %v\n", divide(a, b))
	}
}
