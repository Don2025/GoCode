package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/coloring-a-border/
//------------------------Leetcode Problem 1034------------------------
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	n, m := len(grid), len(grid[0])
	type point struct{ x, y int }
	dirs := [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	borders := []point{}
	cur := grid[row][col]
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}
	var dfs func(int, int)
	dfs = func(x int, y int) {
		vis[x][y] = true
		flag := false //是否为边界
		for _, dir := range dirs {
			newx, newy := x+dir.x, y+dir.y
			if !(newx >= 0 && newx < n && newy >= 0 && newy < m && grid[newx][newy] == cur) {
				flag = true
			} else if !vis[newx][newy] {
				vis[newx][newy] = true
				dfs(newx, newy)
			}
			if flag {
				borders = append(borders, point{x, y})
			}
		}
	}
	dfs(row, col)
	for _, border := range borders {
		grid[border.x][border.y] = color
	}
	return grid
}

//------------------------Leetcode Problem 1034------------------------
/*
 * https://leetcode.cn/problems/coloring-a-border/
 * 执行用时：20ms 在所有Go提交中击败了7.41%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了7.41%的用户
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
		scanner.Scan()
		var row, col, color int
		Sscanf(scanner.Text(), "%d %d %d", &row, &col, &color)
		Printf("Output: %v\n", colorBorder(grid, row, col, color))
	}
}
