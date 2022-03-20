package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func diagonalSum(mat [][]int) int {
	n := len(mat)
	sum, mid := 0, n>>1
	for i := 0; i < n; i++ {
		sum += mat[i][i] + mat[i][n-i-1]
	}
	sum -= mat[mid][mid] * (n & 1)
	return sum
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			mat[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(diagonalSum(mat))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了84.62%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了100.00%的用户
**/
