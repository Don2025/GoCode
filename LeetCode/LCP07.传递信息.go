package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		cnt, _ := strconv.Atoi(input.Text())
		relation := make([][]int, cnt)
		for i := 0; i < cnt; i++ {
			input.Scan()
			relation[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(numWays(n, relation, k))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了84.21%的用户
**/
