package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/get-kth-magic-number-lcci/
func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		x2, x3, x5 := dp[p3]*3, dp[p5]*5, dp[p7]*7
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p3++
		}
		if dp[i] == x3 {
			p5++
		}
		if dp[i] == x5 {
			p7++
		}
	}
	return dp[k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
 * https://leetcode.cn/problems/get-kth-magic-number-lcci/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了72.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", getKthMagicNumber(k))
	}
}
