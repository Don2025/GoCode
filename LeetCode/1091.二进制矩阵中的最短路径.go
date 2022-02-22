package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 1
	}
	type point struct{ x, y int }
	d := []point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	queue := []int{0}
	grid[0][0] = 1
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur/n, cur%n
		for _, p := range d {
			nx, ny := x+p.x, y+p.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				queue = append(queue, nx*n+ny)
				grid[nx][ny] = grid[x][y] + 1
				if nx == m-1 && ny == n-1 {
					return grid[nx][ny]
				}
			}
		}
	}
	return -1
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
		println(shortestPathBinaryMatrix(grid))
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
 * 执行用时：36ms 在所有Go提交中击败了94.66%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了87.38%的用户
**/
