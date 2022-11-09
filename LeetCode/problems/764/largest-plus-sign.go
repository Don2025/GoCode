package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/largest-plus-sign/
//------------------------Leetcode Problem 764------------------------
func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}
	banned := make(map[int]bool)
	for _, mine := range mines {
		banned[mine[0]*n+mine[1]] = true
	}
	for i := 0; i < n; i++ {
		var cnt int
		// left
		for j := 0; j < n; j++ {
			if banned[i*n+j] {
				cnt = 0
			} else {
				cnt++
			}
			dp[i][j] = min(dp[i][j], cnt)
		}
		cnt = 0
		// right
		for j := n - 1; j >= 0; j-- {
			if banned[i*n+j] {
				cnt = 0
			} else {
				cnt++
			}
			dp[i][j] = min(dp[i][j], cnt)
		}
	}
	for i := 0; i < n; i++ {
		cnt := 0
		// up
		for j := 0; j < n; j++ {
			if banned[j*n+i] {
				cnt = 0
			} else {
				cnt++
			}
			dp[j][i] = min(dp[j][i], cnt)
		}
		// down
		cnt = 0
		for j := n - 1; j >= 0; j-- {
			if banned[j*n+i] {
				cnt = 0
			} else {
				cnt++
			}
			dp[j][i] = min(dp[j][i], cnt)
			ans = max(ans, dp[j][i])
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 764------------------------
/*
 * https://leetcode.cn/problems/largest-plus-sign/
 * 执行用时：176ms 在所有Go提交中击败了33.33%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了50.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		mines := make([][]int, m)
		for i := range mines {
			scanner.Scan()
			mines[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", orderOfLargestPlusSign(n, mines))
	}
}
