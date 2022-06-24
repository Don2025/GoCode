package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/minimum-path-sum/
//------------------------Leetcode Problem 64------------------------
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

//------------------------Leetcode Problem 64------------------------
/*
 * https://leetcode.cn/problems/minimum-path-sum/
 * 执行用时：4ms 在所有Go提交中击败了97.81%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", minPathSum(grid))
	}
}
