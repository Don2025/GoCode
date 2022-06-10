package main

import "math"

func cutRope(number int) int {
	if number == 2 {
		return 1
	} else if number == 3 {
		return 2
	}
	dp := make([]int, number+1, math.MinInt32)
	for i := 1; i <= 4; i++ {
		dp[i] = i
	}
	for i := 5; i <= number; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*dp[i-j])
		}
	}
	return dp[number]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 运行时间：8ms 超过0.00%用Go提交的代码
 * 占用内存：948KB 超过16.00%用Go提交的代码
**/
