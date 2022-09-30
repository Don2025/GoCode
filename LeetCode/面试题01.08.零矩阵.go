package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/zero-matrix-lcci/
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

/*
 * https://leetcode.cn/problems/zero-matrix-lcci/
 * 执行用时：4ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了48.76%的用户
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
		setZeroes(matrix)
		Printf("%v\n", matrix)
	}
}
