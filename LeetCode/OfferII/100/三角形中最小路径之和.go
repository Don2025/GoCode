package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/IlPe0q/
// ------------------------剑指 Offer II Problem 100------------------------
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	ans := dp[0]
	for i := 1; i < n; i++ {
		ans = min(ans, dp[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 100------------------------
/*
 * https://leetcode.cn/problems/IlPe0q/
 * 执行用时：4ms 在所有Go提交中击败了86.88%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了76.88%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		triangle := make([][]int, n)
		for i := range triangle {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			triangle[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", minimumTotal(triangle))
		Printf("Input the number of rows of matrix:")
	}
}
