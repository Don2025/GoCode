package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/making-a-large-island/
//------------------------Leetcode Problem 827------------------------
func largestIsland(grid [][]int) (ans int) {
	dirs := []struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	var cnt int
	index := 2
	area := map[int]int{}
	flag := true //  Mark whether the whole grid is land
	var dfs func(int, int)
	dfs = func(i, j int) {
		if visited[i][j] || grid[i][j] == 0 {
			return
		}
		visited[i][j] = true
		grid[i][j] = index
		cnt++
		for _, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] > 0 {
				dfs(x, y)
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				flag = false
			}
			if grid[i][j] == 1 && !visited[i][j] {
				cnt = 0
				index++
				dfs(i, j)
				area[index] = cnt
				index++
			}
		}
	}
	if flag { // If it is all land, return the area of the whole grid
		return n * m
	}
	visitedGrid := make(map[int]bool)
	for i := range grid {
		for j := range grid[i] {
			cnt = 1 // Record the number of islands after connection
			visitedGrid = map[int]bool{}
			if grid[i][j] == 0 {
				for _, dir := range dirs {
					x, y := i+dir.x, j+dir.y
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] > 0 && !visitedGrid[grid[x][y]] {
						cnt += area[grid[x][y]]
						visitedGrid[grid[x][y]] = true
					}
				}
			}
			if cnt > ans {
				ans = cnt
			}
		}
	}
	return
}

//------------------------Leetcode Problem 827------------------------
/*
 * https://leetcode.cn/problems/making-a-large-island/
 * 执行用时：296ms 在所有Go提交中击败了40.79%的用户
 * 占用内存：11.1MB 在所有Go提交中击败了38.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %d\n", largestIsland(grid))
	}
}
