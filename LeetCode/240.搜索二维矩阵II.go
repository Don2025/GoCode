package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	x, y := len(matrix)-1, 0
	for x >= 0 && y < len(matrix[0]) {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			y += 1
		} else {
			x -= 1
		}
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			matrix[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		println(searchMatrix(matrix, target))
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
 * 执行用时：24ms 在所有Go提交中击败了68.37%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了96.31%的用户
**/

