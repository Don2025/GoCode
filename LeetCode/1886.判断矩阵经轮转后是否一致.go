package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		mat := make([][]int, n)
		for i := range mat {
			input.Scan()
			mat[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		n, _ = strconv.Atoi(input.Text())
		target := make([][]int, n)
		for i := range target {
			input.Scan()
			target[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(findRotation(mat, target))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了84.61%的用户
**/
