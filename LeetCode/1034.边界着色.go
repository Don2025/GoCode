package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		row, col, color := arr[0], arr[1], arr[2]
		ans := colorBorder(grid, row, col, color)
		for _, a := range ans {
			for _, x := range a {
				fmt.Printf("%d ", x)
			}
			fmt.Println()
		}
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
 * 执行用时：20ms 在所有Go提交中击败了7.41%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了7.41%的用户
**/
