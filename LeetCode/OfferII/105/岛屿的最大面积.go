package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/ZL6zAn/
// ------------------------剑指 Offer II Problem 105------------------------
func maxAreaOfIsland(grid [][]int) int {
	var maxArea, curArea int
	var dfs func(int, int)
	dfs = func(x int, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			return
		}
		if grid[x][y] == 1 {
			curArea++
			grid[x][y] = 0
		}
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				curArea = 0
				dfs(i, j)
				maxArea = max(maxArea, curArea)
			}
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 105------------------------
/*
 * https://leetcode.cn/problems/ZL6zAn/
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.6MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", maxAreaOfIsland(grid))
	}
}
