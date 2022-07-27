package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/calculate-money-in-leetcode-bank/
//------------------------Leetcode Problem 1716------------------------
func totalMoney(n int) int {
	weeks, days := n/7, n%7
	return 7*weeks*(1+weeks)/2 + 21*weeks + days*(1+2*weeks+days)/2
}

//------------------------Leetcode Problem 1716------------------------
/*
 * https://leetcode.cn/problems/calculate-money-in-leetcode-bank/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了67.74%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", totalMoney(n))
	}
}
