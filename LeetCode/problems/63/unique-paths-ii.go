package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/unique-paths-ii/
// ---------------------------Leetcode Problem 63------------------------
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

// ---------------------------Leetcode Problem 63------------------------
/*
 * https://leetcode.cn/problems/unique-paths-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了91.69%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		obstacleGrid := make([][]int, n)
		for i := range obstacleGrid {
			scanner.Scan()
			obstacleGrid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", uniquePathsWithObstacles(obstacleGrid))
	}
}
