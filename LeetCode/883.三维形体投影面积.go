package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func projectionArea(grid [][]int) int {
	var xyArea, yzArea, zxArea int
	for i := range grid {
		var yzHeight, zxHeight int
		for j := range grid[i] {
			if grid[i][j] > 0 {
				xyArea++
			}
			yzHeight = max(yzHeight, grid[j][i])
			zxHeight = max(zxHeight, grid[i][j])
		}
		yzArea += yzHeight
		zxArea += zxHeight
	}
	return xyArea + yzArea + zxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		grid := make([][]int, n)
		for i := range grid {
			input.Scan()
			grid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(projectionArea(grid))
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
 * 执行用时：4ms 在所有Go提交中击败了90.00%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了100.00%的用户
**/
