package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/largest-local-values-in-a-matrix/
//------------------------Leetcode Problem 2373------------------------
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 1; i < n-1; i++ {
		row := make([]int, n-2)
		for j := 1; j < n-1; j++ {
			max := grid[i][j]
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if grid[r][c] > max {
						max = grid[r][c]
					}
				}
			}
			row[j-1] = max
		}
		ans[i-1] = row
	}
	return ans
}

//------------------------Leetcode Problem 2373------------------------
/*
 * https://leetcode.cn/problems/largest-local-values-in-a-matrix/
 * 执行用时：12ms 在所有Go提交中击败了70.00%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了6.67%的用户
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
		ans := largestLocal(grid)
		fmt.Printf("%v\n", ans)
	}
}
