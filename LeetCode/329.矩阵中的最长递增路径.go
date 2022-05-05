package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return 0
	}
	memo := make([][]int, row)
	for i := range memo {
		memo[i] = make([]int, col)
	}
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if memo[x][y] != 0 {
			return memo[x][y]
		}
		memo[x][y]++
		for _, dir := range dirs {
			newX, newY := x+dir.x, y+dir.y
			if newX >= 0 && newX < row && newY >= 0 && newY < col && matrix[x][y] < matrix[newX][newY] {
				memo[x][y] = max(memo[x][y], dfs(newX, newY)+1)
			}
		}
		return memo[x][y]
	}
	var ans int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ans = max(ans, dfs(i, j))
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			input.Scan()
			matrix[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(longestIncreasingPath(matrix))
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
 * 执行用时：28ms 在所有Go提交中击败了82.33%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了82.33%的用户
**/
