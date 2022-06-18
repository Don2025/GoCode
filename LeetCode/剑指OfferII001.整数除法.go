package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/two-sum/
//------------------------Leetcode Problem 1------------------------
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

//------------------------Leetcode Problem 1------------------------
/*
 * https://leetcode.cn/problems/two-sum/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了100.00%的用户
**/
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Split(input.Text(), "/")
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		Printf("Output: %v\n", divide(a, b))
	}
}
