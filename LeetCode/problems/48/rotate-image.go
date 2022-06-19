package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/rotate-image/
//------------------------Leetcode Problem 48------------------------
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n >> 1); i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

//------------------------Leetcode Problem 48------------------------
/*
 * https://leetcode.cn/problems/rotate-image/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		rotate(matrix)
		Printf("Output: %v\n", matrix)
	}
}
