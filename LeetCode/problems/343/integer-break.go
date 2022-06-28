package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/integer-break/
//------------------------Leetcode Problem 343------------------------
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		var curMax int
		for j := 1; j < i; j++ {
			curMax = max(curMax, j*(i-j), j*dp[i-j])
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(a ...int) int {
	val := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > val {
			val = a[i]
		}
	}
	return val
}

//------------------------Leetcode Problem 343------------------------
/*
 * https://leetcode.com/problems/integer-break/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了93.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", integerBreak(n))
	}
}
