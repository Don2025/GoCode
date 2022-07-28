package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/solving-questions-with-brainpower/
//------------------------Leetcode Problem 2140------------------------
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

//------------------------Leetcode Problem 2140------------------------
/*
 * https://leetcode.cn/problems/solving-questions-with-brainpower/
 * 执行用时：228ms 在所有Go提交中击败了15.37%的用户
 * 占用内存：19.6MB 在所有Go提交中击败了64.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		questions := make([][]int, n)
		for i := range questions {
			scanner.Scan()
			questions[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", mostPoints(questions))
	}
}
