package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/2bCMpM/
// ------------------------剑指 Offer II Problem 107------------------------
func updateMatrix(mat [][]int) [][]int {
	const INF = math.MaxInt32 >> 1
	m, n := len(mat), len(mat[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if mat[i][j] != 0 {
				dist[i][j] = INF
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j >= 1 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}

		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j+1 < n {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------------------剑指 Offer II Problem 107------------------------
/*
 * https://leetcode.cn/problems/2bCMpM/
 * 执行用时：40ms 在所有Go提交中击败了98.97%的用户
 * 占用内存：7MB 在所有Go提交中击败了98.97%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		grid := make([][]int, n)
		for i := range grid {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			grid[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", updateMatrix(grid))
		Printf("Input the number of rows of matrix:")
	}
}
