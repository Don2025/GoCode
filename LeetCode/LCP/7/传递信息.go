package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/chuan-di-xin-xi/
// ------------------------LeetCode Cup Problem 7------------------------
func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < k; i++ {
		for _, r := range relation {
			src, dst := r[0], r[1]
			dp[i+1][dst] += dp[i][src]
		}
	}
	return dp[k][n-1]
}

// ------------------------LeetCode Cup Problem 7------------------------
/*
 * https://leetcode.cn/problems/chuan-di-xin-xi/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了84.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		cnt, _ := strconv.Atoi(scanner.Text())
		relation := make([][]int, cnt)
		for i := 0; i < cnt; i++ {
			scanner.Scan()
			relation[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", numWays(n, relation, k))
	}
}
