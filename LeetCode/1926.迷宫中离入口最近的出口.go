package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func nearestExit(maze [][]byte, entrance []int) int {
	type point struct{ x, y int }
	var next = []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	row, col := len(maze), len(maze[0])
	type queue struct{ x, y, step int }
	q := []queue{{entrance[0], entrance[1], 0}}
	maze[entrance[0]][entrance[1]] = '+'
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		for _, d := range next {
			nx, ny := t.x+d.x, t.y+d.y
			if nx < 0 || nx == row || ny < 0 || ny == col || maze[nx][ny] == '+' {
				continue
			}
			if nx == 0 || ny == 0 || nx == row-1 || ny == col-1 {
				return t.step + 1
			}
			maze[nx][ny] = '+'
			q = append(q, queue{nx, ny, t.step + 1})
		}
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		maze := make([][]byte, n)
		for i := 0; i < n; i++ {
			input.Scan()
			maze[i] = []byte(input.Text())
		}
		input.Scan()
		entrance := stringArrayToIntArray(strings.Fields(input.Text()))
		println(nearestExit(maze, entrance))
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
 * 执行用时：12ms 在所有Go提交中击败了78.95%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了65.79%的用户
**/
