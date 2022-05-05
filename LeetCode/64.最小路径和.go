package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	for i := 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < col; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[row-1][col-1]
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
		grid := make([][]int, n)
		for i := range grid {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(minPathSum(grid))
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
 * 执行用时：4ms 在所有Go提交中击败了97.81%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了100.00%的用户
**/
