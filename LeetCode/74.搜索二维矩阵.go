package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := range matrix {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text()))
		}
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		println(searchMatrix(matrix, target))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了20.90%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了27.07%的用户
**/
