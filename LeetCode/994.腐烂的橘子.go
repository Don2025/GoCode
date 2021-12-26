package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func orangesRotting(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	var queue []int
	var cnt1 int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, i*m+j)
			} else if grid[i][j] == 1 {
				cnt1++
			}
		}
	}
	var time int
	for len(queue) > 0 && cnt1 > 0 {
		time++
		size := len(queue)
		for i := 0; i < size; i++ {
			orange := queue[0]
			queue = queue[1:]
			x, y := orange/m, orange%m
			if x > 0 && grid[x-1][y] == 1 { //上
				cnt1--
				grid[x-1][y] = 2
				queue = append(queue, (x-1)*m+y)
			}
			if x < n-1 && grid[x+1][y] == 1 { //下
				cnt1--
				grid[x+1][y] = 2
				queue = append(queue, (x+1)*m+y)
			}
			if y > 0 && grid[x][y-1] == 1 { //左
				cnt1--
				grid[x][y-1] = 2
				queue = append(queue, x*m+y-1)
			}
			if y < m-1 && grid[x][y+1] == 1 { //右
				cnt1--
				grid[x][y+1] = 2
				queue = append(queue, x*m+y+1)
			}
		}
	}
	if cnt1 == 0 {
		return time
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
		println(orangesRotting(grid))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了95.24%的用户
**/
