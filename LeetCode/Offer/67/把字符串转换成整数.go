package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strings"
)

// https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
// ------------------------剑指 Offer I Problem 67------------------------
func strToInt(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * result
}

// ------------------------剑指 Offer I Problem 67------------------------
/*
 * https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
 * 执行用时：4ms 在所有Go提交中击败了40.17%的用户
 * 占用内存：2MB 在所有Go提交中击败了68.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", strToInt(scanner.Text()))
	}
}
