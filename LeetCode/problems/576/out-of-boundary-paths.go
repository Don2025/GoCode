package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/out-of-boundary-paths/
//------------------------Leetcode Problem 576------------------------
const mod int = 1e9 + 7

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func findPaths(m, n, maxMove, startRow, startColumn int) int {
	var ans int
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	for i := 0; i < maxMove; i++ {
		dpNew := make([][]int, m)
		for j := range dpNew {
			dpNew[j] = make([]int, n)
		}
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				count := dp[j][k]
				if count > 0 {
					for _, dir := range dirs {
						j1, k1 := j+dir.x, k+dir.y
						if j1 >= 0 && j1 < m && k1 >= 0 && k1 < n {
							dpNew[j1][k1] = (dpNew[j1][k1] + count) % mod
						} else {
							ans = (ans + count) % mod
						}
					}
				}
			}
		}
		dp = dpNew
	}
	return ans
}

//------------------------Leetcode Problem 576------------------------
/*
 * https://leetcode.cn/problems/distribute-candies/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.9MB 在所有Go提交中击败了72.66%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m, n, maxMove, startRow, startColumn int
		_, err := Sscanf(scanner.Text(), "%d %d %d %d %d", &m, &n, &maxMove, &startRow, &startColumn)
		if err != nil {
			break
		}
		Printf("Output: %v\n", findPaths(m, n, maxMove, startRow, startColumn))
	}
}
