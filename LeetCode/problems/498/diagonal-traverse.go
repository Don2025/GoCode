package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/diagonal-traverse/
//------------------------Leetcode Problem 498------------------------
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		if i&1 == 1 {
			x := max(0, i-n+1)
			y := min(i, n-1)
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := max(0, i-m+1)
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
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

//------------------------Leetcode Problem 498------------------------
/*
 * https://leetcode.cn/problems/diagonal-traverse/
 * 执行用时：20ms 在所有Go提交中击败了81.36%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了8.88%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		graph := make([][]int, n)
		for i := range graph {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			graph[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findDiagonalOrder(graph))
		Printf("Input the number of rows of matrix:")
	}
}
