package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/max-increase-to-keep-city-skyline/
//------------------------Leetcode Problem 807------------------------
func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax := make([]int, len(grid))
	colMax := make([]int, len(grid))
	for i, row := range grid {
		for j, h := range row {
			rowMax[i] = max(rowMax[i], h)
			colMax[j] = max(colMax[j], h)
		}
	}
	var sum int
	for i, row := range grid {
		for j, h := range row {
			sum += min(rowMax[i], colMax[j]) - h
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 807------------------------
/*
 * https://leetcode.cn/problems/max-increase-to-keep-city-skyline/
 * 执行用时：12ms 在所有Go提交中击败了17.65%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了58.82%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", maxIncreaseKeepingSkyline(grid))
	}
}
