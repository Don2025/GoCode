package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, col)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j >= 1 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[col-1]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		obstacleGrid := make([][]int, n)
		for i := range obstacleGrid {
			input.Scan()
			obstacleGrid[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(uniquePathsWithObstacles(obstacleGrid))
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
 * 占用内存：2.3MB 在所有Go提交中击败了91.69%的用户
**/
