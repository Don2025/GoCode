package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func matrixBlockSum(mat [][]int, k int) [][]int {
	row, col := len(mat), len(mat[0])
	dp := make([][]int, row+1)
	dp[0] = make([]int, col+1)
	ans := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i+1] = make([]int, col+1)
		ans[i] = make([]int, col)
		for j := 0; j < col; j++ {
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + mat[i][j]
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			lx, ly := max(i-k, 0), max(j-k, 0)
			rx, ry := min(i+k, row-1), min(j+k, col-1)
			ans[i][j] = dp[rx+1][ry+1] - dp[lx][ry+1] - dp[rx+1][ly] + dp[lx][ly]
		}
	}
	return ans
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		mat := make([][]int, n)
		for i := range mat {
			input.Scan()
			mat[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := matrixBlockSum(mat, k)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
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
 * 执行用时：8ms 在所有Go提交中击败了78.26%的用户
 * 占用内存：6.1MB 在所有Go提交中击败了54.35%的用户
**/
