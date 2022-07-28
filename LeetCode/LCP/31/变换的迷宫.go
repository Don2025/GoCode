package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/Db3wC1/
// ------------------------LeetCode Cup Problem 31------------------------
func escapeMaze(maze [][]string) bool {
	type point struct{ x, y int }
	dir := []point{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	k, n, m := len(maze), len(maze[0]), len(maze[0][0])
	vis := make([][][][6]bool, k)
	for i := range vis {
		vis[i] = make([][][6]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([][6]bool, m)
		}
	}
	var f func(int, int, int, int) bool
	f = func(t int, x int, y int, s int) bool {
		if x < 0 || x >= n || y < 0 || y >= m || t+n-1-x+m-1-y >= k || vis[t][x][y][s] {
			return false
		}
		if x == n-1 && y == m-1 {
			return true
		}
		vis[t][x][y][s] = true
		if s>>1 == 1 {
			for _, d := range dir {
				if f(t+1, x+d.x, y+d.y, s^6) { // 跳出永久消除的位置
					return true
				}
			}
			return f(t+1, x, y, s)
		}
		if s>>1 == 0 && maze[t][x][y] == '#' && f(t, x, y, s|2) { // 尝试使用永久消除术
			return true
		}
		// 使用永久消除术已无法走到终点，遇到障碍只能使用临时消除术
		if maze[t][x][y] == '#' {
			if s&1 == 1 {
				return false
			}
			s |= 1
		}
		for _, d := range dir {
			if f(t+1, x+d.x, y+d.y, s) {
				return true
			}
		}
		return f(t+1, x, y, s)
	}
	return f(0, 0, 0, 0)
}

// ------------------------LeetCode Cup Problem 31------------------------
/*
 * https://leetcode.cn/problems/Db3wC1/
 * 执行用时：40ms 在所有Go提交中击败了33.33%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了33.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		maze := make([][]string, n)
		for i := range maze {
			scanner.Scan()
			maze[i] = strings.Fields(scanner.Text())
		}
		Printf("Output: %v\n", escapeMaze(maze))
	}
}
