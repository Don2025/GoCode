package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var queue [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	var hasOcean bool
	var point []int
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for len(queue) != 0 {
		point = queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		for i := 0; i < 4; i++ {
			newX, newY := x+dx[i], y+dy[i]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] != 0 {
				continue
			}
			grid[newX][newY] = grid[x][y] + 1
			hasOcean = true
			queue = append(queue, []int{newX, newY})
		}
	}
	if point == nil || !hasOcean {
		return -1
	}
	return grid[point[0]][point[1]] - 1
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
		println(maxDistance(grid))
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
 * 执行用时：48ms 在所有Go提交中击败了57.32%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了21.95%的用户
**/
