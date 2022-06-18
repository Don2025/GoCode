package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		setZeroes(matrix)
		fmt.Printf("%v\n", matrix)
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了83.63%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了81.16%的用户
**/
