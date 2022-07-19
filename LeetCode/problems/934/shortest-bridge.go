package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/shortest-bridge/
//------------------------Leetcode Problem 934------------------------
func shortestBridge(grid [][]int) int {
	n := len(grid)
	var dfs func([][]int, int, int)
	dfs = func(grid [][]int, i int, j int) {
		if i < 0 || i >= n || j < 0 || j >= n || grid[i][j] != 1 {
			return
		}
		grid[i][j] = 2
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	var flag bool
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j)
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	type queue struct{ x, y, step int }
	var q []queue
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, queue{i, j, 0})
			}
		}
	}
	type point struct{ x, y int }
	d := []point{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		for i := 0; i < len(d); i++ {
			nx, ny := t.x+d[i].x, t.y+d[i].y
			if nx < 0 || nx >= n || ny < 0 || ny >= n || grid[nx][ny] == 2 {
				continue
			}
			if grid[nx][ny] == 0 {
				q = append(q, queue{nx, ny, t.step + 1})
				grid[nx][ny] = 2
			} else if grid[nx][ny] == 1 {
				return t.step
			}
		}
	}
	return 0
}

//------------------------Leetcode Problem 934------------------------
/*
 * https://leetcode.cn/problems/shortest-bridge/
 * 执行用时：32ms 在所有Go提交中击败了88.24%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了83.53%的用户
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
		Printf("Output: %v\n", shortestBridge(grid))
	}
}
