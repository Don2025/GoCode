package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/Q91FMA/
// ------------------------剑指 Offer II Problem 93------------------------
func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 2; i < n; i++ {
		for j, k := 0, i-1; j < k; {
			if arr[j]+arr[k] == arr[i] {
				if dp[j][k] == 0 {
					dp[k][i] = 3
				} else {
					dp[k][i] = max(dp[j][k]+1, dp[k][i])
				}
				ans = max(ans, dp[k][i])
				j++
				k--
			} else if arr[j]+arr[k] < arr[i] {
				j++
			} else {
				k--
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 93------------------------
/*
 * https://leetcode.cn/problems/Q91FMA/
 * 执行用时：76ms 在所有Go提交中击败了95.71%的用户
 * 占用内存：18.3MB 在所有Go提交中击败了13.57%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", lenLongestFibSubseq(nums))
		Printf("Input a line of numbers separated by space:")
	}
}
