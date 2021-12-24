package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxAreaOfIsland(grid [][]int) int {
	var dfs func([][]int, int, int) int
	dfs = func(g [][]int, i, j int) int {
		if i < 0 || j < 0 || i >= len(g) || j >= len(g[i]) || g[i][j] == 0 {
			return 0
		}
		g[i][j] = 0
		cnt := 1 + dfs(g, i-1, j) + dfs(g, i, j+1) + dfs(g, i+1, j) + dfs(g, i, j-1)
		return cnt
	}
	var ans int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(grid, i, j))
			}
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
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(maxAreaOfIsland(grid))
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
 * 执行用时：12ms 在所有Go提交中击败了72.91%的用户
 * 占用内存：5MB 在所有Go提交中击败了36.88%的用户
**/
