package main

import (
	"bufio"
	"os"
	"strconv"
)

func maximalSquare(matrix [][]byte) int {
	var maxSide int
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return maxSide
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				curMaxSide := min(row-i, col-j)
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for l := 0; l < k; l++ {
						if matrix[i+k][j+l] == '0' || matrix[i+l][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		matrix := make([][]byte, n)
		for i := 0; i < n; i++ {
			input.Scan()
			matrix[i] = []byte(input.Text())
		}
		println(maximalSquare(matrix))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了11.47%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了99.43%的用户
**/
