package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/projection-area-of-3d-shapes/
//------------------------Leetcode Problem 883------------------------
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

//------------------------Leetcode Problem 883------------------------
/*
 * https://leetcode.cn/problems/projection-area-of-3d-shapes/
 * 执行用时：4ms 在所有Go提交中击败了90.00%的用户
 * 占用内存：3.5MB 在所有Go提交中击败了100.00%的用户
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
		Printf("Output: %d\n", projectionArea(grid))
	}
}
