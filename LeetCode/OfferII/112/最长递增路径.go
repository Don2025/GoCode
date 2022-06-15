package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/fpTFWP/
// ------------------------剑指 Offer II Problem 112------------------------
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, columns := len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	var dfs func(int, int) int
	dfs = func(row, column int) int {
		if memo[row][column] != 0 {
			return memo[row][column]
		}
		memo[row][column]++
		for _, dir := range dirs {
			newRow, newColumn := row+dir[0], column+dir[1]
			if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
				memo[row][column] = max(memo[row][column], dfs(newRow, newColumn)+1)
			}
		}
		return memo[row][column]
	}
	var ans int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// ------------------------剑指 Offer II Problem 112------------------------
/*
 * https://leetcode.cn/problems/fpTFWP/
 * 执行用时：36ms 在所有Go提交中击败了24.73%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了38.71%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", longestIncreasingPath(matrix))
		Printf("Input the number of rows of matrix:")
	}
}
