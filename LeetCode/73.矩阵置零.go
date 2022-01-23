package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)
	for i := range matrix {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := range matrix {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			input.Scan()
			matrix[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		setZeroes(matrix)
		for _, x := range matrix {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
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
 * 执行用时：12ms 在所有Go提交中击败了83.63%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了81.16%的用户
**/
