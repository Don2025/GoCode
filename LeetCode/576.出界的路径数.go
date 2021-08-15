package main

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
