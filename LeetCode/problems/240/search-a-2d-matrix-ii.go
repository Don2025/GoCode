package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/search-a-2d-matrix-ii/
//------------------------Leetcode Problem 240------------------------
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

//------------------------Leetcode Problem 240------------------------
/*
 * https://leetcode.cn/problems/search-a-2d-matrix-ii/
 * 执行用时：24ms 在所有Go提交中击败了68.37%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了96.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			matrix[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", searchMatrix(matrix, target))
	}
}
