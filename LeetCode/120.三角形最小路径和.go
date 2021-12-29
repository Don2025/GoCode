package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		triangle := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			triangle[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(minimumTotal(triangle))
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
 * 执行用时：4ms 在所有Go提交中击败了96.04%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了100.00%的用户
**/
