package main

import (
	"bufio"
	"os"
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		if s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '5' {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(translateNum(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了73.60%的用户
**/
