package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/determine-whether-matrix-can-be-obtained-by-rotation/
//------------------------Leetcode Problem 1886------------------------
func findRotation(mat [][]int, target [][]int) bool {
	var rotate func([][]int)
	rotate = func(matrix [][]int) {
		n := len(matrix) - 1
		for i := 0; i <= n>>1; i++ {
			for j := i; j < n-i; j++ {
				topLeft := matrix[i][j]
				topRight := matrix[j][n-i]
				bottomLeft := matrix[n-j][i]
				bottomRight := matrix[n-i][n-j]
				matrix[i][j] = bottomLeft
				matrix[j][n-i] = topLeft
				matrix[n-j][i] = bottomRight
				matrix[n-i][n-j] = topRight
			}
		}
	}
	var check func([][]int, [][]int) bool
	check = func(m1 [][]int, m2 [][]int) bool {
		for i := range m1 {
			for j := range m1[i] {
				if m1[i][j] != m2[i][j] {
					return false
				}
			}
		}
		return true
	}
	if check(mat, target) {
		return true
	}
	rotate(mat) // 转90度
	if check(mat, target) {
		return true
	}
	rotate(mat) // 转180度
	if check(mat, target) {
		return true
	}
	rotate(mat) // 转270度
	if check(mat, target) {
		return true
	}
	return false
}

//------------------------Leetcode Problem 1886------------------------
/*
 * https://leetcode.cn/problems/determine-whether-matrix-can-be-obtained-by-rotation/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了84.61%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		mat := make([][]int, n)
		for i := range mat {
			scanner.Scan()
			mat[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		target := make([][]int, n)
		for i := range target {
			scanner.Scan()
			target[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findRotation(mat, target))
	}
}
