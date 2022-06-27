package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/ugly-number-ii/
//------------------------Leetcode Problem 264------------------------
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 264------------------------
/*
 * https://leetcode.cn/problems/ugly-number-ii/
 * 执行用时：4ms 在所有Go提交中击败了62.43%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了78.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", nthUglyNumber(n))
	}
}
