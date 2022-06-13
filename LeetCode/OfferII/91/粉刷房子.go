package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/JEj789/
// ------------------------剑指 Offer II Problem 91------------------------
func minCost(costs [][]int) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0] = costs[0]
	for i := 1; i < len(costs); i++ {
		for j := 0; j < 3; j++ {
			prev1 := dp[(i-1)%2][(j+1)%3]
			prev2 := dp[(i-1)%2][(j+2)%3]
			dp[i%2][j] = min(prev1, prev2) + costs[i][j]
		}
	}
	last := (len(costs) - 1) & 1
	return min(dp[last][0], dp[last][1], dp[last][2])
}

func min(a ...int) int {
	val := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < val {
			val = a[i]
		}
	}
	return val
}

// ------------------------剑指 Offer II Problem 91------------------------
/*
 * https://leetcode.cn/problems/JEj789/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了85.12%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		costs := make([][]int, n)
		for i := range costs {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			costs[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", minCost(costs))
		Printf("Input the number of rows of matrix:")
	}
}
