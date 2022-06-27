package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/triangle/
//------------------------Leetcode Problem 120------------------------
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 120------------------------
/*
 * https://leetcode.cn/problems/triangle/
 * 执行用时：4ms 在所有Go提交中击败了96.04%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		triangle := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			triangle[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", minimumTotal(triangle))
	}
}
