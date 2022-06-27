package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/burst-balloons/
//------------------------Leetcode Problem 312------------------------
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		a[i] = nums[i-1]
	}
	a[0], a[len(a)-1] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < len(dp); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+a[i]*a[j]*a[k])
			}
		}
	}
	return dp[0][len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 312------------------------
/*
 * https://leetcode.cn/problems/burst-balloons/
 * 执行用时：40ms 在所有Go提交中击败了79.44%的用户
 * 占用内存：5.4MB 在所有Go提交中击败了42.68%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxCoins(nums))
	}
}
