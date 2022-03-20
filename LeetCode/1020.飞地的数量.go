package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numEnclaves(grid [][]int) int {
	var dfs func([][]int, int, int)
	dfs = func(grid [][]int, i int, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	r, c := len(grid), len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if (i == 0 || j == 0 || i == r-1 || j == c-1) && grid[i][j] == 1 {
				dfs(grid, i, j)
			}
		}
	}
	var ans int
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(numEnclaves(grid))
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
 * 执行用时：48ms 在所有Go提交中击败了91.67%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了54.75%的用户
**/
