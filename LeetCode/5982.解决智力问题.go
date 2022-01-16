package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		next := questions[i][1] + i + 1
		next = min(next, n)
		dp[i] = questions[i][0] + dp[next]
		dp[i] = max(dp[i], dp[i+1])
	}
	return int64(dp[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		questions := make([][]int, n)
		for i := range questions {
			input.Scan()
			questions[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(mostPoints(questions))
	}
}
