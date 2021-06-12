package main

func jumpFloor(number int) int {
	dp := make([]int, number+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

/*
 * 运行时间：9ms 超过33.36%用Go提交的代码
 * 占用内存：904KB 超过16.15%用Go提交的代码
**/
