package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-of-closed-islands/
//------------------------Leetcode Problem 1254------------------------
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

//------------------------Leetcode Problem 1254------------------------
/*
 * https://leetcode.cn/problems/number-of-closed-islands/
 * 执行用时：8ms 在所有Go提交中击败了95.65%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了58.45%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		println(closedIsland(grid))
	}
}
