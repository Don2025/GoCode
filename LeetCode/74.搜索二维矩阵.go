package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return false
	}
	begin, end := 0, row*col-1
	for begin <= end {
		mid := begin + (end-begin)>>1
		if matrix[mid/col][mid%col] > target {
			end = mid - 1
		} else if matrix[mid/col][mid%col] < target {
			begin = mid + 1
		} else {
			return true
		}
	}
	return false
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
 * 执行用时：4ms 在所有Go提交中击败了20.90%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了27.07%的用户
**/
