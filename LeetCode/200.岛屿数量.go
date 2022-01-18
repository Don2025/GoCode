package main

import (
	"bufio"
	"os"
	"strconv"
)

func numIslands(grid [][]byte) int {
	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, r int, c int) {
		row, col := len(grid), len(grid[0])
		grid[r][c] = '0'
		if r-1 >= 0 && grid[r-1][c] == '1' {
			dfs(grid, r-1, c)
		}
		if r+1 < row && grid[r+1][c] == '1' {
			dfs(grid, r+1, c)
		}
		if c-1 >= 0 && grid[r][c-1] == '1' {
			dfs(grid, r, c-1)
		}
		if c+1 < col && grid[r][c+1] == '1' {
			dfs(grid, r, c+1)
		}
	}
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	var cnt int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		grid := make([][]byte, n)
		for i := 0; i < n; i++ {
			input.Scan()
			grid[i] = []byte(input.Text())
		}
		println(numIslands(grid))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了90.66%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了75.77%的用户
**/
