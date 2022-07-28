package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/PLYXKQ/
// ------------------------剑指 Offer II Problem 40------------------------
func maximalRectangle(matrix []string) int {
	var ans int
	if len(matrix) == 0 {
		return ans
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i := range matrix {
		left[i] = make([]int, n)
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ {
		up, down := make([]int, m), make([]int, m)
		var stack []int
		for i := 0; i < m; i++ {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				stack = stack[:len(stack)-1]
			}
			up[i] = -1
			if len(stack) > 0 {
				up[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		stack = []int{}
		for i := m - 1; i >= 0; i-- {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				stack = stack[:len(stack)-1]
			}
			down[i] = m
			if len(stack) > 0 {
				down[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		for i := range left {
			ans = max(ans, (down[i]-up[i]-1)*left[i][j])
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

// ------------------------剑指 Offer II Problem 40------------------------
/*
 * https://leetcode.cn/problems/PLYXKQ/
 * 执行用时：4ms 在所有Go提交中击败了52.17%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了28.26%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		matrix := strings.Fields(scanner.Text())
		Printf("Output: %d\n", maximalRectangle(matrix))
	}
}
