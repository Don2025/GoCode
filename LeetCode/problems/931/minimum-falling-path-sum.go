package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/minimum-falling-path-sum/
//------------------------Leetcode Problem 931------------------------
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == n-1 {
				matrix[i][j] += min(matrix[i-1][j-1], matrix[i-1][j])
			} else {
				matrix[i][j] += min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i-1][j+1]))
			}
		}
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, matrix[n-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 931------------------------
/*
 * https://leetcode.cn/problems/minimum-falling-path-sum/
 * 执行用时：12ms 在所有Go提交中击败了82.42%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了76.56%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", minFallingPathSum(matrix))
	}
}
