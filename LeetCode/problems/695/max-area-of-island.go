package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/max-area-of-island/
//------------------------Leetcode Problem 695------------------------
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

//------------------------Leetcode Problem 695------------------------
/*
 * https://leetcode.cn/problems/max-area-of-island/
 * 执行用时：12ms 在所有Go提交中击败了72.91%的用户
 * 占用内存：5MB 在所有Go提交中击败了36.88%的用户
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
		Printf("Output: %v\n", maxAreaOfIsland(grid))
	}
}
