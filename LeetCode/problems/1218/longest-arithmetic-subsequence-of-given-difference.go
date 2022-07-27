package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/longest-arithmetic-subsequence-of-given-difference/
//------------------------Leetcode Problem 1218------------------------
func longestSubsequence(arr []int, difference int) int {
	var ans int
	dp := map[int]int{}
	for _, x := range arr {
		dp[x] = dp[x-difference] + 1
		ans = max(ans, dp[x])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1218------------------------
/*
 * https://leetcode.cn/problems/longest-arithmetic-subsequence-of-given-difference/
 * 执行用时：96ms 在所有Go提交中击败了82.61%的用户
 * 占用内存：9.1MB 在所有Go提交中击败了47.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		difference, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", longestSubsequence(arr, difference))
	}
}
