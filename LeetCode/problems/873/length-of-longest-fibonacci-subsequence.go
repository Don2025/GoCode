package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/
//------------------------Leetcode Problem 873------------------------
func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	indices := make(map[int]int, n)
	for i, x := range arr {
		indices[x] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i, x := range arr {
		for j := n - 1; j >= 0 && arr[j]<<1 > x; j-- {
			if k, ok := indices[x-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
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

//------------------------Leetcode Problem 873------------------------
/*
 * https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/
 * 执行用时：224ms 在所有Go提交中击败了24.68%的用户
 * 占用内存：17.4MB 在所有Go提交中击败了57.14%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", lenLongestFibSubseq(arr))
	}
}
