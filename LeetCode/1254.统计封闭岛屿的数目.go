package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func closedIsland(grid [][]int) int {
	var dfs func([][]int, int, int) int
	dfs = func(grid [][]int, i int, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return 0
		}
		if grid[i][j] == 1 {
			return 1
		}
		grid[i][j] = 1
		ans := 1
		row := []int{0, 1, 0, -1}
		col := []int{1, 0, -1, 0}
		for k := 0; k < 4; k++ {
			ans = min(ans, dfs(grid, i+row[k], j+col[k]))
		}
		return ans
	}
	var cnt int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				cnt += dfs(grid, i, j)
			}
		}
	}
	return cnt
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
		for i := 0; i < n; i++ {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(closedIsland(grid))
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
 * 执行用时：8ms 在所有Go提交中击败了95.65%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了58.45%的用户
**/
